# Arrays in Go

Arrays are fixed-size group of same type. 

- Here's syntax for declearing array in Go.

```go
var arr[4]int 
//or to decclear values in array

arr2:=[4]int{1,2,3,4}
```

## Slices in Go

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





