# expresso
## Fast, unopinionated, minimalist web framework for [Go](https://go.dev)

<br>

### Examples

#### Hello World

```(go)
package main

import "github.com/pr47h4m/expresso"

func main() {
	app := expresso.DefaultApp()

	app.GET("/", func(ctx *expresso.Context) {
		ctx.Send(expresso.Text{
			Text: "Hello World",
		})
	})

	app.ListenAndServe(":80")
}
```

#### Basic Routing

```(go)
package main

import "github.com/pr47h4m/expresso"

func main() {
	app := expresso.DefaultApp()

	app.GET("/", func(ctx *expresso.Context) {
		ctx.Send(expresso.Text{
			Text: "Hello World",
		})
	})

	app.POST("/", func(ctx *expresso.Context) {
		ctx.Send(expresso.Text{
			Text: "Got a POST request",
		})
	})

	app.PUT("/user", func(ctx *expresso.Context) {
		ctx.Send(expresso.Text{
			Text: "Got a PUT request as /user",
		})
	})

	app.DELETE("/user", func(ctx *expresso.Context) {
		ctx.Send(expresso.Text{
			Text: "Got a DELETE request at /user",
		})
	})

	app.ListenAndServe(":80")
}
```

#### Serve Static

```(go)
package main

import (
	"net/http"

	"github.com/pr47h4m/expresso"
)

func main() {
	app := expresso.DefaultApp()

	app.ServeStatic("/public/*filepath", http.Dir("public"))
	app.ServeStatic("/private/*filepath", http.Dir("private"))

	app.ListenAndServe(":80")
}
```

#### Web Service

```(go)
package main

import (
	"encoding/xml"
	"net/http"

	"github.com/pr47h4m/expresso"
)

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type User struct {
	Name string `json:"name"`
}

var (
	apiKeys = []string{"foo", "bar", "baz"}
	repos   = []Repo{
		{
			Name: "validation_chain",
			URL:  "https://github.com/pr47h4m/validation_chain",
		},
		{
			Name: "expresso",
			URL:  "https://github.com/pr47h4m/expresso",
		},
	}
	users = []User{
		{
			Name: "pr47h4m",
		},
		{
			Name: "jhanvi1061",
		},
	}
	userRepos = map[User][]Repo{
		users[0]: repos,
		users[1]: {},
	}
)

func main() {
	app := expresso.DefaultApp()

	app.GET("/api", ValidateAPIKey, func(ctx *expresso.Context) {
		ctx.Send(expresso.HTML{
			HTML: "<html><head><title>Web Service in Go</title></head><body><h1>Web Service in Go</h1><h3>\npowered by github.com/pr47h4m/expresso</h3></body></html>",
		})
	})

	app.GET("/api/users", ValidateAPIKey, func(ctx *expresso.Context) {
		ctx.Send(expresso.JSON{
			"status": "200",
			"users":  users,
		})
	})

	app.GET("/api/repos", ValidateAPIKey, func(ctx *expresso.Context) {
		ctx.Send(expresso.JSON{
			"status": "200",
			"repos":  repos,
		})
	})

	app.GET("/api/users/:name/repos", ValidateAPIKey, func(ctx *expresso.Context) {
		username := ctx.Params.ByName("name")

		var idx int = -1

		for i := 0; i < len(users); i++ {
			if users[i].Name == username {
				idx = i
				break
			}
		}

		if idx != -1 {
			ctx.Send(expresso.JSON{
				"status": "200",
				"repos":  userRepos[users[idx]],
			})
		} else {
			ctx.Status(http.StatusNotFound).Send(expresso.JSON{
				"status": "404",
				"error":  "username not found",
			})
		}
	})

	app.ListenAndServe(":80")
}

func ValidateAPIKey(ctx *expresso.Context) {
	key := ctx.QueryParams.Get("api-key")

	keyMatched := false

	for i := 0; i < len(apiKeys); i++ {
		if apiKeys[i] == key {
			keyMatched = true
			break
		}
	}

	if keyMatched {
		ctx.Next()
	} else {
		ctx.Status(http.StatusUnauthorized).Formatted(ctx.RawRequest, expresso.Formatted{
			Text: expresso.Text{
				Text: "401 - invalid api key",
			},
			HTML: expresso.HTML{
				HTML: "<html><head><title>401 - invalid api key</title></head><body>401 - invalid api key</body></html>",
			},
			JSON: expresso.JSON{
				"status": 401,
				"error":  "invalid api key",
			},
			XML: expresso.XML{
				XML: struct {
					XMLName xml.Name `xml:"error"`
					Status  int      `xml:"status"`
					Error   string   `xml:"error"`
				}{
					Status: 401,
					Error:  "invalid api key",
				},
			},
			Default: expresso.JSON{
				"status": 401,
				"error":  "invalid api key",
			},
		})
	}
}
```