package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"strings"
)

//go:embed template.go.tpl
var tpl string

type service struct {
	Name      string // Greeter
	FullName  string // helloworld.Greeter
	FilePath  string // api/helloworld/helloworld.proto
	Methods   []*method
	MethodSet map[string]*method
}

func (s *service) execute() string {
	if s.MethodSet == nil {
		s.MethodSet = map[string]*method{}
		for _, m := range s.Methods {
			m := m
			s.MethodSet[m.Name] = m
		}
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(tpl))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return buf.String()
}

// InterfaceName service interface name
func (s *service) InterfaceName() string {
	return s.Name + "HTTPServer"
}

type method struct {
	Name         string
	Num          int
	Request      string
	Reply        string
	Path         string
	Method       string
	Body         string
	ResponseBody string
}

// HandlerName for gin handler name
func (m *method) HandlerName() string {
	return fmt.Sprintf("%s_%d", m.Name, m.Num)
}

// HasPathParams 是否包含路由参数
func (m *method) HasPathParams() bool {
	paths := strings.Split(m.Path, "/")
	for _, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			return true
		}
	}
	return false
}

// initPathParams 转换参数路由 {xx} --> :xx
func (m *method) initPathParams() {
	paths := strings.Split(m.Path, "/")
	for i, p := range paths {
		if len(p) > 0 && (p[0] == '{' && p[len(p)-1] == '}' || p[0] == ':') {
			paths[i] = ":" + p[1:len(p)-1]
		}
	}
	m.Path = strings.Join(paths, "/")
}
