package handler

import (
    "project_init/project_init/utils"
    "time"
    _ "embed"
)

//go:embed templates/licenses/mit.txt
var MitTmpl string

//go:embed templates/licenses/gpl_3.txt
var Gpl3Tmpl string

//go:embed templates/readme.md.tmpl
var ReadmeTmpl string

//========================================= Gitignore template

//go:embed templates/go/gitignore.tmpl
var GoGitIgnoreTmpl string

//go:embed templates/python/gitignore.tmpl
var PythonGitIgnoreTmpl string

//go:embed templates/javascript/gitignore.tmpl
var JavascriptGitIgnoreTmpl string

//go:embed templates/java/gitignore.tmpl
var JavaGitIgnoreTmpl string

//go:embed templates/c/gitignore.tmpl
var CGitIgnoreTmpl string

//========================================= Sample Codes

//go:embed templates/python/sample_code.tmpl
var PythonSampleCode string

//go:embed templates/go/sample_code.tmpl
var GoSampleCode string

//go:embed templates/javascript/sample_code.tmpl
var JavascriptSampleCode string

//go:embed templates/java/sample_code.tmpl
var JavaSampleCode string

//go:embed templates/c/sample_code.tmpl
var CSampleCode string

//==========================================

type Data struct {
    Author string
    Year int
    Name string // name of project
    License string
}

func CreateData(name, license string) Data {
    year := time.Now().Year()
    author := utils.GetGitUsername()

    return Data{Author: author, Year: year, Name: name, License: license}
}
