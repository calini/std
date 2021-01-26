package main

import (
	"github.com/bradleyjkemp/memviz"
	"log"
	"os"
)


type Node struct {
	next *Node
	content int
}

const filename = "preview.dot"

func main() {

	ptr := &Node{
		next:    &Node{
			next:    &Node{
				content: 4,
			},
			content: 2,
		},
		content: 0,
	}

	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	memviz.Map(out, ptr)
}
