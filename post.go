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
	Tags []string
}

func newPost(fileName string, body io.Reader) Post {
	scanner := bufio.NewScanner(body)

	scanner.Scan()
	tagsData := scanner.Text()

	// ignore -----
	scanner.Scan()

	b := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&b, scanner.Text())
	}

	return Post{
		Name: removeExtension(fileName),
		Body: strings.TrimSuffix(b.String(), "\n"),
		Tags: extractTags(tagsData),
	}
}

func extractTags(data string) (tags []string) {
	withoutTag := strings.TrimPrefix(data, "Tags: ")
	return strings.Split(withoutTag, ", ")
}

func removeExtension(fileName string) string {
	return fileName[:len(fileName)-3]
}
