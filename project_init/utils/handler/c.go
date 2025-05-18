package handler


import (
    "os"
    "fmt"
    "path/filepath"
    "project_init/project_init/utils"
)

type CHandler struct {
    *BaseHandler
}

func NewCHandler(cfg *utils.Config, dirs []string) Handler {
    return &CHandler{BaseHandler: &BaseHandler{Cfg: cfg, Dirs: dirs}}
}

func (c *CHandler) Scaffold() error {
    if err := c.BaseHandler.ScaffoldCommon(); err != nil {
        return err
    }

    return c.createLanguageExtras()
}

func (c *CHandler) createLanguageExtras() error {
    var mainFile string = filepath.Join(c.Cfg.Name, "src", "main.c")
    var content string = CSampleCode

    file, err := os.Create(mainFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return err
    }
    fmt.Printf("%s created\n", mainFile)

    return nil
}
