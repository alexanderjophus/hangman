package state

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name        string
		wordToGuess string
		want        state
	}{
		{
			name:        "happy path: ants",
			wordToGuess: "ants",
			want: state{
				wordToGuess:      "ants",
				attemptedLetters: make(map[string]bool),
				attemptsLeft:     6,
				lettersInWord: map[string]bool{
					"a": true,
					"n": true,
					"t": true,
					"s": true,
				},
			},
		},
		{
			name:        "happy path: aaaaaa",
			wordToGuess: "aaaaaa",
			want: state{
				wordToGuess:      "aaaaaa",
				attemptedLetters: make(map[string]bool),
				attemptsLeft:     6,
				lettersInWord:    map[string]bool{"a": true},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.wordToGuess); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %#v\nwant %#v", got, tt.want)
			}
		})
	}
}

func Test_state_ShouldContinue(t *testing.T) {
	goldenState := state{
		wordToGuess:      "ants",
		attemptedLetters: map[string]bool{},
		attemptsLeft:     6,
		lettersInWord: map[string]bool{
			"a": true,
			"n": true,
			"t": true,
			"s": true,
		},
	}

	tests := []struct {
		name  string
		state state
		want  bool
	}{
		{
			name:  "has attempts and not completed guess",
			state: goldenState,
			want:  true,
		},
		{
			name:  "has no attempts left",
			state: func() state { s := goldenState; s.attemptsLeft = 0; return s }(),
			want:  false,
		},
		{
			name: "correctly guessed all letters",
			state: func() state {
				s := goldenState
				s.attemptedLetters = s.lettersInWord
				return s
			}(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.state.ShouldContinue(); got != tt.want {
				t.Errorf("state.ShouldContinue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_state_hasAttemptsLeft(t *testing.T) {
	tests := []struct {
		name         string
		attemptsLeft int
		want         bool
	}{
		{name: "has attempts left", attemptsLeft: 1, want: true},
		{name: "has no attempts left", attemptsLeft: 0, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := state{attemptsLeft: tt.attemptsLeft}
			if got := s.hasAttemptsLeft(); got != tt.want {
				t.Errorf("state.hasAttemptsLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}
