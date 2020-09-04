# YAGPDB CC - JSONify

This custom command allows you to gather the JSON conversion of a message. The current version is `v4.6.5`, scroll down for new features.

## Features:
- Outputs full JSON message content (duh 🙄)
- Outputs message snowflake
- Outputs channel and message ID and hyperlink
- Outputs specified message type (types listed below):

Message Type | Description
------------ | -------------
Message (default) | Default message
Embed | Embed message (with/without content)
Content/Video | Discord Video Embed link (content)
Attachment | GIF/image/file message attachment (content)

## Usage:
(Change the prefix according to your set prefix, for demonstration the prefix is `-`)

`-json 0 <message-ID>` - Outputs JSON message (with `0` acting as `nil`, only runs if message is in the same channel)

`-json <channel-ID> <message-ID>` - Outputs JSON message normally

## Planned Features:
- [ ] `-f` flag: outputs JSON message, ID, snowflake, (etc) in .txt attachment
- [ ] Parse snowflake date (currently removed)
- [ ] Add message type "Emoji" for message emojis-only
- [ ] Add JSON file auto-attachment for larger files.

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
