package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func getWord(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("cannot open file: %w", err)
	}
	defer f.Close()

	words := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	rand.Seed(time.Now().Unix())
	wordToGuess := words[rand.Intn(len(words))]
	return wordToGuess, nil
}

func printGameState(wordToGuess string, attemptedLetters map[string]bool, attemptsLeft int) {
	keys := make([]string, len(attemptedLetters))
	i := 0
	for k := range attemptedLetters {
		keys[i] = k
		i++
	}

	fmt.Printf("Word to guess:   %s\n", wordToGuess)
	fmt.Printf("Letters guessed: %v\n", keys)
	fmt.Printf("Attempts left:   %d\n", attemptsLeft)
	fmt.Println()
}

func getUserInput(reader *bufio.Reader) (input string, err error) {
	fmt.Print("Enter guess: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("can't read input: %w", err)
	}
	text = strings.TrimSpace(text)
	return text, nil
}

func main() {
	wordToGuess, err := getWord("words.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Word to guess: %s\n", wordToGuess)

	attemptedLetters := make(map[string]bool)

	reader := bufio.NewReader(os.Stdin)
	for attemptsLeft := 6; attemptsLeft > 0; {
		text, err := getUserInput(reader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if len(text) != 1 {
			fmt.Printf("Enter only 1 character, got: %s\n", text)
			continue
		}
		if attemptedLetters[text] {
			fmt.Printf("Letter already attempted: %s\n", text)
			continue
		}

		attemptedLetters[text] = true

		attemptsLeft--

		printGameState(wordToGuess, attemptedLetters, attemptsLeft)
	}
}
