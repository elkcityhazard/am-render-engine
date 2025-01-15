package amrenderengine

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_RenderTemplate(t *testing.T) {

	tests := []struct {
		name     string
		useCache bool
		pageName string
		want     string
		got      string
	}{
		{
			name:     "uses cache",
			useCache: true,
			pageName: "test.gohtml",
			want:     "Content Here",
			got:      "",
		},
		{
			name:     "do not use cache",
			useCache: false,
			pageName: "test.gohtml",
			want:     "Content Here",
			got:      "",
		},
		{
			name:     "bad page name",
			useCache: false,
			pageName: "nonexist.gohtml",
			want:     "Content Here",
			got:      "",
		},
		{
			name:     "nil template data",
			useCache: true,
			pageName: "test.gohtml",
			want:     "Content Here",
			got:      "",
		},
	}

	tc, err := mockTC.CreateTemplateCache()

	if err != nil {
		t.Error("Expected no error but got", err.Error())
	}

	if tc == nil {
		t.Error("Expected template cache to not be nil")
	}

	for _, tt := range tests {
		var mockHandler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			err := mockTC.RenderTemplate(w, r, tt.pageName)

			if err != nil && tt.pageName != "nonexist.gohtml" {
				t.Error("expected nil error but got", err.Error())
			}

			if err != nil && tt.pageName == "nonexist.gohtml" {
				if !strings.Contains(err.Error(), "could not") {
					t.Error("expected an error message for malformed template name")
				}
			}

		})

		ts := httptest.NewServer(mockHandler)

		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "nil template data" {
				mockTC.TemplateData = nil
			}
			mockTC.UseCache = tt.useCache
			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", ts.URL, nil)

			mockHandler.ServeHTTP(w, req)

			result := w.Result()

			defer result.Body.Close()

			body, err := io.ReadAll(result.Body)

			if err != nil && tt.pageName != "nonexist.gohtml" {

				t.Error("expected a body but got", err.Error())
			}

			if !strings.Contains(string(body), "Content Here") && tt.pageName != "nonexist.gohtml" {
				t.Error("expected content on page, but got none")
			}

		})
	}

}

func Test_HelperFunctions(t *testing.T) {

	mockTC.AddCSRFToken("CSRFToken")

	if !strings.EqualFold(mockTC.TemplateData.CSRFToken, "CSRFToken") {
		t.Error("expected CSRFToken to be equal")
	}

	mockTC.SetIsProduction(true)

	if !mockTC.TemplateData.IsProduction {
		t.Error("expected mockTC.TemplateData.IsProduction to be true but got", mockTC.TemplateData.IsProduction)
	}

	mockTC.SetIsAuthenticated(true)

	if !mockTC.TemplateData.IsAuthenticated {
		t.Error("expected mockTC.TemplateData.IsAuthenticated to be true but got", mockTC.TemplateData.IsAuthenticated)
	}

	mockTC.SetStringMapEntry("test", "value")

	if !strings.EqualFold(mockTC.TemplateData.StringMap["test"], "value") {
		t.Error("expected a value for StringMap key")
	}

	mockTC.SetIntMapEntry("test", 1)

	if mockTC.TemplateData.IntMap["test"] != 1 {
		t.Error("expected int map key test to equal 1")
	}

	type mockUser struct {
		name string
	}

	fakeUser := mockUser{
		name: "coolgal123",
	}

	mockTC.SetDataMapEntry("user", fakeUser)

	retrieved := mockTC.TemplateData.Data["user"].(mockUser)

	if !strings.EqualFold(retrieved.name, "coolgal123") {
		t.Error("expected a user but got none")
	}
}
