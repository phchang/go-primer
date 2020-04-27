package main

import "fmt"

func main() {

	/*  pointers

	- different than C:
		no pointer arithmetic
	    no pointers to pointers
	- https://blog.golang.org/declaration-syntax
	*/

	// zero value is nil
	var p *int
	fmt.Println("p = ", p)

	someInt := 42
	p = &someInt

	fmt.Println("p = ", p)

	fmt.Println("value of p = ", *p) // deference a pointer

	// mutate through pointer
	*p = 43

	fmt.Println("value of p = ", *p)
	fmt.Println("value of someInt = ", someInt) // deference a pointer

	fmt.Printf("Memory address of p is <%v>, which is the same as someInt <%v>\n", p, &someInt)

	/* structs

	- collection of fields
	- not a class, but as close to a class that Go has
	- no inheritance, but can be achieved similarly by using composition

	*/

	// exporting (by capitalizing the first char of the name) applies to structs as well
	type Coordinate struct {
		Lat  float64
		Long float64
	}

	// can be declared in global, package, and function scope -- only valid inside
	// the scope where declared

	c1 := Coordinate{41.5, -98.3} // ❌ don't do this, use explicit property names
	fmt.Println("c1 = ", c1)

	c2 := Coordinate{Lat: 41.5} // Long is implied to be 0.0
	fmt.Println("c2 = ", c2)

	c3 := Coordinate{} // zero value for Lat and Long
	fmt.Println("c3 = ", c3)

	fmt.Println("lat = ", c1.Lat)
	fmt.Println("long = ", c1.Long)

	// pointers to structs

	cPointer := &c1
	fmt.Println("cPointer.Lat = ", cPointer.Lat) // convenience (don't have to deference)

	/* ---------- arrays and slices ----------  */

	var arr [10]int
	arr[0] = 1
	fmt.Println("arr =", arr)

	// ⚠️ probably will not use arrays directly... use slices

	/* slices
	- dynamically sized, view of the elements of an array
	- can slice a slice (or array) using half-open range... e.g. slc[low:high]
	*/

	var s1 []int = arr[0:1]
	fmt.Println("s1 = ", s1)

	s2 := []int{1, 2, 3} // slice literal
	fmt.Println("s2 = ", s2)

	a10 := [4]int{1, 2, 3, 4}

	s10 := a10[0:2] // slice [1, 2]
	s20 := a10[1:3] // slice [2, 3]

	fmt.Println("s10 = ", s10)
	fmt.Println("s20 = ", s20)

	s10[1] = 22

	fmt.Println("a10 = ", a10)
	fmt.Println("s10 = ", s10)
	fmt.Println("s20 = ", s20)

	// ⚠️ changes to a slice will change the underlying array
	someFunc := func(p []int) {
		p[1] = 33
	}

	someFunc(s10)
	fmt.Println("---------")
	fmt.Println("a10 = ", a10)
	fmt.Println("s10 = ", s10)
	fmt.Println("s20 = ", s20)

	// adding to a slice
	s10 = append(s10, 5) // append(s10, []int{5, 6, 7}...)
	fmt.Println("---------")
	fmt.Println("a10 = ", a10)
	fmt.Println("s10 = ", s10)
	fmt.Println("s20 = ", s20)

	/**
	a10:
	[1, 2, 3, 4]
	[1, 33, 3, 4]
	[1, 33, 5, 4]

	s10:
	[1, 33, 5]
	*/
	s10 = append(s10, 6)

	fmt.Println("---------")
	fmt.Println("a10 = ", a10)
	fmt.Println("s10 = ", s10)
	fmt.Println("s20 = ", s20)

	s10 = append(s10, 7)

	fmt.Println("---------")
	fmt.Println("a10 = ", a10)
	fmt.Println("s10 = ", s10)
	fmt.Println("s20 = ", s20)

	// ⚠️ you should use copy() -- https://blog.golang.org/slices-intro

	// zero value of slice is nil
	var s30 []string
	fmt.Println("len(s30) = ", len(s30))
	fmt.Println("cap(s30) = ", cap(s30))

	fmt.Println("is s30 == nil? ", s30 == nil)

	// use make to create a slice
	s40 := make([]int, 5)
	fmt.Println("s40 = ", s40) // optional 3rd param can be passed to specify capacity

	/* range */

	s100 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, value := range s100 { // order matters on left side of :=
		fmt.Printf("i = %v, value = %v\n", i, value)
	}

	for _, value := range []int{1, 2, 3} { // key can be ignored
		fmt.Println("value = ", value)
	}

	for index := range []int{1, 2, 3} { // if you want only index, value can be ignored
		fmt.Println("index = ", index)
	}

	/* maps */

	// zero value is nil

	var m1 map[string]int = make(map[string]int) // make with map -> length can be omitted
	fmt.Println("m1 = ", m1)

	m2 := map[string]int{
		"balls":   0,
		"strikes": 1,
		"outs":    1,
	}
	fmt.Println("m2 = ", m2)

	fmt.Println("####") // iterating through a map
	for key := range m2 {
		fmt.Println(key)
	}
	fmt.Println("####")

	m2["strikes"] = 2

	delete(m2, "outs")

	pitchCount := m2["pitchCount1"]

	fmt.Println("pitchCount = ", pitchCount)

	pitchCount, ok := m2["pitchCount"]

	if ok {
		fmt.Println("pitchCount was in m2")
	}

	// second parameters cannot be ignored in your own function returns
	//   - error handling becomes complex otherwise
	someFn := func() (int, error) {
		return 0, nil
	}

	someRes, _ := someFn()

	fmt.Println(someRes)

	pitchCounts := 0
	fmt.Println(pitchCounts)

	pitchCounts, err := someFn()
	fmt.Println(pitchCounts)
	fmt.Println(err)

}
