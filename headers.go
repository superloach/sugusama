package sugusama

import "net/http"

func (s *State) getHeaders(resp *http.Response) {
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "csrftoken" {
			s.CSRF = cookie.Value
		}
	}

	for k := range resp.Header {
		switch k {
		case "x-ig-set-www-claim":
			s.WWWClaim = resp.Header.Get(k)
		default:
		}
	}
}

func (s *State) setHeaders(req *http.Request) {
	req.Header.Set("X-IG-App-ID", InstagramWebDesktopFBAppID)
	req.Header.Set("X-CSRFToken", s.CSRF)
	req.Header.Set("X-IG-WWW-Claim", s.WWWClaim)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	c.setHeaders(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	c.getHeaders(resp)

	return resp, nil
}
