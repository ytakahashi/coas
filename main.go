package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ytakahashi/burl/api"
	"github.com/ytakahashi/burl/internal"
)

func main() {
	file := os.Args[1]
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(file)
	if err != nil {
		panic(err)
	}

	apis := api.ParsePaths(swagger.Paths)
	selected := internal.Search(apis)
	url := internal.BuildURL(selected, new(internal.PromptUI))
	fmt.Println(url)
}
