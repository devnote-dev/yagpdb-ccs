# Custom Report CC
A relatively easy-to-use reports CC with reaction mod actions and report logs.

## Features
- Reactions for quick and easy moderation
- (Integrated moderation actions)
- Report Logging¹
- Message logs
- `reportadmin` commands for utilities

## Usage

### Main Command
`-report <@user/ID> <reason>` - Sends the report

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

`-reportadmin reacthelp` - Displays the reactions help page.

## Planned Features

- [ ] Centralise report logs under one DB Key (when finished, it will be key `7`)
- [ ] Interchangeable reasons for mod actions (basically ease of access)

## Other Info
There is a slight issue with reactions when going to the mod action menu, that is simply due to YAGPDB lag, I cannot do anything about that. It will not affect the moderation actions, only make it look weird.

¹: Logs currently only show the date in the report, and when finished, the date and reason for report. I will work on making this more descriptive, such has user who reported & channel.

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
