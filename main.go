package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chrisvdg/GorageRemote/config"
	"github.com/chrisvdg/GorageRemote/controllers"
)

func main() {
	// get app config
	app, err := setupApp()
	if err != nil {
		log.Fatal(err)
	}

	// set routes
	setRoutes(app)

	// run server
	fmt.Printf("Webserver running on port: %d\n", app.ListenPort)
	log.Fatal(http.ListenAndServe(app.ListenPortString(), nil))
}

// SetupApp sets up the app config
func setupApp() (*config.App, error) {
	// setup flags
	port := flag.Int("port", 6060, "Sets port of the web server")
	configPath := flag.String("cfg", "", "Sets config file location for the server")
	dbPath := flag.String("dbpath", "", "Sets the path to the sqlite database")
	flag.Parse()

	// setup
	app := new(config.App)
	if *dbPath != "" {
		cfgApp, err := config.NewApp(*configPath)
		if err != nil {
			return nil, fmt.Errorf("Could not get server config from file: %v", err)
		}
		app = cfgApp
	} else {
		app.ListenPort = uint16(*port)
		app.SqlitePath = *dbPath
	}

	// set tmp file for db if none was provided
	if app.SqlitePath == "" {
		tmpdbfile, err := ioutil.TempFile("", "")
		if err != nil {
			return nil, err
		}
		app.SqlitePath = tmpdbfile.Name()
	}

	// validate
	err := app.Validate()
	if err != nil {
		return nil, err
	}

	return app, nil
}

func setRoutes(app *config.App) {
	// home route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetHome(w, r, app)
	})
}
