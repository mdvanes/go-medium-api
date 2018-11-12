// https://tour.golang.org/basics/3
// https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
// https://medium.com/@IndianGuru/consuming-json-apis-with-go-d711efc1dcf9
// https://mholt.github.io/json-to-go/
// https://medium.com/statuscode/building-a-basic-web-service-to-display-your-medium-blog-posts-on-your-website-using-aws-api-48597b1771c5

// Install dependencies:
// go get github.com/gorilla/mux

// Building
// go build && ./api

// Path of this file
// ~/go/src/github.com/mdvanes/medium-api/api/main.go

package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
)

// our main function
func main() {
	articles = append(articles, Article{ID: "1", Title: "My Title"})

	fmt.Println("Server running, try GET on http://localhost:8000/articles")

	router := mux.NewRouter()
	router.HandleFunc("/articles", GetArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// TODO rename Articles to Publications
func GetArticles(w http.ResponseWriter, r *http.Request) {

	// authentication https://github.com/Medium/medium-api-docs#21-browser-based-authentication

	//clientId := 1;
    //url1 := fmt.Sprintf("https://medium.com/m/oauth/authorize?client_id=%s&scope=basicProfile,listPublications&state=initial&response_type=code&redirect_uri={{redirectUri}}", clientId)

	// https://github.com/Medium/medium-api-docs#32-publications
	// TODO find USERID for @mdworldNL
	url := "https://medium.com/codestar-blog/latest?format=json"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Strip security prefix
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println(bodyString[16:len(bodyString)])
	}
	// Fill the record with the data from the JSON
	//var record Numverify

	// Use json.Decode for reading streams of JSON data
	//if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
	//	log.Println(err)
	//}

	//fmt.Println("Phone No. = ", record.InternationalFormat)
	//fmt.Println("Country   = ", record.CountryName)
	//fmt.Println("Location  = ", record.Location)
	//fmt.Println("Carrier   = ", record.Carrier)
	//fmt.Println("LineType  = ", record.LineType)
	//
	//json.NewEncoder(w).Encode(articles)
}

// Auto generated with https://mholt.github.io/json-to-go/
type Publication struct {
	Data []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
		ImageURL    string `json:"imageUrl"`
	} `json:"data"`
}

