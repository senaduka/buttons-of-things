package main

import "fmt"
import "github.com/senaduka/buttons-of-things/internal/ui"
import "github.com/senaduka/buttons-of-things/internal/server"
import "github.com/gobuffalo/packr"

func main() {
  box := packr.NewBox("./assets")
	s := box.String("config/config.json")

	fmt.Println("Starting on localhost:8080.")
  fmt.Println(s)
	server.Start("/", ui.Index)
}
