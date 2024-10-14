package main

import "fmt"

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}

type Formatter interface {
	Format() string
}

type Bold struct {
	message string
}

type plainText struct {
	message string
}

type Code struct {
	message string
}

func (b Bold) Format() string {
	return "**" + b.message + "**"
}

func (p plainText) Format() string {
	return p.message
}

func (c Code) Format() string {
	return "`" + c.message + "`"
}

func exerciseFour() {
	fmt.Println("--Exercise Four---")
	plain := plainText{message: "Hello, Gophers!"}
	bold := Bold{message: "Hello, Gophers!"}
	codeText := Code{message: "Hello, Gophers!"}

	formatter := []Formatter{plain, bold, codeText}

	for _, m := range formatter {
		fmt.Println(m.Format())
	}
}
