package utils

import (
    "os"
    "fmt"
    "strings"
    "path/filepath"
)


type Go struct{}


func (g *Go) CreateDirectories(projectName string) error {
    var dirs []string = []string {
        filepath.Join(projectName, "cmd"),
        filepath.Join(projectName, "pkg"),
        filepath.Join(projectName, "internal"),
        filepath.Join(projectName, "tests"),
    }

    for _, dir := range dirs {
        err := os.MkdirAll(dir, 0755)
        if err != nil {
            return fmt.Errorf("Failed to create directory %s: %w", dir, err)
        }
        fmt.Printf("%s created\n", dir)
    }

    return nil
}

func (g *Go) AddSampleFiles(projectName string) error {
    sampleFile := filepath.Join(projectName, "cmd", "main.go")
    content := `package main
// sample code, modify it for your purposes
import "fmt"

func main(){
    fmt.Println("Hello World")
}
`
    file, err := os.Create(sampleFile)
    if err != nil {
        return fmt.Errorf("Failed to create go sample file")
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return fmt.Errorf("It couldn't write the sample code")
    }

    fmt.Printf("%s created\n", sampleFile)
    return nil
}

func (g *Go) AddLicense(projectName, license string) error {
    licenseFile := filepath.Join(projectName, "LICENSE")
    file, err := os.Create(licenseFile)
    if err != nil {
        return fmt.Errorf("It couldnt create the LICENSE file")
    }
    defer file.Close()
    
    var content string = ""
    switch strings.ToLower(license) {
        case "gpl-3":
            content = GPL3_LICENSE

        case "mit":
            content = MIT_LICENSE

        default:
            content = GPL3_LICENSE
    }

    _, err = file.WriteString(content) 
    if err != nil {
        return fmt.Errorf("LICENSE couldn't be written")
    }

    return nil
}


func (g *Go) AddGitIgnore(projectName string) error {
    gitignoreFile := filepath.Join(projectName, ".gitignore")
    file, err := os.Create(gitignoreFile)
    if err != nil {
        return fmt.Errorf("It couldnt create the .gitignore file")
    }
    defer file.Close()

    content := `# Binaries for programs and plugins
*.exe
*.dll
*.so
*.dylib

# Output of 'go build'
*.out

# Vendor directory (if not using Go Modules)
vendor/

# Go test binaries
*.test

# VS Code settings
.vscode/

# Mac OS files
.DS_Store
`

    _, err = file.WriteString(content) 
    if err != nil {
        return fmt.Errorf(".gitignore couldn't be written")
    }

    return nil
}
