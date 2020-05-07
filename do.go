package sugusama

import "net/http"

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	c.SetHeaders(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	c.GetHeaders(resp)

	return resp, nil
}
