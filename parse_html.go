package main

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

func main() {
	data := struct {
		ClientUserName string
	}{
		ClientUserName: "Verry",
	}

	s, err := ParseHTML(data, "./test.html")
	fmt.Println("err", err)
	fmt.Println(s)
}

func ParseHTML(data interface{}, filePath string) (result string, err error) {
	templatePath, err := filepath.Abs(filePath)
	if err != nil {
		return
	}

	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return
	}

	buf := new(bytes.Buffer)
	err = tpl.Execute(buf, data)
	if err != nil {
		return
	}

	return buf.String(), nil
}
