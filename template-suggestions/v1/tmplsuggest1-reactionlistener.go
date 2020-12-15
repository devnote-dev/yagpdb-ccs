{{/*
    Template-Suggestions v1 ReactionListener CC

    Made By Devonte#0745 / Naru#6203

    Trigger Type: Reaction (Added and Removed)

    See tmplsuggest-v1.md for more info
*/}}

{{/* THINGS TO CHANGE */}}

{{$suggest_channel := }} {{/* The ID of the suggestion channel */}}

{{$main_channel := }} {{/* The ID of the review channel */}}

{{$cooldown := 120}} {{/* Suggestions quote cooldown (default is 2 minutes) */}}

{{$staff := cslice ROLEID ROLEID}} {{/* The IDs of staff roles (requires at least one) */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if .ReactionAdded}}
    {{if and (eq .Channel.ID $suggest_channel) .Message.Embeds (eq .Message.Author.ID 204255221017214977)}}
        {{$msg := index (getMessage $suggest_channel .Message.ID).Embeds 0|structToSdict}}
        {{if reFind `(?i)suggest(?:ion)?` $msg.Title}}
            {{range $k, $v := $msg}}
                {{if eq (kindOf $v true) "struct"}}{{$msg.Set $k (structToSdict $v)}}{{end}}
            {{end}}
            {{$msg.Author.Set "icon_url" $msg.Author.IconURL}}
            {{if eq .Reaction.Emoji.Name "💬"}}
                {{if not (dbGet .Message.ID "suggest_qcd")}}
                    {{sendMessage $main_channel (complexMessage "content" (print .User.Mention " quoted this suggestion:") "embed" (cembed $msg))}}
                    {{dbSetExpire .Message.ID "suggest_qcd" true $cooldown}}
                    {{deleteMessageReaction nil .Message.ID .User.ID "💬"}}
                {{else}}
                    {{$id := sendMessageRetID $main_channel (print .User.Mention " That suggestion cant be quoted for another " (humanizeDurationSeconds ((dbGet .Message.ID "suggest_qcd").ExpiresAt.Sub currentTime)))}}
                    {{deleteMessage $main_channel $id}}
                {{end}}
            {{else if eq .Reaction.Emoji.Name "🛡"}}
                {{$isStaff := false}}
                {{range $staff}}
                    {{if hasRoleID .}}{{$isStaff = true}}{{end}}
                {{end}}
                {{if $isStaff}}
                    {{dbSetExpire .Message.ID "final" (str .User.ID) 300}}
                    {{addMessageReactions nil .Message.ID "❌"}}
                {{end}}
            {{else if eq .Reaction.Emoji.Name "❌"}}
                {{with (dbGet .Message.ID "final")}}
                    {{if eq (toInt .Value) $.User.ID}}
                        {{deleteMessageReaction nil $.Message.ID $.User.ID "❌"}}
                        {{deleteMessageReaction nil $.Message.ID 204255221017214977 "❌"}}
                        {{dbDel $.Message.ID "final"}}
                    {{end}}
                {{end}}
            {{else if eq .Reaction.Emoji.ID 737036090510278716 737036123767046155}}
                {{if (dbGet .Message.ID "final")}}
                    {{if eq (toInt (dbGet .Message.ID "final").Value) .User.ID}}
                        {{$auth := reFind `\d{17,19}` $msg.Footer.Text|toInt}}
                        {{$total := dict "upvotes" 0 "downvotes" 0}}
                        {{range .Message.Reactions}}
                            {{if eq .Emoji.ID 737036090510278716}}{{$total.Set "upvotes" .Count}}{{end}}
                            {{if eq .Emoji.ID 737036123767046155}}{{$total.Set "downvotes" .Count}}{{end}}
                        {{end}}
                        {{$msg.Set "fields" (cslice (sdict "name" "Final Count" "value" (print "<:check:737036090510278716> " $total.upvotes "\n<:cross:737036123767046155> " $total.downvotes) "inline" true) (sdict "name" "Staff Responsible" "value" (printf "• %s - %s\n• ID %d" .User.Mention .User.String .User.ID) "inline" true))}}
                        {{if eq .Reaction.Emoji.ID 737036090510278716}}
                            {{sendMessage $main_channel (complexMessage "content" (print "<@" $auth "> Your suggestion was approved!") "embed" (cembed $msg))}}
                        {{else if eq .Reaction.Emoji.ID 737036123767046155}}
                            {{sendMessage $main_channel (complexMessage "content" (print "<@" $auth "> Your suggestion was denied (see message below for reason)") "embed" (cembed $msg))}}
                        {{end}}
                        {{deleteMessage $suggest_channel .Message.ID 2}}
                    {{end}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
