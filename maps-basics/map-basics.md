# maps in Go

`Maps` are similer to javascript's object, python's dictionary and ruby's hashes. Maps are a data strcture that provides key-value pairing.

- Zero value of maps is `nil`
- We can create maps by using `make()` function, like this:-

```go
sms:=make(map[string]string)
sms["kdrai"]="hey there how are you"

// or we can put values like this
sms=map[]{
    "sender 1":"message 1",
    "sender 2":"message 2",
}
```

- the `len()` function work in map return total key-value pair.

```go
fmt.Println(len(sms))
//output: 2
```
### Like slices, map hold references

Like slices, map hold reference of underlying data structure. if you pass map to a function as a argument if you change value of map in function it will effect original map too.

#### delete key-value pair

To delete a key-value pair in map it's pretty easy use `delete(map,key)` 

`delete(map,"key")`


#### To check if any key-pair exist in map or not

To check any key-value pair exist in map we can use this syntax,

```go
elem,ok:=m[key]
if ok{
    fmt.Println(elem)
}else{
    fmt.Println("invalid operation")
}
```

- if key exists in map then value of `ok` will be `true` or vice-versa

### Key types

Any type can used as value in the `map`, but keys are more restrictive

- maps keys can be **any type that can be compareable** like, 
  - comparable types are boolean, numeric, string,pointer,channel, and interfaces types and struct and arrays that contain only those types



### Nested

Maps can contain maps, creating a nested structure

example,

```go
map[string]map[string]int
map[rune]map[string]bool
map[int]map[string]string
```
