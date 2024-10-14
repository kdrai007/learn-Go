package main

import (
	"fmt"
	"reflect"
)

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

func exercise_two() {

	typ := reflect.TypeOf(perms{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())

	fmt.Println("Here we go! Exercise Two")
}
