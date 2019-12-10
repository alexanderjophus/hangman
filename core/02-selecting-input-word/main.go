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
	// opens the file
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	// ensure the file is closed
	defer f.Close()

	fmt.Printf("File name: %s\n", f.Name())

	words := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	rand.Seed(time.Now().Unix())
	wordToGuess := words[rand.Intn(len(words))]

	fmt.Printf("Word to guess: %s\n", wordToGuess)
}
