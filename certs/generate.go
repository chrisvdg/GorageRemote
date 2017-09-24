package certs

//go:generate openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 3650 -nodes -subj "/C=BE/L=Gent/O=Gorange/CN=gorage.internal"
