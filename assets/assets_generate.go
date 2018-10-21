// +build ignore

package main

import (
	"log"

	"github.com/senaduka/buttons-of-things/assets"
	"github.com/shurcooL/vfsgen"
)

func main() {
	err := vfsgen.Generate(assets.ConfigFolder, vfsgen.Options{
		PackageName:  "assets",
		BuildTags:    "!dev",
		VariableName: "ConfigFolder",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
