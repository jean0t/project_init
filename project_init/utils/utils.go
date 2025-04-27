package utils

import (
    "fmt"
    "os/exec"
    "path/filepath"
    "os"
    "strings"
    "project_init/project_init/utils/handler"
)

type Config struct {
    Name string
    Language string
    License string
    GitIgnore bool
}


func GitAvailable() bool {
    _, err := exec.LookPath("git")
    return err == nil
}


func GitInit(projectDir string) error {
    cmd := exec.Command("git", "init")
    cmd.Dir = projectDir
    return cmd.Run()
}

func GetGitUserName() string {
    if !GitAvailable() {
        return "<name>"
    }

    var cmd *exec.Cmd = exec.Command("git", "config", "--global", "user.name")
    output, err := cmd.Output()
    if err != nil {
        return "<name>"
    }

    var name string = strings.TrimSpace(string(output))
    if name == "" {
        return "<name>"
    }

    return name
}

func CreateReadme(projectName, license string) error {
    var readmeTemplate string = `# %s

A brief description of what the project does.

## Installation

Instructions to install or build.

## Usage

Examples of how to run the project.

## License

%s © <year> <your name>
`
    readmeFile := filepath.Join(projectName, "README.md")
    file, err := os.Create(readmeFile)
    if err != nil {
        return fmt.Errorf("It couldn't create the README.md file")
    }
    defer file.Close()

    _, err = file.WriteString(fmt.Sprintf(readmeTemplate, projectName, license))
    if err != nil {
        return fmt.Errorf("It couldn't write the README.md contents")
    }

    return nil
}
