{{/*
	 JSON Converter CC (JSONify)
	 Made By Devonte#0745 / Naru#6203
	 Recommended Trigger Type: Command
	 Recommended Trigger: json
	 See README.md for more info on use.
*/}}
 
{{$args:=parseArgs 2 "Syntax is: `-json <#channel/ID> <messageID> [-file:flag]`" (carg "string" "channelID") (carg "int" "messageID") (carg "string" "file flag")}}{{$ver:="4.11"}}{{$content:=""}}{{$fileCheck:=false}}{{$channel:=.Channel.ID}}{{$struct:="Message"}}
 
{{if $args}}
	{{$l:=sendMessageRetID nil (cembed "description" "Converting Message... <a:loading:760219029620523008>")}}
 
	{{/* noticed you could replace the bottom bit with:
	{{$channel =(or (getChannel (or (reFind `\d+` ($args.Get 0)) ($args.Get 0))) .Channel).ID}}
	 or if really wanna be specific about the condition (btw I recommend the previous one cus 1. their channel may be named "0" cus trolls and 2. Shorter): 
	{{$channel =(or (and (ne ($args.Get 0) "0") (getChannel (or (reFind `\d+` ($args.Get 0)) ($args.Get 0))) .Channel).ID}}
	decided not to cus it’s more expensive on the code (uses getChannel and reFind in cases where it could be escaped / not used by the if statement, your choice if ya wanna use them 
*/}}
	{{if ne ($args.Get 0) "0"}}
		{{$channel =(or (getChannel (or (reFind `\d+` ($args.Get 0)) ($args.Get 0))) .Channel).ID}}
	{{end}}
 
	{{/* moved this up to save some chars */}}
	{{$msglink:=print "https://discordapp.com/channels/" .Guild.ID "/" $channel "/" ($args.Get 1)}}
	{{if $msg:=getMessage $channel ($args.Get 1)}}
 
		{{/* added this below, replaced old with statement */}}
		{{with $msg.Embeds}}
			{{$struct =print (title (index . 0).Type) " Embed"}}
		{{else}}
			{{$struct =or (and $msg.Attachments "Attachment Message") (and (eq $msg.Type 6 7 8 9 10 11) (print "System Message: " $msg.Type)) $struct}}
		{{end}}
		{{$time:=div $msg.ID 4194304|mult 1000000|toDuration}}
		{{$json:=json $msg}}
		{{if or (ge (len $json) 2048) (reFind `\[(?:{.*},?){4,}\]` $json)}}
			{{$fileCheck =true}}
			{{$content ="The message you requested was either too big or contained something that would crash the CC. To prevent this, a downloadable attachment version will be sent instead."}}
		{{else if ($args.IsSet 2)}}
			{{if eq (lower ($args.Get 2)) "-f" "-file"}}
				{{$content ="The downloadable file attachment will be sent shortly. 👌"}}
			{{end}}
		{{end}}
		{{sleep 1}}{{/* you were using sleep in both parts of the if statement below so moved it up*/}}
		{{if $fileCheck}}
			{{deleteMessage nil $l 0}}
			{{sendMessage nil (complexMessage
				"content" $content
				"file" (printf "REQUESTED: %s\nGUILD: %s - %d\nSNOWFLAKE: %s\nJSON:\n\n%s" currentTime .Guild.Name .Guild.ID $time $json)
			)}}
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
					(sdict "name" "Size" "value" (print (fdiv (len $json) 1000) "kb") "inline" true)
				)
				"footer" (sdict "text" (print "JSONify v" $ver))
			}}
			{{editMessage nil $l $nembed}}
		{{end}}
	{{else}}
		{{editMessage nil $l (complexMessageEdit "embed" (cembed "title" "Error" "description" "Unkown message. Please try again."))}}
	{{end}}
{{end}}
