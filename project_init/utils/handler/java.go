package handler


import (
    "os"
    "fmt"
    "path/filepath"
    "project_init/project_init/utils"
)

type JavaHandler struct {
    *BaseHandler
}

func NewJavaHandler(cfg *utils.Config, dirs []string) Handler {
    return &JavaHandler{BaseHandler: &BaseHandler{Cfg: cfg, Dirs: dirs}}
}

func (j *JavaHandler) Scaffold() error {
    if err := j.BaseHandler.ScaffoldCommon(); err != nil {
        return err
    }

    return j.createLanguageExtras()
}

func (j *JavaHandler) createLanguageExtras() error {
    var mainFile string = filepath.Join(j.Cfg.Name, "src", "main", "java", "HelloWorld.java")
    var content string = JavaSampleCode

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
