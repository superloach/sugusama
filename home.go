package sugusama

import (
	"io/ioutil"
	"net/http"
)

func (c *Client) Home() (*SharedData, *AdditionalData, error) {
	req, err := http.NewRequest("GET", c.Bases.Web, nil)
	if err != nil {
		return nil, nil, err
	}

	cresp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer cresp.Body.Close()

	body, err := ioutil.ReadAll(cresp.Body)
	if err != nil {
		return nil, nil, err
	}

	shared, err := c.ExtractSharedData(body)
	if err != nil {
		return nil, nil, err
	}

	additional, err := c.ExtractAdditionalData(body)
	if err != nil {
		return nil, nil, err
	}

	return shared, additional, nil
}
