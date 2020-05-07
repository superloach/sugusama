package sugusama

import "errors"

var (
	ErrActivityMissing = errors.New("no activity returned")
)

const (
	ActivityHash = "0f318e8cfff9cc9ef09f88479ff571fb"
)

type ActivityResp struct {
	Data    ActivityData `json:"data"`
	Status  string       `json:"status"`
	Message string       `json:"message,omitempty"`
}

func (c *Client) Activity() (*ActivityResp, error) {
	resp := &ActivityResp{}
	err := c.DoGraphQL(
		ActivityHash,
		map[string]interface{}{
			"id": c.State.Login.UserID,
		},
		&resp,
	)
	if err != nil {
		return nil, err
	}

	if resp.Status != "ok" {
		err := NotOK("activity", resp.Message)
		return nil, err
	}

	return resp, nil
}

func (a *ActivityResp) Activity() (Activity, error) {
	edges := a.Data.User.EdgeActivityCount.Edges
	if len(edges) == 0 {
		return Activity{}, ErrActivityMissing
	}

	return edges[0].Node, nil
}

type ActivityData struct {
	User ActivityUser `json:"user"`
}

type ActivityUser struct {
	EdgeActivityCount ActivityEdgeActivityCount `json:"edge_activity_count"`
}

type ActivityEdgeActivityCount struct {
	Edges []ActivityEdge `json:"edges"`
}

type ActivityEdge struct {
	Node Activity `json:"node"`
}

type Activity struct {
	CommentLikes  int `json:"comment_likes"`
	Comments      int `json:"comments"`
	Likes         int `json:"likes"`
	Relationships int `json:"relationships"`
	UserTags      int `json:"usertags"`
}
