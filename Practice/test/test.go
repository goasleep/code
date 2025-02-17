package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) howOld() int {
	fmt.Printf("internal howOld, addrss is %p \n", &p)
	return p.age
}

func (p *person) growUp() {
	fmt.Printf("internal growUp, addrss is %p \n", p)
	p.age += 1
}

func main() {
	yuting := person{
		name: "yuting",
		age:  18,
	}
	fmt.Printf("person, addrss is %p \n", &yuting)
	// yuting.howOld()
	yuting.growUp()
	fmt.Printf("person, addrss is %p \n", &yuting)

	fmt.Println("------------- 2 -------------")
	yt := &person{
		name: "yuting",
		age:  18,
	}
	fmt.Printf("person, addrss is %p \n", yt)
	// yt.howOld()
	yt.growUp()
	fmt.Printf("person, addrss is %p \n", yt)
}
