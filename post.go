package fstest_spike

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Name        string
	Body        string
	Tags        []string
	Description string
}

func newPost(fileName string, body io.Reader) Post {
	scanner := bufio.NewScanner(body)

	scanAndRead := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	tags := extractTags(scanAndRead())
	description := extractDescription(scanAndRead())
	scanner.Scan() // ignore ----- separator
	postBody := readBody(scanner)

	return Post{
		Name:        removeExtension(fileName),
		Body:        postBody,
		Tags:        tags,
		Description: description,
	}
}

func extractDescription(data string) string {
	return strings.TrimPrefix(data, "Description: ")
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
