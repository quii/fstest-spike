package main

import (
	"fmt"
	blog "github.com/quii/fstest-spike"
	"log"
)

func main() {
	blog, err := blog.New(blog.NewFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blog)
}
