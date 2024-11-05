package templating

import (
	"bytes"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/matthewchivers/dodl/pkg/dateutils"
)

// RenderTemplate renders a template string using the provided data map.
// It returns the rendered string or an error if rendering fails.
func RenderTemplate(templateStr string, data map[string]interface{}, date time.Time) (string, error) {

	data["Now"] = date
	data["Today"] = date
	data["WeekStart"] = dateutils.GetDefaultWeekStartDate(date)

	funcMap := sprig.TxtFuncMap()
	funcMap = addCustomFuncs(funcMap)

	tmpl, err := template.New("template").Funcs(funcMap).Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
