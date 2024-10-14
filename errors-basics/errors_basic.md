# Error in Go

## The Error Interface

Go programs express error with `error` values. An Error is any type that implements the simple built-in [error interface](https://blog.golang.org/error-handling-and-go):

```go
type error interface{
    Error() string
}
```
- When something goes wrong in the function, the function should return an `error` as its last return value. 
- Any code that calls a  function that return an `error` should handle `errors` by testing whether the error is `nil`.

Because errors are just interfaces you can built your custom types that implements the `error` interface.
- Here's an example of `userError` struct that implements `errors` interface.

```go
type userError struct{
    name string
}

func (u userError) Error() string{
    retrun fmt.Sprintf("%v has a problem with their account", e.name)
}
```
- It can be used as an error like this:

```go
func sendSMS(recipient,message string) error{
    if !canSendToUser(recipient){
        return userError{name:recipient}
    }
    ...
}
```
- Go programs handle errors with `errors` values.
- Error-values are any type that implements the simple built-in error-interface.

### The error package

The Go standard library provides an "errors" package that makes it easy to deal with errors

here's an easy implementation:-
```go
var err error = errors.New("something went wrong")
```

### Panic

We learnt previously in Go, to handle error use Error interface. However, there is another way to deal with errors in Go: the `panic` function. 

- When a function calls `panic`, the program crashes and prints a stack trace.

> As a general rule, don't use panic!

The panic function yeets control out of the current function and up the call stack until it reaches a function that defers a recover. If no function calls recover, the goroutine (often the entire program) crashes.

```go
func enrichUser(userID string) User {
    user, err := getUser(userID)
    if err != nil {
        panic(err)
    }
    return user
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()

    // this panics, but the defer/recover block catches it
    // a truly astonishingly bad way to handle errors
    enrichUser("123")
}
```

- I suggest don't use panic/recover, but when there's truly any unrecoverable error  use `log.fatal` to print a message and exit the program.

## Formatting String review

A convenient way to format string in Go is by using  standard library `fmt.Sprintf()` function. It's a string interpolation function , similar to python's `f-strings` 

- The `%v` substring uses the type's default formatting, which is often what you want.

```go
value:=15
str:="hey there! num"
fmt.Sprintf("%v %v",value,str) // hey there! num 15
```

