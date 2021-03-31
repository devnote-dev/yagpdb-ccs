{{/*
    Custom Reports Main CC v4
    
    Made By Devonte#0745 / Naru#6203
    Contributors: DZ#6669, Piter#5960, Lyonhaert#3393
    
    Recommended Trigger Type: Command
    Recommended Trigger     : report

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$logChannel := CHANNEL-ID}} {{/* Channel ID to log reports */}}

{{$ping := 0}} {{/* Role to ping when report (set to 0 for none) */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if .CmdArgs}}
    {{if or .Message.Mentions (reFind `\d{17,19}` (index .CmdArgs 0))}}
        {{if (ge (len .CmdArgs) 3)}}{{$user := ""}}
            {{if .Message.Mentions}}{{$user = index .Message.Mentions 0}}
            {{else}}{{$user = userArg (index .CmdArgs 0)}}{{end}}
            {{if eq $user.ID .User.ID}}{{.User.Mention}} You cant report yourself, **silly**.
            {{else}}
                {{$re := joinStr " " (slice .CmdArgs 1)}}{{$logs := exec "logs"}}{{$hst := ""}}
                {{with (dbGet $user.ID "rhistory")}}
                {{$hst = .Value}}{{if ge (len $hst) 800}}{{$hst = "To Many Reports! View them with the `ra history` cmd if added."}}{{end}}
                {{dbSet .UserID .Key (print .Value "\n" (currentTime.Format "02-01-2006-15:04:05") " :: " $re)}}
                {{else}}{{dbSet $user.ID "rhistory" (print (currentTime.Format "02-01-2006-15:04:05") " :: " $re)}}{{end}}
                {{$report := sdict
                    "author" (sdict "name" (print "New Report from " .User.String) "icon_url" (.User.AvatarURL "256"))
                    "thumbnail" (sdict "url" ($user.AvatarURL "256"))
                    "description" (print "Not reviewed yet. [\u200b](" $user.ID ")")
                    "fields" (cslice
                    (sdict "name" "Report Reason" "value" $re "inline" false)
                    (sdict "name" "Reported User" "value" (print $user.Mention " (ID " $user.ID ")") "inline" false)
                    (sdict "name" "Info" "value" (print "Channel: <#" .Channel.ID "> (ID " .Channel.ID ")\nTime: " (currentTime.Format "Mon 02 Jan 2006 15:04:05") "\n[Message Logs](" $logs ")") "inline" false)
                    (sdict "name" "History" "value" (print "```\n" (or $hst "None recorded") "\n```") "inline" false))
                    "color" 16698149
                    "footer" (sdict "text" "React for options")
                    "timestamp" currentTime}}{{$x := 0}}
                {{if .Message.Attachments}}{{$report.Set "image" (sdict "url" (index .Message.Attachments 0).URL)}}
                {{else if .Message.Embeds}}{{if eq (index .Message.Embeds 0).Type "image"}}
                {{$report.Set "image" (sdict "url" (index .Message.Embeds 0).URL)}}{{end}}{{end}}
                {{deleteTrigger 0}}{{$x := "Uh Oh"}}
                {{if $ping}}{{$x = sendMessageNoEscapeRetID $logChannel (complexMessage "content" (mentionRoleID $ping) "embed" (cembed $report))}}
                {{else}}{{$x = sendMessageRetID $logChannel (cembed $report)}}{{end}}
                User reported to the Staff Team.
                {{sleep 1}}{{addMessageReactions $logChannel $x "✅" "❎" "🛡"}}
            {{end}}
        {{else}}{{.User.Mention}} Your report needs to be longer than **1** word.{{end}}
    {{else}}{{.User.Mention}} You need to specify someone to report.{{end}}
{{else}}Command: `-report @user/ID <reason>`{{"\n"}}Your report must be longer than 1 word. You cant report yourself.{{end}}
