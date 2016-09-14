package main

import (
	"bytes"
	"text/template"
	"fmt"
	"os"
	"io/ioutil"
	"os/user"
)

func main() {
	tmplContent, err := getTemplate(os.Args[1]);
	if err != nil {
		panic(err)
	}
	tmpl, _ := template.New("tmpl").Parse(tmplContent)
	var out bytes.Buffer
	tmpl.Execute(&out, os.Args[2:])
	fmt.Println(out.String())
}

func getTemplate(name string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	dir := usr.HomeDir
	content, err := ioutil.ReadFile(dir + "/.thrombosis-of-the-grainmill/" + name)
	return string(content), err
}
