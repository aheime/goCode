package main

import "fmt"

type Node struct {
	Id   int
	Name string
	Pre  *Node
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
	node.Pre = temp
	if temp.Next != nil {
		temp.Next.Pre = node
	}
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
		if temp.Next != nil {
			temp.Next.Pre = temp
		}

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

func asc(head *Node) {

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

func desc(head *Node) {
	temp := head
	if temp.Next == nil {
		fmt.Println("empty")
		return
	}

	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	for {
		fmt.Printf("[%d,%v]==>", temp.Id, temp.Name)
		temp = temp.Pre
		if temp.Pre == nil {
			break
		}

	}
	fmt.Println()
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

	asc(head)
	desc(head)

	del(head, 2)
	asc(head)

	update(head, 1, "one111")
	update(head, 3, "wefw")

	asc(head)

	//fmt.Println(head)
	//fmt.Println(*head.Next)
	//fmt.Println(*head.Next.Next)
}
