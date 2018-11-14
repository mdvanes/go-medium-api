package main

type SimplePost struct {
	ID                     string `json:"id"`
	Title                  string `json:"title"`
	Author                 string `json:"author"`
	AuthorImg			   string `json:"authorImg"`
	LatestPublishedAt      int    `json:"latestPublishedAt"`
	UniqueSlug             string `json:"uniqueSlug"`
	Paragraphs []struct {
		Text               string `json:"text"`
	} `json:"paragraphs"`
}

type Post struct {
	ID                     string `json:"id"`
	Title                  string `json:"title"`
	CreatorId              string `json:"creatorId"`
	LatestPublishedAt      int    `json:"latestPublishedAt"`
	UniqueSlug     		   string `json:"uniqueSlug"`
	PreviewContent struct {
		BodyModel struct {
			Paragraphs []struct {
				Text       string `json:"text"`
			} `json:"paragraphs"`
		} `json:"bodyModel"`
	} `json:"previewContent"`
}

type User struct {
	UserID            string `json:"userId"`              // userId (equal to Post.creatorId)
	Name              string `json:"name"`                // name (author name)
	Username          string `json:"username"`
	ImageID           string `json:"imageId"`
	BackgroundImageID string `json:"backgroundImageId"`
	Bio               string `json:"bio"`
	TwitterScreenName string `json:"twitterScreenName"`
}

// Auto generated with https://mholt.github.io/json-to-go/ from https://medium.com/codestar-blog/latest?format=json
type LatestResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Posts []Post `json:"posts,omitempty"`
		References struct {
			User map[string]User
		}
	} `json:"payload"`
}
