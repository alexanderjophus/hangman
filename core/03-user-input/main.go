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

func main() {
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	words := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	rand.Seed(time.Now().Unix())
	wordToGuess := words[rand.Intn(len(words))]

	fmt.Printf("Word to guess: %s\n", wordToGuess)

	attemptedLetters := make(map[string]bool)

	reader := bufio.NewReader(os.Stdin)
	for attemptsLeft := 6; attemptsLeft > 0; {
		fmt.Print("Enter guess: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		text = strings.TrimSpace(text)
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
	}
}
