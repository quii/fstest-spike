package fstest_spike

import (
	"fmt"
	"net/http"
	"strings"
)

type Posts []Post

func (p Posts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	postName := strings.TrimPrefix(r.URL.Path, "/")
	post, err := p.find(postName)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	fmt.Fprint(w, post.Body)
}

func (p Posts) find(name string) (Post, error) {
	for _, post := range p {
		if post.Name == name {
			return post, nil
		}
	}
	return Post{}, fmt.Errorf("could not find %q in %q", name, p.postNames())
}

func (p Posts) postNames() string {
	var postNames []string
	for _, post := range p {
		postNames = append(postNames, post.Name)
	}
	return strings.Join(postNames, ",")
}
