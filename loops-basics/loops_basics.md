# Loops in Go

The basic Loop written in Go is like standard C-like syntax: 

```go
for INITIAL;CONDITION:AFTER{
    //Do Something
}
```
example,

```go
for i:=0;i<10:i++{
    fmt.Println(i)
}
```

## Omitting condition from a loop in Go

Loops in go can omit sections of a for loop. 

- For example, `CONDITION` part of loop can be omitted result would loop will run forever

