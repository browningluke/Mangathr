package defaults

import (
	"github.com/browningluke/mangathrV2/internal/logging"
	"os"
	"path/filepath"
)

/*
	Config directory defaults
*/

func ConfigDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logging.Fatalln(err)
	}

	configDir := filepath.Join(homeDir, ".config", "mangathrv2")
	return configDir
}

func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config")
}

/*
	Config defaults
*/

func DatabaseDriver() string {
	return "sqlite"
}

func DatabaseUri() string {
	return filepath.Join(ConfigDir(), "db.sqlite")
}

/*
	Log level
*/

func LogLevel() string {
	return "WARNING"
}
