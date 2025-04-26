package main

import (
    "project_init/project_init/utils"
    "project_init/project_init/cli"
    "flag"
    "testing"
    "os"
    "reflect"
)


func TestCliFlags(t *testing.T) {

    var tests = []struct {
        name string
        args []string
        expected *utils.Config
    }{
        {
            name: "default arguments",
            args: []string {"./project_init", "myproject"},
            expected: &utils.Config{
                Name: "myproject",
                Language: "python",
                License: "gpl-3",
                GitIgnore: false,
            },
        },

        {
            name: "override all arguments",
            args: []string {"./project_init", "-lang", "go", "-license", "mit", "-gitignore", "myproject"},
            expected: &utils.Config{
                Name: "myproject",
                Language: "go",
                License: "mit",
                GitIgnore: true,
            },
        },

    }

    for _, test_case := range tests {
        t.Run(test_case.name, func(t *testing.T) {
            originalArgs := os.Args
            t.Cleanup(func() {
                os.Args = originalArgs
                flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
            })

            os.Args = test_case.args
            flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

            var got *utils.Config = new(utils.Config)
            cli.Cli(got)

            if !reflect.DeepEqual(got, test_case.expected) {
                t.Errorf("CLI produced %+v; want %+v", got, test_case.expected)
            }
        })
    }
}

