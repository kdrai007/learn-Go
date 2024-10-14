package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}

func exercise() {
	user := authenticationInfo{
		username: "mail",
		password: "user",
	}

	fmt.Println(user.getBasicAuth())
}
