package main

import "fmt"

//同个类型实现多个接口

type Sayer interface {
	say(word string)
}
type Mover interface {
	move()
}

type Animal interface {
	Sayer
	Mover
}

type Dog struct {
	name string
}

func (d *Dog) say(word string) {
	fmt.Printf("%v say %v\n", d.name, word)
}

func (d *Dog) move() {
	fmt.Printf("%v move\n", d.name)
}

func main() {
	//1
	//var dog = &Dog{"小明"}
	//var say Sayer = dog
	//var move Mover = dog
	//say.say("wang")
	//move.move()

	//2
	var animal Animal
	var dog = &Dog{"小明"}
	animal = dog
	animal.say("wang")
	animal.move()

}
