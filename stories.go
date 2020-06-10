package sugusama

const (
	StoriesHash = "04334405dbdef91f2c4e207b84c204d7"
)

type StoriesResp struct {
	Data    StoriesData `json:"data"`
	status  string      `json:"status"`
	message string      `json:"message"`
}

type StoriesOpts struct {
	OnlyStories       bool `json:"only_stories"`
	Prefetch          bool `json:"stories_prefetch"`
	VideoDashManifest bool `json:"stories_video_dash_manifest"`
}

func (c *Client) Stories(opts *StoriesOpts) (*StoriesResp, error) {
	if opts == nil {
		opts = &StoriesOpts{
			OnlyStories:       true,
			Prefetch:          false,
			VideoDashManifest: false,
		}
	}

	resp := &StoriesResp{}
	err := c.GraphQL(
		StoriesHash,
		opts,
		&resp,
	)
	if err != nil {
		return nil, err
	}

	if resp.status != "ok" {
		err := NotOK("stories", resp.status, resp.message)
		return nil, err
	}

	return resp, nil
}

func (s *StoriesResp) Stories() []Story {
	edges := s.Data.User.FeedReelsTray.EdgeReelsTrayToReel.Edges
	stories := make([]Story, len(edges))
	for i, edge := range edges {
		stories[i] = edge.Node
	}
	return stories
}

type StoriesData struct {
	User StoriesDataUser `json:"user"`
}

type StoriesDataUser struct {
	FeedReelsTray StoriesDataFeedReelsTray `json:"feed_reels_tray"`
}

type StoriesDataFeedReelsTray struct {
	EdgeReelsTrayToReel StoriesDataEdgeReelsTrayToReel `json:"edge_reels_tray_to_reel"`
}

type StoriesDataEdgeReelsTrayToReel struct {
	Edges []StoriesDataEdge `json:"edges"`
}

type StoriesDataEdge struct {
	Node Story `json:"node"`
}

type Story struct {
	Typename              string      `json:"__typename"`
	HasBestiesMedia       bool        `json:"has_besties_media"`
	HasPrideMedia         bool        `json:"has_pride_media"`
	ID                    string      `json:"id"`
	CanReply              bool        `json:"can_reply"`
	CanReshare            bool        `json:"can_reshare"`
	ExpiringAt            int         `json:"expiring_at"`
	LatestReelMedia       int         `json:"latest_reel_media"`
	Muted                 bool        `json:"muted"`
	SupportsReelReactions interface{} `json:"supports_reel_reactions"`
	Items                 []StoryItem `json:"items"`
	PrefetchCount         int         `json:"prefetch_count"`
	RankedPosition        int         `json:"ranked_position"`
	Seen                  int         `json:"seen"`
	SeenRankedPosition    int         `json:"seen_ranked_position"`
	User                  StoriesUser `json:"user"`
	Owner                 StoriesUser `json:"owner"`
}

type StoryItem struct {
	Typename               string                      `json:"__typename"`
	Audience               string                      `json:"audience"`
	ID                     string                      `json:"id"`
	Dimensions             StoryItemDimensions         `json:"dimensions"`
	StoryViewCount         int                         `json:"story_view_count"`
	EdgeStoryMediaViewers  StoryItemMediaViewers       `json:"edge_story_media_viewers"`
	DisplayResources       []DisplayResource           `json:"display_resources"`
	DisplayURL             string                      `json:"display_url"`
	MediaPreview           string                      `json:"media_preview"`
	GatingInfo             interface{}                 `json:"gating_info"`
	FactCheckOverallRating interface{}                 `json:"fact_check_overall_rating"`
	FactCheckInformation   interface{}                 `json:"fact_check_information"`
	TakenAtTimestamp       int                         `json:"taken_at_timestamp"`
	ExpiringAtTimestamp    int                         `json:"expiring_at_timestamp"`
	StoryCtaURL            interface{}                 `json:"story_cta_url"`
	IsVideo                bool                        `json:"is_video"`
	Owner                  StoriesUser                 `json:"owner"`
	TrackingToken          string                      `json:"tracking_token"`
	TappableObjects        []StoryItemTappableObject   `json:"tappable_objects"`
	StoryAppAttribution    interface{}                 `json:"story_app_attribution"`
	MediaToSponsorUser     StoryItemMediaToSponsorUser `json:"edge_media_to_sponsor_user"`
}

type StoryItemDimensions struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type StoryItemMediaViewers struct {
	Count    int                           `json:"count"`
	Edges    []StoryItemMediaViewersEdge   `json:"edges"`
	PageInfo StoryItemMediaViewersPageInfo `json:"page_info"`
}

type StoryItemMediaViewersEdge struct {
	Node StoriesUser `json:"node"`
}

type StoryItemMediaViewersPageInfo struct {
	HasNextPage bool   `json:"has_next_page"`
	EndCursor   string `json:"end_cursor"`
}

type StoryItemTappableObject struct {
	Typename    string                       `json:"__typename"`
	X           float64                      `json:"x"`
	Y           float64                      `json:"y"`
	Width       float64                      `json:"width"`
	Height      float64                      `json:"height"`
	Rotation    float64                      `json:"rotation"`
	CustomTitle interface{}                  `json:"custom_title"`
	Attribution interface{}                  `json:"attribution"`
	Media       StoryItemTappableObjectMedia `json:"media"`
}

type StoryItemTappableObjectMedia struct {
	ID        string `json:"id"`
	Shortcode string `json:"shortcode"`
}

type StoryItemMediaToSponsorUser struct {
	Edges []interface{} `json:"edges"`
}

type StoriesUser struct {
	Typename   string `json:"__typename"`
	ID         string `json:"id"`
	ProfilePic string `json:"profile_pic_url"`
	Username   string `json:"username"`
	Followed   bool   `json:"followed_by_viewer"`
	Requested  bool   `json:"requested_by_viewer"`
}
