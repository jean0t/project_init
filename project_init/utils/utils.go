package utils

import (
    "fmt"
    "os/exec"
    "path/filepath"
    "os"
)

type Config struct {
    Name string
    Language string
    License string
    GitIgnore bool
}


type LanguageInterface interface {
    CreateDirectories(projectName string) error
    AddSampleFiles(projectName string) error
    AddLicense(projectName, license string) error
    AddGitIgnore(projectName string) error
}


func GetLanguageHandler(language string) (LanguageInterface, error) {
    switch language {
        case "python":
            return &Python{}, nil

        case "go":
            return &Go{}, nil

        default:
            return nil, fmt.Errorf("Unsupported Language: %s", language)
    }
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
