package gc_config

import (
	"os"
	"path/filepath"
)

const (
	configDirName  = "gcraft"
	configFileName = "config.json"
)

func DefaultConfigPath() string {
	if configDir, err := os.UserConfigDir(); err == nil {
		return filepath.Join(configDir, configDirName, configDirName)
	}

	return LocalConfigPath()
}

func LocalConfigPath() string {
	return filepath.Join(filepath.Dir(os.Args[0]), configFileName)
}
