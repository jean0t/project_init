package handler


import (
    "os"
    "os/exec"
    "fmt"
    "path/filepath"
    "project_init/project_init/utils"
)

type JavascriptHandler struct {
    *BaseHandler
}

func NewJavascriptHandler(cfg *utils.Config, dirs []string) Handler {
    return &JavascriptHandler{BaseHandler: &BaseHandler{Cfg: cfg, Dirs: dirs}}
}

func (j *JavascriptHandler) Scaffold() error {
    if err := j.BaseHandler.ScaffoldCommon(); err != nil {
        return err
    }

    return j.createLanguageExtras()
}

func (j *JavascriptHandler) createLanguageExtras() error {
    var mainFile string = filepath.Join(j.Cfg.Name, "src", "main.js")
    var content string = JavascriptSampleCode

    file, err := os.Create(mainFile)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return err
    }

    var npm_init *exec.Cmd = exec.Command("npm", "init", "-y")
    npm_init.Dir = j.BaseHandler.Cfg.Name
    if err = npm_init.Run(); err != nil {
        return err
    }
    fmt.Println("package.json initialized")

    return nil
}
