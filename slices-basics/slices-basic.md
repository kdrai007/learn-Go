# Arrays in Go

Arrays are fixed-size group of same type. 

- Here's syntax for declearing array in Go.

```go
var arr[4]int 
//or to decclear values in array

arr2:=[4]int{1,2,3,4}
```

## Slices in Go

A slice is a descriptor for a contigous segement of an underlying array  and provided an access to a numbered sequence from that Array.

#### How slice work:-

- A slice is essentially a descriptor that holds:  

    1. A pointer to the underlying array.
    2. The length of the slice (number of accessible elements)
    3. The capacity of the slice.(the maximum number of accessible array it can access)


99 times out of 100 I will or you will use slices instead of array with ordered list.

- A slice is dynamically-sized flexible view of the elements in array.
- The zero value of slice is `nil`
- Slices always have an underlying array, though it isn't always specified explicitly


To explicitly create slice on top of array we can use this :-

example,

```go
primes:=[6]int{1,2,3,4,5,6}
mySlices:=primes[1:4]
//output: {2,3,4}
```
syntax,

```
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
```

where, `lowIndex` is inclusive and `highIndex` is exclusive
- Both lowIndex and highIndex can be removed to get full array

### Make

Most of the time we don't need to think about underlying array of a slice. We an create a new slice using the `make` function: 

```go
//func make([]T,len,cap) []T
mySlice:=make([]int,5,9)

// the cpacity arguments is omitted by default to length
mySlice:=make([]int,5)
```
- Slice create with make will fill 0 value of the type.

<details open>
<summary style="font-weight:bold; font-size:1.1rem;">Length</summary>
<br>
<p>The length of the slice is simply the values it contains.It is accessed using the built in <code>len()</code></p>

```go
mySlice:=[]int{1,2,3,4,5}
fmt.Println(len(mySlice))
//output: 5 
```
</details>

<details>
<summary style="font-weight:bold; font-size:1.1rem;">Capacity</summary>
<br>
<p>The cpacity of slice is the number of elements in the underlying array,counting from the last element of the slice. It is accessed using the built-in <code>cap()</code>function</p>

```go
mySlice:=[]int{1,2,3,4,5}
fmt.Println(cap(mySlice))
//output: 5 
```
</details>

### Variadic

Many functions, especially those in the standard library, can take an arbitrary number of final arguments.  This is accomplished by using the "..." syntax in the function signature.

A variadic function receive the variadic arguments as a slice.


```go
func concat(str ...string) string{
    res:=""
    for i=0;i<len(str);i++{
        res+=str[i]
    }
    return res
}
```

#### Spread Operator

The spread function allows us to pass a slice into variadic function. The spread operator consist of three dots following the slice of a function call.

```go
func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
    names := []string{"bob", "sue", "alice"}
    printStrings(names...)
}
```
### Append

A built-in `append` function is used to dynamically append elements to slice.

`func append(slice []Type , elems ...Type)[]Type`

- If the underlying array is not big enough to contain more elements then `append()` will create new underlying array and point the returned slice to it.

here the example,

```go
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
```