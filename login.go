package sugusama

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	ErrLoginNotAuthenticated = errors.New("login not authenticated")
)

func (c *Client) Login(user, pass string) error {
	if c.State.Viewer != nil && c.State.Viewer.Username == user {
		return nil
	}

	values := url.Values{
		"username": []string{user},
		"password": []string{pass},
	}
	data := strings.NewReader(values.Encode())

	u, _ := url.Parse(c.Bases.Web)
	u.Path = "/accounts/login/ajax/"

	req, err := http.NewRequest("POST", u.String(), data)
	if err != nil {
		err = fmt.Errorf("post %q: %w", u, err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	cresp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("do req: %w", err)
		return err
	}
	defer cresp.Body.Close()

	body, err := ioutil.ReadAll(cresp.Body)
	cresp.Body.Close()
	if err != nil {
		err = fmt.Errorf("read all body: %w", err)
		return err
	}

	resp := loginResp{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		err = fmt.Errorf("unmarshal %#q %#v: %w", body, resp, err)
		return err
	}

	if !resp.Authenticated {
		err := ErrLoginNotAuthenticated
		return err
	}

	if resp.Status != "ok" {
		err := NotOK("login", resp.Status, resp.Message)
		return err
	}

	if c.State.Viewer == nil {
		c.State.Viewer = &User{}
	}

	if resp.UserID != "" {
		c.State.Viewer.ID = resp.UserID
	}

	c.State.Viewer.Username = user
	c.State.Viewer.Password = pass

	return nil
}

type loginResp struct {
	Authenticated bool   `json:"authenticated"`
	User          bool   `json:"user"`
	UserID        string `json:"userId"`
	OneTapPrompt  bool   `json:"oneTapPrompt"`
	Status        string `json:"status"`
	Message       string `json:"message"`
}
