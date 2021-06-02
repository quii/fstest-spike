package fstest_spike

import (
	"io/fs"
)

func New(blogFolder fs.ReadDirFS) ([]Post, error) {
	dir, err := blogFolder.ReadDir(".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, file := range dir {
		posts = append(posts, Post{file.Name()})
	}
	return posts, nil
}
