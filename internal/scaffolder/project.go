package scaffolder

import (
	"fmt"
	"os"
)

func GenerateProject(name string) {
	os.Mkdir(name, 0755)
	f, _ := os.Create(fmt.Sprintf("%s/main.go", name))
	f.WriteString(`package main

import "fmt"

func main() {
    fmt.Println("Hello from ` + name + `!")
}

`)
	f.Close()
}
