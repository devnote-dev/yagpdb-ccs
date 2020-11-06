# Custom Report CC
A relatively easy-to-use reports CC with reaction mod actions and report logs.

## Features
- Reactions for quick and easy moderation
- (Integrated moderation actions)
- Message logs
- `reportadmin` commands for utilities

## Usage

### Main Command
`-report <@user/ID> <reason>` - Sends the report

**Reports Interface:**

![Interface Image](https://cdn.discordapp.com/attachments/770405826497740860/774392633032310804/unknown.png)

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

`-reportadmin deleteAllreports <@user/ID>` - Deletes the report history of a specified user.
Aliases: `delall`
**This command has been removed from newer versions due to database issues.**

`-reportadmin reacthelp` - Displays the reactions help page.

## Planned Features

- [ ] `resetreactions`/`rr` option for reportadmin when modaction times out
- [ ] Centralise report logs under one DB Key (when finished, it will be key `7`)
- [ ] Interchangeable reasons for mod actions (basically ease of access)

## Other Info
Reports have been removed in the current update in order to fix the message-embed and reaction bugs. Additionally, permission checks have been added for the moderation actions to clear up default error messages/make them more user-friendly. :)
There is a slight issue with reactions when going to the mod action menu, that is simply due to YAGPDB lag, I cannot do anything about that. It will not affect the moderation actions, only make it look weird. The `sleep` timeout has been increased to try and combat this.

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
