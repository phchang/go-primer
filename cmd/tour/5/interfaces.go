package main

import "fmt"

func main() {
	/* interfaces
	- set of method signatures
	- interfaces are implemented implicitly
		- a type implements an interface by implementing all interface methods
		- no explicit declaration (e.g. 'implements')
	- Stringer is a commonly implemented interface
		- also error
	*/

	type Printer interface {
		Print()
	}

	var printer Printer

	var myStr HyphenatedString
	myStr = "hello"

	myStr.Print()

	printer = myStr

	fmt.Println("printing using interface method")
	printer.Print()

	var lb LivingBeing

	person := Person{}
	lb = &person

	fmt.Println("lb = ", lb)

	// empty interface
	//   - may hold values of any type

	/* type assertions */

	var unknownType interface{}
	unknownType = "some string"

	unknownTypeAsString := unknownType.(string)
	fmt.Println("unknownTypeAsString = ", unknownTypeAsString)

	// unknownTypeAsInt := unknownType.(int) <-- this will panic

	// prevent panic
	unknownTypeAsInt, ok := unknownType.(int)

	if ok {
		fmt.Println("unknownType is indeed an int, with value =", unknownTypeAsInt)
	} else {
		fmt.Println("unknownType is not an int")
	}

	switch v := unknownType.(type) {
	case int:
		fmt.Println("unknownType is an int = ", v)
	case string:
		fmt.Println("unknownType is a string = ", v)
	default:
		fmt.Println("unknownType is neither an int or a string")
	}

}

type LivingBeing interface {
	FullName() string
	ChangeFirstName(string)
}

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return fmt.Sprintf("%v %v", p.FirstName, p.LastName)
}

func (p *Person) ChangeFirstName(name string) {
	p.FirstName = name
}

type HyphenatedString string

func (s HyphenatedString) Print() {
	for i, c := range s {
		fmt.Printf("%c", c)
		if i < len(s)-1 {
			fmt.Print("-")
		}
	}
	fmt.Print("\n")
}
