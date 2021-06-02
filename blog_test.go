package fstest_spike

import (
	"testing"
	"testing/fstest"
)

func TestBlog(t *testing.T) {
	t.Run("it returns a list of blogs from a directory", func(t *testing.T) {

		dirFS := make(fstest.MapFS)
		dirFS["hello-world.md"] = &fstest.MapFile{Data: []byte("hello, world.md")}
		dirFS["hello-world2.md"] = &fstest.MapFile{Data: []byte("hello, world again.md")}

		posts, err := NewBlog(dirFS)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != 2 {
			t.Errorf("expected 2 posts but got %d", len(posts))
		}

		if posts[0].Name != "hello-world.md" {
			t.Errorf("got %q, want %q", posts[0].Name, "hello-world.md")
		}
	})
}

