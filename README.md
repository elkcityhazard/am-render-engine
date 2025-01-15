# am-render-engine
A render engine for go templates

## File Structure

- html
	- files
		- layouts
			- base.gohtml
		- pages
			- home.gohtml
		- partials
			- partial.gohtml
	- html.go


## Setup
in html.go
```
package htmltemplates

import "embed"

//go:embed files

var fileFS embed.FS

func GetFileFS() embed.FS {
	return fileFS
}
```

## Setup Renderer

```
package main

import (
	"log"
	"net/http"
	htmltemplates "test-am-render/html"

	amrenderengine "github.com/elkcityhazard/am-render-engine"
)

type config struct {
	renderer amrenderengine.TemplateCollection
}

func main() {

	var cfg config

	cfg.renderer = *amrenderengine.NewTemplateCollection(htmltemplates.GetFileFS(), "./html/files")

	cfg.renderer.UseCache = true
	cfg.renderer.SetIsProduction(false)
	cfg.renderer.SetIsAuthenticated(false)
	cfg.renderer.SetStringMapEntry("Name", "coolbeans123")

	cfg.renderer.SetDirPaths("files/pages", "files/layouts", "files/partials", "gohtml")

	_, err := cfg.renderer.CreateTemplateCache()

	if err != nil {
		log.Fatalln(err)
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		err := cfg.renderer.RenderTemplate(w, r, "home.gohtml")

		if err != nil {
			log.Fatalln(err)
		}

	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	if err = srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

}
```


