package main

import "fmt"

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(analytics *Analytics, message Message) {
	if message.Success {
		analytics.MessagesSucceeded++
		analytics.MessagesTotal++
	} else {
		analytics.MessagesFailed++
		analytics.MessagesTotal++
	}
	fmt.Printf(`
		current recipient: %s
		Message failed: %d
		successfully send: %d
		totalMessages: %d
	`, message.Recipient, analytics.MessagesFailed, analytics.MessagesSucceeded, analytics.MessagesTotal)
}

func exerciseOne() {
	analytics := Analytics{
		0, 0, 0,
	}
	getMessageText(&analytics, Message{"kdrai", false})
	getMessageText(&analytics, Message{"kdrai", true})
	getMessageText(&analytics, Message{"kdrai", false})
}
