package sugusama

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

var (
	additionalDataRegexp = regexp.MustCompile(`window\.__additionalDataLoaded\('feed',(.+)\);`)

	ErrAdditionalDataNotFound = errors.New("additional data not found")
)

func (c *Client) extractAdditional(body []byte) error {
	ad := additionalData{}

	data := additionalDataRegexp.Find(body)
	if data == nil {
		return ErrAdditionalDataNotFound
	}
	data = data[37 : len(data)-2]

	err := json.Unmarshal(data, &ad)
	if err != nil {
		err = fmt.Errorf("unmarshal ad: %w", err)
		return err
	}

	c.State.processAdditional(ad)

	return nil
}

type additionalData struct {
	User struct {
		*User
		Feed struct {
			Edges []struct {
				Node *Post `json:"node"`
			} `json:"edges"`
			PageInfo PageInfo `json:"page_info"`
		} `json:"edge_web_feed_timeline"`
	} `json:"user"`
}

type Comment struct {
	CreatedAt       int64  `json:"created_at"`
	DidReportAsSpam bool   `json:"did_report_as_spam"`
	ID              string `json:"id"`
	Owner           *User  `json:"owner"`
	Text            string `json:"text"`
	ViewerHasLiked  bool   `json:"viewer_has_liked"`
}

type additionalDataDashInfo struct {
	IsDashEligible    bool   `json:"is_dash_eligible"`
	NumberOfQualities int64  `json:"number_of_qualities"`
	VideoDashManifest string `json:"video_dash_manifest"`
}
