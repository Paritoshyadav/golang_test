package blogposts

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestBlogPost(t *testing.T) {

	fs := fstest.MapFS{"hello world.md": {Data: []byte("Title: Post 1")}, "hello-world.md": {Data: []byte("Title: Post 1")}}
	posts, err := NewPostsFromFs(fs)
	if err != nil {
		t.Fatal(err)
	}
	want := Post{title: "Post 1"}
	got := posts[0]

	assertPost(t, got, want)

}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v posts, wanted %v posts", got, want)
	}

}
