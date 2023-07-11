# Lil

A little module of functions I find helpful for Go.

## Getting Started

Install the module with the following command (experimental URL).

```
go get github.com/codazoda/lil
```

Then use one or more of the following imports.

```
import (
    "github.com/codazoda/lil/env"
    "github.com/codazoda/lil/file"
    "github.com/codazoda/lil/server"
)
```

Now you can call the packages and functions below.

## Packages

### env

Functions for working with environment variables.

**GetEnv(key string, fallback string) ()**
Return the value stored in the envionment variable named `key` or fall back to the string specified in `fallback` if the environment variable is not set. Ex: `env.GetEnv("PORT", "8080")`.

### file

A collection of functions I use frequently when working with files in Go.

**FileToSlice(path string) ([]string, error)**  
Read a file and return all the lines as a slice of strings. Ex: `file.FileToSlice("/tmp/example.txt")`

### server

A collection of functions I use frequently when writing a server in Go.

**GetServerPort() string**  
Read the PORT environment variable and return it or return 8080 if it's not set. Ex: `server.GetServerPort()`
