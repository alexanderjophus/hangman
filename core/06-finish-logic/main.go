package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/trelore/hangman/core/06-finish-logic/state"
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

	gameState := state.New(wordToGuess)
	fmt.Println("Word to guess: ", gameState.MaskedWord())

	reader := bufio.NewReader(os.Stdin)
	for gameState.ShouldContinue() {
		func() {
			defer func() {
				fmt.Println(gameState)
			}()
			text, err := getUserInput(reader)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = gameState.Guess(text)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
	if gameState.Victory() {
		fmt.Println("YOU WIN!!! 🚀")
	} else {
		fmt.Printf("Better luck next time. Your word was '%s'\n", gameState.Word())
		os.Exit(1)
	}
}
