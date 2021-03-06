package env

import (
	"os"
	"strconv"
)

// GetOr takes an envionment variable string and if that env key is not set, returns the default which is passed as a
// seperate argument.
func GetOr(e, d string) string {
	if env := os.Getenv(e); env != "" {
		return env
	}
	return d
}

// GetBoolOr takes an envrionment variable string and converts to a bool.
func GetBoolOr(e string, d bool) bool {
	if env := os.Getenv(e); env != "" {
		b, err := strconv.ParseBool(env)
		if err != nil {
			return d
		}
		return b
	}
	return d
}
