package main

import (
	"fmt"
	"time"
)

func main() {

	/* for loop -> only looping construct
	- traditional for init; condition; post
	- init and post are optional
	- infinite loop: for {}
	*/
	for i := 0; i < 10; i++ {
		// ^-init ^-cond ^-post statement
		fmt.Printf("%v ", i)
	}
	fmt.Println("")

	// init and post are optional -> makes it a while loop
	i := 10
	for i < 20 {
		fmt.Printf("%v ", i)
		i++
	}
	fmt.Println("")

	// infinite loop: for { ... }

	/* switch

	- break is not needed, differs from other languages
	- evaluated from top to button

	*/
	fmt.Println("What day is today?")

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Monday")
	case time.Monday + 1:
		fmt.Println("Tuesday")
	default:
		fmt.Println("It's not Monday or Tuesday")
	}

	fmt.Println("------")

	today := time.Now().Weekday()
	// cases are variable, normally variable is passed in with static cases
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}

	// switch without condition is same as `switch true`, a way of writing long if-else chains
	switch {
	case today == time.Monday:
		fmt.Println("Monday")
	}

	/* Defer */

	deferTest := func() {
		defer func() {
			fmt.Println("inside defer")
		}()

		defer func() {
			fmt.Println("inside second defer")
		}()

		fmt.Println("inside here")
	}

	deferTest()

	/* errors handling
	- no concept of exceptions
	- common convention is to return an error, nil value if no error
	*/
}
