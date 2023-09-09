package file

import (
	"bufio"
	"os"
	"path/filepath"
)

// Alias for backward compatibility
var File = FileToSlice

// Read a file and return all the lines.
func FileToSlice(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Retun the XDG config path
func GetConfigDir(app string) (string, error) {
	configDir := os.Getenv("XDG_CONFIG_HOME")
	var homeDir string

	// If the value of the environment variable is unset, empty, or not an absolute path, use the default
	if configDir == "" || configDir[0:1] != "/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, ".config", app), nil
	}

	// The value of the environment variable is valid; use it
	return filepath.Join(homeDir, ".config", app), nil
}
