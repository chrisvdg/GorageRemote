package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/chrisvdg/GorageRemote/controllers"
)

func main() {
	// init app
	app, err := config.NewApp("./appconfig.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get server config: %v", err)
		return
	}
	
	// TODO: commandline args override

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetHome(w, r, app)
	})

	// run server
	log.Fatal(http.ListenAndServe(app.ListenPortString(), nil))
}
