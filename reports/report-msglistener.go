{{/*
    Custom Reports Message Tracker CC
    
    Made By Devonte#0745 / Naru#6203
    
    Recommended Trigger Type: Regex
    Recommended Trigger     : .*

    NOTE: Make sure this is set to run in the LOG-CHANNEL ONLY.

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$staff := cslice ROLE-ID ROLE-ID}} {{/* A list of roles for people considered Admins. Replace ROLEID accordingly. */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{with (dbGet .User.ID "modreason")}}
    {{$isStaff := false}}
    {{range $staff}}{{if hasRoleID .}}{{$isStaff = true}}{{end}}{{end}}
    {{if $isStaff}}
        {{$_ := sdict .Value}}
        {{if eq (lower $.Message.Content) "cancel"}}
        {{execCC $_.ccid nil 0 (sdict "inv" true "rmid" $_.rmid "wmid" $_.wmid)}}
        {{else}}
        {{$s := execAdmin $_.action (toInt $_.user) $.Message.Content}}
        {{execCC $_.ccid nil 0 (sdict "rmid" $_.rmid "wmid" $_.wmid "action" $_.action "mod" $.User.String "user" $_.user)}}
        {{end}}
        {{dbDel $.User.ID "modreason"}}
    {{end}}
{{end}}
