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

	// ignore ----- separator
	scanner.Scan()

	return Post{
		Name: removeExtension(fileName),
		Body: readBody(scanner),
		Tags: extractTags(tagsData),
	}
}

func readBody(scanner *bufio.Scanner) string {
	b := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&b, scanner.Text())
	}
	return strings.TrimSuffix(b.String(), "\n")
}

func extractTags(data string) (tags []string) {
	withoutTag := strings.TrimPrefix(data, "Tags: ")
	return strings.Split(withoutTag, ", ")
}

func removeExtension(fileName string) string {
	return fileName[:len(fileName)-3]
}
