package main

import (
    "project_init/project_init/utils"
    "project_init/project_init/cli"
    "fmt"
)


func main() {
    var projectConfig *utils.Config = new(utils.Config)
    cli.Cli(projectConfig)

    fmt.Printf("Language: %s\n", projectConfig.Language)
}
