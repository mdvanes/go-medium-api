package main

// Extracted to separate file to test having multiple files
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
		simplePost.AuthorImg = users[item.CreatorId].ImageID
		simplePost.LatestPublishedAt = item.LatestPublishedAt
		simplePost.UniqueSlug = item.UniqueSlug
		simplePost.Paragraphs = item.PreviewContent.BodyModel.Paragraphs
		simplePosts = append(simplePosts, simplePost)
	}
	return simplePosts
}
