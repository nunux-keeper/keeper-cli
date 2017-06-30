package common

import (
	"html/template"
	"io"

	"github.com/nunux-keeper/keeper-cli/api"
)

var profileTmpl = `Profile:
 UID:   {{.Uid}}
 Name:  {{.Name}}
 Date:  {{.Date}}
 Admin: {{.Admin}}
`

var documentTmpl = `Document:
 Id:          {{.Id}}
 Title:       {{.Title}}
 ContentType: {{.ContentType}}
 Content:     {{.Content}}
 Origin:      {{.Origin}}
 Date:        {{.Date}}
 Ghost:       {{.Ghost}}
`

var labelTmpl = `Label:
 Id:    {{.Id}}
 Label: {{.Label}}
 Color: {{.Color}}
 Date:  {{.Date}}
 Ghost: {{.Ghost}}
`

func WriteProfile(profile *api.ProfileResponse, out io.Writer) error {
	tmpl, err := template.New("profile").Parse(profileTmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(out, profile)
}

func WriteDocument(doc *api.DocumentResponse, out io.Writer) error {
	tmpl, err := template.New("document").Parse(documentTmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(out, doc)
}

func WriteLabel(label *api.LabelResponse, out io.Writer) error {
	tmpl, err := template.New("label").Parse(labelTmpl)
	if err != nil {
		return err
	}
	return tmpl.Execute(out, label)
}
