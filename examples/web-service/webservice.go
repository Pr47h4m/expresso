package webservice

import (
	"encoding/xml"

	"github.com/pr47h4m/expresso"
)

type Repo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type User struct {
	XMLName xml.Name `xml:"User"`
	Name    string   `json:"name" xml:"Name"`
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
			Name: "gautam",
		},
	}
	userRepos = map[User][]Repo{
		users[0]: repos,
		users[1]: {},
	}
)

func App() *expresso.App {
	app := expresso.DefaultApp()

	app.GET("/api", ValidateAPIKey, func(ctx *expresso.Context) {
		ctx.Send(expresso.HTML{Content: "<html><head><title>Web Service in Go</title></head><body><h1>Web Service in Go</h1><h3>\npowered by github.com/pr47h4m/expresso</h3></body></html>"})
	})

	app.GET("/api/users", ValidateAPIKey, GetUsers)

	app.POST("/api/users", ValidateAPIKey, CreateUser)

	app.GET("/api/repos", ValidateAPIKey, GetRepos)

	app.GET("/api/users/:name/repos", ValidateAPIKey, GetUserRepos)

	return &app
}
