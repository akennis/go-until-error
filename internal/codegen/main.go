package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func makeGenArgs(uppers string) string {
	res := strings.Builder{}
	for i := 0; i < len(uppers); i++ {
		res.WriteString(string(uppers[i]))
		if i != len(uppers)-1 {
			res.WriteString(", ")
		}
	}
	return res.String()
}

func makeSigArgs(lowers string, uppers string) string {
	res := strings.Builder{}
	for i := 0; i < len(uppers); i++ {
		res.WriteString(string(lowers[i]))
		res.WriteString(" ")
		res.WriteString(string(uppers[i]))
		if i != len(uppers)-1 {
			res.WriteString(", ")
		}
	}
	return res.String()
}

func makeCallArgs(lowers string) string {
	res := strings.Builder{}
	for i := 0; i < len(lowers); i++ {
		res.WriteString(string(lowers[i]))
		if i != len(lowers)-1 {
			res.WriteString(", ")
		}
	}
	return res.String()
}

func codegen() (string, error) {
	gueTemplateSrc :=
		`
func Gue{{.NumArgs}}[{{if .HasArgs}}{{.GenArgs}},{{end}} RT any](f func({{.GenArgs}}) (RT, error), e *error) func({{.GenArgs}}) RT {
	return func({{.SigArgs}}) RT {
		if *e != nil {
			var zVal RT
			return zVal
		}
		rv, err := f({{.CallArgs}})
		if err != nil {
			*e = err
		}
		return rv
	}
}

func Gue0{{.NumArgs}}{{if .HasArgs}}[{{.GenArgs}} any]{{end}}(f func({{.GenArgs}}) error, e *error) func({{.GenArgs}}) {
	return func({{.SigArgs}}) {
		if *e != nil {
			return
		}
		err := f({{.CallArgs}})
		if err != nil {
			*e = err
		}
	}
}
	`
	gueTemplate, err := template.New("gue").Parse(gueTemplateSrc)
	if err != nil {
		return "", err
	}
	lowers := "rstuvwxyz" // end of alphabet to avoid e and f which are used elsewhere
	uppers := strings.ToUpper(lowers)
	type TemplateArgs struct {
		HasArgs  bool
		NumArgs  int
		GenArgs  string
		SigArgs  string
		CallArgs string
	}
	res := strings.Builder{}
	for i := 0; i <= len(lowers); i++ {
		templateArgs := TemplateArgs{
			HasArgs:  i > 0,
			NumArgs:  i,
			GenArgs:  makeGenArgs(uppers[:i]),
			SigArgs:  makeSigArgs(lowers[:i], uppers[:i]),
			CallArgs: makeCallArgs(lowers[:i]),
		}
		var buf bytes.Buffer
		gueTemplate.Execute(&buf, templateArgs)
		res.Write(buf.Bytes())
	}
	return res.String(), nil
}

func main() {
	src, err := codegen()
	if err != nil {
		panic(err)
	}
	fmt.Println("package gue")
	fmt.Println(src)
}
