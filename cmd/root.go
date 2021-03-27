package cmd

import (
	"fmt"
	"os"

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
var buildURL bool

func mainFunc() {
	file := getFile()
	swagger := loadFile(file)
	apis := internal.ParsePaths(swagger.Paths)
	selected := internal.SelectAPI(apis)
	printer := &internal.SimplePrinter{}
	internal.PrintAPIDetails(selected, printer)
	if buildURL {
		ui := &internal.PromptUI{
			PromptRunnerFactory: internal.PromptRunnerFactory{},
		}
		url := internal.BuildURL(selected, ui)
		printer.Print(url)
	}
}

func getFile() string {
	if oasFileArgument != "" {
		return oasFileArgument
	}
	fromConf := readConfig()
	if fromConf == "" {
		fmt.Println("'-f' option is required if config file/option is not specified.")
		os.Exit(1)
	}
	return fromConf
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
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringVarP(&oasFileArgument, "file", "f", "", "OAS3 file path")
	rootCommand.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default: $HOME/.coas/config.yaml)")
	rootCommand.PersistentFlags().BoolVarP(&buildURL, "build", "b", false, "if specified, builds url interactively")
}
