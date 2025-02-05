package amrenderengine

import (
	"bytes"
	"errors"
	"html/template"
	"net/http"
)

func (tmpCol *TemplateCollection) RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string) error {

	tc := make(map[string]*template.Template)

	if tmpCol.UseCache {
		tc = tmpCol.Templates
	} else {
		tmps, err := tmpCol.CreateTemplateCache()

		if err != nil {
			return err
		}

		tc = tmps
	}

	t, ok := tc[tmpl]

	if !ok {
		return errors.New("could not get template from template cache")
	}

	buf := &bytes.Buffer{}

	if tmpCol.TemplateData == nil {
		tmpCol.TemplateData = NewTemplateData()
	}
	//AddTD with r and td here

	err := t.Execute(buf, tmpCol.TemplateData)

	if err != nil {
		return err
	}

	mimetype, err := tmpCol.SetContentType(buf)

	if len(mimetype) > 0 {
		w.Header().Set("Content-Type", mimetype)
	} else {
		w.Header().Set("Content-Type", "text/html")
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		return err
	}
	return nil

}

func (tmpCol *TemplateCollection) AddCSRFToken(csrfToken string) {
	tmpCol.TemplateData.CSRFToken = csrfToken
}

func (tmpCol *TemplateCollection) SetIsProduction(isProduction bool) {
	tmpCol.TemplateData.IsProduction = isProduction
}

func (tmpCol *TemplateCollection) SetIsAuthenticated(isAuthenticated bool) {
	tmpCol.TemplateData.IsAuthenticated = isAuthenticated
}

func (tmpCol *TemplateCollection) SetStringMapEntry(key, val string) {
	tmpCol.TemplateData.StringMap[key] = val
}

func (tmpCol *TemplateCollection) SetIntMapEntry(key string, val int) {
	tmpCol.TemplateData.IntMap[key] = val
}

func (tmpCol *TemplateCollection) SetDataMapEntry(key string, val any) {
	tmpCol.TemplateData.Data[key] = val
}
