package sugusama

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type BroadcastsResp struct {
	Broadcasts []interface{} `json:"broadcasts"`
	status     string        `json:"status"`
	message    string        `json:"message"`
}

func (c *Client) Broadcasts() (*BroadcastsResp, error) {
	err := c.BroadcastsOptions()
	if err != nil {
		err = fmt.Errorf("broadcasts options: %w", err)
		return nil, err
	}

	resp := &BroadcastsResp{}

	u, _ := url.Parse(c.Bases.Live)
	u.Path = "/api/v1/live/reels_tray_broadcasts/"

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
		err := NotOK("broadcasts", resp.status, resp.message)
		return resp, err
	}

	return resp, nil
}

func (c *Client) BroadcastsOptions() error {
	u, _ := url.Parse(c.Bases.Live)
	u.Path = "/api/v1/live/reels_tray_broadcasts/"

	req, err := http.NewRequest("OPTIONS", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("options %q: %w", u, err)
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("do req: %w", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
