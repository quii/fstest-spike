package fstest_spike

import (
	"io/fs"
)

func New(blogFolder fs.FS) (Posts, error) {
	dir, err := fs.ReadDir(blogFolder, ".")
	if err != nil {
		return nil, err
	}
	return getPosts(blogFolder, dir)
}

func getPosts(blogFolder fs.FS, dir []fs.DirEntry) ([]Post, error) {
	var posts []Post
	for _, file := range dir {
		post, err := newPostFromFile(blogFolder, file.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func newPostFromFile(blogFolder fs.FS, fileName string) (Post, error) {
	f, err := blogFolder.Open(fileName)
	defer f.Close()

	if err != nil {
		return Post{}, err
	}

	post := newPost(fileName, f)
	return post, nil
}
