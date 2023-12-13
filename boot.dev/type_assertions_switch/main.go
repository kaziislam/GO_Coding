package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .000005
	}
	return float64(len(e.body)) * .000001
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .000001
	}
	return float64(len(s.body)) * .000003
}

func (i invalid) cost() float64 {
	return 0.0
}

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

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	switch e := e.(type) {
	case email:
		return e.cost() * float64(averageMessagesPerYear)
	case sms:
		return e.cost() * float64(averageMessagesPerYear)
	default:
		return 0.0
	}
}

func test(e expense, averageMessagesPerYear int) {
	address, _ := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, estimateYearlyCost(e, averageMessagesPerYear))
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, estimateYearlyCost(e, averageMessagesPerYear))
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	averageMessagesPerYear := 1000
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	}, averageMessagesPerYear)
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	}, averageMessagesPerYear)
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	}, averageMessagesPerYear)
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	}, averageMessagesPerYear)
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	}, averageMessagesPerYear)
	test(invalid{}, 0)
}
