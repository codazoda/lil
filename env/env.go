package env

import (
	"os"
)

// Read a file and return all the lines.
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = fallback
	}
	return value
}
