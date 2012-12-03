package main

import (
	"tour/tree"
	)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	for {
		v1, ok := <- ch1
		if ok == false {break}
		v2, ok := <- ch2
		if ok == false {break}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	println(Same(tree.New(1), tree.New(2)))
}

