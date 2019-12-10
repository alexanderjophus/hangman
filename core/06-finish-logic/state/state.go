package state

import (
	"fmt"
)

type state struct {
	wordToGuess      string
	attemptedLetters map[string]bool
	attemptsLeft     int

	lettersInWord map[string]bool
}

// New creates a default game state with no attempted letters
// and 5 attempts to guess the right letters
func New(wordToGuess string) state {
	lettersInWord := make(map[string]bool)
	for _, char := range wordToGuess {
		lettersInWord[string(char)] = true
	}

	return state{
		wordToGuess:      wordToGuess,
		attemptedLetters: make(map[string]bool),
		attemptsLeft:     6,
		lettersInWord:    lettersInWord,
	}
}

func (s state) ShouldContinue() bool {
	if !s.hasAttemptsLeft() {
		return false
	}
	for letterToGuess := range s.lettersInWord {
		if !s.attemptedLetters[letterToGuess] {
			return true
		}
	}
	return false
}

func (s state) hasAttemptsLeft() bool {
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
	if !s.lettersInWord[guess] {
		s.attemptsLeft--
	}

	return nil
}

func (s state) Victory() bool {
	return s.hasAttemptsLeft()
}

func (s state) MaskedWord() string {
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

func (s state) Word() string {
	return s.wordToGuess
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
Attempts left:   %d`, s.MaskedWord(), keys, s.attemptsLeft)
}
