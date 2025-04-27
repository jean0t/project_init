package handler

import (
    "project_init/project_init/utils"
    "time"
    _ "embed"
)

//go:embed templates/licenses/mit.txt
var mitTmpl string

//go:embed templates/licenses/gpl_3.txt
var gpl3Tmpl string

//go:embed templates/readme.md.tmpl
var readmeTmpl string


type Data struct {
    Author string
    Year int
    Name string //name of the project
}

func CreateData(name string) Data {
    year := time.Now().Year()
    author := utils.GetGitUsername()

    return Data{Author: author, Year: year, Name: name}
}
