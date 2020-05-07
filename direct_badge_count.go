package sugusama

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type DirectBadgeCountResp struct {
	UserID         int64  `json:"user_id"`
	BadgeCount     int    `json:"badge_count"`
	SeqID          int    `json:"seq_id"`
	BadgeCountAtMs int64  `json:"badge_count_at_ms"`
	Status         string `json:"status"`
	Message        string `json:"message,omitempty"`
}

func (c *Client) DirectBadgeCount(noRaven bool) (*DirectBadgeCountResp, error) {
	resp := &DirectBadgeCountResp{}

	query := url.Values{}
	if noRaven {
		query.Set("no_raven", "1")
	} else {
		query.Set("no_raven", "0")
	}

	u, _ := url.Parse(c.Bases.Web)
	u.Path = "/direct_v2/web/get_badge_count/"
	u.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	cresp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer cresp.Body.Close()

	err = json.NewDecoder(cresp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	if resp.Status != "ok" {
		err := NotOK("direct badge count", resp.Message)
		return nil, err
	}

	return resp, nil
}
