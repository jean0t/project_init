package handler


import (
    "os"
    "path/filepath"
    "project_init/project_init/utils"
)

type PythonHandler struct {
    *BaseHandler
}

func NewPythonHandler(cfg *utils.Config, dirs []string) Handler {
    return &PythonHandler{BaseHandler: &BaseHandler{Cfg: cfg, Dirs: dirs}}
}

func (p *PythonHandler) Scaffold() error {
    if err := p.BaseHandler.ScaffoldCommon(); err != nil {
        return err
    }

    return p.createLanguageExtras()
}

func (p *PythonHandler) createLanguageExtras() error {
    var mainFile string = filepath.Join(p.Cfg.Name, "src", "main.py")
    var content string = PythonSampleCode

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
