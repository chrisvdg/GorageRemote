package webserver

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/chrisvdg/GorageRemote/webserver/controllers"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

// Run runs the webserver
func Run(app *config.App) error {
	// setup cookiestor
	store = sessions.NewCookieStore([]byte(app.CookiestoreSecret))

	fmt.Printf("Webserver running on: localhost:%d\n", app.ListenPort)
	return http.ListenAndServe(app.ListenPortString(), context.ClearHandler(http.DefaultServeMux))
}

// SetRoutes set routes sets the app's routes
func SetRoutes(app *config.App) {
	// home route
	http.Handle("/", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.Home(w, r, app)
	})))

	// auth routes
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		controllers.Auth(w, r, app)
	})
	http.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		controllers.AuthLogout(w, r, app)
	})

	// api routes
	http.Handle("/api", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.API(w, r, app)
	})))

	// administration route
	http.Handle("/admin", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.Admin(w, r, app)
	})))
}
