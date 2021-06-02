package main

import (
	"fmt"
	fstest_spike "github.com/quii/fstest-spike"
	"io/fs"
	"log"
	"os"
)

type CJFS struct {
	rootDir fs.FS
}

func NewCJFS(path string) *CJFS {
	open := os.DirFS(path)
	return &CJFS{rootDir: open}
}

func (C *CJFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (C *CJFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(C.rootDir, name)
}

func main() {
	blog, err := fstest_spike.NewBlog(NewCJFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(blog)
}
