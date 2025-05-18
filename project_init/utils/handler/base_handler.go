package handler


import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "text/template"
    "project_init/project_init/utils"
)

type BaseHandler struct {
    Cfg *utils.Config
    Dirs []string
}

func (b *BaseHandler) ScaffoldCommon() error {
    if err := b.createBaseDirs(); err != nil {
        return err
    }

    if err := b.createReadme(); err != nil {
        return err
    }

    if err := b.createLicense(); err != nil {
        return err
    }

    if b.Cfg.GitIgnore {
        if err := b.createGitIgnore(); err != nil {
            return err
        }
        fmt.Println(".gitignore created")
    }

    if err := b.initGit(); err != nil {
        fmt.Println("Git isn't available, git won't be initialized in the project.")
    }

    return nil
}

func (b *BaseHandler) createLicense() error {
    var licensePath string = filepath.Join(b.Cfg.Name, "LICENSE")
    var data Data = CreateData(b.Cfg.Name, b.Cfg.License)
    var tmpl *template.Template

    switch strings.ToLower(b.Cfg.License) {
        case "gpl-3":
            var err error
            tmpl, err = template.New("gpl-3").Parse(Gpl3Tmpl)
            if err != nil {
                return err
            }

        case "mit":
            var err error
            tmpl, err = template.New("gpl-3").Parse(MitTmpl)
            if err != nil {
                return err
            }

        default:
                return fmt.Errorf("unsupported license type: %s", b.Cfg.License)
    }

    file, err := os.Create(licensePath)
    if err != nil {
        return err
    }
    defer file.Close()

    return tmpl.Execute(file, data)
}

func (b *BaseHandler) createReadme() error {
    var licensePath string = filepath.Join(b.Cfg.Name, "README.md")
    tmpl, err := template.New("readme").Parse(ReadmeTmpl)
    if err != nil {
        return err
    }

    var data Data = CreateData(b.Cfg.Name, b.Cfg.License)
    file, err := os.Create(licensePath)
    if err != nil {
        return err
    }
    defer file.Close()

    return tmpl.Execute(file, data)
}


func (b *BaseHandler) createGitIgnore() error {
    var licensePath string = filepath.Join(b.Cfg.Name, ".gitignore")
    var content string = ""

    switch strings.ToLower(b.Cfg.Language) {
        case "go":
            content = GoGitIgnoreTmpl

        case "golang":
            content = GoGitIgnoreTmpl

        case "c":
            content = CGitIgnoreTmpl

        case "javascript":
            content = JavascriptGitIgnoreTmpl

        case "js":
            content = JavascriptGitIgnoreTmpl

        case "java":
            content = JavaGitIgnoreTmpl

        case "python":
            content = PythonGitIgnoreTmpl

        case "py":
            content = PythonGitIgnoreTmpl

        default:
                return fmt.Errorf(".gitignore for the language %s isn't available", b.Cfg.Language)
    }

    file, err := os.Create(licensePath)
    if err != nil {
        return err
    }
    defer file.Close()
    
    _, err = file.WriteString(content)
    return err
}


func (b *BaseHandler) createBaseDirs() error {
    var paths []string = []string {b.Cfg.Name}

    for _, d := range b.Dirs {
        paths = append(paths, filepath.Join(b.Cfg.Name, d))
    }

    for _, p := range paths {
        if err := os.MkdirAll(p, 0755); err != nil {
            return err
        }

        fmt.Printf("Directory created successfully: %s\n", p)
    }

    return nil
}

func (b *BaseHandler) initGit() error {
    var exists bool = utils.GitAvailable()
    if !exists {
        return fmt.Errorf("Git isn't available")
    }

    var cmd *exec.Cmd = exec.Command("git", "init")
    cmd.Dir = b.Cfg.Name
    if err := cmd.Run(); err != nil {
        return err
    }

    fmt.Println("Git initialized")
    return nil
}
