package main

import (
	"fmt"
	blog "github.com/quii/fstest-spike"
	"log"
	"os"
)

func main() {
	blog, err := blog.New(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blog)
}
