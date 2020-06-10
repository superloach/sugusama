package sugusama

import "errors"

var (
	ErrActivityMissing = errors.New("no activity returned")
)

const (
	activityHash = "0f318e8cfff9cc9ef09f88479ff571fb"
)

type Activity struct {
	CommentLikes  int `json:"comment_likes"`
	Comments      int `json:"comments"`
	Likes         int `json:"likes"`
	Relationships int `json:"relationships"`
	UserTags      int `json:"usertags"`
}

func (c *Client) FetchActivity() error {
	resp := activityResp{}
	err := c.GraphQL(
		activityHash,
		map[string]interface{}{
			"id": c.State.Viewer.ID,
		},
		&resp,
	)
	if err != nil {
		return err
	}

	if resp.Status != "ok" {
		err := NotOK("activity", resp.Status, resp.Message)
		return err
	}

	edges := resp.Data.User.EdgeActivityCount.Edges
	if len(edges) == 0 {
		return ErrActivityMissing
	}

	c.State.Activity = edges[0].Node

	return nil
}

type activityResp struct {
	Data struct {
		User struct {
			EdgeActivityCount struct {
				Edges []struct {
					Node *Activity `json:"node"`
				} `json:"edges"`
			} `json:"edge_activity_count"`
		} `json:"user"`
	} `json:"data"`

	Status  string `json:"status"`
	Message string `json:"message"`
}
