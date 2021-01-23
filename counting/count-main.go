{{/*
    Counting Pagination Reaction CC
    Made By Devonte#0745 / Naru#6203
    Modified by WickedWizard#3588

    Trigger Type: Regex `.*`
    © NaruDevnote 2020-2021 (GNU GPL v3)

    https://github.com/NaruDevnote/yagpdb-ccs
*/}}

{{/*All values below are mandatory*/}}
{{$roleid := 794837085495885824}}{{/*Set to 0 if none*/}}
{{$tracker := false}}
{{$count := true}}{{/*Set this to true if you want normal counting. If you want skip counting, set this to false*/}}
{{$skipcount := true}}{{/*Set this only if you set the above value to false.*/}}
{{/*Above, in $skipcount, if set to true, skip counting by even numbers, if set to false, skip counting by odd numbers.*/}}

{{if not (dbGet 0 "Count")}}
    {{$set := ""}}{{$no := 0}}{{$settings := ""}}
    {{if $count}}
        {{dbSet 0 "Count" (toString -1)}}
        {{$set = dbSet 2 "Count" (toString 100)}}
        {{$settings = "This Server has `Normal Counting` method set up."}}
        {{$no = 0}}
    {{else}}
        {{if $skipcount}}
            {{$settings = "This Server has `Skip Counting` by `Even Numbers` set up."}}
            {{$no = 0}}
            {{dbSet 0 "Count" (toString -2)}}
            {{$set = dbSet 2 "Count" (toString 100)}}
        {{else}}
            {{$settings = "This Server has `Skip Counting` by `Odd Numbers` set up."}}
            {{$no = 1}}
            {{dbSet 0 "Count" (toString -1)}}
            {{$set = dbSet 2 "Count" (toString 99)}}
        {{end}}
    {{end}}
    {{$reason := "No RoleID Set."}}
    {{if $roleid}}
        {{$reason = print "RoleID has been set to <@&" $roleid ">"}}
    {{end}}
    {{dbSet 1 "Count" (toString 0)}}
    {{dbSet 3 "Count" (sdict "Tracker" $tracker "Count/SkipCount" $count "SkipCount Data" $skipcount "Channel ID" .Channel.ID)}}
    {{editChannelTopic nil (print "Next Goal: " (dbGet 2 "Count").Value " One member must let go of another before counting. \n**___DO NOT EDIT YOUR MESSAGE, OR ELSE 😈___**")}}
    {{$embed := cembed 
        "title" "Counting System"
        "description" (print "I have got to know that this is your first time using this system. \nPlease set this to run in your counting channel only, or else, it will break. \n\n**Server Settings:- \n---" $settings "---** \n\nYour counting number has been set to `" $no "`. \n\n" $reason "\nPlease do not change the configuration after starting again. This will break the system.")
        "footer" (sdict "text" "You can delete this message, but I am against doing so. Rather, pin this message.")
        "color" (randInt 1111111 9999999)}}
    {{sendMessage nil $embed}}
    {{deleteTrigger 0}}
{{else}}
    {{if not .ExecData}}
        {{if dbGet .User.ID "Banned"}}
            {{sendDM "You have been banned from Counting by <@" (toInt (dbGet .User.ID "Banned").Value) ">. Please enquire your staff for the reason."}}
        {{else}}
            {{if not (eq .Message.Type 6 7 8 9 10 11)}}
                {{$topic := toInt (reFind `\d+` .Channel.Topic)}}
                {{$number := toInt (dbGet 0 "Count").Value}}
                {{$user := toInt (dbGet 1 "Count").Value}}
                {{$goal := toInt (dbGet 2 "Count").Value}}
                {{$cmd := toInt (reFind `\A\d+\z` .Message.Content)}}
                {{if eq .User.ID $user}}
                    {{deleteTrigger 0}}
                    {{sendDM "You cannot count twice in a row."}}
                {{else}}
                    {{if $count}}
                        {{if eq (add $number 1) $cmd}}
                            {{if $roleid}}
                                {{giveRoleID .User.ID $roleid}}
                                {{takeRoleID $user $roleid}}
                            {{end}}
                            {{dbSet 1 "Count" (toString .User.ID)}}
                            {{$incr := add $number 1}}
                            {{dbSet 0 "Count" (toString $incr)}}
                            {{if not (lt $cmd $topic)}}
                                {{editChannelTopic nil (print "Next Goal: " (mult $goal 2) " One member must let go of another before counting. \n**___DO NOT EDIT YOUR MESSAGE, OR ELSE 😈___**")}}
                                {{dbSet 2 "Count" (toString (mult $goal 2))}}
                            {{end}}
                            {{if $tracker}}
                                {{$lb := dbIncr .User.ID "Counter_Tracker" 1}}
                            {{end}}
                        {{else}}
                            {{sendDM "That is **___not___** the next number, learn how to count. :)"}}
                            {{deleteTrigger 0}}
                        {{end}}
                    {{else}}
                        {{if eq (add $number 2) $cmd}}
                            {{if $roleid}}
                                {{giveRoleID .User.ID $roleid}}
                                {{takeRoleID $user $roleid}}
                            {{end}}
                            {{dbSet 1 "Count" (toString .User.ID)}}
                            {{$incr := add $number 2}}
                            {{dbSet 0 "Count" (toString $incr)}}
                            {{if not (lt $cmd $topic )}}
                                {{editChannelTopic nil (print "Next Goal: " (mult $goal 2) " One member must let go of another before counting. \n**___DO NOT EDIT YOUR MESSAGE, OR ELSE 😈___**")}}
                                {{dbSet 2 "Count" (toString (mult $goal 2))}}
                            {{end}}
                            {{if $tracker}}
                                {{$lb := dbIncr .User.ID "Counter_Tracker" 1}}
                            {{end}}
                        {{else}}
                            {{sendDM "That is **___not___** the next number, learn how to count. :)"}}
                            {{deleteTrigger 0}}
                        {{end}}
                    {{end}}
                {{end}}
                {{scheduleUniqueCC .CCID nil 10 "clean" "clean"}}
            {{end}}
        {{end}}
    {{else}}
        {{$clean := execAdmin "clean 100 204255221017214977 -nopin"}}
    {{end}}
{{end}}
