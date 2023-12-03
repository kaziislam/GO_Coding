package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func main() {
	hourOfDay := time.Now().Hour()
	fmt.Println(hourOfDay)
	greeting, err := getGreetings(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(greeting)
}


func getGreetings(hour int) (string, error)  {
	var message string

	if hour < 7 {
		err := errors.New("Too early for greetings!")
		return message, err
	}

	if hour < 12 {
		message = "Good morning!"
	} else if hour < 18 {
		message = "Good Afternoon!"
	} else {
		message = "Good Evening!"
	}

	return message, nil

}