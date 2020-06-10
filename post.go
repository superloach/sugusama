package sugusama

type Post struct {
	Typename             string                 `json:"__typename"`
	AccessibilityCaption string                 `json:"accessibility_caption"`
	Attribution          interface{}            `json:"attribution"`
	CommentsDisabled     bool                   `json:"comments_disabled"`
	DashInfo             additionalDataDashInfo `json:"dash_info"`

	Dimensions struct {
		Height int64 `json:"height"`
		Width  int64 `json:"width"`
	} `json:"dimensions"`

	DisplayResources []DisplayResource `json:"display_resources"`
	DisplayURL       string            `json:"display_url"`

	RawComments struct {
		Count int64 `json:"count"`

		Edges []struct {
			Node *Comment `json:"node"`
		} `json:"edges"`

		PageInfo PageInfo `json:"page_info"`
	} `json:"edge_media_preview_comment"`

	RawLikes struct {
		Count int64 `json:"count"`

		Edges []struct {
			Node *User `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_preview_like"`

	RawCaptions struct {
		Edges []struct {
			Node struct {
				Text string `json:"text"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_caption"`

	RawSponsorUsers struct {
		Edges []interface{} `json:"edges"`
	} `json:"edge_media_to_sponsor_user"`

	RawTaggedUsers struct {
		Edges []struct {
			Node struct {
				User *User   `json:"user"`
				X    float64 `json:"x"`
				Y    float64 `json:"y"`
			} `json:"node"`
		} `json:"edges"`
	} `json:"edge_media_to_tagged_user"`

	RawChildren struct {
		Edges []struct {
			Node *Post `json:"node"`
		} `json:"edges"`
	} `json:"edge_sidecar_to_children"`

	FactCheckInformation       interface{} `json:"fact_check_information"`
	FactCheckOverallRating     interface{} `json:"fact_check_overall_rating"`
	FollowHashtagInfo          interface{} `json:"follow_hashtag_info"`
	GatingInfo                 interface{} `json:"gating_info"`
	ID                         string      `json:"id"`
	IsVideo                    bool        `json:"is_video"`
	Location                   interface{} `json:"location"`
	MediaPreview               string      `json:"media_preview"`
	Owner                      *User       `json:"owner"`
	Shortcode                  string      `json:"shortcode"`
	TakenAtTimestamp           int64       `json:"taken_at_timestamp"`
	TrackingToken              string      `json:"tracking_token"`
	VideoURL                   string      `json:"video_url"`
	VideoViewCount             int64       `json:"video_view_count"`
	ViewerCanReshare           bool        `json:"viewer_can_reshare"`
	ViewerHasLiked             bool        `json:"viewer_has_liked"`
	ViewerHasSaved             bool        `json:"viewer_has_saved"`
	ViewerHasSavedToCollection bool        `json:"viewer_has_saved_to_collection"`
	ViewerInPhotoOfYou         bool        `json:"viewer_in_photo_of_you"`
}
