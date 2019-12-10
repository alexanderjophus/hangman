## Functions

### Creating functions

As all good languages, [Go has functions](https://tour.golang.org/basics/4), this lesson explores breaking down our current functionality into smaller more testable components.

The basic function signature is

```go
func FunctionName() {...}
```

We declare a function with `func`, it has the name `FunctionName`, and it would be called elsewhere in the code by `FunctionName()`.
This function takes no parameters and returns nothing.

Looking at the `getWord` function we created;

```go
func getWord(filename string) (string, error) {...}
```

This function is called `getWord`, as the first character is lower case the function is not exported outside of the package.
Other files in the same package (directory) can call this, but nothing external can call it.
It takes a single parameter called `filename` which is a `string`.
The function returns two parameters, a `string` and an `error`.

In this example we have named the parameter, but not the returned values. We can either name or not name either.
However, we cannot name some of the parameters but not others, similarly we can't name some of the return values and not others.

It is best practice to return either 0 or 1 error, it is also best practice for it to be the last value.
We now need to return values, we can do this with the `return` value, our return statement must return all values declared in the function signature.
When returning an `error` it is best practice to return zero values for other return values, you can see this on line 16 where we are returning a string as well as an error;

```go
return "", fmt.Errorf("cannot open file: %w", err)
```

When returning in a non-error state we can return `nil` for the `error` value, as we can see at the end of the `getWord` function.

### By value or reference

Go passes by value by default. Though you can pass a reference, the `getUserInput` function has the signature

```go
func getUserInput(reader *bufio.Reader) (input string, err error) {...}
```

If we want to reference a variable we can use use `&`, fortunately when we create a new reader it is returned by reference anyway, so we don't need to make a reference to it ourselves.

### Excercises
- We created a function to print the game state, however it is not implemented ideally. We would rather have the `printGameState` function return a string where the calling function can decide what to do with it. Try modifying the function to achieve this.