{{/*
    Custom Reports Admins CC v3
    
    Made By Devonte#0745 / Naru#6203
    
    Recommended Trigger Type: Regex
    Recommended Trigger     : \A-r(?:eport)?a(?:dmin)?(?:\s+|\z)
    
    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
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
            {{else}}Unknown Report. Check that you have the correct ID.{{end}}
        {{else}}{{"Insufficient arguments.\nUsage: `-reportadmin reopen <messageID> <reason>` - Reason must be longer than 1 word."}}{{end}}
    {{else if eq $cmd "reacthelp"}}
        {{sendMessage nil (cembed "title" "Reactions Help" "description" "✅ - Marks a report as finished / done (solved without mod action).\n❎ - Marks a report as ignored or false report (no mod action)\n\n🛡 - Displays the reactions for mod actions:\n\n❌ - cancels mod action and returns to default reaction menu\n⚠ - Warns the user\n🔇 - Mutes the user (server default time)\n👢 - Kicks the user\n🔨 - ~~YEETS~~ Bans the user\n\nAll reasons for mod actions can be configured in the reaction-listener code." "color" 0xfe3025)}}
    {{else if eq $cmd "resetreactions" "rr"}}
        {{if eq (len .CmdArgs) 2}}
            {{if ($msg := getMessage $logChannel (index .CmdArgs 1))}}
                {{if and $msg.Embeds $msg.Author.Bot (eq $msg.Author.ID 204255221017214977)}}
                    {{if (reFind `(?i)report` (index $msg.Embeds 0).Author.Name)}}
                        {{deleteAllMessageReactions $logChannel (index .CmdArgs 1)}}
                        {{addMessageReactions $logChannel (index .CmdArgs 1) "✅" "❎" "🛡"}}
                        Successfully reset the reactions!
                        {{deleteTrigger 6}}{{deleteResponse 6}}
                    {{else}}This message cannot be verified as a report message. Please try again or contact support (Github source).{{end}}
                {{else}}Invalid report message ID.{{end}}
            {{else}}Unknown Report. Check that you have the correct ID.{{end}}
        {{else}}{{"Insufficient arguments.\nUsage: `-reportadmin resetreactions <messageID>`"}}{{end}}
    {{else if eq $cmd "history"}}
        {{if eq (len .CmdArgs) 2}}
            {{if ($user := (userArg (index .CmdArgs 1)))}}
                {{$hst := "< No Report History >"}}
                {{with (dbGet $user.ID "rhistory")}}{{$hst = .Value}}{{end}}
                {{sendMessage nil (cembed "author" (sdict "name" (print "Report History of " $user.String) "icon_url" ($user.AvatarURL "256")) "description" (print "```\n" $hst "\n```") "color" 0xfe3025 "timestamp" currentTime)}}
            {{else}}Unknown user specified. Please try again.{{end}}
        {{else}}{{"Insufficient arguments.\nUsage: `-reportadmin history <@User/ID>`"}}{{end}}
    {{else if eq $cmd "deletehistory" "delhistory"}}
        {{if eq (len .CmdArgs) 2}}
            {{if ($user := (userArg (index .CmdArgs 1)))}}
                {{if (dbGet $user.ID "rhistory")}}
                    {{dbDel $user.ID "rhistory"}}
                    Successfully deleted the report history of `{{$user.String}}`!
                    {{deleteTrigger 6}}{{deleteResponse 6}}
                {{else}}This user does not have any previous reports.{{end}}
            {{else}}Unknown user specified. Please try again.{{end}}
        {{else}}{{"Insufficient arguments.\nUsage: `-reportadmin deletehistory <@User/ID>`"}}{{end}}
    {{else}}{{sendMessage nil (cembed "title" "Report Admin Commands" "description" "Command Aliases: `ra, radmin, reporta`\n\n`-ra reopen <messageID> <reason>` - Reopens a closed report.\n\n`-ra resetreactions <messageID>` - Resets the reactions on a report\nAliases: `rr`\n\n`-ra history <@User/ID>` - Displays the report history of a user\n\n`-ra deletehistory <@User/ID>` - Deletes the report history of a user\nAliases: `delhistory`\n\n`-ra reacthelp` - Reactions Help page." "color" 0xfe3025)}}{{end}}
{{else}}{{sendMessage nil (cembed "title" "Report Admin Commands" "description" "Command Aliases: `ra, radmin, reporta`\n\n`-ra reopen <messageID> <reason>` - Reopens a closed report.\n\n`-ra resetreactions <messageID>` - Resets the reactions on a report\nAliases: `rr`\n\n`-ra history <@User/ID>` - Displays the report history of a user\n\n`-ra deletehistory <@User/ID>` - Deletes the report history of a user\nAliases: `delhistory`\n\n`-ra reacthelp` - Reactions Help page." "color" 0xfe3025)}}{{end}}
