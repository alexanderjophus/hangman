## User input

The next key part to creating a CLI is accepting user input.

### Getting user input

To read user input we could use the `bufio` and `os` packages. this allows us to create a reader object on input into the CLI as such
```go
reader := bufio.NewReader(os.Stdin)
``` 
There's a couple of ways to read input [tutorial edge](https://tutorialedge.net/golang/reading-console-input-golang/) has a couple of other examples.
An even more complex example can be found on a [pacman tutorial](https://github.com/danicat/pacgo/tree/master/step02) demonstrating we can use different terminal modes for input.

```go
// try the excercise at the bottom to make this code snippet even clearer.
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
```

We use the `ReadString` function on our `reader` struct, this allows us to read input _until_ we reach the delimiter character.
As our requirements are to have the user enter one character at a time, this introduces potential defects unless we're careful.
Again, we check for errors and in the case of an error, we log a fatal message and terminate the app.
We could choose to handle this differently, and log out the error and `continue`.

Once we have our input, we then trim the string of any whitespace - and importantly reassign that back to the `text` variable, we could also create a new variable, but it's simpler this way as there's less stuff to maintain in our heads.
Lastly we ensure the length is one, as this denotes a user error we can simply print out to the terminal the issue with their input, and use the keyword `continue` to complete this iteration of the loop and go back to the top of the loop.
For further reading on flow control within go, the [golang tour](https://tour.golang.org/flowcontrol/1) is great (or for anything the golang tour covers).

### Creating a game loop

As we want users to continually guess, it makes sense to have a `for` loop counting down the number of guesses a user can have.
We handle this very naïvely as a start, we initially give the user 6 attempts, and so long as they have attempts left we let them guess.

```go
for attemptsLeft := 6; attemptsLeft > 0; {
    // here we can put the reading of the users input that we covered above
    ...
}
```

Referring back to the previous lesson, we can see this for loop follows the C style convention, however it omits the increment (or decrement) of the control variable, however we still need the trailing `;` so the compiler knows what's what.


### Tracking users guesses

Now we have a way have getting a guess, the next thing we want to is track the guess and decrement the counter if it is a new guess.

A developers natural intuition is to lookup sets, Go does not have any set implementation, instead Go encourages developers to use maps.
We are able to emulate sets, and we'll look how after a brief introduction to Go's maps.

First we want to create and instantiate a map, we can do this with the `make` function, similar to `append` it's a built in function of Go's.
As an excercise try omitting `make` and see what happens when you try to insert an item into it.
The make function initializes the map and allows us to use it.
A short note on maps, the syntax `map[string]int` is the syntax for a map that has `string`s as keys and `int`s as values.

With that in mind, we could emulate a set using the value type of `struct{}` which is just an empty structure.
Often it is simpler and makes for more readable code to instead us `map[string]bool` to achieve this, as it allows us to make simple conditionals based off the return value.
We can do this because the zero value of a `bool` is false, and if we attempt to retrieve a value that doesn't exist, we would be returned the zero value of false.

Adding the following allows us to handle users input.

```go
// creates the map - we need to do this before the loop, otherwise we'd re-assign the variable to an empty map every iteration
attemptedLetters := make(map[string]bool)
```

Within our for loop, after getting user data we can also add;

```go
if attemptedLetters[text] {
    fmt.Printf("Letter already attempted: %s\n", text)
    continue
}

attemptedLetters[text] = true
```

Here we see two usages of maps, the first within the `if` statement that retrieves the value of that key within the map.
If the letter already exists, we would like to inform the user and not decrement their number of attempts remaining, so we use `continue` to go back to the top of the loop.
If the letter hasn't been guessed before, we assign that key with the value of `true`, so we can keep track of what has been seen.

### Excercises
- Try printing out the state of the dictionary after each successful guess.
- Change the way we use the reader such that the user can only enter one character at a time.
