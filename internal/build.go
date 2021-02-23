package internal

import (
	"fmt"
	"strings"

	"github.com/ytakahashi/burl/api"
)

// InputHandler handles user input
type InputHandler interface {
	readInput(parameter api.Parameter) string
}

// BuildURL builds url from an api
func BuildURL(target api.API, handler InputHandler) string {
	fmt.Println(target.ToText())
	path := target.Path
	for _, p := range target.GetPathParameters() {
		result := handler.readInput(p)
		path = replacePathParam(path, p, result)
	}

	query := ""
	for _, p := range target.GetQueryParameters() {
		result := handler.readInput(p)
		query = appendQueryString(query, p, result)
	}

	var url string
	if query == "" {
		url = path
	} else {
		url = fmt.Sprintf("%s?%s", path, query)
	}
	return url
}

func replacePathParam(path string, param api.Parameter, value string) string {
	target := fmt.Sprintf("{%s}", param.Name)
	return strings.Replace(path, target, value, 1)
}

func appendQueryString(query string, param api.Parameter, value string) string {
	if value == "" {
		return query
	}
	q := fmt.Sprintf("%s=%s", param.Name, value)
	if query == "" {
		return q
	}
	return fmt.Sprintf("%s&%s", query, q)
}
