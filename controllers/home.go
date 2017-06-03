package controllers

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
)

// GetHome handles GET request on home route
func GetHome(w http.ResponseWriter, r *http.Request, app config.App) {
	fmt.Fprintln(w, "Hello world!")
}
