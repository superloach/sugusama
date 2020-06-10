package sugusama

import (
	"io/ioutil"
	"net/http"
)

func (c *Client) FetchHome() error {
	req, err := http.NewRequest("GET", c.Bases.Web, nil)
	if err != nil {
		return err
	}

	cresp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer cresp.Body.Close()

	body, err := ioutil.ReadAll(cresp.Body)
	if err != nil {
		return err
	}

	_ = c.extractShared(body)
	_ = c.extractAdditional(body)

	return nil
}
