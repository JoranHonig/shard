package common

import (
	"bytes"
	"text/template"
)

type Issue struct {
	Title       string
	Description string
	Function    string
	Type        string
	Address     string
	Debug       string
}

func (i *Issue) String() string {
	templ, err := template.New("IssueTemplate").Parse(
		"== {{.Title}} ==\n" +
			"Function: {{.Function}} \n" +
			"Type: {{.Type}} \n" +
			"Description: \n" +
			"{{.Description}}")
	if err != nil {
		return ""
	}
	var output bytes.Buffer
	if err := templ.Execute(&output, i); err != nil {
		return ""
	}
	return output.String()
}
