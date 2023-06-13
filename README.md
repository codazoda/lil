# Lil

A little module of functions I find helpful for Go.

## Modules

### lil/server

A collection of functions I use frequently when writing a server in Go.

_GetServerPort() string_  
Read the PORT environment variable and return it or return 8080 if it's not set.

### lil/file

A collection of functions I use frequently when working with files in Go.

_FileToSlice(path string) ([]string, error)_  
Read a file and return all the lines as a slice of strings.
