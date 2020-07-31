package sugusama

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type DirectBadgeCountResp struct {
	UserID         int64  `json:"user_id"`
	BadgeCount     int    `json:"badge_count"`
	SeqID          int    `json:"seq_id"`
	BadgeCountAtMS int64  `json:"badge_count_at_ms"`
	status         string `json:"status"`
	message        string `json:"message"`
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
		err = fmt.Errorf("get %q: %w", u, err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	cresp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("do req: %w", err)
		return nil, err
	}
	defer cresp.Body.Close()

	err = json.NewDecoder(cresp.Body).Decode(&resp)
	if err != nil {
		err = fmt.Errorf("decode resp: %w", err)
		return nil, err
	}

	if resp.status != "ok" {
		err := NotOK("direct badge count", resp.status, resp.message)
		return nil, err
	}

	return resp, nil
}
