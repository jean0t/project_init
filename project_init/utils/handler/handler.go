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
            return NewGoHandler(cfg, []string{"cmd", "pkg", "internal"}), nil

        case "golang":
            return NewGoHandler(cfg, []string{"cmd", "pkg", "internal"}), nil

        case "c":
            return NewCHandler(cfg, []string{"src"}), nil

        case "java":
            return NewJavaHandler(cfg, []string{"src", "src/main", "src/main/java"}), nil

        case "javascript":
            return NewJavascriptHandler(cfg, []string{"src"}), nil

        case "js":
            return NewJavascriptHandler(cfg, []string{"src"}), nil

        case "python":
            return NewPythonHandler(cfg, []string{"src", "tests", "docs"}), nil

        case "py":
            return NewPythonHandler(cfg, []string{"src", "tests", "docs"}), nil

        default:
            return nil, fmt.Errorf("Unsupported language %q", cfg.Language)
    }
}
