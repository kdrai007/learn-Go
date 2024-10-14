# Go Basics

Go is compiled language

### Go is statically typed 

Go enforces *static typing* meaning variables types are known before the code runs.

## Type Sizes

- Whole Numbers(No Decimal)

```
int int8 int32 int64
```

- positive numbers (No Decimal)

```
uint uint8 uint32 uint64 uintptr
```

- Signed decimal numbers  

```
float32 float64
```

- Imaginary numbers (rarely used)

```
complex64 complex128
```

The size (8,16,32,64,128,etc) represents how many *bits* in memory will be used to store the variable. 

- The "default" `int` and `uint` types refer to their respective 32 or 64 bit sizes depending on the environment of the user.

### Converting Between Types

- Some type can be easily converted like this:

```go
temperaturefloat :=88.26
termperatureInt=int64(termperaturefloat)
```

### Strings

#### Concatenating Strings

Two strings can be concatenated with `+` operator,But the compiler will not allow you to concatenate a string varible with an `int` or `float64`

### Runes and Strings Encoding

In many programming languages a `char` is a single byte. Using `ASCII` encoding, the standard for c programming language, we can represent 128 characters with 7 bit. This is enough for English alphabet, numbers and special characters.

- Go uses UTF-8 encoding, which is variable length-encoding
- Go strings can handle charcters from any language.

#### Ignore Return Values

suppose function return 2 value x,y and you want to ignore y for some purpose,

`x,_=fun()`


### Defer

The `defer` keywoard is fairly unique feature of Go, It allows a function to be executed automatically just beore it's enclosing function returns

### Block scope

Unlike, `python` which is *function-scoped* `Go` is *block-scoped*

### Closures

A `closure` is a function that reference variables outside of the function body. 

- The function may access and assign values to that variables 

```go
func concatter() func(string) string{
    strPoint:=""
    return func (word string) string{
        strPoint+= word+ " "
        return strPoint
    }
}
```
