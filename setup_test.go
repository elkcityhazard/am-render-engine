package amrenderengine

import (
	"os"
	"testing"

	"github.com/elkcityhazard/am-render-engine/mocks"
)

var mockTC *TemplateCollection

func TestMain(m *testing.M) {

	mockTC = NewTemplateCollection(mocks.GetMockTemplateFS(), "./mocks/templates/")
	mockTC.SetIsProduction(false)
	mockTC.SetStringMapEntry("Test", "hello world")
	mockTC.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")
	mockTC.UseCache = true

	mockTC.CreateTemplateCache()

	os.Exit(m.Run())
}

func GetRenderer() *TemplateCollection {
	return mockTC
}
