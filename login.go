package sugusama

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrLoginNotAuthenticated = errors.New("login not authenticated")
)

type LoginResp struct {
	Authenticated bool   `json:"authenticated"`
	User          bool   `json:"user"`
	UserID        string `json:"userId"`
	OneTapPrompt  bool   `json:"oneTapPrompt"`
	Status        string `json:"status"`
	Message       string `json:"message,omitempty"`
}

func (c *Client) Login(user, pass string) (*LoginResp, error) {
	if c.State.Login != nil {
		return c.State.Login, nil
	}

	resp := &LoginResp{}

	values := url.Values{
		"username": []string{user},
		"password": []string{pass},
	}
	data := strings.NewReader(values.Encode())

	u, _ := url.Parse(c.Bases.Web)
	u.Path = "/accounts/login/ajax/"

	req, err := http.NewRequest("POST", u.String(), data)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	cresp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer cresp.Body.Close()

	err = json.NewDecoder(cresp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	if !resp.Authenticated {
		err := ErrLoginNotAuthenticated
		return nil, err
	}

	if resp.Status != "ok" {
		err := NotOK("login", resp.Message)
		return nil, err
	}

	c.State.Login = resp

	return resp, nil
}
