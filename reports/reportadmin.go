{{/*
    Custom Reports Admins CC v2
    
    Made By Devonte#0745 / Naru#6203
    
    Recommended Trigger Type: Regex
    Recommended Trigger     : \A-r(?:eport)?a(?:dmin)?(?:\s+|\z)
*/}}

{{/* THINGS TO CHANGE */}}

{{$logChannel := }} {{/* Channel ID to log reports */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if .CmdArgs}}
    {{$cmd := index .CmdArgs 0}}
    {{if eq $cmd "reopen"}}
        {{if ge (len .CmdArgs) 4}}
            {{$reason := joinStr " " (slice .CmdArgs 2)}}
            {{if ($report := index (getMessage $logChannel (index .CmdArgs 1)).Embeds 0|structToSdict)}}
                {{range $k, $v := $report}}{{if eq (kindOf $v true) "struct"}}{{$report.Set $k (structToSdict $v)}}{{end}}{{end}}
                {{$rUser := (userArg (reReplace `\(|\)` (reFind `\d{17,19}` $report.Description) "")).ID}}
                {{with $report}}
                    {{.Set "description" (print "Report **reopened** by " $.User.String "\n**Reason:** " $reason " [\u200b](" $rUser ")")}}
                    {{.Set "color" 0xfeb225}}
                    {{.Author.Set "icon_url" .Author.IconURL}}
                {{end}}
                {{editMessage $logChannel (index .CmdArgs 1) (complexMessageEdit "embed" $report)}}
                {{dbSet 7 "reopen" true}}
                {{addMessageReactions $logChannel (index .CmdArgs 1) "✅" "❎" "🛡"}}
                Successfully reopened the report! Reactions will be added back shortly.
                {{deleteTrigger 6}}{{deleteResponse 6}}
            {{else}}
                Unknown Report. Check that you have the correct ID.
            {{end}}
        {{else}}
            {{print .User.Mention " Something went wrong with that command! Make sure you format it like this:\n`-reportadmin reopen <messageID> <reason>` - Reason must be 2 words or more."}}
        {{end}}
    {{else if eq $cmd "reacthelp"}}
        {{sendMessage nil (cembed "title" "Reactions Help" "description" "✅ - Marks a report as finished / done (solved without mod action).\n❎ - Marks a report as ignored or false report (no mod action)\n\n🛡 - Displays the reactions for mod actions:\n\n❌ - cancels mod action and returns to default reaction menu\n⚠ - Warns the user\n🔇 - Mutes the user (server default time)\n👢 - Kicks the user\n🔨 - ~~YEETS~~ Bans the user\n\nAll reasons for mod actions can be configured in the reaction-listener code." "color" 0xfe3025)}}
    {{else}}
        Unknown command. Please try again.
    {{end}}
{{else}}
    {{sendMessage nil (cembed "title" "Report Admin Commands" "description" "Command Aliases: `ra, radmin, reporta`\n\n`-ra reopen <messageID> <reason>` - Reopens a closed report.\n`-ra reacthelp` - Reactions Help page." "color" 0xfe3025)}}
{{end}}
