{{/*
    Custom Reports CC
    
    Made By Devonte#0745 / Naru#6203
    Contributors: DZ#6669, Piter#5960
    
    Recommended Trigger Type: Command
    Recommended Trigger     : report
*/}}

{{/* THINGS TO CHANGE */}}
{{$logChannel := }} {{/* CHANNEL TO LOG REPORTS */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}
{{$ping := 0}}
{{if .CmdArgs}}
    {{if or .Message.Mentions (reFind `\d{17,19}` (index .CmdArgs 0))}}
        {{if (ge (len .CmdArgs) 4)}}{{$user := ""}}
            {{if .Message.Mentions}}
                {{$user = index .Message.Mentions 0}}
            {{else}}
                {{$user = userArg (index .CmdArgs 0)}}
            {{end}}
            {{if eq $user.ID .User.ID}}
                {{print .User.Mention ", you cant report yourself."}}
            {{else}}
                {{$re := joinStr " " (slice .CmdArgs 1)}}
                {{$logs := exec "logs"}}{{$ht := "No reports"}}{{$old := ""}}{{$new := ""}}
                {{if (dbGet .User.ID "reports")}}
                    {{range (dbGetPattern .User.ID "reports" 5 0) -}}
                        {{- $ht = (reFind `^([a-zA-Z]{3,}-\d+-\d+-\d+:\d+:\d+)` (str .Value)) -}}
                    {{- end}}
                {{end}}
                {{$report := cembed
                    "author" (sdict "name" (print "New Report from " .User.String) "icon_url" (.User.AvatarURL "256"))
                    "thumbnail" (sdict "url" ($user.AvatarURL "256"))
                    "description" (print "Not reviewed yet. [\u200b](" $user.ID ")")
                    "fields" (cslice
                    (sdict "name" "Reason" "value" $re "inline" false)
                    (sdict "name" "Reported User" "value" (print $user.Mention " (ID " $user.ID ")") "inline" false)
                    (sdict "name" "Info" "value" (print "<#" .Channel.ID "> (ID " .Channel.ID ") - [Message Logs](" $logs ")\nTime - `" (currentTime.Format "Mon 02 Jan 15:04:05") "`") "inline" false)
                    (sdict "name" "History" "value" (str $ht) "inline" false))
                    "color" 16698149
                    "footer" (sdict "text" "React for options")
                    "timestamp" currentTime}}{{$x := 0}}
                {{if $ping}}
                    {{$x = sendMessageNoEscapeRetID $logChannel (complexMessage "content" (mentionRoleID $ping) "embed" $report)}}
                {{else}}
                    {{$x = sendMessageRetID $logChannel $report}}
                {{end}}
                {{deleteTrigger 0}}{{"User reported to the Staff Team."}}
                {{sleep 1}}{{addMessageReactions $logChannel $x "✅" "❎" "🛡"}}
                {{with (dbGet .User.ID "reports").Value}}
                    {{$old = .}}
                    {{$new = joinStr "\n" $old (print (currentTime.Format "Mon-02-01-15:04:05") " " $re)}}
                {{else}}
                    {{dbSet .User.ID "reports" (print (currentTime.Format "Mon-02-01-15:04:05") " " $re)}}
                {{end}}
            {{end}}
        {{else}}
            {{print .User.Mention ", your report needs to be longer than **2** words."}}
        {{end}}
    {{else}}
        {{print .User.Mention ", you need to specify someone to report."}}
    {{end}}
{{else}}
    {{"Command: `-report @user/ID <reason>`\n\nYour report must be longer than 2 words. You cant report yourself."}}
{{end}}
