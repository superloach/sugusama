package sugusama

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GraphQL(hash string, variables interface{}, ret interface{}) error {
	u, err := url.Parse(c.Bases.Web)
	if err != nil {
		err = fmt.Errorf("parse %q: %w", c.Bases.Web, err)
		return err
	}
	u.Path = "/graphql/query/"

	vars, err := json.Marshal(variables)
	if err != nil {
		err = fmt.Errorf("marshal %v: %q", variables, err)
		return err
	}

	u.RawQuery = url.Values{
		"query_hash": []string{hash},
		"variables":  []string{string(vars)},
	}.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		err = fmt.Errorf("get %q: %w", u, err)
		return err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		err = fmt.Errorf("do req: %w", err)
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		err = fmt.Errorf("decode ret: %w", err)
		return err
	}

	return nil
}
