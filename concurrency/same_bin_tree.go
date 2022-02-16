 package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	RWalk(t, ch)
	close(ch)
}

func RWalk(t *tree.Tree, ch chan int){
	if t != nil {
        RWalk(t.Left, ch)
        ch <- t.Value
        RWalk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values
// it works only for sorted trees.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		a1, ok1 := <-ch1
		a2, ok2 := <-ch2
		
		if !ok1 && !ok2 {
			break
		}
		//если разная длина или значения не совпадают
		if ok1 != ok2 || a1 != a2 {
			return false
		}
		
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
    fmt.Println(Same(tree.New(1), tree.New(1)))
    fmt.Println(Same(tree.New(2), tree.New(1)))
}
