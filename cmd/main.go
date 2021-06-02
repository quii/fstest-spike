package main

import (
	blog "github.com/quii/fstest-spike"
	"log"
	"net/http"
	"os"
)

func main() {
	blog, err := blog.New(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", blog)
}
