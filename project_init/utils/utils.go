package utils

import (
    "fmt"
    "os/exec"
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
