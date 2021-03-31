{{/*
    Suggestions v1 ReactionListener CC

    Made By Devonte#0745 / Naru#6203

    Trigger Type: Reaction - Added and Removed

    See tmplsuggest-v1.md for more info

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$suggest_channel := CHANNEL-ID}} {{/* The ID of the suggestion channel */}}

{{$main_channel := CHANNEL-ID}} {{/* The ID of the review channel */}}

{{$cooldown := 120}} {{/* Suggestions quote cooldown (default is 2 minutes) */}}

{{$staff := cslice ROLE-ID ROLE-ID}} {{/* The IDs of staff roles (requires at least one) */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if .ReactionAdded}}
    {{if and (eq .Channel.ID $suggest_channel) .Message.Embeds (eq .Message.Author.ID 204255221017214977)}}
        {{$msg := index (getMessage $suggest_channel .Message.ID).Embeds 0|structToSdict}}
        {{if reFind `(?i)suggest(?:ion)?` $msg.Title}}
            {{range $k, $v := $msg}}{{if eq (kindOf $v true) "struct"}}{{$msg.Set $k (structToSdict $v)}}{{end}}{{end}}
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
                {{range $staff}}{{if hasRoleID .}}{{$isStaff = true}}{{end}}{{end}}
                {{if $isStaff}}
                    {{dbSetExpire .Message.ID "final" (str .User.ID) 300}}
                    {{addMessageReactions nil .Message.ID ":nonegrey:796925441855586374"}}
                {{else if (eq ($a := reFind `\d{17,19}` $msg.Footer.Text|toInt|userArg).ID .User.ID)}}
                    {{dbSetExpire .Message.ID "delete" true 300}}
                    {{addMessageReactions nil .Message.ID ":nonegrey:796925441855586374"}}
                {{end}}
            {{else if eq .Reaction.Emoji.ID 796925441855586374}}
                {{with (dbGet .Message.ID "final")}}
                    {{if eq (toInt .Value) $.User.ID}}
                        {{deleteMessageReaction nil $.Message.ID $.User.ID ":nonegrey:796925441855586374"}}
                        {{deleteMessageReaction nil $.Message.ID 204255221017214977 ":nonegrey:796925441855586374"}}
                        {{dbDel $.Message.ID "final"}}
                    {{end}}
                {{end}}
            {{else if eq .Reaction.Emoji.ID 796925441771438080 796925441490681889}}
                {{if ($db := dbGet .Message.ID "final")}}
                    {{if eq (toInt $db.Value) .User.ID}}
                        {{$auth := reFind `\d{17,19}` $msg.Footer.Text|toInt|userArg}}
                        {{$total := dict "upvotes" 0 "downvotes" 0}}
                        {{range .Message.Reactions}}
                        {{if eq .Emoji.ID 796925441771438080}}{{$total.Set "upvotes" .Count}}{{end}}
                        {{if eq .Emoji.ID 796925441490681889}}{{$total.Set "downvotes" .Count}}{{end}}
                        {{end}}
                        {{$msg.Set "fields" (cslice (sdict "name" "Final Count" "value" (print "<:checkgreen:796925441771438080> " $total.upvotes "\n<:crossred:796925441490681889> " $total.downvotes) "inline" true) (sdict "name" "Staff Responsible" "value" (printf "• %s - %s\n• ID %d" .User.Mention .User.String .User.ID) "inline" true))}}
                        {{if eq .Reaction.Emoji.ID 796925441771438080}}
                        {{sendMessage $main_channel (complexMessage "content" (print $auth.Mention " Your suggestion was approved!") "embed" (cembed $msg))}}
                        {{else if eq .Reaction.Emoji.ID 796925441490681889}}
                        {{sendMessage $main_channel (complexMessage "content" (print $auth.Mention " Your suggestion was denied (see message below for reason)") "embed" (cembed $msg))}}
                        {{end}}
                        {{deleteMessage $suggest_channel .Message.ID 1}}
                    {{end}}
                {{else if and (eq .Reaction.Emoji.ID 796925441490681889) (dbGet .Message.ID "delete")}}
                    {{$auth := reFind `\d{17,19}` $msg.Footer.Text|toInt|userArg}}
                    {{if eq $auth.ID .User.ID}}
                        {{sendMessage $main_channel (cembed "description" (print .User.Mention " deleted their suggestion!\nSuggestion ID: " (reFind `\d+` $msg.Title)) "color" 0xff2e2e "timestamp" currentTime)}}
                        {{deleteMessage $suggest_channel .Message.ID 1}}
                    {{end}}
                {{end}}
            {{end}}
        {{end}}
    {{end}}
{{end}}
