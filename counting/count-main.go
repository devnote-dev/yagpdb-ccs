{{/*
    Counting Main CC

    Made By Devonte#0745 / Naru#6203

    Recommended Trigger Type: Regex
    Recommended Trigger     : .*

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{/* THINGS TO CHANGE */}}

{{$roleID := 0}} {{/* Role ID of the role to give, leave as 0 for none */}}

{{$tracker := false}} {{/* Individual user count tracker, change to "true" to enable */}}

{{/* ACTUAL CODE - DO NOT TOUCH */}}

{{if not (dbGet 5 "c_count")}}
    {{dbSet 5 "c_count" "0"}}
    {{dbSet 5 "c_counter" (str .User.ID)}}
    {{$rm := "No count role has been set."}}{{if $roleID}}{{$rm = print "The count role has been set to <@&" $roleID ">. You will get this role after counting more than once."}}{{end}}
    {{sendMessage nil (cembed
        "title" "Welcome to Counting!"
        "description" (printf "I see this is your first time using this CC! I have set the count to **0** and the user to **%s**! %s\nMake sure that this CC only runs in this channel (<#%d>), otherwise the CC will run in too many channels and break. Have fun!" .User.String $rm .Channel.ID)
        "color" (randInt 111111 999999)
        "footer" (sdict "text" "You can delete this message now :)"))}}
{{else}}
    {{if not .ExecData}}
        {{$cUser := (toInt (dbGet 5 "c_counter").Value)}}
        {{$cCount := (toInt (dbGet 5 "c_count").Value)}}
        {{if ($n := reFind `\A\d+\z` (joinStr " " .Args))}}
            {{if ne .User.ID $cUser}}
                {{if eq (toInt $n) (add $cCount 1)}}
                    {{if $roleID}}
                        {{takeRoleID $cUser $roleID}}{{addRoleID $roleID}}
                    {{end}}
                    {{dbSet 5 "c_count" $n}}
                    {{dbSet 5 "c_counter" (str .User.ID)}}
                    {{if $tracker}}
                        {{$_ := dbIncr .User.ID "count_tracker" 1}}
                    {{end}}
                {{else}}{{deleteTrigger 0}}{{end}}
            {{else}}
                {{deleteTrigger 0}}{{sendDM "You cant count twice in a row!"}}
            {{end}}
        {{else}}
            {{deleteTrigger 0}}
            {{cancelScheduledUniqueCC .CCID "countclr"}}
            {{scheduleUniqueCC .CCID nil 12 "countclr" true}}
        {{end}}
    {{else}}{{$_ := execAdmin "clear 100 204255221017214977 -nopin"}}{{end}}
{{end}}
