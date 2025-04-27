package handler


import (
    "os"
    "path/filepath"
    "project_init/project_init/utils"
)

type GoHandler struct {
    *BaseHandler
}

func NewGoHandler(cfg *utils.Config, dirs []string) Handler {
    return &GoHandler{BaseHandler: &BaseHandler{Cfg: cfg, Dirs: dirs}}
}

func (g *GoHandler) createLanguageExtras() error {
    var mainFile string = filepath.Join(p.Cfg.Name, "cmd", "main.go")
    var content string = `package main
import "fmt"

func main() {
    fmt.Println("Hello World!")
}`
    file, err := os.Create(mainFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return err
    }

    return nil
}
