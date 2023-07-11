package server

import (
	"os"
	"testing"
)

func TestGetServerPort(t *testing.T) {
	name := "GetServerPort()"

	// Verify that we got the default port of 8080 if PORT is not set
	os.Unsetenv("PORT")
	wantPort := "8080"
	gotPort := GetServerPort()
	if gotPort != wantPort {
		t.Fatalf(`%v: want port of %v got %v`, name, wantPort, gotPort)
	}

	// Verify that we got a custom port of 9090 when setting PORT to 9090
	t.Setenv("PORT", "9090")
	wantPort = "9090"
	gotPort = GetServerPort()
	if gotPort != wantPort {
		t.Fatalf(`%v: want port of %v got %v`, name, wantPort, gotPort)
	}

}
