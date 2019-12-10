package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/trelore/hangman/core/05-structs/state"
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

	gameState := state.New(wordToGuess)

	reader := bufio.NewReader(os.Stdin)
	for gameState.HasAttemptsLeft() {
		text, err := getUserInput(reader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = gameState.Guess(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(gameState)
	}
}
