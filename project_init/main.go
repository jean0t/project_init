package main

import (
    "project_init/project_init/utils/handler"
    "project_init/project_init/utils"
    "project_init/project_init/cli"
    "fmt"
    "os"
)


func main() {
    var projectConfig *utils.Config = new(utils.Config)
    cli.Cli(projectConfig)

    langHandler, err := handler.GetLanguageHandler(projectConfig)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    if err := langHandler.Scaffold(); err != nil {
        fmt.Println("error: ", err)
        os.Exit(1)
    }


    fmt.Printf("Project %s with %s language created successfully\n", projectConfig.Name, projectConfig.Language)
}
