package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person person
	ltk    bool
}

func (r secretAgent) canKill(name string) bool {
	return r.person.firstName == name
}

func main() {
	p1 := person{
		firstName: "Gustavo",
		lastName:  "Santos",
	}

	p2 := person{
		firstName: "Rosana",
		lastName:  "Santos",
	}

	p3 := secretAgent{
		person: p2,
		ltk:    true,
	}

	fmt.Println(p3.canKill("Gustavo"))
	fmt.Println(p3.canKill("Rosana"))

	fmt.Println(p1.firstName, p2.firstName)
	fmt.Println(p3)
}
