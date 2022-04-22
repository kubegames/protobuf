package main

import (
	"bytes"
	_ "embed"
	"html/template"
	"strings"
)

//go:embed template.go.tpl
var tpl string

type Message struct {
	Name string
}

func (m *Message) Execute() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("xorm").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, m); err != nil {
		panic(err)
	}
	return buf.String()
}
