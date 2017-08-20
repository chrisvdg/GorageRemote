package webserver

import (
	"fmt"
	"net/http"
)

func checkAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check login
		if checklogin(w, r) {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/auth?error=not%20authenticated", 303)
		return
	})
}

func checklogin(w http.ResponseWriter, r *http.Request) bool {
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
