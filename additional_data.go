package sugusama

type AdditionalData struct {
	User AdditionalDataUser `json:"user"`
}

func (c *Client) ExtractAdditionalData(body []byte) (*AdditionalData, error) {
	return nil, nil
}

type AdditionalDataUser struct {
	Feed          AdditionalDataFeed `json:"edge_web_feed_timeline,omitempty"`
	Blocked       bool               `json:"blocked_by_viewer"`
	Followed      bool               `json:"followed_by_viewer"`
	FullName      string             `json:"full_name"`
	HasBlocked    bool               `json:"has_blocked_viewer"`
	ID            string             `json:"id"`
	Private       bool               `json:"is_private"`
	Verified      bool               `json:"is_verified"`
	ProfilePicURL string             `json:"profile_pic_url"`
	Requested     bool               `json:"requested_by_viewer"`
	Restricted    bool               `json:"restricted_by_viewer"`
	Username      string             `json:"username"`
}

type AdditionalDataDisplayResource struct {
	ConfigHeight int64  `json:"config_height"`
	ConfigWidth  int64  `json:"config_width"`
	Src          string `json:"src"`
}

type AdditionalDataPreviewComments struct {
	Count    int64                              `json:"count"`
	Edges    []AdditionalDataPreviewCommentEdge `json:"edges"`
	PageInfo AdditionalDataPageInfo             `json:"page_info"`
}

type AdditionalDataPreviewLikes struct {
	Count int64                           `json:"count"`
	Edges []AdditionalDataPreviewLikeEdge `json:"edges"`
}

type AdditionalDataPreviewCommentNode struct {
	CreatedAt       int64              `json:"created_at"`
	DidReportAsSpam bool               `json:"did_report_as_spam"`
	ID              string             `json:"id"`
	Owner           AdditionalDataUser `json:"owner"`
	Text            string             `json:"text"`
	ViewerHasLiked  bool               `json:"viewer_has_liked"`
}

type AdditionalDataToCaption struct {
	Edges []AdditionalDataToCaptionEdge `json:"edges"`
}

type AdditionalDataToTaggedUser struct {
	Edges []AdditionalDataToTaggedUserEdge `json:"edges"`
}

type AdditionalDataToChildren struct {
	Edges []AdditionalDataToChildrenEdge `json:"edges"`
}

type AdditionalDataFeed struct {
	Edges    []AdditionalDataPostEdge `json:"edges"`
	PageInfo AdditionalDataPageInfo   `json:"page_info"`
}

type AdditionalDataToSponsorUser struct {
	Edges []interface{} `json:"edges"`
}

type AdditionalDataPageInfo struct {
	EndCursor   string `json:"end_cursor"`
	HasNextPage bool   `json:"has_next_page"`
}

type AdditionalDataDimensions struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
}

type AdditionalDataDashInfo struct {
	IsDashEligible    bool   `json:"is_dash_eligible"`
	NumberOfQualities int64  `json:"number_of_qualities"`
	VideoDashManifest string `json:"video_dash_manifest"`
}

type AdditionalDataToCaptionEdge struct {
	Node AdditionalDataToCaptionNode `json:"node"`
}

type AdditionalDataToTaggedUserEdge struct {
	Node AdditionalDataToTaggedUserNode `json:"node"`
}

type AdditionalDataToChildrenEdge struct {
	Node AdditionalDataToChildrenNode `json:"node"`
}

type AdditionalDataPostEdge struct {
	Node AdditionalDataPost `json:"node"`
}

type AdditionalDataPreviewLikeEdge struct {
	Node AdditionalDataUser `json:"node"`
}

type AdditionalDataPreviewCommentEdge struct {
	Node AdditionalDataPreviewCommentNode `json:"node"`
}

type AdditionalDataToCaptionNode struct {
	Text string `json:"text"`
}

type AdditionalDataPost struct {
	Typename                   string                          `json:"__typename"`
	AccessibilityCaption       string                          `json:"accessibility_caption"`
	Attribution                interface{}                     `json:"attribution"`
	CommentsDisabled           bool                            `json:"comments_disabled"`
	DashInfo                   AdditionalDataDashInfo          `json:"dash_info"`
	Dimensions                 AdditionalDataDimensions        `json:"dimensions"`
	DisplayResources           []AdditionalDataDisplayResource `json:"display_resources"`
	DisplayURL                 string                          `json:"display_url"`
	PreviewComments            AdditionalDataPreviewComments   `json:"edge_media_preview_comment"`
	PreviewLikes               AdditionalDataPreviewLikes      `json:"edge_media_preview_like"`
	ToCaption                  AdditionalDataToCaption         `json:"edge_media_to_caption"`
	ToSponsorUser              AdditionalDataToSponsorUser     `json:"edge_media_to_sponsor_user"`
	ToTaggedUser               AdditionalDataToTaggedUser      `json:"edge_media_to_tagged_user"`
	ToChildren                 AdditionalDataToChildren        `json:"edge_sidecar_to_children"`
	FactCheckInformation       interface{}                     `json:"fact_check_information"`
	FactCheckOverallRating     interface{}                     `json:"fact_check_overall_rating"`
	FollowHashtagInfo          interface{}                     `json:"follow_hashtag_info"`
	GatingInfo                 interface{}                     `json:"gating_info"`
	ID                         string                          `json:"id"`
	IsVideo                    bool                            `json:"is_video"`
	Location                   interface{}                     `json:"location"`
	MediaPreview               string                          `json:"media_preview"`
	Owner                      AdditionalDataUser              `json:"owner"`
	Shortcode                  string                          `json:"shortcode"`
	TakenAtTimestamp           int64                           `json:"taken_at_timestamp"`
	TrackingToken              string                          `json:"tracking_token"`
	VideoURL                   string                          `json:"video_url"`
	VideoViewCount             int64                           `json:"video_view_count"`
	ViewerCanReshare           bool                            `json:"viewer_can_reshare"`
	ViewerHasLiked             bool                            `json:"viewer_has_liked"`
	ViewerHasSaved             bool                            `json:"viewer_has_saved"`
	ViewerHasSavedToCollection bool                            `json:"viewer_has_saved_to_collection"`
	ViewerInPhotoOfYou         bool                            `json:"viewer_in_photo_of_you"`
}

type AdditionalDataToChildrenNode struct {
	Typename             string                          `json:"__typename"`
	AccessibilityCaption string                          `json:"accessibility_caption"`
	DashInfo             AdditionalDataDashInfo          `json:"dash_info"`
	Dimensions           AdditionalDataDimensions        `json:"dimensions"`
	DisplayResources     []AdditionalDataDisplayResource `json:"display_resources"`
	DisplayURL           string                          `json:"display_url"`
	ToTaggedUser         AdditionalDataToTaggedUser      `json:"edge_media_to_tagged_user"`
	FollowHashtagInfo    interface{}                     `json:"follow_hashtag_info"`
	ID                   string                          `json:"id"`
	IsVideo              bool                            `json:"is_video"`
	TrackingToken        string                          `json:"tracking_token"`
	VideoURL             string                          `json:"video_url"`
	VideoViewCount       int64                           `json:"video_view_count"`
}

type AdditionalDataToTaggedUserNode struct {
	User AdditionalDataUser `json:"user"`
	X    float64            `json:"x"`
	Y    float64            `json:"y"`
}
