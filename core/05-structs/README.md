## Structs

Structs are one of Gos answer to [object oriented programming](https://golang.org/doc/faq#Is_Go_an_object-oriented_language). In this lesson, we're going to dive into creating a game state struct, and maintaining its state though functions. We're also going to use export vs non-exported functions.

First thing to note is we created a separate package called `state`, naming things is hard, but we'll stick with it for now. The package name and directory name don't necessarily have to line up, but it's normally a good idea if they do.

In this package we have declared a `struct`, and given it some key fields that we will use to maintain the state of the game.

```go
type state struct {
	wordToGuess      string
	attemptedLetters map[string]bool
	attemptsLeft     int
}
```

Remember from previous lessons that lower case dictates a resource is not exported. This means from outside this package (our main package, for example), we can't directly create `state`, we also can't directly manipulate the fields within a state struct.

If our struct was fully exported, we could create new instances of it as we wished. Depending on the code you're creating, you may want to expose lots of things or not very many things. The [core libraries `time`](https://golang.org/src/time/time.go) package has a very good example of not exposing fields of the struct, but allowing manipulation of it (side note, look how readable and commented the core library is).

```go
// state.go
type State struct {
	WordToGuess      string
	AttemptedLetters map[string]bool
	AttemptsLeft     int
}

// main.go
s := State{
    WordToGuess: "word",
    ... // other fields here
}
```

For users of the `state` package, we have written a `New` function, this requires a word, but then creates a few default values for the user for more mundane things like creating the map of already attempted letters.

```go
func New(wordToGuess string) state {
	return state{
		wordToGuess:      wordToGuess,
		attemptedLetters: make(map[string]bool),
		attemptsLeft:     6,
	}
}
```

It is always a good idea to document any exported resources of a package, most IDEs will remind users to do this too.

### Methods

Methods are very similar to functions, the only difference is methods are on a struct.
Assume we want to create a function that returns the state of the guessed word, we could do this a couple of different ways;

```go
// the function way
func ShowMaskedWord(s state) string {
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

// The method way
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
```

The method way is preferred in this example, however the differences are subtle and nuanced. Typically it's a good idea to follow any OO instincts you have.

To explain the method above as well, we iterate through every character in the word we have to guess, if we have guessed the letter, we add it to a string, if not we add an underscore, we then return the string we built.

Note as well, we call fields on the struct by using the dot notation of `s.<field>`, similarly we can call methods the same way.

### Interfaces

In Go, an interface defines the desired behaviour a struct can implement. The benefit of this is functions can take in interfaces as an argument, and so long as your struct implements the desired behaviour then you can use the function with your struct as the argument.

A general guideline of creating functions is to [accept interfaces, return structs](https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8), in general this allows your functions to be more flexible. The internet has a lot of resources for [further information](https://medium.com/rungo/interfaces-in-go-ab1601159b3a) on interfaces in go.

A simple interface in the core library we can implement is the `Stringer` interface;

```go
type Stringer interface {
    String() string
}
```
So long as our `state` function has a `String` function which takes no arguments and returns a `string`, we have implemented the `Stringer` interface.
A naïve implementation may just return a string with the word with the letters we have discovered to far in the game, the set of guessed letters, and the number of attempts left.

```go
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
```

It's important to know that the `Sprint` functions behave almost identically to `Print`, just the return as a string instead of printing to the terminal.

### Exercises
- Modify our new function such that it has a set (or set-adjacent) containing all letters used in our given word.
- Modify our `String()` function such that the list of letters guessed is consistently alphabetical (hint: lookup the `sort` interface)
