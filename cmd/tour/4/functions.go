package main

import "fmt"

func main() {
	/* functions as values, first order types */

	sayHello := func(fn func(string) string, name string) string {
		return fn(name)
	}

	helloInEnglish := func(name string) string {
		return name + " says hello"
	}

	helloInSpanish := func(name string) string {
		return name + " says hola"
	}

	english := sayHello(helloInEnglish, "sung")
	spanish := sayHello(helloInSpanish, "paul")

	fmt.Println("english: ", english)
	fmt.Println("spanish: ", spanish)

	//////////////////////////////

	type Sayer func(string) string

	var helloInEmoji Sayer

	helloInEmoji = func(name string) string {
		return name + " says ✋"
	}

	emoji := sayHello(helloInEmoji, "john")

	fmt.Println("emoji: ", emoji)

	// function on a function
	var g Greeter
	g = func(s string) string {
		return emoji
	}
	fmt.Println("g = ", g)
	g.Say()

	/* methods
	- no classes in Go, but methods on types
	- methods are functions with a receiver argument
	*/

	fmt.Println("---------- METHODS ----------")

	person := Person{
		FirstName: "John",
		LastName:  "Smith",
	}

	fmt.Println(person.FullName())

	person.ChangeFirstName("Larry") // maybe discuss pointer indirection, if the question comes up

	// choosing value vs pointer receiver
	//  - modify value of receiver
	//  - avoid copying the value on each method call

	fmt.Println(person.FullName())

	var myStr HyphenatedString
	myStr = "hello"

	myStr.Print()

	//var someInterface SomeInterface = SomeImplementation{}

}

type Greeter func(string) string

func (g Greeter) Say() {
	fmt.Println("inside Say")
}

////////////////////////

type Person struct {
	FirstName string
	LastName  string
}

// mixed pointer and value is not recommended
// - this is an issue with interface types

type SomeInterface interface {
	A()
	B()
}

type SomeImplementation struct {
}

func (s SomeImplementation) A() {
}

func (s SomeImplementation) B() {
}

func (p Person) FullName() string {
	return fmt.Sprintf("%v %v", p.FirstName, p.LastName)
}

func (p *Person) ChangeFirstName(name string) {
	p.FirstName = name
}

///////////////////////////

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
