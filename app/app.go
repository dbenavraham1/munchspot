package app

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"time"
)

type App struct {
	Server struct {
		// Host is the local machine IP Address to bind the HTTP Server to
		Host string `yaml:"host"`

		// Port is the local machine TCP Port to bind the HTTP Server to
		Port    string `yaml:"port"`

		Timeout struct {
			// Server is the general server timeout to use
			// for graceful shutdowns
			Server time.Duration `yaml:"server"`

			// Write is the amount of time to wait until an HTTP server
			// write opperation is cancelled
			Write time.Duration `yaml:"write"`

			// Read is the amount of time to wait until an HTTP server
			// read operation is cancelled
			Read time.Duration `yaml:"read"`

			// Read is the amount of time to wait
			// until an IDLE HTTP session is closed
			Idle time.Duration `yaml:"idle"`
		} `yaml:"timeout"`
	} `yaml:"server"`
}

func NewApp(appPath string) (*App, error) {
	app := &App{}

	// Open app file
	file, err := os.Open("app.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&app); err != nil {
		return nil, err
	}

	return app, nil
}

func ValidateAppPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

func ParseFlags() (string, error) {
	var appPath string

	// Create a cli flag called "-app" to override default app configuration
	flag.StringVar(&appPath, "app", "./app.yml", "path to app configuration file")

	flag.Parse()

	if err := ValidateAppPath(appPath); err != nil {
		return "", err
	}

	return appPath, nil
}
