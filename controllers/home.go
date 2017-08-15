package controllers

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
)

// GetHome handles GET request on home route
func GetHome(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}

// GETAuth handles GET request on auth route
func GETAuth(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}

// PostAuth handles POST request on auth route
func PostAuth(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}

// GetAdmin handles GET request for administation page
func GetAdmin(w http.ResponseWriter, r *http.Request, app *config.App) {
	fmt.Fprintln(w, "Hello world!")
}
