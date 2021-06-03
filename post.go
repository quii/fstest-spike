package fstest_spike

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Name string
	Body string
}

func newPost(fileName string, body io.Reader) Post {
	scanner := bufio.NewScanner(body)

	// ignore the first few lines, we'll parse these into meta data later
	scanner.Scan()
	scanner.Scan()

	b := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&b, scanner.Text())
	}

	return Post{
		Name: removeExtension(fileName),
		Body: strings.TrimSuffix(b.String(), "\n"),
	}
}

func removeExtension(fileName string) string {
	return fileName[:len(fileName)-3]
}
