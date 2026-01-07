package addon

import (
	"net/url"
	"strings"
)

func appendAuthToken(urlStr string, token string) string {
	if token == "" {
		return urlStr
	}

	isStremioProtocol := strings.HasPrefix(urlStr, "stremio://")
	urlToParse := urlStr
	if isStremioProtocol {
		urlToParse = "http://" + strings.TrimPrefix(urlStr, "stremio://")
	}

	u, err := url.Parse(urlToParse)
	if err != nil {
		return urlStr
	}

	query := u.Query()
	query.Set("token", token)
	u.RawQuery = query.Encode()

	result := u.String()
	if isStremioProtocol {
		result = "stremio://" + strings.TrimPrefix(result, "http://")
	}

	return result
}
