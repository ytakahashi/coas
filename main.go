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

	res := api.ParsePaths(swagger.Paths)

	j, _ := json.Marshal(res)
	fmt.Print(string(j))
}
