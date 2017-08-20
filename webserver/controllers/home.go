package controllers

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
)

// Home handles a request on home route
func Home(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}

// Admin handles a request for the home administation page
func Admin(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}

// API handles a request for the home API page
func API(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}
