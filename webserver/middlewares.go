package webserver

import (
	"fmt"
	"net/http"
)

func checkAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check login
		if checklogin(w, r) > 0 {
			next.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/auth?error=not%20authenticated", 301)
		return
	})
}

func checklogin(w http.ResponseWriter, r *http.Request) int {
	// check login
	s, err := store.Get(r, "authentication")
	if err != nil {
		fmt.Println(err)
	}
	id := s.Values["id"]
	if id != nil && id.(int) > 0 {
		return id.(int)
	}
	return -1
}
