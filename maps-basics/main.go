package main

import "fmt"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	contact := make(map[string]user)
	for i := 0; i < len(names); i++ {
		contact[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}

	}
	return contact, nil
}

type user struct {
	name        string
	phoneNumber int
}

func countInstance(str string) {
	names := map[string]int{}
	missingNames := []string{}
	if _, ok := names[str]; !ok {
		missingNames = append(missingNames, str)
	}
	fmt.Println("missinga names: ", missingNames)
}
func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		} else {
			validUsers[user] = 1
		}
	}
	fmt.Println(validUsers)
}

func main() {
	// monthDays := make(map[string]int)
	// monthDays["january"] = 31
	// monthDays["fabruary"] = 28
	// monthDays["march"] = 31
	// monthDays["aprill"] = 30
	// fmt.Println(monthDays)
	// countInstance("kdrai")

	//write tests for getCounts function

	// getCounts([]string{"cersei", "tyrian", "jaime", "cersei"},
	// 	map[string]int{"cersei": 0, "jaime": 0})

	// fmt.Println(getUserMap([]string{"a", "b", "c"}, []int{1, 2, 3}))

	// exerciseOne()
	exerciseTwo()
}
