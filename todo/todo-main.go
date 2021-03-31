{{/*
    Todo Main CC

    Made By Devonte#0745 / Naru#6203

    Recommended Trigger Type: Command
    Recommended Trigger:      todo

    © NaruDevnote 2020-2021 (GNU GPL v3)
    https://github.com/devnote-dev/yagpdb-ccs
*/}}

{{with .ExecData}}
    {{with (dbGet (toInt .) "todo")}}
    {{$list := cslice.AppendSlice .Value}}{{$new := cslice}}
    {{range $i, $e := $list}}
        {{$e := sdict $e}}
        {{if $e.d}}
        {{$dtt := toDuration $e.d|currentTime.Add}}
        {{if currentTime.After $dtt}}
        {{$new = $new.Append $e}}
        {{end}}{{else}}{{$new = $new.Append $e}}{{end}}
    {{end}}{{if len $new}}{{dbSet .UserID .Key $new}}{{else}}{{dbDel .UserID .Key}}{{end}}{{end}}
{{else}}
    {{if .CmdArgs}}
        {{$sub := index .CmdArgs 0|lower}}
        {{if eq $sub "a" "add"}}
            {{if ge (len .CmdArgs) 2}}
                {{$c := slice .CmdArgs 1|joinStr " "}}{{$d := ""}}{{$db := ""}}
                {{with reFind `\s+-d\s+.+\z` $c}}
                {{$d = reReplace `\s+-d\s+` . ""}}{{$c = reReplace `\s+-d\s+.+\z` $c ""}}
                {{end}}
                {{if gt (len $c) 1200}}Your task is too big! Try shortening it!{{else}}
                {{with toDuration $d}}{{$d = str .}}{{else}}{{$d = ""}}{{end}}
                {{with (dbGet .User.ID "todo")}}
                {{$db = cslice.AppendSlice .Value}}
                {{$db = $db.Append (sdict "c" $c "d" $d)}}
                {{dbSet .UserID .Key $db}}
                Task added to the todo list!
                {{else}}
                {{dbSet .User.ID "todo" (cslice (sdict "c" $c "d" $d))}}
                Task added to the todo list!
                {{end}}
                {{if $d}}
                {{$ := exec "remindme" $d (print "Task:\n" $c)}}{{execCC .CCID nil (toDuration $d).Seconds (str .User.ID)}}
                {{end}}{{end}}
            {{else}}No message specified{{"\n"}}Usage: `todo add <content> [-d <duration>]`{{end}}
        {{else if eq $sub "d" "del"}}
            {{if eq (len .CmdArgs) 2}}
                {{with (dbGet .User.ID "todo")}}
                    {{$list := cslice.AppendSlice .Value}}
                    {{$id := index $.CmdArgs 1}}
                    {{if eq (lower $id) "all"}}
                    {{dbDel .UserID .Key}}All tasks removed from your todo list!
                    {{else}}{{$new := cslice}}
                    {{if reFind `[^\d]+` $id}}Invalid task ID "{{$id}}"{{else}}{{$id = toInt $id}}
                    {{if not (len $list)}}You don't have any tasks on your todo list!{{else}}
                    {{$v := true}}{{range $i, $_ := $list}}{{if eq $i $id}}{{$v = false}}{{end}}{{end}}
                    {{if $v}}No task by the ID of {{$id}}{{else}}
                    {{range $i, $e := $list}}{{if ne $i $id}}{{$new = $new.Append .}}{{end}}{{end}}
                    {{if len $new}}{{dbSet .UserID .Key $new}}{{else}}{{dbDel .UserID .Key}}{{end}}
                    Task number `{{$id}}` was removed from your todo list!
                    {{end}}{{end}}{{end}}{{end}}
                {{else}}You don't have any tasks on your todo list!{{end}}
            {{else}}No ID Specified. You can view them with the `todo l/list` command.{{end}}
        {{else if eq $sub "l" "list"}}
            {{with (dbGet .User.ID "todo")}}
                {{$list := cslice.AppendSlice .Value}}{{$c := ""}}
                {{if len $list}}
                {{range $i, $e := $list}}
                {{$e = sdict $e}}
                {{$c = print $c (printf "ID: %d\nTask: %s\nExpires %s\n\n" $i $e.c (or (and $e.d (print (toDuration $e.d|humanizeDurationSeconds) " from now")) "Never"))}}
                {{end}}{{printf "%s's Todo List:\n```\n%s\n```" $.User.Username $c}}
                {{else}}You don't have a todo list! You can create one with the `todo a/add` command.{{end}}
            {{else}}You don't have a todo list! You can create one with the `todo a/add` command.{{end}}
        {{else}}Unknown Subcommand "{{$sub}}"{{"\n"}}You can view subcommands with the `todo` command.{{end}}
    {{else}}
**Todo List Commands:**
```
todo a/add <content> [-d <duration>]
todo d/del <ID>
todo d/del all
todo l/list
```
    {{end}}
{{end}}
