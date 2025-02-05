package amrenderengine

import (
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_SetContentType(t *testing.T) {

	html := `
	<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8">
		<title>My First Webpage</title>
	  </head>
		<p>Hello World</p>
	  <body>
		<h1>Hello World!</h1><p>Hello World</p><p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<section>Hello world</section>
		<h1>Hello World!</h1><p>Hello World</p><p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<section>Hello world</section>
		<section>Hello world</section>
		<h1>Hello World!</h1><p>Hello World</p><p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<p>Hello World</p>
		<section>Hello world</section>
	  </body>
	</html>
	`

	ts := httptest.NewRequest("GET", "/", nil)

	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		buf := &bytes.Buffer{}

		tmpl, err := template.New("test").Parse(html)

		if err != nil {
			t.Error("expected no error but got", err.Error())
		}

		err = tmpl.Execute(buf, nil)

		if err != nil {
			t.Error(err)
		}

		mimetype, err := mockTC.SetContentType(buf)

		if err != nil {
			t.Error("expected no error but got", err.Error())
		}

		if !strings.Contains(mimetype, "text/html") {
			t.Error("expected text/html mimetype but got ", mimetype)
		}

		_, err = buf.WriteTo(w)

		w.Header().Set("Content-Type", mimetype)

		if err != nil {
			t.Error("expected no error but got", err.Error())
		}

	})

	recorder := httptest.NewRecorder()

	mockHandler.ServeHTTP(recorder, ts)

	if !strings.Contains(recorder.Result().Header["Content-Type"][0], "text/html") {
		t.Error("expected Content Type to contain text/html but got ", recorder.Result().Header)
	}
}
