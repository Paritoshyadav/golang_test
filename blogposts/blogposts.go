package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeparator         = "Tags: "
)

type Post struct {
	title       string
	description string
	tags        []string
	body        string
}

func NewPostsFromFs(filesystem fs.FS) ([]Post, error) { //os.DirFS("posts")

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
	scanner := bufio.NewScanner(file)

	readMetaLine := func(tag string) string {
		scanner.Scan()

		return strings.TrimPrefix(scanner.Text(), tag)

	}

	post := Post{title: readMetaLine(titleSeparator), description: readMetaLine(descriptionSeparator), tags: strings.Split(readMetaLine(tagSeparator), ", "), body: getBody(scanner)}

	return post, nil

}

func getBody(scanner *bufio.Scanner) string {
	scanner.Scan() //igonre a line
	data := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&data, scanner.Text())
	}
	return strings.TrimSuffix(data.String(), "\n")
}
