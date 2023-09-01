package fserver

import (
	"flag"
)

// InitFlag init the default flags
func InitFlag(versionFunc, aboutFunc func()) (exit bool, port int) {
	var (
		printVersion bool
		printAbout   bool
	)

	if versionFunc != nil {
		flag.BoolVar(&printVersion, "v", false, "print the version info")
	}
	if aboutFunc != nil {
		flag.BoolVar(&printAbout, "about", false, "print the about info")
	}
	flag.IntVar(&port, "port", 9099, "the HTTP server port")

	flag.Parse()

	if printVersion {
		versionFunc()
		exit = true
		return exit, port
	}

	if printAbout {
		aboutFunc()
		exit = true
		return exit, port
	}

	return exit, port
}
