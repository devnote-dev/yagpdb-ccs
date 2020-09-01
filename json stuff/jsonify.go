{{/*
    Made By Devonte#0745 / Naru#6203

    Recommended Trigger Type: Command
    Recommended Trigger: json

    See README.md for more info on use.
*/}}
{{$args := parseArgs 2 "Syntax is: `-json <channelID> <messageID>`" (carg "int" "channelID") (carg "int" "messageID")}}
{{$struct := "message"}}{{$ver := "4.6.5"}}
{{if and ($args.IsSet 0) ($args.IsSet 1)}}
    {{$x := sendMessageRetID nil (cembed "description" "JSON-ifying message...")}}
    {{if eq ($args.Get 0) 0}}
        {{$msg := getMessage nil ($args.Get 1)}}
        {{if $msg}}
            {{if $msg.Embeds}}
                {{if or (reFind `(https:\/\/)?(www\.)?(youtube\.com/|youtu\.be/)(watch\?v=)?.+[a-z+A-Z+0-9]` $msg.Content) (reFind `(https:\/\/)?(www\.)?twitch\.tv/.+[a-z+A-Z+0-9]` $msg.Content)}}
                    {{$struct = "Content/Video"}}
                {{else}}
                    {{$struct = "embed"}}
                {{end}}
            {{else if and $msg.Attachments}}
                {{range $msg.Attachments}}
                    {{if reFind `(?i)\.(txt|jpg|jpeg|png|gif|gifv|tif|tiff)$` .Filename}}
                        {{$struct = "Attachment"}}
                    {{end}}
                {{end}}
            {{end}}
            {{$time := div $msg.ID 4194304|mult 1000000|toDuration}}
            {{$jsonmsg := json $msg}}
            {{$msglink := (print "https://discordapp.com/channels/" .Guild.ID "/" .Channel.ID "/" ($args.Get 1))}}
            {{$embed := cembed
                "author" (sdict "name" (print "Triggered by " .User.String) "icon_url" (.User.AvatarURL "256"))
                "title" "JSON Output"
                "fields" (cslice
                    (sdict "name" "Channel" "value" (print "<#" .Channel.ID ">\n(ID " .Channel.ID ")") "inline" true)
                    (sdict "name" "Message ID" "value" (print "`" ($args.Get 1) "`\n[Click here](" $msglink ") to go to message.") "inline" true)
                    (sdict "name" "Message Type" "value" (print "`" $struct "`") "inline" true)
                    (sdict "name" "Snowflake Created At" "value" (humanizeDurationSeconds (currentTime.Sub ($time|.DiscordEpoch.Add))) "inline" false)
                )
                "description" (print "```" $jsonmsg "```")
                "footer" (sdict "text" (print "JSONify v" $ver))}}
            {{sleep 1}}
            {{editMessage nil $x (complexMessageEdit "embed" $embed)}}
        {{else}}
            {{$nomsg := cembed "title" "Error" "description" "Unkown message. Please try again."}}
            {{sleep 1}}
            {{editMessage nil $x (complexMessageEdit "embed" $nomsg)}}
        {{end}}
    {{else}}
        {{$msg := getMessage ($args.Get 0) ($args.Get 1)}}
        {{if $msg}}
            {{if $msg.Embeds}}
                {{if or (reFind `(https:\/\/)?(www\.)?(youtube\.com/|youtu\.be/)(watch\?v=)?.+[a-z+A-Z+0-9]` $msg.Content) (reFind `(https:\/\/)?(www\.)?twitch\.tv/.+[a-z+A-Z+0-9]` $msg.Content)}}
                    {{$struct = "Content/Video"}}
                {{else}}
                    {{$struct = "embed"}}
                {{end}}
            {{else if and $msg.Attachments}}
                {{range $msg.Attachments}}
                    {{if reFind `(?i)\.(txt|jpg|jpeg|png|gif|gifv|tif|tiff)$` .Filename}}
                        {{$struct = "Attachment"}}
                    {{end}}
                {{end}}
            {{end}}
            {{$time := div $msg.ID 4194304|mult 1000000|toDuration}}
            {{$jsonmsg := json $msg}}
            {{$msglink := (print "https://discordapp.com/channels/" .Guild.ID "/" ($args.Get 0) "/" ($args.Get 1))}}
            {{$embed := cembed
                "author" (sdict "name" (print "Triggered by " .User.String) "icon_url" (.User.AvatarURL "256"))
                "title" "JSON Output"
                "fields" (cslice
                    (sdict "name" "Channel" "value" (print "<#" ($args.Get 0) ">\n(ID " ($args.Get 0) ")") "inline" true)
                    (sdict "name" "Message ID" "value" (print "`" ($args.Get 1) "`\n[Click here](" $msglink ") to go to message.") "inline" true)
                    (sdict "name" "Message Type" "value" (print "`" $struct "`") "inline" true)
                    (sdict "name" "Snowflake Created At" "value" (humanizeDurationSeconds (currentTime.Sub ($time|.DiscordEpoch.Add))) "inline" false)
                )
                "description" (print "```" $jsonmsg "```")
                "footer" (sdict "text" (print "JSONify v" $ver))}}
            {{sleep 1}}
            {{editMessage nil $x (complexMessageEdit "embed" $embed)}}
        {{else}}
            {{$nomsg := cembed "title" "Error" "description" "Unkown message. Please try again."}}
            {{sleep 1}}
            {{editMessage nil $x (complexMessageEdit "embed" $nomsg)}}
        {{end}}
    {{end}}
{{end}}
