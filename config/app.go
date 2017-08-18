package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// App represents data useable throughout the app
type App struct {
	AppConfigPath     string `json:"-"`
	ListenPort        uint16 `json:"port"`
	SqlitePath        string `json:"dbpath"`
	CookiestoreSecret string `json:"cookie_store_secret"`
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

// ListenPortString returns string of ListenPort
func (app *App) ListenPortString() string {
	portstr := fmt.Sprintf(":%v", app.ListenPort)
	return portstr
}

// Validate validate the App fields
func (app *App) Validate() error {
	if app.ListenPort <= 0 {
		return fmt.Errorf("listenport was not provided")
	}
	if len(app.SqlitePath) <= 0 {
		return fmt.Errorf("database file was not provided")
	}
	if len(app.CookiestoreSecret) <= 0 {
		return fmt.Errorf("cookiestore secret was not provided")
	}
	return nil
}
