package main

import (
	"github.com/bradleyjkemp/memviz"
	"github.com/calini/std/structs/collection/list"
	"log"
	"os"
)

func main() {

	lst := new(list.LinkedList)

	out, err := os.Create("level0.dot")
	if err != nil {
		log.Fatal(err)
	}
	memviz.Map(out, lst)
	out.Close()

	// Add Elem
	lst.Add(1)
	out, err = os.Create("level1.dot")
	if err != nil {
		log.Fatal(err)
	}
	memviz.Map(out, lst)
	out.Close()

	// Add Elem
	lst.Add(2)
	out, err = os.Create("level2.dot")
	if err != nil {
		log.Fatal(err)
	}
	memviz.Map(out, lst)
	out.Close()

	// Add Elem
	lst.Add(3)
	out, err = os.Create("level3.dot")
	if err != nil {
		log.Fatal(err)
	}
	memviz.Map(out, lst)
	out.Close()

	err = lst.AddAtIndex(3, 69)
	if err != nil {
		log.Fatal(err)
	}
	out, err = os.Create("level4.dot")
	if err != nil {
		log.Fatal(err)
	}
	memviz.Map(out, lst)
	out.Close()

	lst.AddAll(6,7,6,7,6)
	out, err = os.Create("level5.dot")
	if err != nil {
		log.Fatal(err)
	}
	memviz.Map(out, lst)
	out.Close()

}


// Vector, guaranteed that at most one element has 2k+1 occurences
// Tell me which one it is, O(n) time, O(1) mem
// -> Xor

// swap without a glass
// a, b
//
//a = a ^ b
//b = a ^ b
//a = a ^ b


// what do you see on the screen at the end
int main() {
	int i = 0;
	cout << (i++) << (++i) << (i++);
	return 0;
}

$ 0 2 2; -> Value of i at the end = 3;
