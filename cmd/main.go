package main

import (
	fstest_spike "github.com/quii/fstest-spike"
	"os"
)

func main() {
	postsFolder := os.DirFS("posts")
	fstest_spike.NewBlog(postsFolder)
}
