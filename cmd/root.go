package cmd

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/spf13/cobra"
	"github.com/ytakahashi/coas/internal"
)

var rootCommand = &cobra.Command{
	Use:   "coas",
	Short: "Command line tool to see OAS3 file.",
	Run: func(cmd *cobra.Command, args []string) {
		mainFunc()
	},
}

func Execute() {
	rootCommand.Execute()
}

var oasFileArgument string

func mainFunc() {
	file := getFile()
	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromFile(file)
	if err != nil {
		panic(err)
	}

	apis := internal.ParsePaths(swagger.Paths)
	selected := internal.SelectAPI(apis)
	url := internal.BuildURL(selected, new(internal.PromptUI))
	fmt.Println(url)
}

func getFile() string {
	if oasFileArgument == "" {
		panic("error file")
	}
	return oasFileArgument
}

func loadFile(oasFile string) *openapi3.Swagger {
	loader := openapi3.NewSwaggerLoader()
	loader.IsExternalRefsAllowed = true
	swagger, err := loader.LoadSwaggerFromFile(oasFile)
	if err != nil {
		panic(err)
	}
	return swagger
}

func init() {
	rootCommand.PersistentFlags().StringVarP(&oasFileArgument, "file", "f", "", "OAS3 file path")
}
