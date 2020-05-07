package sugusama

import "net/http"

func (c *Client) GetHeaders(resp *http.Response) {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			c.State.CSRF = cookie.Value
		}
	}

	for k := range resp.Header {
		switch k {
		case "x-ig-set-www-claim":
			c.State.WWWClaim = resp.Header.Get(k)
		default:
		}
	}
}

func (c *Client) SetHeaders(req *http.Request) {
	req.Header.Set("X-IG-App-ID", "936619743392459")
	req.Header.Set("X-CSRFToken", c.State.CSRF)
	req.Header.Set("X-IG-WWW-Claim", c.State.WWWClaim)
}
