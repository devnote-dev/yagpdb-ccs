{{/*
    JSON Converter CC (JSONify)

    Made By Devonte#0745 / Naru#6203

    Recommended Trigger Type: Command
    Recommended Trigger: json

    See README.md for more info on use.
*/}}

{{$args := parseArgs 2 "Syntax is: `-json <#channel/ID> <messageID> [-file:flag]`" (carg "string" "channelID") (carg "int" "messageID") (carg "string" "file flag")}}
{{$struct := "Message"}}{{$ver := "4.11"}}{{$c := ""}}{{$fcpt := false}}
{{$channel := .Channel.ID}}{{$msglink := (print "https://discordapp.com/channels/" .Guild.ID "/" .Channel.ID "/" ($args.Get 1))}}

{{if $args}}
    {{$l := sendMessageRetID nil (cembed "description" "Converting Message... <a:loading:760219029620523008>")}}
    {{if not (eq ($args.Get 0) "0")}}
        {{$channel = (getChannel (reReplace `<#|>` ($args.Get 0) "")).ID}}
        {{$channel = (getChannel $channel).ID}}
        {{$msglink = (print "https://discordapp.com/channels/" .Guild.ID "/" $channel "/" ($args.Get 1))}}
    {{end}}
    {{if ($msg := getMessage $channel ($args.Get 1))}}
        {{with $msg}}
            {{if .Embeds}}
                {{$struct = (print (title (index .Embeds 0).Type) " Embed")}}
            {{else if .Attachments}}
                {{$struct = "Attachment Message"}}
            {{else if eq .Type 6 7 8 9 10 11}}
                {{$struct = (print "System Message: " .Type)}}
            {{end}}
        {{end}}
        {{$time := div $msg.ID 4194304|mult 1000000|toDuration}}
        {{$json := json $msg}}
        {{if or (ge (len $json) 2048) (reFind `\[(?:{.*},?){4,}\]` $json)}}
            {{$fcpt = true}}
            {{$c = "The message you requested was either too big or contained something that would crash the CC. To prevent this, a downloadable attachment version will be sent instead."}}
        {{else if ($args.IsSet 2)}}
        {{if eq (str ($args.Get 2)) "-f" "-file"}}
            {{$c = "The downloadable file attachment will be sent shortly. 👌"}}
        {{end}}
    {{end}}
        {{if $fcpt}}
            {{$f := (printf "REQUESTED: %s\nGUILD: %s - %d\nSNOWFLAKE: %s\nJSON:\n\n%s" currentTime .Guild.Name .Guild.ID $time $json)}}
            {{sleep 1}}
            {{deleteMessage nil $l 0}}
            {{sendMessage nil (complexMessage "content" $c "file" $f)}}
        {{else}}
            {{$nembed := cembed
                "author" (sdict "name" (print "Triggered by " .User.String) "icon_url" (.User.AvatarURL "256"))
                "title" "JSON Output"
                "description" (print "```" $json "```")
                "fields" (cslice
                    (sdict "name" "Channel" "value" (print "<#" $channel ">\n(ID " $channel ")") "inline" true)
                    (sdict "name" "Message ID" "value" (print ($args.Get 1) "\n[Click here](" $msglink ") to go to message.") "inline" true)
                    (sdict "name" "Message Type" "value" $struct "inline" true)
                    (sdict "name" "Snowflake (Age)" "value" (humanizeDurationSeconds (currentTime.Sub ($time|.DiscordEpoch.Add))) "inline" true)
                    (sdict "name" "Size" "value" (print (fdiv (len $json) 1000) "kb") "inline" true))
                "footer" (sdict "text" (print "JSONify v" $ver))}}
            {{sleep 1}}
            {{editMessage nil $l (complexMessageEdit "embed" $nembed)}}
        {{end}}
    {{else}}
        {{editMessage nil $l (complexMessageEdit "embed" (cembed "title" "Error" "description" "Unkown message. Please try again."))}}
    {{end}}
{{end}}
