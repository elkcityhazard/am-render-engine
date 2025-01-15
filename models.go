package templates

import (
	"embed"
	"html/template"
	"net/url"
)

type TemplateCollection struct {
	EmbeddedFS             embed.FS
	FileExt                string
	PathToTemplates        string
	Templates              map[string]*template.Template
	TemplateFuncs          template.FuncMap
	TemplateData           *TemplateData
	UseCache               bool
	TemplatePageDirPath    string
	TemplateLayoutDirPath  string
	TemplatePartialDirPath string
}

type TemplateData struct {
	StringMap       map[string]string
	BoolMap         map[string]bool
	IntMap          map[string]int
	Data            map[string]any
	IsProduction    bool
	IsAuthenticated bool
	CSRFToken       string
	Form            TemplateForm
	FlashMsg        string
	WarnMsg         string
	ErrorMsg        string
}

type TemplateForm struct {
	Form   url.Values
	Errors TemplateFormErrors
}

type TemplateFormErrors map[string][]string
