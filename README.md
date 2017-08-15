# Gorage Remote

Golang based web app / API for the Raspberry Pi to open/close an electric garage door

# Setup

The app server can be configured in 2 ways, by cli flag arguments or by a JSON config file.

## CLI flags

* -cfg \<file location\>: Sets config file location for the server
* -port \<port number\>: Sets port of the web server
* -dbpath \<file location\>: Sets the path of the sqlite database

When -cfg is set, the other flags will be ignored and the config will be taken from the provided JSON config file

## Config file

```js
{
    "port": 8080, // Sets port of the web server
    "dbpath": "./gorage.db" // Sets the path of the sqlite database
}
```

## Default values

When no port is provided, port 6060 will be used.
If no dbpath is provided, a file will be created in the tmp directory and will run there. This is not saved so every time the app is relaunched it will create a new db file and start from the default database.

If the file defined by dbpath does not exists, an empty sqlite database file will be created.
If the file is empty, the app will set the tables and fields and inserts a default user into the database.

The default credentials are:  
* username: admin
* password: Gorage123

It is recommended to change the password after setting up the webserver through the web app.