package blogposts

import (
	"reflect"
	"testing"
	"testing/fstest"
)

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
	secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
)

func TestBlogPost(t *testing.T) {

	fs := fstest.MapFS{"hello world.md": {Data: []byte(firstBody)}, "hello-world.md": {Data: []byte(secondBody)}}
	posts, err := NewPostsFromFs(fs)
	if err != nil {
		t.Fatal(err)
	}
	want := Post{title: "Post 1", description: "Description 1", tags: []string{"tdd", "go"}, body: `Hello
World`}
	got := posts[0]

	assertPost(t, got, want)

}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v posts, wanted %v posts", got, want)
	}

}
