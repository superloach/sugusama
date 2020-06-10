package sugusama

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	*State

	Client *http.Client
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

	err = c.FetchHome()
	if err != nil {
		return nil, err
	}

	return c, nil
}
