package internal

import (
	"log"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/ytakahashi/coas/api"
)

// SelectAPI returns an api in a given array.
func SelectAPI(apis []api.API) api.API {
	index, err := fuzzyfinder.Find(
		apis,
		func(i int) string {
			return apis[i].ToText()
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return apis[i].BuildDetailedDescription(w)
		}))
	if err != nil {
		log.Fatal(err)
	}
	return apis[index]
}
