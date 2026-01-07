package static

import (
	_ "embed"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//go:embed configure.html
var configure []byte

func HandleConfigure(c *fiber.Ctx) error {
	c.Response().Header.Add("Cache-control", "max-age=86400, public")
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)

	content := string(configure)
	authToken := os.Getenv("STREAMX_AUTH_TOKEN")
	if authToken != "" {
		content = strings.Replace(content, "const AUTH_TOKEN = ''", "const AUTH_TOKEN = '"+authToken+"'", 1)
	}

	return c.SendString(content)
}
