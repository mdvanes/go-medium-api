package main

// TODO extract to separate file to test having multiple files
func ConvertPosts(posts []Post, users map[string]User) []SimplePost {
	// Convert Post to SimplePost
	var simplePosts []SimplePost

	// TODO use mapping
	// range posts returns index, item
	for _, item := range posts {
		var simplePost SimplePost
		simplePost.ID = item.ID
		simplePost.Title = item.Title
		simplePost.Author = users[item.CreatorId].Name
		simplePost.LatestPublishedAt = item.LatestPublishedAt
		simplePost.Paragraphs = item.PreviewContent.BodyModel.Paragraphs
		simplePosts = append(simplePosts, simplePost)
	}
	return simplePosts
}
