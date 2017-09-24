package config

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gorilla/securecookie"
)

var (
	defaultApp = &App{
		ListenPort:  443,
		TLSCertPath: "./certs/cert.pem",
		TLSKeyPath:  "./certs/key.pem",
	}
)

// App represents data useable throughout the app
type App struct {
	AppConfigPath     string  `json:"-"`
	ListenPort        uint16  `json:"port"`
	TLSCertPath       string  `json:"tls_cert_path"`
	TLSKeyPath        string  `json:"tls_key_path"`
	SqlitePath        string  `json:"db_path"`
	CookiestoreSecret string  `json:"cookie_store_secret"`
	DB                *sql.DB `json:"-"`
}

// NewApp returns the app data from provided json file
func NewApp(path string) (*App, error) {
	app := new(App)

	// Get data from provided json file
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		errstr := fmt.Sprintf("Error while reading config file: %v", err)
		return nil, errors.New(errstr)
	}
	json.Unmarshal(raw, app)
	app.AppConfigPath = path

	return app, nil
}

// FillEmptyWithDefault fills empty values with default ones
func (app *App) FillEmptyWithDefault() error {
	// set default port of none provided
	if app.ListenPort == 0 {
		app.ListenPort = defaultApp.ListenPort
	}

	// set tmp file for db if none was provided
	if app.SqlitePath == "" {
		tmpdbfile, err := ioutil.TempFile("", "")
		if err != nil {
			return err
		}
		app.SqlitePath = tmpdbfile.Name()
	}

	// generate cookiestor secret if not provided or using cli
	if len(app.CookiestoreSecret) <= 0 {
		app.CookiestoreSecret = string(securecookie.GenerateRandomKey(64))
	}

	// set tls paths if none was provided
	if app.TLSCertPath == "" {
		app.TLSCertPath = defaultApp.TLSCertPath
	}
	if app.TLSKeyPath == "" {
		app.TLSKeyPath = defaultApp.TLSKeyPath
	}

	return nil
}

// ListenPortString returns string of ListenPort
func (app *App) ListenPortString() string {
	portstr := fmt.Sprintf(":%v", app.ListenPort)
	return portstr
}
