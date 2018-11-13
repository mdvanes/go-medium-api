package main

import "testing"

func TestConvertPosts(t *testing.T) {
	var dummyPosts []Post

	var post Post
	post.ID = "1"
	post.Title = "My Title"
	post.CreatorId = "123"
	post.LatestPublishedAt = 123456

	dummyPosts = append(dummyPosts, post)

	var dummyUsers = make(map[string]User)

	var user User
	user.Name = "An Author"

	dummyUsers["123"] = user

	convertedPosts := ConvertPosts(dummyPosts, dummyUsers)
	if len(convertedPosts) != 1 {
		t.Errorf("ConvertPosts was incorrect, got: %d, want: %d.", len(convertedPosts), 1)
	}
}
