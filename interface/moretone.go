package main

import "fmt"

//多个类型实现同个接口

type Eater interface {
	eat()
}

type Duck struct {
	name string
}

type Cat struct {
	name string
	*Fish
}

type Fish struct {
	name string
}

func (d *Duck) eat() {
	fmt.Printf("%v\n", d.name)
}

func (c *Cat) eat() {
	fmt.Printf("%v\n", c.name)
}

func (f *Fish) eat() {
	fmt.Printf("%v\n", f.name)
}

func main() {
	var x Eater
	var duck = &Duck{"鸭紫"}
	var cat = &Cat{"肥猫", &Fish{"小鱼儿"}}
	x = duck
	x.eat()
	x = cat
	x.eat()
	x = cat.Fish
	x.eat()
}
