package cli

import (
    "os"
    "fmt"
    "flag"
    "project_init/project_init/utils"
)


func Cli(config *utils.Config) {
    flag.StringVar(&config.Language, "lang", "python", "Programming Language to the project")
    flag.StringVar(&config.License, "license", "gpl-3", "License to use (MIT or GPL-3)")
    flag.BoolVar(&config.GitIgnore, "gitignore", false, "Include a .gitignore file")

    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] project_name \n", os.Args[0])
        fmt.Println("\nOptions:")
        flag.PrintDefaults()
    }

    flag.Parse()

    var args []string = flag.Args()
    if len(args) < 1 {
        fmt.Println("Error: project_name is missing")
        flag.Usage()
        os.Exit(1)
    }

    if args[0] == "version" {
        fmt.Println("version 2.3")
        os.Exit(0)
    }

    config.Name = args[0]
}

