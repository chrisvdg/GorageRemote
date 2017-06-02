package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// App represents data useable throughout the app
type App struct {
	AppConfigPath string
	DBConfigPath  string `json:"db_config_path"`
	ListenPort    uint16 `json:"port"`
}

func NewApp(path string) (App, error) {
	var app App

	// Get data from provided json file
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		errstr := fmt.Sprintf("Error while reading config file: %v", err)
		return app, errors.New(errstr)
	}
	json.Unmarshal(raw, &app)
	app.AppConfigPath = path

	return app, nil
}

// ListenPortString returns string of ListenPort
func (app *App) ListenPortString() string {
	portstr := fmt.Sprintf(":%v", app.ListenPort)
	return portstr
}
