{{/*
    Counting Pagination Reaction CC

    Made By Devonte#0745 / Naru#6203

    Trigger Type: Reaction (Added Only)

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{if .ReactionAdded}}
    {{if and (eq .Message.Author.ID 204255221017214977) .ReactionMessage.Embeds}}
        {{$msg := (index (getMessage nil .Message.ID).Embeds 0)|structToSdict}}
        {{if and $msg.Title (reFind `(?i)Counting Leaderboard` $msg.Title)}}
            {{range $k, $v := $msg}}
                {{if eq (kindOf $v true) "struct"}}
                    {{$msg.Set $k (structToSdict $v)}}
                {{end}}
            {{end}}
            {{$page := toInt (reFind `\d+` $msg.Footer.Text)}}
            {{$list := ""}}{{$rank := 1}}{{$r := false}}
            {{if eq .Reaction.Emoji.Name "➡"}}
                {{if not (reFind `No entries` $msg.Description)}}
                    {{$rank = mult $page 10}}{{$r = true}}
                    {{$msg.Footer.Set "text" (print "Page: " (add $page 1))}}
                {{end}}
            {{else if eq .Reaction.Emoji.Name "⬅"}}
                {{if ne $page (or 0 1)}}
                    {{$rank = (div $page 10)|roundCeil|toInt}}{{$r = true}}
                    {{$msg.Footer.Set "text" (print "Page: " (sub $page 1))}}
                {{end}}
            {{else if eq .Reaction.Emoji.Name "🗑"}}
                {{deleteMessage nil .Message.ID 0}}{{end}}
            {{if $r}}
                {{range (dbTopEntries `count\_tracker` 10 $rank) -}}
                    {{- $rank = add $rank 1 -}}
                    {{- $list = print $list (printf "%-4d %4d\t %-16s" $rank (toInt .Value) .User.String) "\n" -}}
                {{- else -}}
                    {{- $list = "No entries to display." -}}
                {{- end}}
                {{$msg.Set "description" (print "```\n# -- Count -- User\n" $list "\n```")}}
                {{editMessage nil .Message.ID (complexMessageEdit "embed" $msg)}}
                {{deleteMessageReaction nil .Message.ID .User.ID "⬅" "➡"}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
