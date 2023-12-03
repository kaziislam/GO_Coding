package main

import "fmt"

type potluck struct {
	name string
	number int
	food menu
}

type menu struct {
	rice string
	meat string
	dessert string
}

func test(p potluck){
	fmt.Printf(`%s are coming to the party total %v guests; and bringing %s, %s and %s.`,
		p.name,
		p.number,
		p.food.rice,
		p.food.meat,
		p.food.dessert,
	)
	fmt.Println()
}

func main() {
	test(potluck{
		name: "Pallab and family",
		number: 3,
		food: menu{
			rice: "polaw",
			meat: "chicken",
			dessert: "brownies",
		},
	})
	test(potluck{
		name: "Shah and family",
		number: 4,
		food: menu{
			rice: "rice",
			meat: "goat",
			dessert: "sweets",
		},
	})
	test(potluck{
		name: "Anwar and family",
		number: 4,
		food: menu{
			rice: "yellow rice",
			meat: "beef",
			dessert: "muffins",
		},
	})
}
