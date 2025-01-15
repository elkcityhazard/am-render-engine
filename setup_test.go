package templates

import (
	"os"
	"testing"

	"github.com/elkcityhazard/git-line-counter/pkg/templates/mocks"
)

var mockTC *TemplateCollection

func TestMain(m *testing.M) {

	mockTC = NewTemplateCollection(mocks.GetMockTemplateFS(), "./pkg/templates/mocks/templates")
	mockTC.SetDirPaths("templates/pages", "templates/layouts", "templates/partials", "gohtml")
	mockTC.UseCache = true

	os.Exit(m.Run())
}
