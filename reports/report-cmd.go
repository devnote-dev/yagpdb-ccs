{{/*
    Custom Reports Main CC v3
    
    Made By Devonte#0745 / Naru#6203
    Contributors: DZ#6669, Piter#5960, Lyonhaert#3393
    
    Recommended Trigger Type: Command
    Recommended Trigger     : report
*/}}

{{/* THINGS TO CHANGE */}}

{{$logChannel := }} {{/* Channel ID to log reports */}}

{{$ping := 0}} {{/* Role to ping when report (set to 0 for none) */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if .CmdArgs}}
    {{if or .Message.Mentions (reFind `\d{17,19}` (index .CmdArgs 0))}}
        {{if (ge (len .CmdArgs) 3)}}{{$user := ""}}
            {{if .Message.Mentions}}
                {{$user = index .Message.Mentions 0}}
            {{else}}
                {{$user = userArg (index .CmdArgs 0)}}
            {{end}}
            {{if eq $user.ID .User.ID}}
                {{print .User.Mention ", you cant report yourself."}}
            {{else}}
                {{$re := joinStr " " (slice .CmdArgs 1)}}{{$logs := exec "logs"}}{{$hst := ""}}
                {{if (dbGet $user.ID "rhistory")}}
                    {{range (dbGetPattern $user.ID "rhistory%" 7 0)}}
                        {{$hst = .Value}}
                    {{end}}
                    {{dbSet $user.ID "rhistory" (print (dbGet $user.ID "rhistory").Value "\n" (currentTime.Format "02-01-2006-15:04:05") " :: " $re)}}
                {{else}}
                    {{dbSet $user.ID "rhistory" (print (currentTime.Format "02-01-2006-15:04:05") " :: " $re)}}
                {{end}}
                {{$report := cembed
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
                {{if $ping}}
                    {{$x = sendMessageNoEscapeRetID $logChannel (complexMessage "content" (mentionRoleID $ping) "embed" $report)}}
                {{else}}
                    {{$x = sendMessageRetID $logChannel $report}}
                {{end}}
                {{deleteTrigger 0}}{{"User reported to the Staff Team."}}
                {{sleep 2}}{{addMessageReactions $logChannel $x "✅" "❎" "🛡"}}
            {{end}}
        {{else}}
            {{print .User.Mention ", your report needs to be longer than **1** word."}}
        {{end}}
    {{else}}
        {{print .User.Mention ", you need to specify someone to report."}}
    {{end}}
{{else}}
    {{"Command: `-report @user/ID <reason>`\nYour report must be longer than 1 word. You cant report yourself."}}
{{end}}
