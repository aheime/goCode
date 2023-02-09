package main

import "fmt"

type Node struct {
	Id   int
	Name string
	Next *Node
}

func add(head *Node, node *Node) {
	//fmt.Println("node_add_start=", node)
	//fmt.Println("head_add_start=", head)

	if head.Next == nil {
		head.Id = node.Id
		head.Name = node.Name
		head.Next = head
		return
	}
	temp := head

	for {
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}

	temp.Next = node
	node.Next = head
	//fmt.Println("head=", head)
	//fmt.Println("temp=", temp)
	//fmt.Println("node_add_end=", node)
	//fmt.Println("head_add_end=", head)

}

func list(head *Node) {

	temp := head

	if temp.Next == nil {
		fmt.Println("empty")
		return
	}

	for {
		fmt.Printf("[%d,%v]==>", temp.Id, temp.Name)
		//fmt.Println("temp=", temp)
		//fmt.Println("next=", temp.Next)
		//fmt.Println("head=", head)
		if temp.Next == head {
			break
		}
		temp = temp.Next
	}
	fmt.Println()

}

func Del(head *Node, id int) *Node {

	temp := head
	helper := head

	if temp.Next == nil {
		return head
	}

	for {
		if helper.Next == head {
			break
		}
		//指向最后一个节点
		helper = helper.Next
	}

	if temp.Next == head {
		//只有一个节点
		temp.Next = nil
		return head
	}

	for {
		if temp.Next == head {
			//找到最后一个节点
			if temp.Id == id {
				//比较最后一个
				helper.Next = temp.Next
			}
			break
		}

		if temp.Id == id {
			if temp == head {
				//删除的头结点
				head = temp.Next
			}
			helper.Next = temp.Next
			break
		}

		temp = temp.Next
		helper = helper.Next

	}

	return head
}

func main() {

	head := &Node{}

	node1 := &Node{
		Id:   1,
		Name: "tom1",
	}
	node2 := &Node{
		Id:   2,
		Name: "tom2",
	}
	node3 := &Node{
		Id:   3,
		Name: "tom3",
	}
	//fmt.Println("node0=", node1)
	//fmt.Println("head0=", head)

	add(head, node1)
	//fmt.Println("node1=", node1)
	//fmt.Println("head1=", head)
	add(head, node2)
	//fmt.Println(head)
	add(head, node3)

	//fmt.Println(head.Next)
	//fmt.Println(node2)
	list(head)
	//fmt.Println(head)

	head = Del(head, 2)
	list(head)

	head = Del(head, 1)
	list(head)

}
