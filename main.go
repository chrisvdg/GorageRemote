package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/chrisvdg/GorageRemote/db"

	"github.com/chrisvdg/GorageRemote/config"
	server "github.com/chrisvdg/GorageRemote/webserver"
	"github.com/gorilla/securecookie"
)

func main() {
	// get app config
	app, err := setupApp()
	if err != nil {
		log.Fatalf("could not set up app: %s\n", err)
	}

	// set routes
	server.SetRoutes(app)

	// run server
	log.Fatal(server.Run(app))
}

// SetupApp sets up the app config
func setupApp() (*config.App, error) {
	// setup flags
	port := flag.Int("port", 6060, "Sets port of the web server")
	configPath := flag.String("cfg", "", "Sets config file location for the server")
	dbPath := flag.String("dbpath", "", "Sets the path to the sqlite database")
	flag.Parse()

	// setup by config file or cli args
	app := new(config.App)
	if *configPath != "" {
		fmt.Println("using config file")
		cfgApp, err := config.NewApp(*configPath)
		if err != nil {
			return nil, fmt.Errorf("Could not get server config from file: %v", err)
		}
		app = cfgApp
	} else {
		fmt.Println("using cli args")
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

	// generate cookiestor secret if not provided or using cli
	if len(app.CookiestoreSecret) <= 0 {
		app.CookiestoreSecret = string(securecookie.GenerateRandomKey(64))
	}

	// validate
	err := app.Validate()
	if err != nil {
		return nil, err
	}

	// setup db
	app.DB, err = db.NewDB(app.SqlitePath)
	if err != nil {
		return nil, err
	}

	return app, nil
}
