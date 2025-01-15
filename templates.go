package templates

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"path/filepath"
)

func NewTemplateCollection(emb embed.FS, tmpPath string) *TemplateCollection {
	return &TemplateCollection{
		EmbeddedFS:             emb,
		PathToTemplates:        tmpPath,
		Templates:              make(map[string]*template.Template),
		TemplateFuncs:          nil,
		TemplateData:           NewTemplateData(),
		UseCache:               true,
		TemplatePageDirPath:    "",
		TemplateLayoutDirPath:  "",
		TemplatePartialDirPath: "",
		FileExt:                "",
	}
}

func NewTemplateData() *TemplateData {
	return &TemplateData{
		StringMap:       make(map[string]string),
		BoolMap:         make(map[string]bool),
		IntMap:          make(map[string]int),
		Data:            make(map[string]any),
		IsProduction:    false,
		IsAuthenticated: false,
		CSRFToken:       "",
		Form:            *NewTemplateForm(),
		FlashMsg:        "",
		WarnMsg:         "",
		ErrorMsg:        "",
	}

}

func NewTemplateForm() *TemplateForm {
	return &TemplateForm{
		Form:   nil,
		Errors: make(TemplateFormErrors),
	}
}

func (tmpCol *TemplateCollection) SetDirPaths(pages, layouts, partials, ext string) {

	tmpCol.TemplatePageDirPath = pages
	tmpCol.TemplateLayoutDirPath = layouts
	tmpCol.TemplatePartialDirPath = partials
	tmpCol.FileExt = ext

}

func (tmpCol *TemplateCollection) CreateTemplateCache() (map[string]*template.Template, error) {

	arePathsSet := tmpCol.CheckForTemplatePathsAndExt()

	if !arePathsSet {
		return nil, errors.New("missing template data")
	}

	if tmpCol.UseCache {
		tc, err := tmpCol.BuildTemplateCacheFromEmbed()

		if err != nil {
			return nil, err
		}

		tmpCol.Templates = tc
	} else {
		tc, err := tmpCol.BuildTemplateCacheFromHost()

		if err != nil {
			return nil, err
		}

		tmpCol.Templates = tc

	}

	return tmpCol.Templates, nil

}

func (tmpCol *TemplateCollection) BuildTemplateCacheFromEmbed() (map[string]*template.Template, error) {

	fs := tmpCol.EmbeddedFS

	pages, err := fs.ReadDir(tmpCol.TemplatePageDirPath)
	if err != nil {
		return nil, err
	}

	for i := range pages {

		tmpl, err := template.New(pages[i].Name()).Funcs(tmpCol.TemplateFuncs).ParseFS(fs, fmt.Sprintf("%s/%s", tmpCol.TemplatePageDirPath, pages[i].Name()))

		if err != nil {
			return nil, err
		}

		layouts, err := fs.ReadDir(tmpCol.TemplateLayoutDirPath)

		if err != nil {
			return nil, err
		}

		if len(layouts) > 0 {
			for _, l := range layouts {
				tmpl, err = tmpl.ParseFS(fs, fmt.Sprintf("%s/%s", tmpCol.TemplateLayoutDirPath, l.Name()))
				if err != nil {
					return nil, err
				}
			}
		}

		partials, err := fs.ReadDir(tmpCol.TemplatePartialDirPath)
		if err != nil {
			return nil, err
		}

		if len(partials) > 0 {
			for _, p := range partials {
				tmpl, err = tmpl.ParseFS(fs, fmt.Sprintf("%s/%s", tmpCol.TemplatePartialDirPath, p.Name()))
			}
		}

		tmpCol.Templates[pages[i].Name()] = tmpl

	}

	return tmpCol.Templates, nil

}

func (tmpCol *TemplateCollection) BuildTemplateCacheFromHost() (map[string]*template.Template, error) {

	pages, err := filepath.Glob(fmt.Sprintf("%s/%s/*.%s", tmpCol.PathToTemplates, tmpCol.TemplatePageDirPath, tmpCol.FileExt))

	if err != nil {
		return nil, err
	}

	for i := range pages {
		name := filepath.Base(pages[i])

		ts, err := template.New(name).Funcs(tmpCol.TemplateFuncs).ParseFiles(pages[i])

		if err != nil {
			return nil, err
		}

		layoutMatches, err := filepath.Glob(
			fmt.Sprintf("%s/%s/*.%s",
				tmpCol.PathToTemplates,
				tmpCol.TemplateLayoutDirPath,
				tmpCol.FileExt))

		if err != nil {
			return nil, err
		}

		if len(layoutMatches) > 0 {
			ts, err = ts.ParseGlob(
				fmt.Sprintf("%s/%s/*.%s",
					tmpCol.PathToTemplates,
					tmpCol.TemplateLayoutDirPath,
					tmpCol.FileExt))
			if err != nil {
				return nil, err
			}
		}

		partialMatches, err := filepath.Glob(
			fmt.Sprintf("%s/%s/*.%s",
				tmpCol.PathToTemplates,
				tmpCol.TemplatePartialDirPath,
				tmpCol.FileExt))

		if err != nil {
			return nil, err
		}

		if len(partialMatches) > 0 {
			ts, err = ts.ParseGlob(
				fmt.Sprintf("%s/%s/*.%s",
					tmpCol.PathToTemplates,
					tmpCol.TemplatePartialDirPath,
					tmpCol.FileExt))
			if err != nil {
				return nil, err
			}
		}

		tmpCol.Templates[name] = ts

	}

	return tmpCol.Templates, nil

}

// CheckForTemplatePathsAndExt make sure that the paths to the template directories are set in *TemplateCollection
// As well as a well as a FileExt
// it returns a bool
func (tmpCol *TemplateCollection) CheckForTemplatePathsAndExt() bool {
	if tmpCol.TemplateLayoutDirPath == "" || tmpCol.TemplatePageDirPath == "" || tmpCol.TemplatePartialDirPath == "" || tmpCol.FileExt == "" {
		return false
	}
	return true
}
