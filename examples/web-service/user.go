package webservice

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/pr47h4m/expresso"
)

func GetUsers(ctx *expresso.Context) {
	ctx.Send(expresso.JSON{
		Data: map[string]interface{}{
			"status": "200",
			"users":  users,
		},
	})
}

func CreateUser(ctx *expresso.Context) {
	l := ctx.Logger()
	cType := strings.Split(ctx.Request.Headers.Get("Content-Type"), ";")[0]

	user := User{}
	l.Debug("body: " + string(ctx.Request.Body))
	l.Debug("query params: " + string(ctx.QueryParams.Encode()))
	switch cType {
	case "application/json":
		if err := json.Unmarshal(ctx.Request.Body, &user); err != nil {
			l.Error("error unmarshalling json: " + err.Error())
			ctx.SendStatus(400)
			return
		}
	case "application/xml":
		if err := xml.Unmarshal(ctx.Request.Body, &user); err != nil {
			l.Error("error unmarshalling xml: " + err.Error())
			ctx.SendStatus(400)
			return
		}
	case "application/x-www-form-urlencoded":
		user.Name = ctx.RawRequest.Form.Get("name")
	case "multipart/form-data":
		if err := ctx.RawRequest.ParseMultipartForm(1 << 20); err != nil {
			l.Error("error parsing multipart form: " + err.Error())
			ctx.SendStatus(400)
			return
		}
		user.Name = ctx.RawRequest.FormValue("name")
	default:
		l.Error("invalid content type: " + cType)
		ctx.SendStatus(400)
		return
	}

	if user.Name != "" {
		users = append(users, user)
		userRepos[user] = []Repo{}
		ctx.Status(http.StatusCreated).Send(expresso.JSON{
			Data: map[string]interface{}{
				"status": "201",
				"user":   user,
			},
		})
	}
}
