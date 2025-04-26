package utils

import "fmt"

type Config struct {
    Name string
    Language string
    License string
    GitIgnore bool
}

type LanguageInterface interface {
    CreateDirectories(projectName string) error
    AddSampleFiles(projectName string) error
    AddLicense() error
    AddGitIgnore() error
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

