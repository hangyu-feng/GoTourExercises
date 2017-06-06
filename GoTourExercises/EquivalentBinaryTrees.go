package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	t := Tree{Value: 3}
	t.insert(2)
	t.insert(5)
	t.insert(1)
	ch := make(chan int)
	go Walk(&t, ch)
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	fmt.Println(t.Value)
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	Walk(t1, c1)
	Walk(t2, c2)
	for i := range c1 {
		j := <-c2
		if i != j {
			return false
		}
	}
	return true
}

// New(n int) 构造一个随机结构的有n项的二叉树
func New(n int) Tree {
	rand.Int()
	return Tree{}
}

// Tree.insert method 将一个int插入一个binary search tree
func (t *Tree) insert(n int) {
	if t == nil {
		t = &Tree{Value: 3}
	} else if n <= t.Value {
		t.Left.insert(n)
	} else {
		t.Right.insert(n)
	}
}
