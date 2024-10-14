# Interface in Go

Interfaces are just collection of methods signatures. A type "implements" an interface if it has methods that match the interface method's signature.

- **Interfaces are implemented implicitly** 
A type implements an interface by implementing its methods. Unlike in many other languages, there is no explicit declaration of intent, there is no "implements" keyword.

for example,

```go
type shape interface{
    area() float64
    perimeter() float64
}

type rect struct{
    height,width float64
}

func (r rect) area() float64{
    return r.height*r.width
}

func (r rect) perimeter() float64{
    return 2*r.height+2*r.width
}

```
- In the following example, a *shape* must be able to return `area` and `perimeter` 

### The empty Interface

The interface type specifes zero methods is known as *empty interface*

`interface{}`

- An empty interface can hold value of any type 
- An empty interface are used by codes that handle unknown value type

### Multiple Interface

A type can implment many inerface in Go. for example, the empty interface, `interface{}`, is always implemented by every type because it has no requirments

### Name your Interface Arguments

consider this example,

```go
type copier interface{
    Copy(string string) int 
}
```

- Based on just the code can you deduce what kind of string you should pass as arguments
- So, here's an feature you can name your arguments in interface lik this:


```go
type copier interface{
    Copy(sourceFile string, destinationFile string) (fileSize int) 
}
```

- Here it is , now it's pretty easy to understand what type of arguments you have to pass to the function 

## Type assertions in Go

When working with interfaces in Go, Every once in a while you'll need access to underlying type of an interface value. You can cast an interface to its underlying type using *type-assertions*

for example,

```go
func main(){
    var i interface{}="hello"
    c:=i.(string)

    fmt.Println(c) // output: hello
}
```
- In the given example, first we assign interface value to `i` with `string` `hello`
- then, we use `i.(string)` to assign undelying value of interface to `c`
- then we print the output

we careful when you use type-assertions without taking any precautions like, If there isn't any `string` in interface then this will cause `panik` to prevent that we use `ok` ideoms like this :- 

```go
func main(){
    var i interface{}="hello"
    c,ok:=i.(string)

    if ok{
        fmt.Println(c) // output: hello
     }else{
       // whatever you want to do  
     }
}
```

### Type Switches in Go

A *type switch* makes it easy to do several  types assertions in a series.

- A type switch is similer to the normal switch statement, but the cases specify types instead of values

example,

```go
switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
```

## Clean Interface

Keeping code simple and clean is every programmer dream,but anytime you work with abstraction in code,simple become complex pretty quickly 

- Some rules to keep in mind for keeping interfaces clean


1. Keep interfaces small
2. Interfaces should have no knowledge of satisfying types
3. Interfaces are not classes


<details open>
    <summary style="font-weight: bold;font-size:1.2rem;margin-bottom:10px">1. keep Interfaces small</summary>
    <p style="margin-bottom:10px;">If you learn anything from this sentence, make it this: keep interfaces small! interfaces are meant to define minimal behaviour to represnt idea or concept</p>

<p>Here an example from the standard HTTP package of a larger interface :-</p>

```go
type File interface {
    io.Closer
    io.Reader
    io.Seeker
    Readdir(count int) ([]os.FileInfo, error)
    Stat() (os.FileInfo, error)
}
```
<small> - That's a good example of defining minimal behaviour</small>
</details>


<details>
    <summary style="font-weight: bold;font-size:1.2rem;margin-bottom:10px">2. Interfaces should Have No knowledge of Satisfying Types</summary>
    <p style="margin-bottom:10px;">An interface should define what's necessary  for others types to classify as a member of interface. They shouldn't be aware of any types that happen to satisfy the interface at design time.</p>

<p> for example,</p>

```go
type car interface{
    Color() string
    Speed() float64
    isFireTruck() bool
}
```
<p><code>Color()</code> and <code>Speed()</code> make perfect sense,thye are methods confined to the scope of a car. but is <code>isFireTruck()</code> is really necessary to add in interface. </p>


</details>

<details>
    <summary style="font-weight: bold;font-size:1.2rem;margin-bottom:10px">3. Interfaces are not classes</summary>
    <ul>
    <li>Interfaces are not classes, they are slimmer</li>
    <li>Interfaces don't have constructors or deconstructors that require that data is created or destroyed</li>
    <li>Interfaces aren't hierarchical by nature ,though they are syntactic sugar to create interfaces that happen to be supersets of other interfaces </li>
    <li>Interfaces define functions signatures, but not underlying behaviour. Making an interface often won't DRY up your code in regards to struct methods.</li>
    </ul>
</details>


