package webserver

import (
	"fmt"
	"net/http"
	"os"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/chrisvdg/GorageRemote/webserver/controllers"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

// Run runs the webserver
func Run(app *config.App) error {
	// setup tls
	err := tlsConfig(app)
	if err != nil {
		return err
	}

	// setup cookiestor
	store = sessions.NewCookieStore([]byte(app.CookiestoreSecret))

	fmt.Printf("Webserver running on: localhost:%d\n", app.ListenPort)
	return http.ListenAndServeTLS(app.ListenPortString(), app.TLSCertPath, app.TLSKeyPath, context.ClearHandler(http.DefaultServeMux))
}

// SetRoutes set routes sets the app's routes
func SetRoutes(app *config.App) {
	// public assets
	fs := onlyFiles{http.Dir("webserver/assets")}
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(fs)))

	// home route
	http.Handle("/", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.Home(w, r, app)
	})))

	// auth routes
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		controllers.Auth(w, r, app, store)
	})
	http.HandleFunc("/auth/logout", func(w http.ResponseWriter, r *http.Request) {
		controllers.AuthLogout(w, r, app, store)
	})

	// api routes
	http.Handle("/api", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.API(w, r, app)
	})))
	http.Handle("/api/actionsocket", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.ActionSocket(w, r, app)
	})))

	// administration route
	http.Handle("/admin", checkAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		controllers.Admin(w, r, app)
	})))
}

type onlyFiles struct {
	fs http.FileSystem
}

func (fs onlyFiles) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return emptydir{f}, nil
}

type emptydir struct {
	http.File
}

func (f emptydir) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}
