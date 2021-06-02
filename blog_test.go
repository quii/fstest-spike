package fstest_spike

import (
	"github.com/matryer/is"
	"testing"
	"testing/fstest"
)

func TestBlog(t *testing.T) {
	is := is.New(t)

	dirFS := fstest.MapFS{
		"hello-world.md":  {Data: []byte("hello, world")},
		"hello-world.md2": {Data: []byte("hello, world again")},
	}

	posts, err := New(dirFS)
	is.NoErr(err)

	t.Run("it returns a post for each (valid) file", func(t *testing.T) {
		is.Equal(len(posts), len(dirFS))
	})

	t.Run("it parses the first post correctly", func(t *testing.T) {
		is.Equal(posts[0].Name, "hello-world.md")
		is.Equal(posts[0].Body, "hello, world")
	})
}
