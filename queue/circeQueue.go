package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	MaxSize int
	Items   []int //只能村4-1个
	Head    int
	Tail    int
}

func (t *Queue) Push(val int) (err error) {
	if t.IsFull() {
		return errors.New("queue is full")
	}
	t.Items[t.Tail] = val
	//将tail指向下一个位置
	t.Tail = (t.Tail + 1) % t.MaxSize
	fmt.Println("Tail=", t.Tail)
	return
}

func (t *Queue) Pop() (val int, err error) {

	if t.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	val = t.Items[t.Head]
	//将head指向下一个位置
	t.Head = (t.Head + 1) % t.MaxSize
	fmt.Println("Head=", t.Head)
	return
}

func (t *Queue) Size() int {
	return (t.Tail - t.Head + t.MaxSize) % t.MaxSize
}

func (t *Queue) Show() {
	size := t.Size()
	if size == 0 {
		fmt.Println("当前队列为空")
	}
	temp := t.Head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", temp, t.Items[temp])
		temp = (temp + 1) % t.MaxSize
	}
}

func (t *Queue) IsEmpty() bool {
	return t.Head == t.Tail
}

func (t *Queue) IsFull() bool {
	return (t.Tail+1)%t.MaxSize == t.Head

}

func main() {

	var size = 4
	var queue = Queue{
		MaxSize: size,
		Items:   make([]int, size),
	}

	for {
		var key string
		fmt.Scanln(&key)
		switch key {
		case "push":
			var value int
			fmt.Scanln(&value)
			err := queue.Push(value)
			if err != nil {
				fmt.Println(err)
			}
		case "pop":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(val)
		case "show":
			queue.Show()
		case "tail":
			fmt.Println(queue.Tail)
		case "head":
			fmt.Println(queue.Head)
		case "size":
			fmt.Println(queue.Size())

		}
	}

}
