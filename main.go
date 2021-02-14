package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ytakahashi/api-builder/api"
	"github.com/ytakahashi/api-builder/internal"
)

func main() {
	file := os.Args[1]
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(file)

	if err != nil {
		panic(err)
	}

	servers := api.ParseServers(swagger.Servers)
	apis := api.ParsePaths(swagger.Paths)

	serversJSON, _ := json.Marshal(servers)
	apisJSON, _ := json.Marshal(apis)
	fmt.Println(string(serversJSON))
	fmt.Println(string(apisJSON))
	fmt.Println()

	selected := internal.Search(apis)
	fmt.Println("Selected: ", selected.ToText())
}
