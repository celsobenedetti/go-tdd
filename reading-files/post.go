package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLineClearMetadata := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	readTags := func() []string {
        tags := strings.Split(readLineClearMetadata(tagsSeparator), ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		return tags
	}

	readBody := func() string {
		scanner.Scan() // discard body separator line

		buf := bytes.Buffer{}
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")

	}

	return Post{
		Title:       readLineClearMetadata(titleSeparator),
		Description: readLineClearMetadata(descriptionSeparator),
		Tags:        readTags(),
		Body:        readBody(),
	}, nil
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)
