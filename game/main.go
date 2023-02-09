package main

import "fmt"

type Boy struct {
	Id   int
	Next *Boy
}

func add(nums int) (head *Boy) {

	head = &Boy{}
	CurrentBoy := &Boy{}

	for i := 1; i <= nums; i++ {
		boy := &Boy{Id: i}
		if i == 1 {
			head = boy
			CurrentBoy = boy
			CurrentBoy.Next = head
		} else {
			CurrentBoy.Next = boy
			CurrentBoy = boy
			CurrentBoy.Next = head
		}
	}

	return
}

func show(head *Boy) {

	CurBoy := head

	if CurBoy.Next == nil {
		fmt.Println("请先添加")
		return
	}
	for {
		fmt.Printf("[%d]->", CurBoy.Id)
		if CurBoy.Next == head {
			break
		}

		CurBoy = CurBoy.Next
	}

	fmt.Println()
}

func play(head *Boy, start int, count int) {

	tail := head

	for {
		if tail.Next == head {
			break
		}
		tail = tail.Next
	}

	for i := 1; i <= start-1; i++ {
		head = head.Next
		tail = tail.Next
	}

	for {
		for j := 1; j <= count-1; j++ {
			head = head.Next
			tail = tail.Next
		}
		fmt.Printf("%d号出列\n", head.Id)
		head = head.Next
		tail.Next = head

		if head == tail {
			fmt.Printf("%d号最后出列", head.Id)
			break
		}
	}

}

func main() {

	head := add(50)

	show(head)

	play(head, 2, 3)

}
