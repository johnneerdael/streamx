package addon

import "testing"

func TestAppendAuthToken(t *testing.T) {
	tests := []struct {
		name     string
		urlStr   string
		token    string
		expected string
	}{
		{
			name:     "empty token",
			urlStr:   "http://localhost/manifest.json",
			token:    "",
			expected: "http://localhost/manifest.json",
		},
		{
			name:     "no existing params",
			urlStr:   "http://localhost/manifest.json",
			token:    "test-token",
			expected: "http://localhost/manifest.json?token=test-token",
		},
		{
			name:     "existing params",
			urlStr:   "http://localhost/manifest.json?foo=bar",
			token:    "test-token",
			expected: "http://localhost/manifest.json?foo=bar&token=test-token",
		},
		{
			name:     "stremio protocol",
			urlStr:   "stremio://localhost/manifest.json",
			token:    "test-token",
			expected: "stremio://localhost/manifest.json?token=test-token",
		},
		{
			name:     "stremio protocol with params",
			urlStr:   "stremio://localhost/abc/manifest.json?foo=bar",
			token:    "test-token",
			expected: "stremio://localhost/abc/manifest.json?foo=bar&token=test-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := appendAuthToken(tt.urlStr, tt.token)
			if actual != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, actual)
			}
		})
	}
}
