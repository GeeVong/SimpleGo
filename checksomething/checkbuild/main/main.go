package main

import (
	"embed"
	"fmt"
	"github.com/no-src/fserver"
)

var WebDist embed.FS

func main() {

	fmt.Println("heelo")
	fserver.Run(8001, "/build", WebDist)
}
