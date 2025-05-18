package cli

import (
    "os"
    "fmt"
    "flag"
    "project_init/project_init/utils"
)

var version bool

func Cli(config *utils.Config) {
    flag.StringVar(&config.Language, "lang", "python", "Programming Language to the project")
    flag.StringVar(&config.License, "license", "gpl-3", "License to use (MIT or GPL-3)")
    flag.BoolVar(&config.GitIgnore, "gitignore", false, "Include a .gitignore file")
    flag.BoolVar(&version, "v", false, "Outputs the version of the program")

    flag.Usage = func() {
        fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] project_name \n", os.Args[0])
        fmt.Fprintf(flag.CommandLine.Output(), "Languages supported: python, java, javascript, c, go\n")
        fmt.Println("\nOptions:")
        flag.PrintDefaults()
    }

    flag.Parse()

    if version {
        fmt.Println("version 2.4")
        os.Exit(0)
    }

    var args []string = flag.Args()
    if len(args) < 1 {
        fmt.Println("Error: project_name is missing")
        flag.Usage()
        os.Exit(1)
    }


    config.Name = args[0]
}

