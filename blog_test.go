package fstest_spike

import (
	"github.com/matryer/is"
	"testing"
	"testing/fstest"
)

func TestBlog(t *testing.T) {
	is := is.New(t)

	const (
		firstPostBody = `Tags: go, tdd
---
the content
has newlines`
		secondPostBody = `Tags: football, sport
---
football!`
	)

	dirFS := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstPostBody)},
		"hello-world2.md": {Data: []byte(secondPostBody)},
	}

	posts, err := New(dirFS)
	is.NoErr(err)

	t.Run("it returns a post for each (valid) file", func(t *testing.T) {
		is.Equal(len(posts), len(dirFS))
	})

	t.Run("it parses the first post correctly", func(t *testing.T) {
		t.Run("it removes the extension from the title", func(t *testing.T) {
			is.Equal(posts[0].Name, "hello world")
		})
		expectedContent := `the content
has newlines`
		is.Equal(posts[0].Body, expectedContent)
	})
}
