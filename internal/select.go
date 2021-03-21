package internal

import (
	"fmt"
	"log"

	text "github.com/MichaelMure/go-term-text"
	"github.com/ktr0731/go-fuzzyfinder"
)

// SelectAPI returns an api in a given array.
func SelectAPI(apis []API) API {
	index, err := fuzzyfinder.Find(
		apis,
		func(i int) string {
			return apis[i].ToText()
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return buildPreviewText(&apis[i], w)
		}))
	if err != nil {
		log.Fatal(err)
	}
	return apis[index]
}

func buildPreviewText(api *API, width int) string {
	preview := fmt.Sprintf("OperationID: %s\n", api.OperationID)

	if api.Description != "" {
		description, _ := text.Wrap(api.Description, width/2-5)
		preview += "\n" + description
	}
	preview += api.PrintParameters()
	return preview
}
