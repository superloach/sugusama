package sugusama

import (
	"fmt"
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
		err = fmt.Errorf("new jar: %w", err)
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
		err = fmt.Errorf("fetch home: %w", err)
		return nil, err
	}

	return c, nil
}
