package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type TemplateData struct {
	Title string
	Name  string
}

var rootDir, _ = os.Getwd()
var templatePath = filepath.Join(rootDir, "/templates/*")

var myTemplates = template.Must(template.ParseGlob(templatePath))

func serve(res http.ResponseWriter, req *http.Request) {

	myVars := TemplateData{"My Website Title", "My Website Heading111"}
	myTemplates.Execute(res, myVars)

}

func main() {
	fmt.Println(rootDir)
	http.HandleFunc("/", serve)
	http.ListenAndServe(":3000", nil)
}
