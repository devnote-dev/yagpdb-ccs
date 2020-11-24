# JSONify CC

This custom command allows you to gather the JSON conversion of a message. The current version is `v4.15`, scroll down for new/current features.

## Features:
- Outputs full JSON message content (duh 🙄)
- Displays message snowflake
- Displays channel and message ID and hyperlink
- Optional downloadable attachment
- Displays message size
- Optional formatted JSON output
- Outputs specified message type (types listed below):

Message Type | Description
------------ | -------------
Default Message | Normal Discord Message
Pinned Message | Channel-message Pinned Message
Join Message | Member Join Message
Boost Message | Server Boost Message
Tier 1 Boost Message | 1st Tier Message (2 boosts)
Tier 2 Boost Message | 2nd Tier Message (15 boosts)
Tier 3 Boost Message | 3rd Tier Message (30 boosts)
Followed Channel Message | Server Followed Channel Message
Reply Message | Message Reply

**Please read [Other Info](#Other-Info) for information on join, boost, and relpy messages!**

## Usage:
(Change the prefix according to your set prefix, for demonstration the prefix is `-`)

`-json 0 <Message-ID>` - Outputs JSON message (with `0` acting as `nil`, only runs if message is in the same channel)

`-json <Channel> <Message-ID>` - Outputs JSON message, with channel being either: channel name (text), channel mention, or channel ID.

**Flag: `-j`** - Formats the JSON in several lines.

**Flag: `-f`** - Outputs the JSON in a .txt file.

## Planned Features:
- [ ] Update formatting Regex

## Other Info
The new `-j` flag for formatting is still being tested, the regex for it may be altered in future versions.

Join and boost messages from the server/Discord have no content and are not always recognised as it's designated type. I presume this is a Discord bug as it has worked for some users, so it may be fixed when **API v8** is released. Reply messages still identify as TYPE0, however the Discord Developers documentation was updated showing that reply messages will go under it's own type, TYPE19, when **API v8** is released. Until then, the code for reply messages will stay put, then when the new API rolls out, it should work flawlessly.

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