// Auto generated with https://mholt.github.io/json-to-go/ from https://medium.com/codestar-blog/latest?format=json
type LatestResponse struct {
	Success bool `json:"success"`
	Payload struct {
		Collection struct {
			ID               string   `json:"id"`
			Name             string   `json:"name"`
			Slug             string   `json:"slug"`
			Tags             []string `json:"tags"`
			CreatorID        string   `json:"creatorId"`
			Description      string   `json:"description"`
			ShortDescription string   `json:"shortDescription"`
			Image            struct {
				ImageID        string `json:"imageId"`
				Filter         string `json:"filter"`
				BackgroundSize string `json:"backgroundSize"`
				OriginalWidth  int    `json:"originalWidth"`
				OriginalHeight int    `json:"originalHeight"`
				Strategy       string `json:"strategy"`
				Height         int    `json:"height"`
				Width          int    `json:"width"`
			} `json:"image"`
			Metadata struct {
				FollowerCount int   `json:"followerCount"`
				ActiveAt      int64 `json:"activeAt"`
			} `json:"metadata"`
			Virtuals struct {
				Permissions struct {
					CanPublish                      bool `json:"canPublish"`
					CanPublishAll                   bool `json:"canPublishAll"`
					CanRepublish                    bool `json:"canRepublish"`
					CanRemove                       bool `json:"canRemove"`
					CanManageAll                    bool `json:"canManageAll"`
					CanSubmit                       bool `json:"canSubmit"`
					CanEditPosts                    bool `json:"canEditPosts"`
					CanAddWriters                   bool `json:"canAddWriters"`
					CanViewStats                    bool `json:"canViewStats"`
					CanSendNewsletter               bool `json:"canSendNewsletter"`
					CanViewLockedPosts              bool `json:"canViewLockedPosts"`
					CanViewCloaked                  bool `json:"canViewCloaked"`
					CanEditOwnPosts                 bool `json:"canEditOwnPosts"`
					CanBeAssignedAuthor             bool `json:"canBeAssignedAuthor"`
					CanEnrollInHightower            bool `json:"canEnrollInHightower"`
					CanLockPostsForMediumMembers    bool `json:"canLockPostsForMediumMembers"`
					CanLockOwnPostsForMediumMembers bool `json:"canLockOwnPostsForMediumMembers"`
				} `json:"permissions"`
				IsSubscribed             bool   `json:"isSubscribed"`
				IsNewsletterSubscribed   bool   `json:"isNewsletterSubscribed"`
				MemberOfMembershipPlanID string `json:"memberOfMembershipPlanId"`
				IsEnrolledInHightower    bool   `json:"isEnrolledInHightower"`
				IsEligibleForHightower   bool   `json:"isEligibleForHightower"`
			} `json:"virtuals"`
			Logo struct {
				ImageID        string `json:"imageId"`
				Filter         string `json:"filter"`
				BackgroundSize string `json:"backgroundSize"`
				OriginalWidth  int    `json:"originalWidth"`
				OriginalHeight int    `json:"originalHeight"`
				Strategy       string `json:"strategy"`
				Height         int    `json:"height"`
				Width          int    `json:"width"`
			} `json:"logo"`
			TwitterUsername string `json:"twitterUsername"`
			Sections        []struct {
				Type                     int `json:"type"`
				CollectionHeaderMetadata struct {
					Title           string `json:"title"`
					Description     string `json:"description"`
					BackgroundImage struct {
					} `json:"backgroundImage"`
					LogoImage struct {
					} `json:"logoImage"`
					Alignment int `json:"alignment"`
					Layout    int `json:"layout"`
				} `json:"collectionHeaderMetadata,omitempty"`
				PostListMetadata struct {
					Source  int           `json:"source"`
					Layout  int           `json:"layout"`
					Number  int           `json:"number"`
					PostIds []interface{} `json:"postIds"`
				} `json:"postListMetadata,omitempty"`
			} `json:"sections"`
			Favicon struct {
				ImageID        string `json:"imageId"`
				Filter         string `json:"filter"`
				BackgroundSize string `json:"backgroundSize"`
				OriginalWidth  int    `json:"originalWidth"`
				OriginalHeight int    `json:"originalHeight"`
				Strategy       string `json:"strategy"`
				Height         int    `json:"height"`
				Width          int    `json:"width"`
			} `json:"favicon"`
			ColorPalette struct {
				DefaultBackgroundSpectrum struct {
					ColorPoints []struct {
						Color string `json:"color"`
						Point int    `json:"point"`
					} `json:"colorPoints"`
					BackgroundColor string `json:"backgroundColor"`
				} `json:"defaultBackgroundSpectrum"`
				HighlightSpectrum struct {
					ColorPoints []struct {
						Color string `json:"color"`
						Point int    `json:"point"`
					} `json:"colorPoints"`
					BackgroundColor string `json:"backgroundColor"`
				} `json:"highlightSpectrum"`
			} `json:"colorPalette"`
			NavItems                    []interface{} `json:"navItems"`
			ColorBehavior               int           `json:"colorBehavior"`
			InstantArticlesState        int           `json:"instantArticlesState"`
			AcceleratedMobilePagesState int           `json:"acceleratedMobilePagesState"`
			AmpLogo                     struct {
				ImageID        string `json:"imageId"`
				Filter         string `json:"filter"`
				BackgroundSize string `json:"backgroundSize"`
				OriginalWidth  int    `json:"originalWidth"`
				OriginalHeight int    `json:"originalHeight"`
				Strategy       string `json:"strategy"`
				Height         int    `json:"height"`
				Width          int    `json:"width"`
			} `json:"ampLogo"`
			Header struct {
				Title           string `json:"title"`
				Description     string `json:"description"`
				BackgroundImage struct {
				} `json:"backgroundImage"`
				LogoImage struct {
				} `json:"logoImage"`
				Alignment int `json:"alignment"`
				Layout    int `json:"layout"`
			} `json:"header"`
			Type string `json:"type"`
		} `json:"collection"`
		Posts []struct {
			ID                     string `json:"id"`
			VersionID              string `json:"versionId"`
			CreatorID              string `json:"creatorId"`
			HomeCollectionID       string `json:"homeCollectionId"`
			Title                  string `json:"title"`
			DetectedLanguage       string `json:"detectedLanguage"`
			LatestVersion          string `json:"latestVersion"`
			LatestPublishedVersion string `json:"latestPublishedVersion"`
			HasUnpublishedEdits    bool   `json:"hasUnpublishedEdits"`
			LatestRev              int    `json:"latestRev"`
			CreatedAt              int64  `json:"createdAt"`
			UpdatedAt              int64  `json:"updatedAt"`
			AcceptedAt             int    `json:"acceptedAt"`
			FirstPublishedAt       int64  `json:"firstPublishedAt"`
			LatestPublishedAt      int64  `json:"latestPublishedAt"`
			Vote                   bool   `json:"vote"`
			ExperimentalCSS        string `json:"experimentalCss"`
			DisplayAuthor          string `json:"displayAuthor"`
			Virtuals               struct {
				StatusForCollection string `json:"statusForCollection"`
				AllowNotes          bool   `json:"allowNotes"`
				PreviewImage        struct {
					ImageID        string `json:"imageId"`
					Filter         string `json:"filter"`
					BackgroundSize string `json:"backgroundSize"`
					OriginalWidth  int    `json:"originalWidth"`
					OriginalHeight int    `json:"originalHeight"`
					Strategy       string `json:"strategy"`
					Height         int    `json:"height"`
					Width          int    `json:"width"`
				} `json:"previewImage"`
				WordCount        int     `json:"wordCount"`
				ImageCount       int     `json:"imageCount"`
				ReadingTime      float64 `json:"readingTime"`
				Subtitle         string  `json:"subtitle"`
				UserPostRelation struct {
					UserID                             string  `json:"userId"`
					PostID                             string  `json:"postId"`
					ReadAt                             int     `json:"readAt"`
					ReadLaterAddedAt                   int     `json:"readLaterAddedAt"`
					VotedAt                            int     `json:"votedAt"`
					CollaboratorAddedAt                int     `json:"collaboratorAddedAt"`
					NotesAddedAt                       int     `json:"notesAddedAt"`
					SubscribedAt                       int     `json:"subscribedAt"`
					LastReadSectionName                string  `json:"lastReadSectionName"`
					LastReadVersionID                  string  `json:"lastReadVersionId"`
					LastReadAt                         int64   `json:"lastReadAt"`
					LastReadParagraphName              string  `json:"lastReadParagraphName"`
					LastReadPercentage                 float64 `json:"lastReadPercentage"`
					ViewedAt                           int64   `json:"viewedAt"`
					PresentedCountInResponseManagement int     `json:"presentedCountInResponseManagement"`
					ClapCount                          int     `json:"clapCount"`
					SeriesUpdateNotifsOptedInAt        int     `json:"seriesUpdateNotifsOptedInAt"`
					QueuedAt                           int     `json:"queuedAt"`
					SeriesFirstViewedAt                int     `json:"seriesFirstViewedAt"`
					PresentedCountInStream             int     `json:"presentedCountInStream"`
					SeriesLastViewedAt                 int     `json:"seriesLastViewedAt"`
					AudioProgressSec                   int     `json:"audioProgressSec"`
				} `json:"userPostRelation"`
				PublishedInCount        int           `json:"publishedInCount"`
				UsersBySocialRecommends []interface{} `json:"usersBySocialRecommends"`
				NoIndex                 bool          `json:"noIndex"`
				Recommends              int           `json:"recommends"`
				IsBookmarked            bool          `json:"isBookmarked"`
				Tags                    []struct {
					Slug      string `json:"slug"`
					Name      string `json:"name"`
					PostCount int    `json:"postCount"`
					Metadata  struct {
						PostCount  int `json:"postCount"`
						CoverImage struct {
							ID             string `json:"id"`
							OriginalWidth  int    `json:"originalWidth"`
							OriginalHeight int    `json:"originalHeight"`
							IsFeatured     bool   `json:"isFeatured"`
						} `json:"coverImage"`
					} `json:"metadata"`
					Type string `json:"type"`
				} `json:"tags"`
				SocialRecommendsCount int `json:"socialRecommendsCount"`
				ResponsesCreatedCount int `json:"responsesCreatedCount"`
				Links                 struct {
					Entries []struct {
						URL  string `json:"url"`
						Alts []struct {
							Type int    `json:"type"`
							URL  string `json:"url"`
						} `json:"alts"`
						HTTPStatus int `json:"httpStatus"`
					} `json:"entries"`
					Version     string `json:"version"`
					GeneratedAt int64  `json:"generatedAt"`
				} `json:"links"`
				IsLockedPreviewOnly bool          `json:"isLockedPreviewOnly"`
				TakeoverID          string        `json:"takeoverId"`
				MetaDescription     string        `json:"metaDescription"`
				TotalClapCount      int           `json:"totalClapCount"`
				SectionCount        int           `json:"sectionCount"`
				ReadingList         int           `json:"readingList"`
				Topics              []interface{} `json:"topics"`
			} `json:"virtuals"`
			Coverless                  bool   `json:"coverless"`
			Slug                       string `json:"slug"`
			TranslationSourcePostID    string `json:"translationSourcePostId"`
			TranslationSourceCreatorID string `json:"translationSourceCreatorId"`
			IsApprovedTranslation      bool   `json:"isApprovedTranslation"`
			InResponseToPostID         string `json:"inResponseToPostId"`
			InResponseToRemovedAt      int    `json:"inResponseToRemovedAt"`
			IsTitleSynthesized         bool   `json:"isTitleSynthesized"`
			AllowResponses             bool   `json:"allowResponses"`
			ImportedURL                string `json:"importedUrl"`
			ImportedPublishedAt        int    `json:"importedPublishedAt"`
			Visibility                 int    `json:"visibility"`
			UniqueSlug                 string `json:"uniqueSlug"`
			PreviewContent             struct {
				BodyModel struct {
					Paragraphs []struct {
						Name      string        `json:"name"`
						Type      int           `json:"type"`
						Text      string        `json:"text"`
						Markups   []interface{} `json:"markups"`
						Alignment int           `json:"alignment"`
					} `json:"paragraphs"`
					Sections []struct {
						StartIndex int `json:"startIndex"`
					} `json:"sections"`
				} `json:"bodyModel"`
				IsFullContent bool   `json:"isFullContent"`
				Subtitle      string `json:"subtitle"`
			} `json:"previewContent"`
			License                           int    `json:"license"`
			InResponseToMediaResourceID       string `json:"inResponseToMediaResourceId"`
			CanonicalURL                      string `json:"canonicalUrl"`
			ApprovedHomeCollectionID          string `json:"approvedHomeCollectionId"`
			NewsletterID                      string `json:"newsletterId"`
			WebCanonicalURL                   string `json:"webCanonicalUrl"`
			MediumURL                         string `json:"mediumUrl"`
			MigrationID                       string `json:"migrationId"`
			NotifyFollowers                   bool   `json:"notifyFollowers"`
			NotifyTwitter                     bool   `json:"notifyTwitter"`
			NotifyFacebook                    bool   `json:"notifyFacebook"`
			ResponseHiddenOnParentPostAt      int    `json:"responseHiddenOnParentPostAt"`
			IsSeries                          bool   `json:"isSeries"`
			IsSubscriptionLocked              bool   `json:"isSubscriptionLocked"`
			SeriesLastAppendedAt              int    `json:"seriesLastAppendedAt"`
			AudioVersionDurationSec           int    `json:"audioVersionDurationSec"`
			SequenceID                        string `json:"sequenceId"`
			IsNsfw                            bool   `json:"isNsfw"`
			IsEligibleForRevenue              bool   `json:"isEligibleForRevenue"`
			IsBlockedFromHightower            bool   `json:"isBlockedFromHightower"`
			DeletedAt                         int    `json:"deletedAt"`
			LockedPostSource                  int    `json:"lockedPostSource"`
			HightowerMinimumGuaranteeStartsAt int    `json:"hightowerMinimumGuaranteeStartsAt"`
			HightowerMinimumGuaranteeEndsAt   int    `json:"hightowerMinimumGuaranteeEndsAt"`
			FeatureLockRequestAcceptedAt      int    `json:"featureLockRequestAcceptedAt"`
			MongerRequestType                 int    `json:"mongerRequestType"`
			LayerCake                         int    `json:"layerCake"`
			SocialTitle                       string `json:"socialTitle"`
			SocialDek                         string `json:"socialDek"`
			EditorialPreviewTitle             string `json:"editorialPreviewTitle"`
			EditorialPreviewDek               string `json:"editorialPreviewDek"`
			CurationEligibleAt                int    `json:"curationEligibleAt"`
			Type                              string `json:"type"`
		} `json:"posts"`
		TagName     interface{} `json:"tagName"`
		WriterNames []string    `json:"writerNames"`
		References  struct {
			User struct {
				FiveEd06Eb2199C struct {
					UserID            string `json:"userId"`
					Name              string `json:"name"`
					Username          string `json:"username"`
					CreatedAt         int64  `json:"createdAt"`
					ImageID           string `json:"imageId"`
					BackgroundImageID string `json:"backgroundImageId"`
					Bio               string `json:"bio"`
					TwitterScreenName string `json:"twitterScreenName"`
					FacebookAccountID string `json:"facebookAccountId"`
					AllowNotes        int    `json:"allowNotes"`
					MediumMemberAt    int    `json:"mediumMemberAt"`
					IsNsfw            bool   `json:"isNsfw"`
					Type              string `json:"type"`
				} `json:"5ed06eb2199c"`
				Three1A789732F21 struct {
					UserID            string `json:"userId"`
					Name              string `json:"name"`
					Username          string `json:"username"`
					CreatedAt         int64  `json:"createdAt"`
					ImageID           string `json:"imageId"`
					BackgroundImageID string `json:"backgroundImageId"`
					Bio               string `json:"bio"`
					TwitterScreenName string `json:"twitterScreenName"`
					FacebookAccountID string `json:"facebookAccountId"`
					AllowNotes        int    `json:"allowNotes"`
					MediumMemberAt    int    `json:"mediumMemberAt"`
					IsNsfw            bool   `json:"isNsfw"`
					Type              string `json:"type"`
				} `json:"31a789732f21"`
				NineEb88B7110D9 struct {
					UserID            string `json:"userId"`
					Name              string `json:"name"`
					Username          string `json:"username"`
					CreatedAt         int64  `json:"createdAt"`
					ImageID           string `json:"imageId"`
					BackgroundImageID string `json:"backgroundImageId"`
					Bio               string `json:"bio"`
					TwitterScreenName string `json:"twitterScreenName"`
					FacebookAccountID string `json:"facebookAccountId"`
					AllowNotes        int    `json:"allowNotes"`
					MediumMemberAt    int    `json:"mediumMemberAt"`
					IsNsfw            bool   `json:"isNsfw"`
					Type              string `json:"type"`
				} `json:"9eb88b7110d9"`
				B8149197A62 struct {
					UserID            string `json:"userId"`
					Name              string `json:"name"`
					Username          string `json:"username"`
					CreatedAt         int64  `json:"createdAt"`
					ImageID           string `json:"imageId"`
					BackgroundImageID string `json:"backgroundImageId"`
					Bio               string `json:"bio"`
					TwitterScreenName string `json:"twitterScreenName"`
					FacebookAccountID string `json:"facebookAccountId"`
					AllowNotes        int    `json:"allowNotes"`
					MediumMemberAt    int    `json:"mediumMemberAt"`
					IsNsfw            bool   `json:"isNsfw"`
					Type              string `json:"type"`
				} `json:"b8149197a62"`
			} `json:"User"`
			Collection struct {
				Ea5Cff13E3C9 struct {
					ID               string   `json:"id"`
					Name             string   `json:"name"`
					Slug             string   `json:"slug"`
					Tags             []string `json:"tags"`
					CreatorID        string   `json:"creatorId"`
					Description      string   `json:"description"`
					ShortDescription string   `json:"shortDescription"`
					Image            struct {
						ImageID        string `json:"imageId"`
						Filter         string `json:"filter"`
						BackgroundSize string `json:"backgroundSize"`
						OriginalWidth  int    `json:"originalWidth"`
						OriginalHeight int    `json:"originalHeight"`
						Strategy       string `json:"strategy"`
						Height         int    `json:"height"`
						Width          int    `json:"width"`
					} `json:"image"`
					Metadata struct {
						FollowerCount int   `json:"followerCount"`
						ActiveAt      int64 `json:"activeAt"`
					} `json:"metadata"`
					Virtuals struct {
						Permissions struct {
							CanPublish                      bool `json:"canPublish"`
							CanPublishAll                   bool `json:"canPublishAll"`
							CanRepublish                    bool `json:"canRepublish"`
							CanRemove                       bool `json:"canRemove"`
							CanManageAll                    bool `json:"canManageAll"`
							CanSubmit                       bool `json:"canSubmit"`
							CanEditPosts                    bool `json:"canEditPosts"`
							CanAddWriters                   bool `json:"canAddWriters"`
							CanViewStats                    bool `json:"canViewStats"`
							CanSendNewsletter               bool `json:"canSendNewsletter"`
							CanViewLockedPosts              bool `json:"canViewLockedPosts"`
							CanViewCloaked                  bool `json:"canViewCloaked"`
							CanEditOwnPosts                 bool `json:"canEditOwnPosts"`
							CanBeAssignedAuthor             bool `json:"canBeAssignedAuthor"`
							CanEnrollInHightower            bool `json:"canEnrollInHightower"`
							CanLockPostsForMediumMembers    bool `json:"canLockPostsForMediumMembers"`
							CanLockOwnPostsForMediumMembers bool `json:"canLockOwnPostsForMediumMembers"`
						} `json:"permissions"`
						IsSubscribed             bool   `json:"isSubscribed"`
						IsNewsletterSubscribed   bool   `json:"isNewsletterSubscribed"`
						MemberOfMembershipPlanID string `json:"memberOfMembershipPlanId"`
						IsEnrolledInHightower    bool   `json:"isEnrolledInHightower"`
						IsEligibleForHightower   bool   `json:"isEligibleForHightower"`
					} `json:"virtuals"`
					Logo struct {
						ImageID        string `json:"imageId"`
						Filter         string `json:"filter"`
						BackgroundSize string `json:"backgroundSize"`
						OriginalWidth  int    `json:"originalWidth"`
						OriginalHeight int    `json:"originalHeight"`
						Strategy       string `json:"strategy"`
						Height         int    `json:"height"`
						Width          int    `json:"width"`
					} `json:"logo"`
					TwitterUsername string `json:"twitterUsername"`
					Sections        []struct {
						Type                     int `json:"type"`
						CollectionHeaderMetadata struct {
							Title           string `json:"title"`
							Description     string `json:"description"`
							BackgroundImage struct {
							} `json:"backgroundImage"`
							LogoImage struct {
							} `json:"logoImage"`
							Alignment int `json:"alignment"`
							Layout    int `json:"layout"`
						} `json:"collectionHeaderMetadata,omitempty"`
						PostListMetadata struct {
							Source  int           `json:"source"`
							Layout  int           `json:"layout"`
							Number  int           `json:"number"`
							PostIds []interface{} `json:"postIds"`
						} `json:"postListMetadata,omitempty"`
					} `json:"sections"`
					Favicon struct {
						ImageID        string `json:"imageId"`
						Filter         string `json:"filter"`
						BackgroundSize string `json:"backgroundSize"`
						OriginalWidth  int    `json:"originalWidth"`
						OriginalHeight int    `json:"originalHeight"`
						Strategy       string `json:"strategy"`
						Height         int    `json:"height"`
						Width          int    `json:"width"`
					} `json:"favicon"`
					ColorPalette struct {
						DefaultBackgroundSpectrum struct {
							ColorPoints []struct {
								Color string `json:"color"`
								Point int    `json:"point"`
							} `json:"colorPoints"`
							BackgroundColor string `json:"backgroundColor"`
						} `json:"defaultBackgroundSpectrum"`
						HighlightSpectrum struct {
							ColorPoints []struct {
								Color string `json:"color"`
								Point int    `json:"point"`
							} `json:"colorPoints"`
							BackgroundColor string `json:"backgroundColor"`
						} `json:"highlightSpectrum"`
					} `json:"colorPalette"`
					NavItems                    []interface{} `json:"navItems"`
					ColorBehavior               int           `json:"colorBehavior"`
					InstantArticlesState        int           `json:"instantArticlesState"`
					AcceleratedMobilePagesState int           `json:"acceleratedMobilePagesState"`
					AmpLogo                     struct {
						ImageID        string `json:"imageId"`
						Filter         string `json:"filter"`
						BackgroundSize string `json:"backgroundSize"`
						OriginalWidth  int    `json:"originalWidth"`
						OriginalHeight int    `json:"originalHeight"`
						Strategy       string `json:"strategy"`
						Height         int    `json:"height"`
						Width          int    `json:"width"`
					} `json:"ampLogo"`
					Header struct {
						Title           string `json:"title"`
						Description     string `json:"description"`
						BackgroundImage struct {
						} `json:"backgroundImage"`
						LogoImage struct {
						} `json:"logoImage"`
						Alignment int `json:"alignment"`
						Layout    int `json:"layout"`
					} `json:"header"`
					Type string `json:"type"`
				} `json:"ea5cff13e3c9"`
			} `json:"Collection"`
		} `json:"references"`
		Paging struct {
			Previous struct {
				Limit int    `json:"limit"`
				From  string `json:"from"`
			} `json:"previous"`
			Next struct {
				Limit int    `json:"limit"`
				To    string `json:"to"`
			} `json:"next"`
		} `json:"paging"`
	} `json:"payload"`
	V int    `json:"v"`
	B string `json:"b"`
}

type Article struct {
	ID        string   `json:"id,omitempty"`
	Title     string   `json:"title,omitempty"`
}

var articles []Article
