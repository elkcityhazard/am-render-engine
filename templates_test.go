package amrenderengine

import (
	"testing"

	"github.com/elkcityhazard/am-render-engine/mocks"
)

func Test_NewTemplateCollection(t *testing.T) {

	mockTC := NewTemplateCollection(mocks.GetMockTemplateFS(), "./mocks/templates")

	if mockTC == nil {
		t.Error("expected mock template collection to not be nil")
	}

}

func Test_NewTemplateData(t *testing.T) {

	mockTD := NewTemplateData()

	if mockTD == nil {
		t.Error("expected mock template data to not be nil")
	}
}

func Test_NewTemplateForm(t *testing.T) {

	mockTF := NewTemplateForm()

	if mockTF == nil {
		t.Error("expected mock template form to not be nil")
	}
}

func Test_SetDirPaths(t *testing.T) {

	mockTC := NewTemplateCollection(mocks.GetMockTemplateFS(), "./mocks/templates")

	mockTC.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")

	if mockTC.TemplatePageDirPath == "" {
		t.Error("expected mockTC TemplatePageDirPath to be templates/pages but got", mockTC.TemplatePageDirPath)
	}
	if mockTC.TemplateLayoutDirPath == "" {
		t.Error("expected mockTC TemplateLayoutDirPath to be templates/layouts but got", mockTC.TemplateLayoutDirPath)
	}
	if mockTC.TemplatePartialDirPath == "" {
		t.Error("expected mockTC TemplatePartialDirPath to be templates/partials but got", mockTC.TemplatePartialDirPath)
	}
	if mockTC.FileExt == "" {
		t.Error("expected mockTC FileExt to be gohtml but got", mockTC.TemplatePageDirPath)
	}

}

func Test_CreateTemplateCache(t *testing.T) {
	mockTC := NewTemplateCollection(mocks.GetMockTemplateFS(), "pkg/templates/mocks/templates")

	mockTC.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")

	tc, err := mockTC.CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

	if tc == nil {
		t.Error("expected non-nil template cache")
	}

	mockTC.TemplatePageDirPath = ""
	mockTC.TemplateLayoutDirPath = ""
	mockTC.TemplatePartialDirPath = ""
	mockTC.FileExt = ""

	tc, err = mockTC.CreateTemplateCache()

	if err == nil {
		t.Error("Expected an error due to missing paths, but got nil")
	}

	mockNoCacheTC := NewTemplateCollection(mocks.GetMockTemplateFS(), "pkg/templates/mocks/templates")
	mockNoCacheTC.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")
	mockNoCacheTC.UseCache = false

	noCacheTC, err := mockNoCacheTC.CreateTemplateCache()

	if err != nil {
		t.Error(err)
	}

	if noCacheTC == nil {
		t.Error("expected non-nil template cache")
	}

}

func Test_BuildTemplateCacheFromHost(t *testing.T) {
	mockTC := NewTemplateCollection(mocks.GetMockTemplateFS(), "pkg/mocks/templates")
	mockTC.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")
	mockTC.UseCache = false

	tc, err := mockTC.BuildTemplateCacheFromHost()

	if err != nil {
		t.Error("expected nil error but got", err.Error())
	}

	if tc == nil {
		t.Error("expected template cache to not be nil")
	}

}
