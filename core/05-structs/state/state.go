package state

import (
	"fmt"
)

type state struct {
	wordToGuess      string
	attemptedLetters map[string]bool
	attemptsLeft     int
}

// New creates a default game state with no attempted letters
// and 5 attempts to guess the right letters
func New(wordToGuess string) state {
	return state{
		wordToGuess:      wordToGuess,
		attemptedLetters: make(map[string]bool),
		attemptsLeft:     6,
	}
}

func (s state) HasAttemptsLeft() bool {
	return s.attemptsLeft > 0
}

func (s *state) Guess(guess string) error {
	if len(guess) != 1 {
		return fmt.Errorf("enter only 1 character, got: %s", guess)
	}

	if s.attemptedLetters[guess] {
		return fmt.Errorf("letter already attempted: %s", guess)
	}

	s.attemptedLetters[guess] = true
	s.attemptsLeft--

	return nil
}

func (s state) ShowMaskedWord() string {
	ret := ""
	for _, char := range s.wordToGuess {
		if s.attemptedLetters[string(char)] {
			ret += string(char)
		} else {
			ret += "_"
		}
	}
	return ret
}

func (s state) String() string {
	keys := make([]string, len(s.attemptedLetters))
	i := 0
	for k := range s.attemptedLetters {
		keys[i] = k
		i++
	}

	return fmt.Sprintf(`
Word to guess:   %s
Letters guessed: %v
Attempts left:   %d`, s.ShowMaskedWord(), keys, s.attemptsLeft)
}
