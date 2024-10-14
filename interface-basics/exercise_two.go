package main

import "fmt"

// func getExpenseReport(e expense) {
// 	report, ok := e.(email)
// 	if ok {
// 		fmt.Printf("Email cost: $%v\n", report.cost())
// 		return
// 	} else {
// 		report := e.(sms)
// 		fmt.Printf("Sms cost: $%v\n", report.cost())
// 		return
// 	}
// }

func getExpenseReport(e expense) {
	switch v := e.(type) {
	case email:
		fmt.Printf("%T is email\n", v)
	case sms:
		fmt.Printf("%T is sms\n", v)
	default:
		fmt.Printf("%T what's this\n", v)
	}
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func exerciseTwo() {
	// e := email{
	// 	isSubscribed: false,
	// 	body:         "Hello, this is a test email",
	// 	toAddress:    "kd@gmail.com",
	// }
	// ex := expense(e)
	// getExpenseReport(ex)

	s := sms{
		isSubscribed:  false,
		body:          "Hello, this is a test sms",
		toPhoneNumber: "1234567890",
	}
	ex := expense(s) //
	getExpenseReport(ex)

}
