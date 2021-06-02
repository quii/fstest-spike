package fstest_spike

import (
	"io/fs"
)

func NewBlog(blogFolder fs.ReadDirFS) ([]Post, error) {
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
