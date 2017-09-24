package webserver

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/chrisvdg/GorageRemote/config"
)

const (
	// how long self signed certificates are valid in days
	certLifespan = 365
	// tls min version
	tlsMinVersion = tls.VersionTLS12
)

// tlsConfig returns a TLS config from provided Zedis config
func tlsConfig(app *config.App) error {
	// if key and cert files don't exist, generate them
	certExists, err := exists(app.TLSCertPath)
	if err != nil {
		return err
	}
	keyExists, err := exists(app.TLSKeyPath)
	if err != nil {
		return err
	}

	if !certExists || !keyExists {
		return fmt.Errorf("certificate and key not found\nGenerating in code not implemented yet...\nRun 'go generate' in the certs directory")
	}

	return nil
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
