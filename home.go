package sugusama

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) FetchHome() error {
	req, err := http.NewRequest("GET", c.Bases.Web, nil)
	if err != nil {
		err = fmt.Errorf("get %q: %w", c.Bases.Web, err)
		return err
	}

	cresp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("do req: %w", err)
		return err
	}
	defer cresp.Body.Close()

	body, err := ioutil.ReadAll(cresp.Body)
	if err != nil {
		err = fmt.Errorf("readall body: %w", err)
		return err
	}

	_ = c.extractShared(body)
	_ = c.extractAdditional(body)

	return nil
}
