# Pointers in Go

A variable is a named location in memory of value it stores.We can manipulate the value of a variable by assigning new value or performing operations on it. 

### A Pointer is a variable

A pointer is a variable that stores the memory address of another variable. This means that a pointer "points to" the location where original value is stored.

* The value `*` syntax defines a pointer: 

```go
var p *int
```

- The "&" operator generates a pointer to its operand.

```go
myString:="hello"
myStringPtr:=&myString
```

## References

it's possible to create pointer without assigning any value

`var p *int`

- Initally `nil` is assigned to it becuase it's not  pointing to any memory location.
- Empty pointers are called *nil pointers*

Instead of creating empty pointer it's common to use `&` operator to get a pointer to it's operand

```go
myString:="hello"
myStringPtr:=&myString
//myStringPtr stores the memory location of myString variable

fmt.Printf("value of my string %s\n", myStringPtr)
// output: value of my string 014x5060
```

As we can see compiler give some random vaue instead of "hello" string  that's because we are printing the variable which stores the memory location of the variale not the value itself.

### DeReferences

The `*` Dereferences the pointer to get the original value.

#### Pass by Reference

Generally, functions in go passed by value, meaning functions receives copy of non-composite types:

```go
func increment(x int) {
    x++
    fmt.Println(x)
    // 6
}


func main() {
    x := 5
    increment(x)
    fmt.Println(x)
    // 5
}
```
One of the common use case of pointer is to pass value in function as  reference so, if the value is change it will reflect on original variable as well.

```go
func increment(x *int) {
    x++
    fmt.Println(x)
    // 6
}


func main() {
    x := 5
    increment(&x)
    fmt.Println(x)
    // 6
}
```

### Nil Pointers

Pointers can be very dengerous 

if a pointer points to nothing ( the zero value of the pointer type) then dereferencing it will cause *runtime error* (a panic) that crashes the program. Generally speaking whenever you try to use a pointer check if it's `nil` or not

### Pointer receiver

A receiver type on a method can be a pointer.

Methods with pointer receier can modify the value to which pointer points.Since methods often used to change the value of reciever so it's common to use pointer receiver then value reciver.

- pointer receiver

```go
type car struct{
    color string
}

func (c *car) setColor(color string){
    c.color=color
}

func main(){
    car:=car{"white"}
    // color=white
    car.setColor("blue")
    // color=blue
}
```

- non-pointer receiver

```go
type car struct{
    color string
}

func (c car) setColor(color string){
    c.color=color
}

func main(){
    car:=car{"white"}
    // color=white
    car.setColor("blue")
    // color=white
}
```






