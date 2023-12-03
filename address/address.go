package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone string `json: "phone"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Enter address: ")
	address, _ := reader.ReadString('\n')
	address = address[:len(address)-1]

	fmt.Print("Phone number: ")
	phone, _ := reader.ReadString('\n')
	phone = phone[:len(phone)-1]

	person := Person{Name: name, Address: address, Phone: phone}

	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(personJSON))
}
