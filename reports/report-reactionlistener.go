{{/*
    Custom Reports ReactionListener CC v4
    
    Made By Devonte#0745 / Naru#6203
    Contributors: DZ#6669, Piter#5960, WickedWizard#3588
    
    Recommended Trigger Type: Reaction - Added Only

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$staff := cslice ROLE-ID ROLE-ID}} {{/* A list of roles for people considered Admins. Replace ROLEID accordingly. */}}

{{$logChannel := CHANNEL-ID}} {{/* Channel ID to log reports */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{with .ExecData}}
    {{$v := sdict .}}
    {{if $v.inv}}Cancelled the Mod Action.{{deleteResponse 5}}
        {{deleteMessage $logChannel (toInt $v.wmid) 0}}
        {{deleteAllMessageReactions $logChannel (toInt $v.rmid)}}
        {{addMessageReactions $logChannel (toInt $v.rmid) "✅" "❎" "🛡"}}
    {{else}}
        {{$report := index (toInt $v.rmid|getMessage $logChannel).Embeds 0|structToSdict}}
        {{range $k, $v := $report}}{{if eq (kindOf $v true) "struct"}}{{$report.Set $k (structToSdict $v)}}{{end}}{{end}}
        {{$user := (userArg (reReplace `\(|\)` (reFind `\d{17,19}` $report.Description) "")).ID}}
        {{with $report}}
            {{.Set "color" 0xfe3025}}
            {{.Set "description" (print "Mod-Action: **" $v.action "** by " $v.mod " [\u200b](" $v.user ")")}}
            {{.Set "timestamp" currentTime}}
            {{.Author.Set "icon_url" .Author.IconURL}}
        {{end}}
        {{editMessage $logChannel (toInt $v.rmid) (complexMessageEdit "embed" $report)}}
        {{deleteAllMessageReactions $logChannel (toInt $v.rmid)}}
        {{deleteMessage $logChannel (toInt $v.wmid) 0}}
    {{end}}{{dbDel 7 "modaction"}}{{dbDel 7 "reopen"}}
{{else}}
    {{$isStaff := false}}
    {{if .ReactionAdded}}
        {{if and .Message.Author.Bot (eq .Channel.ID $logChannel)}}
            {{if or (dbGet 7 "reopen") (not .Message.EditedTimestamp)}}
                {{range $staff}}{{if hasRoleID .}}{{$isStaff = true}}{{end}}{{end}}
                {{if $isStaff}}
                    {{$mod := (userArg .User.ID).String}}
                    {{$report := index (getMessage nil .Message.ID).Embeds 0|structToSdict}}
                    {{range $k, $v := $report}}{{if eq (kindOf $v true) "struct"}}{{$report.Set $k (structToSdict $v)}}{{end}}{{end}}
                    {{$user := (userArg (reReplace `\(|\)` (reFind `\d{17,19}` $report.Description) "")).ID}}
                    {{if eq .Reaction.Emoji.Name "✅"}}
                        {{with $report}}
                            {{.Set "color" 0x83fe25}}
                            {{.Set "description" (print "Report marked **Done** by " $mod " [\u200b](" $user ")")}}
                            {{.Set "timestamp" currentTime}}
                            {{.Author.Set "icon_url" .Author.IconURL}}
                        {{end}}
                        {{editMessage nil .Message.ID (complexMessageEdit "embed" $report)}}
                        {{deleteAllMessageReactions nil .Message.ID}}
                        {{dbDel 7 "reopen"}}
                    {{else if eq .Reaction.Emoji.Name "❎"}}
                        {{with $report}}
                            {{.Set "color" 0xfeb225}}
                            {{.Set "description" (print "Report marked **Ignored** by " $mod " [\u200b](" $user ")")}}
                            {{.Set "timestamp" currentTime}}
                            {{.Author.Set "icon_url" .Author.IconURL}}
                        {{end}}
                        {{editMessage nil .Message.ID (complexMessageEdit "embed" $report)}}
                        {{deleteAllMessageReactions nil .Message.ID}}
                        {{dbDel 7 "reopen"}}
                    {{else if eq .Reaction.Emoji.Name "🛡"}}
                        {{deleteAllMessageReactions nil .Message.ID}}
                        {{addMessageReactions nil .Message.ID "❌" "⚠" "🔇" "👢" "🔨"}}
                        {{dbSetExpire 7 "modaction" true 300}}
                    {{else if eq .Reaction.Emoji.Name "❌" "⚠" "🔇" "👢" "🔨"}}
                        {{if (dbGet 7 "modaction")}}{{$action := ""}}
                            {{$p := exec "viewperms 204255221017214977"}}
                            {{if eq .Reaction.Emoji.Name "⚠"}}
                                {{if (reFind `(?i)ManageMessages` $p)}}
                                    {{$_ := sendMessageRetID nil (cembed "description" "Type out your reason for warning. Type \"cancel\" to cancel. You have 30 seconds." "color" 16698149)}}
                                    {{dbSetExpire .User.ID "modreason" (sdict "ccid" .CCID "rmid" (str .Message.ID) "wmid" (str $_) "user" (str $user) "action" "warn") 30}}
                                {{else}}{{.User.Mention}} unable to warn that user: Missing Permissions `ManageMessages`.{{end}}
                            {{else if eq .Reaction.Emoji.Name "🔇"}}
                                {{if (reFind `(?i)KickMembers` $p)}}
                                    {{$_ := sendMessageRetID nil (cembed "description" "Type out your reason for muting. Type \"cancel\" to cancel. You have 30 seconds." "color" 16698149)}}
                                    {{dbSetExpire .User.ID "modreason" (sdict "ccid" .CCID "rmid" (str .Message.ID) "wmid" (str $_) "user" (str $user) "action" "mute") 30}}
                                {{else}}{{.User.Mention}} unable to warn that user: Missing Permissions `KickMembers`.{{end}}
                            {{else if eq .Reaction.Emoji.Name "👢"}}
                                {{if (reFind `(?i)KickMembers` $p)}}
                                    {{$_ := sendMessageRetID nil (cembed "description" "Type out your reason for kicking. Type \"cancel\" to cancel. You have 30 seconds." "color" 16698149)}}
                                    {{dbSetExpire .User.ID "modreason" (sdict "ccid" .CCID "rmid" (str .Message.ID) "wmid" (str $_) "user" (str $user) "action" "kick") 30}}
                                {{else}}{{.User.Mention}} unable to warn that user: Missing Permissions `KickMembers`.{{end}}
                            {{else if eq .Reaction.Emoji.Name "🔨"}}
                                {{if (reFind `(?i)BanMembers` $p)}}
                                    {{$_ := sendMessageRetID nil (cembed "description" "Type out your reason for banning. Type \"cancel\" to cancel. You have 30 seconds." "color" 16698149)}}
                                    {{dbSetExpire .User.ID "modreason" (sdict "ccid" .CCID "rmid" (str .Message.ID) "wmid" (str $_) "user" (str $user) "action" "ban") 30}}
                                {{else}}{{.User.Mention}} unable to warn that user: Missing Permissions `KickMembers`.{{end}}
                            {{else if eq .Reaction.Emoji.Name "❌"}}
                                {{deleteAllMessageReactions nil .Message.ID}}
                                {{addMessageReactions nil .Message.ID "✅" "❎" "🛡"}}
                            {{end}}
                        {{else}}
                            {{deleteAllMessageReactions nil .Message.ID}}
                            {{addMessageReactions nil .Message.ID "✅" "❎" "🛡"}}
                        {{end}}
                    {{end}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
