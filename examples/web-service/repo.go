package webservice

import (
	"net/http"

	"github.com/pr47h4m/expresso"
)

func GetRepos(ctx *expresso.Context) {
	ctx.Send(expresso.JSON{
		Data: map[string]interface{}{
			"status": "200",
			"repos":  repos,
		},
	})
}

func GetUserRepos(ctx *expresso.Context) {
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
			Data: map[string]interface{}{
				"status": "200",
				"repos":  userRepos[users[idx]],
			},
		})
	} else {
		ctx.Status(http.StatusNotFound).Send(expresso.JSON{
			Data: map[string]interface{}{
				"status": "404",
				"error":  "username not found",
			},
		})
	}
}
