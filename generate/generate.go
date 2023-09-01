//go:generate go run generate.go

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Running go generate...")
	file, err := os.Create("generated.go")
	if err != nil {
		fmt.Println("Failed to create generated.go:", err)
		return
	}
	defer file.Close()

	constantValue := `package main

const HelloWorld = "Hello, World!"
`

	_, err = file.WriteString(constantValue)
	if err != nil {
		fmt.Println("Failed to write constant value:", err)
		return
	}

	fmt.Println("generated.go generated successfully!")
}
