package state_test

import (
	"testing"

	"github.com/trelore/hangman/extra/testing/state"
)

func TestGameFlow(t *testing.T) {
	wordToGuess := "aardvark"
	tts := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "successful game",
			fn: func(t *testing.T) {
				s := state.New(wordToGuess)
				badGuess := "too many letters"
				if err := s.Guess(badGuess); err.Error() != "enter only 1 character, got: too many letters" {
					t.Errorf("want: %s, got %s", err.Error(), "enter only 1 character, got: too many letters")
				}

				if word := s.MaskedWord(); word != "________" {
					t.Errorf("checking masked initial state, want: '________', got: '%s'", word)
				}

				if err := s.Guess("a"); err != nil {
					t.Error("should be able to insert letter 'a'")
				}

				if word := s.MaskedWord(); word != "aa___a__" {
					t.Errorf("checking masked initial state, want: 'aa___a__', got: '%s'", word)
				}
			},
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, tt.fn)
	}
}
