{{/*
    Suggestions v1 Main CC

    Made By Devonte#0745 / Naru#6203

    Recommended Trigger Type: Regex
    Recommended Trigger     : \A

    NOTE: Make sure this only runs in the suggestion channel (see tmplsuggest-v1.md for more info)

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$threshold := 2}} {{/* The number of words needed to be accepted as a suggestion */}}

{{$time := 1800}} {{/* Suggestion Cooldown in seconds (default is 30 minutes) */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if and (gt (len .Args) $threshold) (not (dbGet .User.ID "suggest_cooldown"))}}
    {{$col := 16777215}}{{$p := 0}}{{$r := .Member.Roles}}
    {{range .Guild.Roles}}
        {{if and (in $r .ID) (.Color) (lt $p .Position)}}
        {{$p = .Position}}{{$col = .Color}}{{end}}
    {{end}}
    {{$embed := sdict
        "author" (sdict "name" .User.String "icon_url" (.User.AvatarURL "256"))
        "title" (print "Suggestion #" (dbIncr 0 "suggest_count" 1))
        "description" .Message.Content
        "color" $col
        "footer" (sdict "text" (print "Submit suggestions by typing in suggestions | User: " .User.ID))
        "timestamp" currentTime}}
    {{if .Message.Attachments}}{{$embed.Set "image" (index .Message.Attachments 0)}}{{end}}
    {{$id := sendMessageRetID nil (cembed $embed)}}{{addMessageReactions nil $id ":checkgreen:796925441771438080" ":crossred:796925441490681889" "💬" "🛡"}}
    {{dbSetExpire .User.ID "suggest_cooldown" true $time}}{{deleteTrigger 0}}
{{else}}{{deleteTrigger 0}}{{end}}
