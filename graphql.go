package sugusama

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func (c *Client) GraphQL(hash string, variables interface{}, ret interface{}) error {
	u, err := url.Parse(c.Bases.Web)
	if err != nil {
		return err
	}
	u.Path = "/graphql/query/"

	vars, err := json.Marshal(variables)
	if err != nil {
		return err
	}

	u.RawQuery = url.Values{
		"query_hash": []string{hash},
		"variables":  []string{string(vars)},
	}.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return err
	}

	return nil
}
