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

### Sample Game / API Flow
codenames start
  "Red Spymaster, it is your turn to offer a clue."
  show field (with codenames, positions, reveals)
codenames clue bugs 2
  "All right, red team. Your clue is 'bugs 2'. Feel free to guess a card."
  show field
codenames guess 18
  "Nice! You found a red field agent!"
  show field
  "You have max two more guesses."
codenames guess 17
  "Nope! It was a bystander! Red's turn has ended."
  "Blue spymaster, please offer a clue."
codenames clue animals 4
  "All right, blue team. Your clue is 'animals 2'. Feel free to guess a card."
codenames guess 1
  "Nice! You found a blue field agent!"
  show field
  "You have max three more guesses."
codenames guess 7
  "Whoops! You found a red field agent!"
  show field
  "Blue's turn has ended."
  "Red spymaster, please offer a clue."
codenames clue bob 6
  "All right, red team. Your clue is 'bob 6'. Feel free to guess a card."

// assassination outcome
codenames guess 8
  "Oh no! You found an assassin. You dead."
  "Blue wins!"
(any command from here doesn't work except `codenames start`)

// winning outcome
codenames guess 7
  "Nice! You found the final red field agent! Red wins!"
(command from here doesn't work except `codenames start`)

### Codenames API v2
`codenames config`
`codenames field`
-->