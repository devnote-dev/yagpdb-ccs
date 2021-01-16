# Tags CCs
A fully functional, customisable, and reliable tags system for your server. :D

## Features
- Integrated individual tag colours (yes colo**u**rs, I'm British)
- Easier editing
- Compatibility with `snippets` (see [Other Info](#Other-Info))
- Better DM formatting
- Faster fetching and saving tags

## Usage
The base prefix for the tag commands are a semi-colon `;`, this is embedded throughout the main custom command. Remember that the trigger for `tags-main.go` is **Starts With**, not **commmand**.

`;tagname` - Sends the specified tag

`;tag add <name/aliases:text> [color:decimal] [image:URL] <content>` - Adds a tag with the specified fields. Content is set as a required argument if no others are provided, but if others are, you can leave it blank.

`;tag edit <full-name> [new-aliases:text] [color:decimal] [imaage:URL] [new-content]` - Edits an existing tag with the specified fields.

`;tag del/delete <full-name>` - Deletes an existing tag. `<full-name>` has to be the full tag name, not an alias of that tag.

`;tag search <name>` - Searches for a tag based on the name or alias provided. This returns all similar or exact tag names.

`;tag list` - Lists all the tags in the server.

`;tag dump <name>` - Sends the tag entry in JSON format.

`;tag help` - Sends the Discord embed version of the available commands.

The `add`, `edit`, and `del/delete` sub-commands can only be used by the assigned roles in the main custom command. If you don't have any of those roles, the command is ignored. Additionally, you cant name you tag "tag", this is blocked in the code and will send a "No Special Characters" error. With the `add` and `edit` commands, if you want to skip or ignore a field, you can put quotes `""` in it's space and it will use it's existing or default value. For color, you can either put `""` or `0`, both will default to the black embed color.

### Example Usage:
![Example Add-Edit](https://cdn.discordapp.com/attachments/783061830842974280/790358250809065542/cBUdMGPG8p.gif)

### Reactions
- 📱 - DMs the tag content to the user (mobile-Discord friendly)
- 🗑 - Deletes the tag message.

## Planned Features
- [ ] Support for hex codes for `color` field

## Other Info
Regarding compatibility; if you previously used the **snippets** custom command from the YAG support server and don't want to rewrite them, you can use the [**tags-mod.go**](https://github.com/NaruDevnote/yagpdb-ccs/blob/updates/tags/tags-mod.go) version of the CC which has been modified to be compatible with the old snippets. Please keep in mind that this version of the CC will not be regularly updated, the main CC is more important. Below is the structure for tags which you should probably use if you are going to edit the CC.

```json
{
    "author":"tag author",
    "color":0,
    "content":"tag content here",
    "image":"image-url"
}
```

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
