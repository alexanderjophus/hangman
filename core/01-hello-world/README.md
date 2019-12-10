## Hello world example

There are a few things we must have to create a Go executable.
The entry point for every single Go executable is the `main` function within `package main`.
As we begin to build more complex systems, we may split these out into individual packages, such that a package handles one area of functionality and handles it well.

To keeping things simple, we will only be concerned with `package main` for the first few lessons.
This means when we run our Go executable either through compiling the code and running the binary, or through `go run`, our executable will begin on line 5 at the code;
```go
func main() {
    ...
}
```

For our first lesson, we're going to print the string `Hello, world!` to the terminal.
Thankfully the [Go core library](https://golang.org/pkg/) is incredibly rich containing packages like `time`, `log`, and `fmt`, providing a wealth of usability without having to depend on many third party libraries (so much so we should be able to create a hangman game through the CLI).

The most common way to print out to the terminal is to use the `fmt.Print` family of commands.

To achieve this we must first import the `fmt` library into our file, we can do this with 
```go
import "fmt"
```
this allows us to use the exported resources of that package.
An important note on exported vs unexported resources in packages, there is no need for keywords like `public`, `export`, or `pub`, to export a resource within Go, that resources name simply has to start with an upper case.
For example `MyFunction` would be exported, `myFunction` would not be (side note, gofmt encourages CamelCase).

With `fmt` imported, we can call `Println` (as it is exported), which takes in a variadic array of elements and prints them all out, space delimited, to the terminal, ending with a new line.
Meaning 
```go
fmt.Println("Hello, world!")
```
achieves the same as
```go
fmt.Println("Hello,", "world!")
```
note the lack of space within the quotes in the second example.

### Comments

Last thing to note in this file, the `//` syntax denotes a comment, go also supports the `/*..*/` format.