# Template Suggestion CC Version 1
The first version of the template suggestion CCs. The CCs part of this version are configurable to your server and can be used by some of the custom commands in the **suggest utilities** folder (check for eligibility).

## Features
- Direct type-to-suggest system
- Proper attachment support
- Embeds :P
- Interactive reactions menu

## Usage
This system is based around a two-channel suggestion system, in which there is an open channel that suggestions can be posted in, slowmode optional with the cooldown (see first GIF below), and a secondary channel for staff to review suggestions and add comments below (see second GIF below).

### Main Suggestion CC

Use with the recommended trigger type `Regex` and trigger `\A`. **You must** make sure to limit the CC to only run in the suggestion channel. If you don't, it will spam in all active channels and possibly break.

![Suggesting Example](https://cdn.discordapp.com/attachments/783061830842974280/788234625288765470/xMXNbtaAoC.gif)

### Reaction Menu
- 💬 - Quotes the suggestion in the main suggestions channel / review channel
- 🛡 - Sets the reactions into "final" review mode where suggestions are either approved or denied. Requires at least one of the staff roles set in the CC code.
    - shield + check - Approves the suggestion
    - shield + cross - Denies the suggestion
    - ❌ - Cancels the review

![Review Example](https://cdn.discordapp.com/attachments/783061830842974280/788237355084677180/zOZREyzlYk.gif)

*If you find any bugs or issues, feel free to PR an issue or fix, or contact me through the YAGPDB Support Server*
