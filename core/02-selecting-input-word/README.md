## Selecting a word

### Opening a file

To open a file, we have a [couple of options](https://gobyexample.com/reading-files) available to us. In this example, I have arbitrarily opted for `os`. To take a look at opening a file in isolation;

```go
import (
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Open("words.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    fmt.Printf("File name: %s\n", f.Name())
}
```

There's a couple of new things to point out here, firstly as we have multiple imports we can use a shorthand method of imports, which is `import (...)`, and we can import multiple packages without repeating ourselves. Another small thing to note is the syntax of `:=` to initialise and assign a value to a variable, we could do this in [more steps](https://dzone.com/articles/go-for-beginners-part-2-declaring-variables-in-go-1).

We're going to naïvely open a file called `words.txt`, the location of this file is relative to where we run the go code, we could create an absolute value. `os.Open` returns both a pointer to a file descriptor and an error.

A very common pattern in Go is to return an error as the last return value, then immediately check and handle it. We can revisit this later in more depth, for now we're going to check if the error value is nil, and if it is not (so there _is_ an error), we're going to call `log.Fatal(err)` with the error, this will print the error to the terminal and immediately terminate.

If our call to `os.Open` did not return an error, we can safely use `f`. This is a [pointer](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back) to a file descriptor, `f` is _not_ the data in the file.
Importantly with files it is **always** a good idea to close them once you are done reading them.
The function `f.Close()` will take care of this for us.
You may notice the use of `defer`, this means `f.Close()` becomes a deferred function, a deferred function is _deferred_ until the surrounding function returns. This means our `f.Close()` is only declared once, and is called regardless of how many exit branches our code has.

We are using the `Printf` function from `fmt`, this is subtly different from Println in that we can create our own format. The function takes a string as the first argument, and a variadic array to populate formatting verbs from the first argument. `%s` specifies to put string in that place, `%d` would specify a base 10 int. Arguments are also used in order. We will revisit this in future lessons, for further reading in the meantime there are a [few](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) [cheatsheets](https://gobyexample.com/string-formatting), as well as [docs with examples](https://golang.org/pkg/fmt/)

### Reading a file

Next we will read from the file and put all the read words into a slice.

```go
...
words := []string{}
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    words = append(words, strings.TrimSpace(scanner.Text()))
}
...
```

In the first line we create and instantiate an empty slice of strings.
To break the syntax down further `words` is the variable name, `:=` is shorthand syntax for creating the variable `words` and initialising it to the value of the right hand side.
On the right hand side, `[]string` is declaring a slice or type `string`, inside `{}` we can put initial values, for example we could have done `[]string{"a"}`, and the slice would contain a singular string with the value `a`.

After this we create a scanner object to allow us to read the contents of the file.
The `for` syntax can be declared in a few different ways, such as typical C like for loop
```go
for i := 0; i < 10 ; i++ {
    ...
}
```
In Go we can omit the initialiser and the increment of the control variable as well as the semi-colons. This leaves us with a for loop containing just the conditional statement.
In our example `scanner.Scan()` returns a boolean value letting us know if there are more values to be gotten from the scanner.

Lastly in this step, we can append each line to our slice of potential words.
Go has a built in `append` function, this can be a source of confusion to a couple of new starters. `append` returns the updated slice, therefor we need to assign that back to our slice `words`.
Beyond this, we are using the `strings` packages function `TrimSpace` to ensure there are no whitespace or new line characters in our text.

### Choosing a word at random

Thankfully Go has a `rand` library or two we can use. There's either the lightweight `math/rand` library which takes in a seed and is typically used when 'kind of random' is good enough, there's also `crypto/rand` which is a bit more heavy than our use case requires.

```go
rand.Seed(time.Now().Unix())
wordToGuess := words[rand.Intn(len(words))]
```

We first make sure that random has a seed (otherwise we coukd be picking the same word over and over), then after that we have a big one liner that we can break down.

Taking the inner most block `len(words)`, returns the length of our slice of words to choose from.
`rand.Intn(len(words))` returns a number between 0 and the inputted argument, in our case the length of the slice. We could have assigned the length of the slice to a variable and passed it into this function.
`words[rand.Intn(len(words))]` then returns the element of the slice at the given index, had we used `words[1]` it would have returned the second element in the array (go slices are 0 based). Similar to the other call, we could have assigned the random number to a variable and used that instead.

### Simple Debugging

Lastly, as a simple debugging tip, we will print out the word in plain text while we create this. If you can use the debugger in your IDE, then this may not be as useful.