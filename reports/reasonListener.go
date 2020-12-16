{{/*
    
    This is the Message Listener CC made by WickedWizard#3588
  
    Trigger:- Regex `.*`
    
    This listens to the Moderator's reason and adds it to the Kick/Warn/Mute/Ban DM.
    
    Make sure your mods have the permission to type in the Reports Channel.
    
 */}}
 
 {{/*Configuration Values*/}}
 {{$logchannel := CHANNEL_ID}}{{/*Same as the other CC's*/}}
 {{/*Configuration Values End*/}}
 
{{if eq .Channel.ID $logchannel}}
{{dbSet 0 "warn_kick_mute_ban" .Cmd}}
{{end}}
