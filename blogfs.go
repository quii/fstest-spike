package fstest_spike

import (
	"io/fs"
	"os"
)

type BlogFS struct {
	rootDir fs.FS
}

func NewFS(path string) *BlogFS {
	open := os.DirFS(path)
	return &BlogFS{rootDir: open}
}

func (b *BlogFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (b *BlogFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return fs.ReadDir(b.rootDir, name)
}
