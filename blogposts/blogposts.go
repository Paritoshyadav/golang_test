package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	title string
}

func NewPostsFromFs(filesystem fs.FS) ([]Post, error) {

	dir, _ := fs.ReadDir(filesystem, ".")
	var posts []Post
	for _, f := range dir {

		post, _ := getPost(filesystem, f.Name())

		posts = append(posts, post)

	}

	return posts, nil

}

func getPost(filesystem fs.FS, f string) (Post, error) {

	file, err := filesystem.Open(f)
	if err != nil {
		return Post{}, err

	}
	defer file.Close()

	return newPost(file)

}

func newPost(file io.Reader) (Post, error) {
	postData, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err

	}
	post := Post{title: string(postData)[7:]}

	return post, nil

}
