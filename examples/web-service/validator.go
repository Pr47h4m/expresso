package webservice

import (
	"encoding/xml"
	"net/http"

	"github.com/pr47h4m/expresso"
)

func ValidateAPIKey(ctx *expresso.Context) {
	key := ctx.Headers.Get("api-key")
	if key == "" {
		ctx.Info("api-key not found in headers, checking query params")
		key = ctx.QueryParams.Get("api-key")
	}

	keyMatched := false

	for i := 0; i < len(apiKeys); i++ {
		if apiKeys[i] == key {
			keyMatched = true
			break
		}
	}

	if keyMatched {
		ctx.Debug("api-key matched")
		ctx.Next()
	} else {
		ctx.Status(http.StatusUnauthorized).Formatted(ctx.RawRequest, expresso.Formatted{
			Text: &expresso.Text{
				Content: "401 - invalid api key",
			},
			HTML: &expresso.HTML{
				Content: "<html><head><title>401 - invalid api key</title></head><body>401 - invalid api key</body></html>",
			},
			JSON: &expresso.JSON{
				Data: map[string]interface{}{
					"status": 401,
					"error":  "invalid api key",
				},
			},
			XML: &expresso.XML{
				Data: struct {
					XMLName xml.Name `xml:"error"`
					Status  int      `xml:"status"`
					Error   string   `xml:"error"`
				}{
					Status: 401,
					Error:  "invalid api key",
				},
			},
			YAML: &expresso.YAML{
				Data: map[string]interface{}{
					"name":    "John Doe",
					"age":     30,
					"address": "123 Main St",
					"skills":  []string{"Go", "Python", "JavaScript"},
				},
			},
			Default: expresso.JSON{
				Data: map[string]interface{}{
					"status": 401,
					"error":  "invalid api key",
				},
			},
		})
	}
}
