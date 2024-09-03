package servestatic

import (
	"net/http"

	"github.com/pr47h4m/expresso"
)

func App() *expresso.App {
	app := expresso.DefaultApp()
	app.ServeStatic("/public/*filepath", http.Dir("serve-static/public"))
	app.ServeStatic("/private/*filepath", http.Dir("serve-static/private"))
	return &app
}
