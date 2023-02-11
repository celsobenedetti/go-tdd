package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/celso-patiri/go-tdd/reading-files"
)

func TestNewBlogPosts(t *testing.T) {

	postsBodys := []string{
		`Title: Post1
Description: Description 1
Tags: go, tdd
---
Hello
World!`,

		`Title: Post2
Description: Description 2
rust, borrow-checker
---
Ola
Mundo!`,
	}

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(postsBodys[0])},
		"hello-world2.md": {Data: []byte(postsBodys[1])},
	}

	posts, err := blogposts.NewBlogPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post1",
		Description: "Description 1",
		Tags:        []string{"go", "tdd"},
		Body: `Hello
World!`,
	})
}

func assertPost(t testing.TB, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

type StubFailingFS struct{}

func (s *StubFailingFS) Open() (fs.File, error) {
	return nil, errors.New("I have failed my family")
}
