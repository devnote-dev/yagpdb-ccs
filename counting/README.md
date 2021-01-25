# Counting System CCs
This is an updated and more efficient version of **TimCampy118#5636**'s original CCs (you can find them in the [YAGPDB Support Server](https://discord.com/invite/4udtcA5)). Credits go to him for the original code.

## Features
- Only 1 CC required (others are optional accessories)
- Faster counting
- Optional counting role
- Simpler code (easier to edit)
- Paginated leaderboard embed (less lag)
- Additional counting commands
- Auto-Channel Edit Topic
- Skip Counting by Even or Odd Numbers
- Auto-schedule clean of commands (stuff that shouldn't be in counting)

***__Sometimes you might find that the Main Counting CC Breaks after a Switch from One Method to Another. There is nothing I can do for this, as Discord has a Channel Topic Edit Rate Limit of 2 per 10 Minutes. To fix this, just wait for 10minutes and then try again.__***

## Usage
This is primarily for the count commands, usage for the main command is just... count. Remember to limit the channels the main CC can run in to the count channel only!!!

In the reaction CC, the `$tracker` variable must be switched to `true` if you are going to use the count commands. The default is `false` in favour of users that don't want to use the commands or waste DB space.

`-count help` - Displays the command help embed in Discord.

`-count [@User/ID]` - Displays the count of a specified user, or the triggering user if no one is specified.

`-count info <@User/ID>` - Displays all the counting info of the specified user.

`-count leaderboard` - Displays an embed with the top counters, 1 to 10 per page.
Aliases: `leaderb`, `lb`

`-count next` - Displays the next number to count.

`-count set <number:int>` - Edits the current count to the specified int (whole number). Requires the **Manage Messages** permission to use.

`-count ban <@User/ID>` - Bans an user from Counting. RoleID Needed. Requires the **Manage Messages** permission to use.

`-count unban <@User/ID>` - Unbans an user from Counting. Requires the **Manage Messages** permission to use.

`-count shifthelp` - Shows the Help Menu for Shifting from One Type of Counting to another.

`-count shift` - Shifts from One Type of Counting to another. Needs to be run twice within 5 minutes for confirmation. Requires the **Administrator** permission to use.

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
