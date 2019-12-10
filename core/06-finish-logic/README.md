## Finishing the game

### Final touches

We have a few final elements to finish our game, we need to;
- Display the state of the game regardless of good or bad input
- Handle correct guesses correctly - they shouldn't reduce the amount of attempts remaining
- Inform the user whether they have won or lost
- Remove our debugging statement stating what the word is

### Displaying the state of the game

This example is likely overcomplicated, however we're going to explore `defer` functions a bit more.
It's worth noting `defer` is normally used for things like closing a network connection, closing a file, etc.
Most of these examples do not take in parameters, as they're methods on the struct itself, like our `f.Close()`.
A [small gotcha](https://blog.learngoprogramming.com/gotchas-of-defer-in-go-1-8d070894cb01) of `defer` statements is that the arguments are determined when the deferred function is put onto the stack, so to print the variable after the game loop, we would want to create something like;

```go
defer func() {
    fmt.Println(gameState)
}()
```

As the defer function does not take any parameters, gameState is only determined once the defer statement begins executing.
This is also a good example demonstrating that we can do a few things in the same defer statement, though again, simple is better and defer statements are normally closing a file or a connection.

The flaw in this logic though, is deferred functions are only called when the parent function is returning.
Our parent function is `main`, we need to print this after ever iteration in the game loop.
The simple solution to this is to put all the logic into a parent function, this is a common pattern when combined with [goroutines](https://gobyexample.com/goroutines), which we may explore in a future lesson.
In the meantime, we can do the following;

```go
for gameState.ShouldContinue() {
    func() {
        defer func() {
            fmt.Println(gameState)
        }()
        // game logic goes here
    }()
}
```
This guarantees we would print the state of the game if the user guesses incorrectly, _or_ correctly.

### Handle correct guesses

In the last lesson there was an excercise of creating a field containing a map of all the letters, we can use this in our `Guess` method we created in the last lesson. To do this we can create a map in our `New` function, and have that return a map, similar to our guessed letters one of `map[sting]bool`, this allows us to only decrement the guessed counter fairly simply as such;

```go
if !s.lettersInWord[guess] {
    s.attemptsLeft--
}
```

### Inform the user of win/loss

Lastly, we need to inform the use if they won or lost.
We can create a method on the `state` struct that returns a boolean if a victory condition has been reached.
Using that after the game loop has finished we are able to tell a user that they were successful (with emojis too!).
Alternatively, we can use the `os` package to exit the CLI with an specific error code, i.e. `os.Exit(1)` (though we may want to print a friendly message first).

### Excercises
- If the user fails, try telling them how many letters got/were missing (i.e. "You got 4/7 letters correctly")
  - Try printing out the number as a decimal to 2 decimal places
- Fix the bug in the `Victory` method, to ensure the word has also been guessed correctly.
- Extend the CLI to allow a user to try again with a new word
