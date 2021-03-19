package main

import (
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ytakahashi/coas/internal"
)

func main() {
	file := os.Args[1]
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(file)
	if err != nil {
		panic(err)
	}

	apis := internal.ParsePaths(swagger.Paths)
	selected := internal.SelectAPI(apis)
	url := internal.BuildURL(selected, new(internal.PromptUI))
	fmt.Println(url)
}
