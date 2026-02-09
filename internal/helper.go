package internal

import (
	"net/url"
)

func buildURL(host string, path string, queries []map[string]string) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}

	u.Path = path

	q := u.Query()
	for _, m := range queries {
		for key, val := range m {
			if val != "" {
				q.Set(key, val)
			}
		}
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}
