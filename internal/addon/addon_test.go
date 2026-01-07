package addon

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHandleGetManifest(t *testing.T) {
	tests := []struct {
		name      string
		authToken string
		expected  string
	}{
		{
			name:      "without token",
			authToken: "",
			expected:  "http://example.com/logo",
		},
		{
			name:      "with token",
			authToken: "test-token",
			expected:  "http://example.com/logo?token=test-token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			add := New(WithAuthToken(tt.authToken))
			
			app.Get("/manifest.json", add.HandleGetManifest)
			app.Get("/logo", add.HandleLogo)

			req := httptest.NewRequest("GET", "/manifest.json", nil)
			resp, _ := app.Test(req)

			if resp.StatusCode != 200 {
				t.Fatalf("expected status 200, got %d", resp.StatusCode)
			}

			var manifest Manifest
			json.NewDecoder(resp.Body).Decode(&manifest)

			if manifest.Logo != tt.expected {
				t.Errorf("expected logo %s, got %s", tt.expected, manifest.Logo)
			}
		})
	}
}
