package fstest_spike_test

import (
	"github.com/matryer/is"
	blog "github.com/quii/fstest-spike"
	"testing"
	"testing/fstest"
)

func TestBlog(t *testing.T) {
	is := is.New(t)

	const (
		firstPostBody = `Tags: go, tdd
Description: a post on TDD
---
the content
has newlines`
		secondPostBody = `Tags: football, sport
Description: a post on England's squad news
---
football!`
	)

	dirFS := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstPostBody)},
		"hello-world2.md": {Data: []byte(secondPostBody)},
	}

	posts, err := blog.New(dirFS)
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
		is.Equal(posts[0].Tags, []string{"go", "tdd"})
		is.Equal(posts[0].Description, "a post on TDD")
	})
}
