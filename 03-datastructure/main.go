package main

import (
	"fmt"

	"github.com/zh-liang-cn/learning-go/03-datastructure/ds"
)

func main() {
	q := ds.NewLLQueue("first", "second", "third")

	q.Add("fourth")
	fmt.Println(q.ToString())

	fmt.Println("get:", q.Get())
	fmt.Println(q.ToString())

	q.Add("fifth")
	fmt.Println(q.ToString())

	q.Add("sixth")
	fmt.Println(q.ToString())

	fmt.Println("get:", q.Get())
	fmt.Println(q.ToString())

	fmt.Println("get:", q.Get())
	fmt.Println(q.ToString())

	fmt.Println("get:", q.Get())
	fmt.Println(q.ToString())

	fmt.Println("get:", q.Get())
	fmt.Println(q.ToString())

	fmt.Println("get:", q.Get())
	fmt.Println(q.ToString())

}
