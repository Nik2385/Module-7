package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func changePerson(person **Person) {
	*person = &Person{
		name: "Vladimir",
		age:  25,
	}
	//fmt.Println("func:", *person.name)
}

func main() {

	person := &Person{
		name: "Ivan",
		age:  30,
	}

	fmt.Println(person.name)
	changePerson(&person)
	fmt.Println(person.name)

}

