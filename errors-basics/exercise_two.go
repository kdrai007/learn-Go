package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	charLen := len(status)
	if charLen == 0 {
		return errors.New("Message can't be empty")
	} else if charLen > 140 {
		return errors.New("Message can't be more than 140 characters")
	}
	return nil
}

// Helper function to run tests manually
func runTests() {
	// Test cases
	tests := []struct {
		name          string
		status        string
		expectedError error
	}{
		{
			name:          "Empty status",
			status:        "",
			expectedError: errors.New("Message can't be empty"),
		},
		{
			name:          "Status exceeding 140 characters",
			status:        "This status exceeds the 140-character limit by having a lot of unnecessary words. Let's keep typing until we go beyond the threshold. Almost there. Done!",
			expectedError: errors.New("Message can't be more than 140 characters"),
		},
		{
			name:          "Valid status",
			status:        "This is a valid status message.",
			expectedError: nil,
		},
	}

	// Loop through the tests
	for _, test := range tests {
		fmt.Printf("Running test: %s\n", test.name)

		// Call validateStatus
		err := validateStatus(test.status)

		// Check the result
		if err != nil && test.expectedError != nil && err.Error() == test.expectedError.Error() {
			fmt.Println("PASS")
		} else if err == nil && test.expectedError == nil {
			fmt.Println("PASS")
		} else {
			fmt.Printf("FAIL - Expected: %v, Got: %v\n", test.expectedError, err)
		}
		fmt.Println("-----------------------------------")
	}
}
func exerciseTwo() {
	fmt.Println("--------exercise two---------")
	runTests()
}
