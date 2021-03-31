{{/*
    Counting Commands CC

    Made By Devonte#0745 / Naru#6203

    Trigger Type: Regex
    Trigger     : \A\-count(?:\s+|\z)

    NOTE: If your prefix is NOT a "-" you need to change the \A\- at the start of the regex to your prefix.
    EXAMPLE: \A\? if the prefix was ?

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{if .CmdArgs}}
    {{$c := lower (index .CmdArgs 0)}}
    {{if eq $c "next"}}
        Next number: {{or (add (toInt (dbGet 5 "c_count").Value) 1) 0}}
    {{else if eq $c "lb" "leaderb" "leaderboard"}}
        {{$list := ""}}{{$rank := 1}}
        {{range (dbTopEntries `count\_tracker` 10 0) -}}
            {{- $list = print $list (printf "%-4d %4d\t %-16s" $rank (toInt .Value) .User.String) "\n" -}}
            {{- $rank = add $rank 1 -}}
        {{- else -}}{{- $list = "No entries to display." -}}{{- end}}
        {{$embed := cembed
            "title" "Counting Leaderboard"
            "description" (print "```\n# -- Count -- User\n" $list "\n```")
            "color" 816464
            "footer" (sdict "text" "Page: 1")
            "timestamp" currentTime}}
        {{$id := sendMessageRetID nil $embed}}{{addMessageReactions nil $id "⬅" "➡" "🗑"}}
    {{else if eq $c "help" "commands"}}
        {{$embed := cembed
            "title" "Counting Help"
            "description" "`count help/commands` - Shows this message\n\n`count [@User/ID]` - Shows the count of a user, or your own if no user is specified\n`count info <@User/ID>` - Shows the counting info of a specified user\n`count lb/leaderboard` - Shows the counting leaderboard\n\n`count next` - Shows the next number to count\n`count set <number>` - Changes the current count (mods only)"
            "color" 816464}}
        {{sendMessage nil $embed}}
    {{else if eq $c "set"}}
        {{if (reFind `ManageMessages` (exec "viewperms"))}}
            {{$args := parseArgs 2 "Usage: `count set <number>`" (carg "string" "") (carg "int" "")}}
            {{dbSet 5 "c_count" (str ($args.Get 1))}}
            {{.User.Mention}} Set the count to {{$args.Get 1}}
        {{else}}Missing permissions: `ManageMessages`
            {{deleteTrigger}}{{deleteResponse}}{{end}}
    {{else if eq $c "info"}}
        {{$args := parseArgs 2 "Usage: `count info <@User/ID>`" (carg "string" "") (carg "member" "")}}
        {{$u := ($args.Get 1).User}}
        {{with (dbGet $u.ID "count_tracker")}}
            {{$col := 16777215}}{{$p := 0}}{{$r := ($args.Get 1).Roles}}
            {{range $.Guild.Roles}}{{if and (in $r .ID) (.Color) (lt $p .Position)}}{{$p = .Position}}{{$col = .Color}}{{end}}{{end}}
            {{$embed := cembed
                "author" (sdict "name" $u.String "icon_url" ($u.AvatarURL "256"))
                "description" (print "• Current Count: " .Value "\n• Started Counting: " (humanizeTimeSinceDays .CreatedAt) " ago\n• Last Counted: " (humanizeTimeSinceDays .UpdatedAt) " ago")
                "color" $col
                "timestamp" currentTime}}
            {{sendMessage nil $embed}}
        {{else}}{{$u.Username}} has not counted yet.{{end}}
    {{else}}
        {{if ($u := getMember $c)}}
            {{with (dbGet $u.User.ID "count_tracker")}}
                {{$u.User.Username}} has counted {{.Value}} time(s)!
            {{else}}{{$u.User.Username}} has not counted yet.{{end}}
        {{else}}Invalid member mention/ID.{{end}}
    {{end}}
{{else}}
    {{with (dbGet .User.ID "count_tracker")}}
        You have counted {{.Value}} time(s)!
    {{else}}You haven't counted yet.{{end}}
{{end}}
