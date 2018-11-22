package main

import (
	"bytes"
	"html/template"
)

const htmlTmpl = `
// GENERATED FILE -- DO NOT EDIT
//
//go:generate $GOPATH/src/istio.io/istio/galley/tools/gen-meta/gen-meta.sh kube pkg/metadata/kube/types.go
//

package kube

import (
	"istio.io/istio/galley/pkg/kube"
	"istio.io/istio/galley/pkg/kube/converter"
	"istio.io/istio/galley/pkg/metadata"
)

// Types in the schema.
var Types *kube.Schema

func init() {
	b := kube.NewSchemaBuilder()
{{range .Resources}}
	b.Add(kube.ResourceSpec{
		Kind:       "{{.Kind}}",
		ListKind:   "{{.ListKind}}",
		Singular:   "{{.Singular}}",
		Plural:     "{{.Plural}}",
		{{- if .ShortNamesStr }}
		ShortNames: {{.ShortNamesStr | raw}},
		{{- end}}
		Version:    "{{.Version}}",
		Group:      "{{.Group}}",
		Target:     metadata.Types.Get("type.googleapis.com/{{.Proto}}"),
		Converter:  converter.Get("{{ if .Converter }}{{.Converter}}{{ else }}identity{{end}}"),
    })
{{end}}
	Types = b.Build()
}
`

func applyHTMLTemplate(m *metadata) ([]byte, error) {
	t := template.New("tmpl")
	t = t.Funcs(template.FuncMap{"raw": raw})

	t2, err := t.Parse(htmlTmpl)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err = t2.Execute(&b, m); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// raw directly returns the string value and does not escape it.
func raw(value string) interface{} {
	return template.HTML(value)
}
