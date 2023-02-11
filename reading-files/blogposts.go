package blogposts

import (
	"io/fs"
)

func NewBlogPostsFromFS(fileSystem fs.FS) (posts []Post, err error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	for _, f := range dir {
		post, err := readPostFile(fileSystem, f.Name())

		if err != nil {
			return nil, err //TODO: needs clarification. Should we totally fail if one file has error?
		}

		posts = append(posts, post)
	}
	return posts, nil
}

func readPostFile(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, nil
	}
	defer postFile.Close()

	return newPost(postFile)
}
