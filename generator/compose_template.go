package generator

var composeTemplate = `package compose

import (
	businessClient "github.com/rancher/types/client/cloud/v3"
	managementClient "github.com/rancher/types/client/management/v3"
)

type Config struct {
	Version string %BACK%yaml:"version,omitempty"%BACK%

	// Management Client
	{{range .managementSchemas}}
    {{- if . | hasPost }}{{.CodeName}}s map[string]managementClient.{{.CodeName}} %BACK%json:"{{.PluralName}},omitempty" yaml:"{{.PluralName}},omitempty"%BACK%
{{end}}{{end}}

	// Business Client
	{{range .businessSchemas}}
	{{- if . | hasGet }}{{.CodeName}}s map[string]businessClient.{{.CodeName}} %BACK%json:"{{.PluralName}},omitempty" yaml:"{{.PluralName}},omitempty"%BACK%
{{end}}{{end}}
`