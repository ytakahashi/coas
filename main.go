package main

import (
	"encoding/json"
	"fmt"
	"os"

	openapi3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/ytakahashi/api-builder/api"
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
}
