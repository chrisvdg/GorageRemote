package webserver

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/chrisvdg/GorageRemote/webserver/controllers"
)

// Run runs the webserver
func Run(app *config.App) error {
	fmt.Printf("Webserver running on port: %d\n", app.ListenPort)
	return http.ListenAndServe(app.ListenPortString(), nil)
}

// SetRoutes set routes sets the app's routes
func SetRoutes(app *config.App) {
	// home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.Home(w, r, app)
	})
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		controllers.Admin(w, r, app)
	})

	// auth routes
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		controllers.Auth(w, r, app)
	})
	http.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		controllers.AuthLogout(w, r, app)
	})

	// api routes
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		controllers.API(w, r, app)
	})
}
