{{/*
    Tags ReactionListener CC

    Made By Devonte#0745 / Naru#6203
    Contributors: Maverick Wolf#1010, TheHDCrafter#0001

    Recommended Trigger Type: Reaction (Added and Removed)

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/NaruDevnote/yagpdb-ccs
*/}}

{{if .ReactionAdded}}
    {{if and .Message.Embeds (eq .Message.Author.ID 204255221017214977)}}
        {{$msg := index .Message.Embeds 0|structToSdict}}
        {{if (reFind `\ATag(?::.+|\sList|\sSearch Results|\sHelp)\b` $msg.Title)}}
            {{range $k, $v := $msg}}
                {{if eq (kindOf $v true) "struct"}}{{$msg.Set $k (structToSdict $v)}}{{end}}
            {{end}}
            {{if $msg.Footer}}
                {{if and (eq .Reaction.Emoji.Name "📱") (reFind `React with 📱 to be DMed a mobile version\.\x{200b}` $msg.Footer.Text)}}
                    {{sendDM (printf "\n**%s**\n%s\n%s\n%s" $msg.Title $msg.Description (or $msg.Image (sdict "URL" "")).URL (reReplace `\)\n.+` $msg.Footer.Text ")"))}}
                    {{deleteAllMessageReactions nil .Message.ID}}{{addMessageReactions nil .Message.ID "📱"}}
                {{else if and (eq .Reaction.Emoji.Name "🗑") (reFind `React with 🗑 to delete this message\.` $msg.Footer.Text)}}
                    {{deleteMessage nil .Message.ID 0}}
                {{else if and (eq .Reaction.Emoji.Name "◀" "▶") (reFind `React with 🗑 to delete this message\.\x{200b}\nPage: \d+` $msg.Footer.Text)}}
                    {{$list := ""}}{{$skip := 0}}
                    {{$page := (reFind `\d+` $msg.Footer.Text|toInt)}}
                    {{if eq .Reaction.Emoji.Name "▶"}}
                        {{if not (reFind `No Further Entries` $msg.Description)}}
                            {{$msg.Footer.Set "text" (print "React with 🗑 to delete this message.\u200b\nPage: " (add $page 1))}}
                            {{$skip = mult $page 10}}
                            {{range (dbTopEntries `tag\_%` 10 $skip)}}{{$list = print $list "\n`" (slice .Key 4) "`"}}{{end}}
                            {{if $list}}
                                {{$msg.Set "description" $list}}
                                {{editMessage nil .Message.ID (complexMessageEdit "embed" $msg)}}
                                {{deleteMessageReaction nil .Message.ID .User.ID "▶"}}
                            {{end}}
                        {{end}}
                    {{else if eq .Reaction.Emoji.Name "◀"}}
                        {{if ne $page 1}}
                            {{$msg.Footer.Set "text" (print "React with 🗑 to delete this message.\u200b\nPage: " (sub $page 1))}}
                            {{$skip = sub (mult (add $page -1) 10) 10}}
                            {{range (dbTopEntries `tag\_%` 10 $skip)}}{{$list = print $list "\n`" (slice .Key 4) "`"}}{{end}}
                            {{if $list}}
                                {{$msg.Set "description" $list}}
                                {{editMessage nil .Message.ID (complexMessageEdit "embed" $msg)}}
                                {{deleteMessageReaction nil .Message.ID .User.ID "◀"}}
                            {{end}}
                        {{end}}
                    {{end}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
