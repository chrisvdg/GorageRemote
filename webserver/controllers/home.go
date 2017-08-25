package controllers

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
)

// Home handles a request on home route
func Home(w http.ResponseWriter, r *http.Request, app *config.App) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "webserver/static/home.html")
	} else {
		http.Redirect(w, r, "/", 303)
	}
}

// Admin handles a request for the home administation page
func Admin(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}
