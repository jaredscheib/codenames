## Codenames API v1

`codenames`
-->

```
Welcome to Codenames!

Usage: codenames <command>

where <command> can be as follows:
<API>
```

`codenames start`
--> checks for game in progress, if so then prompts for destroy to start again
--> otherwise, creates a new game
--> reports whose turn it is

`codenames clue <word> <number>`
--> set your turn's clue

`codenames guess <position>`
--> validity
--> remaining guesses on turn

`codenames next`
--> ends your turn early

`codenames clues`
--> history of clues (incl. current)

`codenames end`
--> deletes current game

### Codenames API v2
`codenames config`
`codenames field`
-->