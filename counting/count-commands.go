{{/*

    Counting Pagination Reaction CC
    Made By Devonte#0745 / Naru#6203
    Modified By: WickedWizard#3588

    Trigger Type: Regex `\A(\-\s?|<@!?204255221017214977>\s*)count(?:\s+|\z)`

    Replace Prefix if its not `-`

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/NaruDevnote/yagpdb-ccs

*/}}

{{/*Config Vars Start*/}}
{{$cheaterrole := 0}}{{/*ID of your Cheater role, not the temporary one. **Mandatory***/}}
{{$modlog := 0}}{{/*ID of the ModLog Channel. Optional*/}}
{{/*Config Vars End*/}}

{{/*DON'T TOUCH     RAWR :)*/}}
{{if .CmdArgs}}
    {{$cmd := lower (index .CmdArgs 0)}}
    {{$col := 16777215}}{{$p := 0}}{{$r := .Member.Roles}}{{range .Guild.Roles}}{{if and (in $r .ID) (.Color) (lt $p .Position)}}{{$p = .Position}}{{$col = .Color}}{{end}}{{end}}
    {{$check := sdict (dbGet 3 "Count").Value}}
    {{$tracker := $check.Get "Tracker"}}
    {{$count := $check.Get "Count/SkipCount"}}
    {{$skipcount := $check.Get "SkipCount Data"}}
    {{$channel := $check.Get "Channel ID"}}
    {{if eq $cmd "help" "command" "commands"}}
        {{$embed := cembed
            "title" (joinStr "" "Counting Help")
            "description" "`count help/commands` - Shows this message\n\n`count [@User/ID]` - Shows the count of a user, or your own if no user is specified\n`count info <@User/ID>` - Shows the counting info of a specified user\n`count lb/leaderboard` - Shows the counting leaderboard\n\n`count next` - Shows the next number to count\n`count set <number>` - Changes the current count (mods only) \n`count ban <@User/ID>` - Bans a User from Counting (mods only) \n`count unban <@User/ID>` - Unbans a user from Counting (mods only) \n`count shift` - Shifts from one method of counting to another \n`count shifthelp` - Shows Help for shifting from one method of Counting to another"
            "color" $col
            "timestamp" currentTime}}
        {{sendMessage nil $embed}}
    {{else if eq $cmd "info"}}
        {{$args := parseArgs 1 "" (carg "string" "command") (carg "userid" "ID/Mention")}}
        {{$user := ""}}
        {{if $args.IsSet 1}}
            {{$user = userArg ($args.Get 1)}}
        {{else}}
            {{$user = userArg (.User.ID)}}
        {{end}}
        {{if $tracker}}
            {{$info := cembed
                "title" (joinStr "" "Info of " $user.String)
                "description" (print "• Current Count: " (dbGet $user.ID "Counter_Tracker").Value "\n• Started Counting: " (humanizeTimeSinceDays (dbGet $user.ID "Counter_Tracker").CreatedAt) "ago \n• Last Counted: " (humanizeTimeSinceDays (dbGet $user.ID "Counter_Tracker").UpdatedAt) " ago")
                "color" $col
                "timestamp" currentTime}}
            {{sendMessage nil $embed}}
        {{else}}
            You haven't enabled Tracker in the Main Counting CC, so I cannot retrieve this value.
        {{end}}
    {{else if eq $cmd "next"}}
        {{$words := "The next number is:"}}{{$n := ""}}
        {{if $count}}
            {{$n = add 1 (dbGet 0 "Count").Value}}
        {{else}}
            {{$n = add 2 (dbGet 0 "Count").Value}}
        {{end}}
        {{$words := "The next number is:"}} {{$n}}
    {{else if eq $cmd "set"}}
        {{if reFind `ManageMessages` (exec "viewperms")}}
            {{$args := parseArgs 2 "" (carg "string" "Command") (carg "int" "Number to Set to")}}
            {{if $skipcount}}
                {{dbSet 0 "Count" (toString ($args.Get 1))}}
                {{if $modlog}}
                    {{sendMessage $modlog (print "Counting number has been set to " ($args.Get 1))}}
                {{end}}
            {{else}}
                {{if $skipcount}}
                    {{if not (mod ($args.Get 1) 2)}}
                        Set the number to {{$args.Get 1}}
                        {{if $modlog}}
                            {{sendMessage $modlog (print "Counting number has been set to " ($args.Get 1))}}
                        {{end}}
                    {{else}}
                        Cannot set {{$args.Get 1}} as the number because, the system is set to allow even numbers only.
                    {{end}}
                {{else}}
                    {{if (mod ($args.Get 0) 2)}}
                        Set the number to {{$args.Get 1}}
                        {{if $modlog}}
                            {{sendMessage $modlog (print "Counting number has been set to " ($args.Get 1))}}
                        {{end}}
                    {{else}}
                        Cannot set {{$args.Get 1}} as the number because, the system is set to allow odd numbers only.
                    {{end}}
                {{end}}
            {{end}}
        {{else}}
            You lack the perms `ManageMessages` to run this command
        {{end}}
    {{else if eq $cmd "ban"}}
        {{if reFind `ManageMessages` (exec "viewperms")}}
            {{$args := parseArgs 2 "" (carg "string" "Command") (carg "userid" "ID/Mention")}}
            {{$user := userArg ($args.Get 1)}}
            {{giveRoleID $user.ID $cheaterrole}}
            {{dbSet $user.ID "Banned" (toString .User.ID)}}
            Banned {{$user.Mention}}
            {{if $modlog}}
                {{sendMessage $modlog (print .User.Mention " banned " $user.Mention " from counting")}}
            {{end}}
        {{else}}
            You lack the perms `ManageMessages` to run this command
        {{end}}
    {{else if eq $cmd "unban"}}
        {{if reFind `ManageMessages` (exec "viewperms")}}
            {{$args := parseArgs 2 "" (carg "string" "Command") (carg "userid" "ID/Mention")}}
            {{$user := userArg ($args.Get 1)}}
            {{if dbGet $user.ID "Banned"}}
                {{dbDel $user.ID "Banned"}}
                {{takeRoleID $user.ID $cheaterrole}}
                Unbanned {{$user.Mention}}
                {{if $modlog}}
                    {{sendMessage $modlog (print .User.Mention " unbanned " $user.Mention " from counting")}}
                {{end}}
            {{else}}
                This user has not been banned.
            {{end}}
        {{else}}
            You lack the perms `ManageMessages` to run this command
        {{end}}
    {{else if eq $cmd "leaderboard" "lb" "leaderb" "lboard"}}
        {{$list := ""}}{{$rank := 1}}
        {{if $tracker}}
            {{range dbTopEntries "Counter_Tracker" 10 0}}
                {{- $list = print $list (printf "%-4d %4d\t %-16s" $rank (toInt .Value) .User.String) "\n" -}}
                {{- $rank = add $rank 1 -}}
            {{- else -}}
                {{- $list = "No entries to display." -}}
            {{- end -}}
        {{else}}
            {{$list = "This server does not have Tracker enabled, so I cannot get the Leaderboard"}}
        {{end}}
        {{$embed := cembed
            "title" "Counting Leaderboard"
            "description" (print "```\n# -- Count -- User\n" $list "\n```")
            "color" 816464
            "footer" (sdict "text" "Page: 1")
            "timestamp" currentTime}}
        {{$id := sendMessageRetID nil $embed}}
        {{addMessageReactions nil $id "⬅" "➡" "🗑"}}
    {{else if eq $cmd "shifthelp"}}
        {{$embed := cembed
            "title" "Shifting from One Counting Method to Another"
            "description" "Do you want to shift from one counting method to another? \nLike from Normal to Skip or Skip to Normal or any other combination? \n\n**__THIS WILL RESET ALL YOUR DATABASES, EXCLUDING THE TRACKER__** \nRun `shift` twice if you're sure of this."
            "timestamp" currentTime
            "color" $col
            "footer" (sdict "text" "Warning:- This cannot be reverted at any cost.")}}
        {{sendMessage nil $embed}}
    {{else if eq $cmd "shift"}}
        {{if reFind `Administrator` (exec "viewperms")}}
            {{if not (dbGet 0 "Shift")}}
                {{sendMessage nil "Are you sure that you want to shift from one counting system to another? \nIf yes, run this command in the next 5 minutes. If you don't know about this, use `shifthelp` for more info."}}
                {{dbSetExpire 0 "Shift" true 300}}
            {{else}}
                {{$channel := sdict (dbGet 3 "Count").Value}}
                {{dbDel 0 "Count"}}
                {{dbDel 1 "Count"}}
                {{dbDel 2 "Count"}}
                {{dbDel 3 "Count"}}
                {{dbDel 0 "Shift"}}
                {{sendMessage nil "All the Count Databases, except the Tracker have been reset. To set it up again, run the Main Counting Command"}}
                {{editChannelTopic ($channel.Get "Channel ID") "This System needs to be Setup again."}}
                {{if $modlog}}
                    {{sendMessage $modlog (print "The counting system has been reset to shift from one method to another. Executed by " .User.Mention)}}
                {{end}}
            {{end}}
        {{else}}
            {{sendMessage nil "You lack the perms `Administrator` to run this command."}}
        {{end}}
    {{else}}
        {{if $tracker}}
            {{if ($user := getMember $cmd)}}
                {{with (dbGet $user.User.ID "Counter_Tracker")}}
                    {{$user.User.Username}} has counted {{.Value}} time(s)!
                {{else}}
                    {{$user.User.Username}} has not counted yet.
                {{end}}
            {{else}}
                Invalid member mention/ID.
            {{end}}
        {{else}}
            This server does not have Tracker enabled, so I cannot get your stats.
        {{end}}
    {{end}}
{{else}}
    {{$check := sdict (dbGet 3 "Count").Value}}
    {{if ($check.Get "Tracker")}}
        {{with (dbGet .User.ID "Counter_Tracker")}}
            You have counted {{.Value}} time(s)!
        {{else}}
            You haven't counted yet.
        {{end}}
    {{else}}
        This server does not have Tracker enabled, so I cannot get your stats.
    {{end}}
{{end}}
