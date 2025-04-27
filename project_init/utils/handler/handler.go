package handler


import (
    "fmt"
    "strings"
    "project_init/project_init/utils"
)

type Handler interface {
    Scaffold() error
}

func GetLanguageHandler(cfg *utils.Config) (Handler, error) {
    switch strings.ToLower(cfg.Language) {
        case "go":
            return NewGoHandler(cfg), nil

        case "python":
            return NewPythonHandler(cfg), nil

        default:
            return nil, fmt.Errorf("Unsupported language %q", cfg.Language)
    }
}
