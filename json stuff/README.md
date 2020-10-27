# YAGPDB CC - JSONify

This custom command allows you to gather the JSON conversion of a message. The current version is `v4.11`, scroll down for new features.

## Features:
- Outputs full JSON message content (duh 🙄)
- Displays message snowflake
- Displays channel and message ID and hyperlink
- Optional downloadable attachment
- Displays message size
- Outputs specified message type (types listed below):

Message Type | Description
------------ | -------------
Message (default) | Default message
\<TYPE> Embed | Embed message
Attachment Message | GIF/image/file message attachment (content)
System Message \<TYPE> | Discord System message

## Usage:
(Change the prefix according to your set prefix, for demonstration the prefix is `-`)

`-json 0 <message-ID>` - Outputs JSON message (with `0` acting as `nil`, only runs if message is in the same channel)

`-json <channel-ID> <message-ID>` - Outputs JSON message normally with specified channel **ID**

`-json <#Channel> <message-ID>` - Outputs JSON message normally with specified channel **mention**

`-json <#Channel/ID> <message-ID> -f` - Outputs JSON attachment file with contents

`-f` will automatically trigger if the requested message contains elements that could crash the CC, or if the requested message is too large.

## Planned Features:
- [x] `-f` flag: outputs JSON message, ID, snowflake, (etc) in .TXT attachment
- [x] Parse snowflake date (should work now)
- [ ] Add message type "Emoji" for message emojis-only
- [x] Add JSON file auto-attachment for larger files.

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
