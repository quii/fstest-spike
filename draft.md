# Learn Go with tests - Working with files

We've been asked to create a package that needs to convert a given folder of blog posts and return a collection of `Post` which represents each file parsed with information about its contents

### Example data

hello world.md
```
Title: Hello, TDD world!
Description: First post on our wonderful blog
Tags: tdd, go
---

# Hello world!

The body of posts starts after the `---` and is in markdown
```

### Expected data

```go
type Post struct {
	Title, Description, Body string
	Tags []string
}
```

## Iterative, test-driven development

As always we'll take an iterative approach where we're always taking simple, safe steps toward our goal. 

This requires us to break up our work into iterative steps, but we should be careful not to fall in to the trap of taking a "bottom up" approach. That might involve us say making some kind of abstraction that is only validated once we stick everything together. This is _not_ iterative! Iterative means we work in "thin" vertical slices, keeping scope small but useful and validated.

# Notes
- There are functions like os.Open that we can use out of the box, but reaching in to the global file system makes testing a little harder. What would be preferable is if there is some kind of abstraction around a file system we can use so that we can inject it like any other dependency. That'll make it simpler to test and simpler to write. 