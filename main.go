package main

import "fmt"

// Version returns the SemVer for this app.
func Version() string {
	return "v0.0.1"
}

func main() {
	fmt.Println(Version())
}
