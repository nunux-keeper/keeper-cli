package common

import (
	"encoding/json"
	"os"
	"text/template"
)

// VERSION Version output
const VERSION = `Client:
 Version:      {{.Client.Version}}
 API version:  {{.Client.APIVersion}}
 Go version:   {{.Client.GoVersion}}
 OS/Arch:      {{.Client.Os}}/{{.Client.Arch}}

Server:
 Version:      {{.Server.Version}}
 API version:  {{.Server.APIVersion}}
`

// SERVER_INFOS Server infos output
var SERVER_INFOS = `Server informations:
 Nb. users     {{.NbUsers}}
 Nb. documents {{.NbDocuments}}
`

// PROFILE Profile output
var PROFILE = `Profile:
 UID:   {{.Uid}}
 Name:  {{.Name}}
 Date:  {{.Date}}
 Admin: {{.Admin}}
`

// DOCUMENT Document output
var DOCUMENT = `Document:
 Id:          {{.Id}}
 Title:       {{.Title}}
 ContentType: {{.ContentType}}
 Content:     {{.Content}}
 Origin:      {{.Origin}}
 Date:        {{.Date}}
 Ghost:       {{.Ghost}}
`

// LABEL Label output
var LABEL = `Label:
 Id:    {{.Id}}
 Label: {{.Label}}
 Color: {{.Color}}
 Date:  {{.Date}}
 Ghost: {{.Ghost}}
`

// USER User output
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

// JOBS_INFO Job infos output
var JOBS_INFO = `Jobs informations:
 Nb. inactive  {{.InactiveCount}}
 Nb. complete  {{.CompleteCount}}
 Nb. active    {{.ActiveCount}}
 Nb. failed    {{.FailedCount}}
 Work time     {{.WorkTime}}
`

// JOB Job details output
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

// WEBHOOK Webhook details output
var WEBHOOK = `Webhook:
 Id:    {{.ID}}
 URL:   {{.URL}}
 Secret: {{.Secret}}
 Labels: {{.Labels}}
 Events: {{.Events}}
 Creation date:  {{.CreationDate}}
 Modification date:  {{.ModificationDate}}
`

// WriteCmdResponse Write command response as readable text
func WriteCmdResponse(v interface{}, tmpl string, outputAsJSON bool) error {
	if outputAsJSON {
		return WriteCmdJSONResponse(v)
	}
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(os.Stdout, v)
}

// WriteCmdJSONResponse Write command response as JSON output
func WriteCmdJSONResponse(v interface{}) error {
	return json.NewEncoder(os.Stdout).Encode(v)
}
