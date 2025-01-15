package mocks

import "embed"

//go:embed templates

var mockTemplateFS embed.FS

func GetMockTemplateFS() embed.FS {
	return mockTemplateFS
}
