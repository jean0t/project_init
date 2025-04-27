package main

import (
    "project_init/project_init/utils"
    "project_init/project_init/cli"
    "fmt"
    "os"
)


func main() {
    var projectConfig *utils.Config = new(utils.Config)
    cli.Cli(projectConfig)

    langHandler, err := utils.GetLanguageHandler(projectConfig.Language)
    if err != nil {
        fmt.Println(err)
    }

    err = langHandler.CreateDirectories(projectConfig.Name)
    if err != nil {
        fmt.Println("Error creating directories: ", err)
        os.Exit(1)
    }

    err = langHandler.AddSampleFiles(projectConfig.Name)
    if err != nil {
        fmt.Println("Error creating sample files: ", err)
        os.Exit(1)
    }

    err = langHandler.AddLicense(projectConfig.Name, projectConfig.License)
    if err != nil {
        fmt.Println("Error creating LICENSE: ", err)
        os.Exit(1)
    }
    fmt.Printf("%s/%s created\n", projectConfig.Name, "LICENSE")

    err = langHandler.AddGitIgnore(projectConfig.Name)
    if err != nil {
        fmt.Println("Error creating .gitignore: ", err)
        os.Exit(1)
    }
    fmt.Printf("%s/%s created\n", projectConfig.Name, ".gitignore")

    if utils.GitAvailable() {
        utils.GitInit(projectConfig.Name)
        fmt.Println("Git was initialized")
    }

    fmt.Printf("Project %s with %s language created successfully\n", projectConfig.Name, projectConfig.Language)
}
