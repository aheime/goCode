package main

import "fmt"

type Node struct {
	Id   int
	Name string
	Next *Node
}

func add(head *Node, node *Node) {

	temp := head
	for true {
		if temp.Next == nil {
			break
		} else if temp.Next.Id >= node.Id {
			break
		}
		temp = temp.Next
	}

	node.Next = temp.Next
	temp.Next = node
}

func del(head *Node, id int) {
	var flag = false
	temp := head
	for true {
		if temp.Next == nil {
			break
		} else if temp.Next.Id == id {
			flag = true
			break
		}
		temp = temp.Next
	}

	if flag {
		temp.Next = temp.Next.Next
	}
}

func update(head *Node, id int, name string) {
	var flag = false
	temp := head
	for true {
		if temp.Next == nil {
			break
		} else if temp.Next.Id == id {
			flag = true
			break
		}
		temp = temp.Next
	}

	if flag {
		temp.Next.Name = name
	}
}

func list(head *Node) {

	temp := head
	if temp.Next == nil {
		fmt.Println("empty")
		return
	}
	for {
		fmt.Printf("[%d,%v]==>", temp.Next.Id, temp.Next.Name)

		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
	println()
}

func main() {
	var head = &Node{}

	node1 := &Node{
		Id:   1,
		Name: "one",
	}

	node2 := &Node{
		Id:   2,
		Name: "two",
	}

	node3 := &Node{
		Id:   3,
		Name: "three",
	}

	node4 := &Node{
		Id:   2,
		Name: "two2",
	}
	add(head, node1)
	add(head, node2)
	add(head, node3)
	add(head, node4)

	list(head)

	del(head, 2)
	list(head)

	update(head, 1, "one111")
	update(head, 3, "wefw")

	list(head)

	//fmt.Println(head)
	//fmt.Println(*head.Next)
	//fmt.Println(*head.Next.Next)
}
