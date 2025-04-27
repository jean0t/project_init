package handler


import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "project_init/project_init/utils"
)

type BaseHandler struct {
    Cfg *utils.Config
    Dirs []string
}

func (b *BaseHandler) Scaffold() error {
    if err := b.createBaseDirs(); err != nil {
        return err
    }

    if err = b.createReadme(); err != nil {
        return err
    }

    if err = b.createLicense(); err != nil {
        return err
    }

    if err = b.initGit(); err != nil {
        return err
    }

    if err = b.createLanguageExtras(); err != nil {
        return err
    }

    return nil
}

func (b *BaseHandler) createLicense() error {
    var licensePath string = filepath.Join(b.Cfg.Name, "LICENSE")
    var content string = "test"
    file, err := os.Create(licensePath)
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

func (b *BaseHandler) createReadme() error {
    var licensePath string = filepath.Join(b.Cfg.Name, "README.md")
    var content string = "test"
    file, err := os.Create(licensePath)
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
    _, err := exec.LookPath("git")
    if err != nil {
        return err
    }

    var cmd *exec.Cmd = exec.Command("git", "init")
    cmd.Dir = b.Cfg.Name
    if err = cmd.Run(); err != nil {
        return err
    }

    fmt.Println("Git initialized")
    return nil
}

func (b *BaseHandler) createLanguageExtras() error {
    return fmt.Errorf("Implementation not created")
}
