{{/*
	Custom Mute by Devonte#0745 / Naru#6203
	Conrtributors: DZ#6669, Maverick Wolf#0001, Satty#9361, ShADowNIGHT#1779
	tysm for the help guys :)
	
	Trigger Type: Command, Trigger: CmdName
	Remember to change everything in []
*/}}

{{if not .ExecData}}
	{{$args := parseArgs 2 "Syntax is: `-[CmdName] <user> <duration:seconds>`" (carg "user" "user to be muted") (carg "duration" "duration of mute")}}
	{{$user := $args.Get 0}}
	{{$delay := $args.Get 1}}
	{{ giveRoleID $user [CustomMuteRole] }}
	{{ print "Muted " $user.ID " for: " (toInt ($delay.Seconds)) " seconds." }}
	{{execCC .CCID nil $delay.Seconds $user}}
{{else}}
	{{ takeRoleID .ExecData.ID [CustomMuteRole] 0 }}
	{{ print .ExecData.String " has been unmuted: duration expired." }}
{{end}}
