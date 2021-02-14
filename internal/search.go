package internal

import (
	"fmt"
	"log"
	"strings"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/ytakahashi/api-builder/api"
)

// Search returns an api in a given array.
func Search(apis []api.API) api.API {
	index, err := fuzzyfinder.Find(
		apis,
		func(i int) string {
			return apis[i].ToText()
		},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}

			param := ""
			if apis[i].Parameters != nil {
				paramTexts := []string{}
				for _, p := range apis[i].Parameters {
					paramTexts = append(paramTexts, p.ToText())
				}
				param = fmt.Sprintf(`
				Parameters:
				- %s`,
					strings.Join(paramTexts, "\n- "),
				)
			}
			preview := fmt.Sprintf(
				`
				OperationID: %s
				%s
				`,
				apis[i].OperationID,
				apis[i].Description,
			)

			if param != "" {
				preview += param
			}
			return preview
		}))
	if err != nil {
		log.Fatal(err)
	}
	return apis[index]
}
