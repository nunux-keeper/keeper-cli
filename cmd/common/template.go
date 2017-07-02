package common

import (
	"encoding/json"
	"os"
	"text/template"
)

const VERSION = `Client:
 Version:      {{.Client.Version}}
 API version:  {{.Client.APIVersion}}
 Go version:   {{.Client.GoVersion}}
 OS/Arch:      {{.Client.Os}}/{{.Client.Arch}}

Server:
 Version:      {{.Server.Version}}
 API version:  {{.Server.APIVersion}}
`

var SERVER_INFOS = `Server informations:
 Nb. users     {{.NbUsers}}
 Nb. documents {{.NbDocuments}}
`

var PROFILE = `Profile:
 UID:   {{.Uid}}
 Name:  {{.Name}}
 Date:  {{.Date}}
 Admin: {{.Admin}}
`

var DOCUMENT = `Document:
 Id:          {{.Id}}
 Title:       {{.Title}}
 ContentType: {{.ContentType}}
 Content:     {{.Content}}
 Origin:      {{.Origin}}
 Date:        {{.Date}}
 Ghost:       {{.Ghost}}
`

var LABEL = `Label:
 Id:    {{.Id}}
 Label: {{.Label}}
 Color: {{.Color}}
 Date:  {{.Date}}
 Ghost: {{.Ghost}}
`

var USER = `User:
 Id:            {{.Id}}
 UID:           {{.Uid}}
 Name:          {{.Name}}
 Date:          {{.Date}}
 Nb. documents: {{.NbDocuments}}
 Nb. labels:    {{.NbLabels}}
 Nb. sharing:   {{.NbSharing}}
 Storage usage: {{.StorageUsage}}
`

var JOBS_INFO = `Jobs informations:
 Nb. inactive  {{.InactiveCount}}
 Nb. complete  {{.CompleteCount}}
 Nb. active    {{.ActiveCount}}
 Nb. failed    {{.FailedCount}}
 Work time     {{.WorkTime}}
`

var JOB = `Job:
 Id:         {{.Id}}
 Type:       {{.Type}}
 Priority:   {{.Priority}}
 Progress:   {{.Progress}}
 State:      {{.State}}
 Created at  {{.CreatedAt}}
 Updated at: {{.UpdatedAt}}
 Duration:   {{.Duration}}
 Params:     {{.Data}}
`

func WriteCmdResponse(v interface{}, tmpl string, outputAsJson bool) error {
	if outputAsJson {
		return WriteCmdJsonResponse(v)
	}
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, v)
}

func WriteCmdJsonResponse(v interface{}) error {
	return json.NewEncoder(os.Stdout).Encode(v)
}
