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
