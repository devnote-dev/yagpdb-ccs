# Custom Report CC
A relatively easy-to-use reports CC with reaction mod actions and report logs.

## Features
- Reactions for quick and easy moderation
- (Integrated moderation actions)
- Configurable Reports Ping
- Report Logging
- Message logs
- `reportadmin` commands for utilities

## Usage

### Main Command
`-report <@User/ID> <reason>` - Sends the report

You **must** disable the built-in `report` command with command overrides to prevent separate executions and whatnot. If you need help with setting up command overrides, join the [YAGPDB Support Server](https://discord.com/invite/4udtcA5) and ask for help.

**Report Interface:**

![Interface](https://cdn.discordapp.com/attachments/783061830842974280/788258776707760128/unknown.png)

### Reaction Menu
- ✅ - Marks a report as done/completed (no need for mod actions)
- ❎ - Marks a report as ignored, for example false reports, etc

- 🛡 - Displays the mod actions react menu
    - ❌ - Returns to the main reactions menu
    - ⚠ - Executes a warning on the user
    - 🔇 - Executes a mute on the user
    - 👢 - Executes kick command on the user
    - 🔨 - Executes a ban on the user

*Note: All moderation actions are executed on the user being **reported**, not the user who reported them. Additionally, there currently isn't a way to change the action messages unless you physically change the code in the ReactionListener CC (at the top).*

### Report Admin Commands
`-reportadmin` - Displays the commands message
Aliases: `ra, radmin, reporta`

`-reportadmin reopen <messageID> <reason>` - Reopens a closed report. A reason is required for this, it must be longer than 2 words.

`-reportadmin resetreactions <messageID>` - Resets the reactions on a report message if unresponsive.
Aliases: `rr`

`-reportadmin history <@User/ID>` - Displays the report history of the specified user.

`-reportadmin deletehistory <@User/ID>` - Deletes the entire report history of a user. **This cannot be undone.**
Aliases: `delhistory`

`-reportadmin reacthelp` - Displays the reactions help page.

## Planned Features

- [ ] Interchangeable reasons for mod actions (basically ease of access)

## Other Info
There is a slight issue with reactions when going to the mod action menu, that is simply due to YAGPDB lag, I cannot do anything about that. It will not affect the moderation actions, only make it look weird.

The newest version (v3) now has functional report history logging (thoroughly tested) as well as the new reportadmin commands. Logs are currently stored in reported-user's individual DB entries, if this method remains stable, the DBKey7 plan will be scrapped, and I will optimise this instead. :)

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
