package sugusama

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	Client *http.Client
	State  *State
	Bases  *Bases
}

func NewClient(bases *Bases) (*Client, error) {
	c := &Client{}

	c.Client = &http.Client{}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	c.Client.Jar = jar

	c.State = &State{}

	if bases == nil {
		bases = DefaultBases
	}
	c.Bases = bases

	_, _, err = c.Home()
	if err != nil {
		return nil, err
	}

	return c, nil
}
