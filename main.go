package main

import "fmt"
import "github.com/senaduka/buttons-of-things/internal/ui"
import "github.com/senaduka/buttons-of-things/internal/server"

func main() {
	fmt.Println("Starting on localhost:8080.")
	server.Start("/", ui.Index)
}
