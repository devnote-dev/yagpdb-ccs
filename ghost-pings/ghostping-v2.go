{{/*
    Ghost-Ping Detector CC v2

    Made By Devonte#07545 / Naru#6203

    Recommended Trigger Type: Regex
    Recommended Trigger     : .*
    Optional Trigger        : <@!?\d{17,19}>

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$dp := false}} {{/* Change to "true" for double exec-check */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if .ExecData}}
    {{$mentions := ""}}{{$ping := false}}
    {{if ($m := getMessage nil .ExecData.message)}}
    {{if $m.Mentions}}{{else}}{{$ping = true}}{{end}}
    {{else}}{{$ping = true}}{{end}}
    {{if $ping}}
    {{range .ExecData.mentions}}{{$mentions = print $mentions "<@" .ID ">, "}}{{end}}
    {{$col := 16777215}}{{$p := 0}}{{$r := .Member.Roles}}
    {{range .Guild.Roles}}
    {{if and (in $r .ID) (.Color) (lt $p .Position)}}{{$p = .Position}}{{$col = .Color}}{{end}}
    {{end}}

    {{/* Embed Construct */}}
    {{$embed := cembed
        "description" (print "**Channel:** <#" .Channel.ID ">\n**Message:**\n" .ExecData.content "\n\n**Mentioned Users:** " $mentions "\n**Logs:** [Message Logs](" (exec "logs") ")")
        "color" $col
        "footer" (sdict "text" "Detector triggered")
        "timestamp" currentTime}}

    {{/* Content Format */}}
    {{$msgContent := print "Ghost ping detected from <@" .ExecData.author ">!"}}

    {{sendMessage nil (complexMessage "content" $msgContent "embed" $embed)}}
    {{else}}
    {{if $dp}}{{execCC .CCID nil 5 (sdict "message" .Message.ID "author" .Message.Author.ID "mentions" .Message.Mentions "content" .Message.Content)}}{{end}}
    {{end}}
{{else}}
    {{if .Message.Mentions}}{{execCC .CCID nil 5 (sdict "message" .Message.ID "author" .Message.Author.ID "mentions" .Message.Mentions "content" .Message.Content)}}{{end}}
{{end}}
