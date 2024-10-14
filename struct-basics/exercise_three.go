package main

import "fmt"

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Membership
	Name string
}

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if u.MessageCharLimit >= messageLength {
		return message, true
	}
	return "", false
}

func newUser(name string, membershipType string) User {
	user := User{}
	if membershipType == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}
	user.Type = membershipType
	user.Name = name
	return user
}

func exercise_three() {
	// Create a new user with a premium membership
	// and print the user's name and message character limit
	premiumUser := newUser("John Doe", "free")
	fmt.Println("User Name:", premiumUser.Name)
	fmt.Println("Type: ", premiumUser.Type)
	fmt.Println("Limit: ", premiumUser.MessageCharLimit)

	message, canSend := premiumUser.SendMessage("Hello, World!", 120)
	fmt.Printf("Message: %s, Can Send: %t\n", message, canSend)
}
