{{/*
    Ghost-Ping Detector CC v1

    Made By Devonte#0745 / Naru#6203

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

    {{/* Message to send when a ping is detected: */}}
    {{sendMessage nil (print "Ghost ping detected by <@" .ExecData.author "> - " $mentions)}}

    {{else}}
    {{if $dp}}{{execCC .CCID nil 5 (sdict "message" .Message.ID "author" .Message.Author.ID "mentions" .Message.Mentions)}}{{end}}
    {{end}}
{{else}}
    {{if .Message.Mentions}}{{execCC .CCID nil 5 (sdict "message" .Message.ID "author" .Message.Author.ID "mentions" .Message.Mentions)}}{{end}}
{{end}}
