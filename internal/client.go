package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// moodleClient is a direct HTTP REST client for the Moodle Web Services API.
type moodleClient struct {
	httpClient *http.Client
	siteURL    string
	token      string
	restful    bool
}

// newMoodleClient constructs a new moodleClient.
func newMoodleClient(siteURL, token string, restful bool) *moodleClient {
	return &moodleClient{
		httpClient: &http.Client{},
		siteURL:    strings.TrimRight(siteURL, "/"),
		token:      token,
		restful:    restful,
	}
}

// call invokes a Moodle Web Service function with the given parameters and returns the decoded JSON response.
// Standard mode: POST to {siteURL}/webservice/rest/server.php with wsfunction, wstoken, moodlewsrestformat=json
// RESTful mode:  POST to {siteURL}/webservice/restful/server.php/{wsfunction}?wstoken=...&moodlewsrestformat=json
func (c *moodleClient) call(ctx context.Context, wsfunction string, params map[string]string) (any, error) {
	var endpoint string
	form := url.Values{}

	if c.restful {
		endpoint = fmt.Sprintf("%s/webservice/restful/server.php/%s", c.siteURL, wsfunction)
		q := url.Values{}
		q.Set("wstoken", c.token)
		q.Set("moodlewsrestformat", "json")
		endpoint = endpoint + "?" + q.Encode()
	} else {
		endpoint = fmt.Sprintf("%s/webservice/rest/server.php", c.siteURL)
		form.Set("wstoken", c.token)
		form.Set("moodlewsrestformat", "json")
		form.Set("wsfunction", wsfunction)
	}

	for k, v := range params {
		form.Set(k, v)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("moodle call %s: build request: %w", wsfunction, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("moodle call %s: http: %w", wsfunction, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("moodle call %s: read body: %w", wsfunction, err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("moodle call %s: status %d: %s", wsfunction, resp.StatusCode, string(body))
	}

	// Detect Moodle exception envelope
	var raw any
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("moodle call %s: decode json: %w", wsfunction, err)
	}

	if m, ok := raw.(map[string]any); ok {
		if errCode, hasErr := m["errorcode"]; hasErr {
			msg, _ := m["message"].(string)
			return nil, fmt.Errorf("moodle error %v: %s", errCode, msg)
		}
		if ex, hasEx := m["exception"]; hasEx {
			msg, _ := m["message"].(string)
			return nil, fmt.Errorf("moodle exception %v: %s", ex, msg)
		}
	}

	return raw, nil
}

// callToMap is a convenience wrapper that expects a map[string]any response.
func (c *moodleClient) callToMap(ctx context.Context, wsfunction string, params map[string]string) (map[string]any, error) {
	raw, err := c.call(ctx, wsfunction, params)
	if err != nil {
		return nil, err
	}
	if m, ok := raw.(map[string]any); ok {
		return m, nil
	}
	// Some functions return nothing (null) on success
	return map[string]any{}, nil
}

// callToSlice is a convenience wrapper that expects a []any response.
func (c *moodleClient) callToSlice(ctx context.Context, wsfunction string, params map[string]string) ([]any, error) {
	raw, err := c.call(ctx, wsfunction, params)
	if err != nil {
		return nil, err
	}
	if s, ok := raw.([]any); ok {
		return s, nil
	}
	return []any{}, nil
}
