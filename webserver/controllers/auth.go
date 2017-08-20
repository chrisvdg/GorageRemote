package controllers

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/chrisvdg/GorageRemote/db"
	e "github.com/chrisvdg/GorageRemote/entities"
	"github.com/gorilla/sessions"
)

// Auth handles a request on auth route
func Auth(w http.ResponseWriter, r *http.Request, app *config.App, store *sessions.CookieStore) {
	if r.Method == "GET" {
		// check login
		// TODO, replace by middelware
		auth := checklogin(w, r, store)
		if auth {
			http.Redirect(w, r, "/", 303)
			return
		}
		// serve static admin login page
		http.ServeFile(w, r, "static/login.html")

	} else if r.Method == "POST" {
		// handle login form
		u := new(e.User)

		if r.FormValue("username") != "" {
			u.Name = r.FormValue("username")
		} else {
			http.Redirect(w, r, "/auth?error=enter%20username", 303)
			return
		}
		if r.FormValue("password") != "" {
			u.Password = r.FormValue("password")
		} else {
			http.Redirect(w, r, "/auth?error=enter%20password", 303)
			return
		}

		// check with db
		err := db.CheckPassword(app.DB, u)
		if err != nil {
			http.Redirect(w, r, "/auth?error=failed%20to%20authenticate", 303)
			return
		}
		// create session and redirect
		s, err := store.Get(r, "authentication")
		if err != nil {
			fmt.Println(err)
		}
		s.Values["id"] = u.Name
		s.Values["auth"] = true
		s.Save(r, w)
		http.Redirect(w, r, "/?message=login%20success", 303)

	} else {
		http.Redirect(w, r, "/", 303)
	}
}

// AuthLogout handles a logout request
func AuthLogout(w http.ResponseWriter, r *http.Request, app *config.App, store *sessions.CookieStore) {
	// end session
	s, _ := store.Get(r, "authentication")
	delete(s.Values, "id")
	s.Options.MaxAge = -1
	s.Save(r, w)
	http.Redirect(w, r, "/auth?message=logout%20success", 303)
}

func checklogin(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) bool {
	// check login
	s, err := store.Get(r, "authentication")
	if err != nil {
		fmt.Println(err)
	}
	if auth, ok := s.Values["auth"].(bool); ok {
		return auth
	}
	return false
}
