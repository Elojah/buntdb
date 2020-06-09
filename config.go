package buntdb

import (
	"errors"
)

// Config is buntdb structure config.
type Config struct {
	Path string
}

// Equal returns is both configs are equal.
func (c Config) Equal(rhs Config) bool {
	return c.Path == rhs.Path
}

// Dial set the config from a config namespace.
func (c *Config) Dial(fileconf interface{}) error {
	fconf, ok := fileconf.(map[string]interface{})
	if !ok {
		return errors.New("namespace empty")
	}

	cPath, ok := fconf["path"]
	if !ok {
		return errors.New("missing key path")
	}
	if c.Path, ok = cPath.(string); !ok {
		return errors.New("key path invalid. must be string")
	}

	return nil
}
