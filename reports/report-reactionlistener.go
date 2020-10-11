{{/*
    Custom Reports ReactionListener CC
    
    Made By Devonte#0745 / Naru#6203
    Contributors: DZ#6669, Piter#5960
    
    Recommended Trigger Type: Reaction - Added Only
*/}}

{{/* THINGS TO CHANGE */}}
{{$staff := cslice ROLEID ROLEID}} {{/* A list of roles for people considered Admins. Replace ROLEID accordingly. */}}

{{$logChannel := }} {{/* Channel ID to log reports */}}

{{/* Messages / reasons for mod actions */}}
{{$warn := "Please follow the rules in this server."}}
{{$mute := "Please follow the rules in this server."}}
{{$kick := "Try again next time."}}
{{$ban := "Don't break the rules here"}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{$isStaff := false}}
{{if .ReactionAdded}}
    {{if and .ReactionMessage.Author.Bot (eq .Reaction.ChannelID $logChannel)}}
        {{if or (dbGet 7 "reopen") (not .ReactionMessage.EditedTimestamp)}}
            {{range (getMember $.Reaction.UserID).Roles}}
                {{if in $staff .}}
                    {{$isStaff = true}}
                {{end}}
            {{end}}
            {{if $isStaff}}
                {{$mod := (userArg .Reaction.UserID).String}}
                {{$report := index (getMessage nil .ReactionMessage.ID).Embeds 0|structToSdict}}
                    {{range $k, $v := $report}}
                        {{if eq (kindOf $v true) "struct"}}
                            {{$report.Set $k (structToSdict $v)}}
                        {{end}}
                    {{end}}
                    {{$user := (userArg (reReplace `\(|\)` (reFind `\d{17,19}` $report.Description) "")).ID}}
                {{if eq .Reaction.Emoji.Name "✅"}}
                    {{with $report}}
                        {{.Set "color" 0x83fe25}}
                        {{.Set "description" (print "Report marked **Done** by " $mod " [\u200b](" $user ")")}}
                        {{.Set "timestamp" currentTime}}
                        {{.Author.Set "icon_url" .Author.IconURL}}
                    {{end}}
                    {{editMessage nil .ReactionMessage.ID (complexMessageEdit "embed" $report)}}
                    {{deleteAllMessageReactions nil .ReactionMessage.ID}}
                    {{dbDel 7 "reopen"}}
                {{else if eq .Reaction.Emoji.Name "❎"}}
                    {{with $report}}
                        {{.Set "color" 0xfeb225}}
                        {{.Set "description" (print "Report marked **Ignored** by " $mod " [\u200b](" $user ")")}}
                        {{.Set "timestamp" currentTime}}
                        {{.Author.Set "icon_url" .Author.IconURL}}
                    {{end}}
                    {{editMessage nil .ReactionMessage.ID (complexMessageEdit "embed" $report)}}
                    {{deleteAllMessageReactions nil .ReactionMessage.ID}}
                    {{dbDel 7 "reopen"}}
                {{else if eq .Reaction.Emoji.Name "🛡"}}
                    {{deleteAllMessageReactions nil .ReactionMessage.ID}}
                    {{addMessageReactions nil .ReactionMessage.ID "❌" "⚠" "🔇" "👢" "🔨"}}
                    {{dbSetExpire 7 "modaction" true 300}}
                {{else if eq .Reaction.Emoji.Name "❌" "⚠" "🔇" "👢" "🔨"}}
                    {{if (dbGet 7 "modaction")}}
                        {{$action := ""}}
                        {{if eq .Reaction.Emoji.Name "⚠"}}
                            {{$s := execAdmin "warn" $user $warn}}
                            {{$action = "warned"}}
                        {{else if eq .Reaction.Emoji.Name "🔇"}}
                            {{$s := execAdmin "mute" $user $mute}}
                            {{$action = "muted"}}
                        {{else if eq .Reaction.Emoji.Name "👢"}}
                            {{$s := execAdmin "kick" $user $kick}}
                            {{$action = "kicked"}}
                        {{else if eq .Reaction.Emoji.Name "🔨"}}
                            {{$s := execAdmin "ban" $user $ban}}
                            {{$action = "banned"}}
                        {{else if eq .Reaction.Emoji.Name "❌"}}
                            {{deleteAllMessageReactions nil .ReactionMessage.ID}}
                            {{addMessageReactions nil .ReactionMessage.ID "✅" "❎" "🛡"}}
                        {{end}}
                        {{if $action}}
                            {{with $report}}
                                {{.Set "color" 0xfe3025}}
                                {{.Set "description" (print "ModAction: **" $action "** by " $mod " [\u200b](" $user ")")}}
                                {{.Set "timestamp" currentTime}}
                                {{.Author.Set "icon_url" .Author.IconURL}}
                            {{end}}
                            {{editMessage nil .ReactionMessage.ID (complexMessageEdit "embed" $report)}}
                            {{deleteAllMessageReactions nil .ReactionMessage.ID}}
                            {{dbDel 7 "modaction"}}{{dbDel 7 "reopen"}}
                        {{end}}
                    {{else}}
                        {{deleteAllMessageReactions nil .ReactionMessage.ID}}
                        {{addMessageReactions nil .ReactionMessage.ID "✅" "❎" "🛡"}}
                    {{end}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
