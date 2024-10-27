package main

import "fmt"

func suggestFriends(username string, friendsList map[string][]string) []string {
	userFriend, ok := friendsList[username]
	// if there is no user in list then return null
	if !ok {
		return nil
	}
	friendDict := make(map[string]bool)
	for _, fr := range userFriend {
		friendDict[fr] = true
	}

	distinctSuggestFriendList := make(map[string]bool)
	for _, fr := range userFriend {
		if f, ok := friendsList[fr]; ok {
			for _, cfr := range f {
				if cfr != username && !friendDict[cfr] {
					distinctSuggestFriendList[cfr] = true
				}
			}
		}
	}
	suggestFriendList := []string{}
	for key := range distinctSuggestFriendList {
		suggestFriendList = append(suggestFriendList, key)
	}

	if len(suggestFriendList) == 0 {
		return nil
	}

	return suggestFriendList
}

func exerciseTwo() {
	fmt.Println("-------Exercise 2---------")

	friendships := map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	}
	suggestFriends("Alice", friendships)
}
