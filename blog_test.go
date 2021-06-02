package fstest_spike

import (
	"testing"
	"testing/fstest"
)

func TestBlog(t *testing.T) {
	t.Run("it returns a list of blogs from a directory", func(t *testing.T) {

		dirFS := fstest.MapFS{
			"hello-world.md": {Data: []byte("hello, world")},
			"hello-world.md2": {Data: []byte("hello, world again")},
		}

		posts, err := New(dirFS)

		if err != nil {
			t.Fatal(err)
		}

		t.Run("it gets all the files from the directory", func(t *testing.T) {
			got := len(posts)
			want := len(dirFS)
			if got != want {
				t.Errorf("expected %d posts but got %d", want, got)
			}
		})

		t.Run("it parses the first post correctly", func(t *testing.T) {
			if posts[0].Name != "hello-world.md" {
				t.Errorf("got %q, want %q", posts[0].Name, "hello-world.md")
			}

			if posts[0].Body != "hello, world" {
				t.Errorf("got %q, want %q", posts[0].Body, "hello, world")
			}
		})

	})
}
