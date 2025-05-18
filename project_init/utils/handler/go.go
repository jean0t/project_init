package handler


import (
    "os"
    "os/exec"
    "fmt"
    "path/filepath"
    "project_init/project_init/utils"
)

type GoHandler struct {
    *BaseHandler
}

func NewGoHandler(cfg *utils.Config, dirs []string) Handler {
    return &GoHandler{BaseHandler: &BaseHandler{Cfg: cfg, Dirs: dirs}}
}

func (g *GoHandler) Scaffold() error {
    if err := g.BaseHandler.ScaffoldCommon(); err != nil {
        return err
    }

    return g.createLanguageExtras()
}

func (g *GoHandler) createLanguageExtras() error {
    var mainFile string = filepath.Join(g.Cfg.Name, "cmd", "main.go")
    var content string = GoSampleCode

    file, err := os.Create(mainFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return err
    }

    var go_mod *exec.Cmd = exec.Command("go", "mod", "init", g.BaseHandler.Cfg.Name)
    go_mod.Dir = g.BaseHandler.Cfg.Name
    if err = go_mod.Run(); err != nil {
        return err
    }
    fmt.Println("go.mod initialized")

    return nil
}
