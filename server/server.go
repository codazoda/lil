package server

import "os"

// Return the port to use as a string. Use the PORT environment variable if it
// exists otherwise default to 8080.
func GetServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
