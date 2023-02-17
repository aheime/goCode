package main

import "fmt"

type Student1 struct {
	*Person
	id   int
	addr string
}

type Person struct {
	name string
	sex  string
	age  int
}

func (p Person) say() {
	p.name = "ss"
	fmt.Println(p)
	fmt.Printf("%p\n", &p)
}

func (p *Person) eat() {
	p.name = "ee"
	fmt.Printf("%p\n", p)
}

func main() {
	s1 := Student1{&Person{"11", "man", 18}, 1, "bj"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)

	s1.say()
	fmt.Println(s1.Person)

	s1.eat()
	fmt.Println(s1.Person)
}
