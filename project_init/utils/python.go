package utils

import (
    "os"
    "fmt"
    "strings"
    "path/filepath"
)


type Python struct{}


func (p *Python) CreateDirectories(projectName string) error {
    var dirs []string = []string {
        filepath.Join(projectName, "src"),
        filepath.Join(projectName, "tests"),
        filepath.Join(projectName, "docs"),
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

func (p *Python) AddSampleFiles(projectName string) error {
    sampleFile := filepath.Join(projectName, "src", "main.py")
    content := `# sample code, modify it for your purposes
def main():
    print("Hello World!")

if __name__ == "__main__":
    main()
`
    file, err := os.Create(sampleFile)
    if err != nil {
        return fmt.Errorf("Failed to create python sample file")
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return fmt.Errorf("It couldn't write the sample code")
    }
    
    fmt.Printf("%s created\n", sampleFile)
    return nil
}

func (p *Python) AddLicense(projectName, license string) error {
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


func (p *Python) AddGitIgnore(projectName string) error {
    gitignoreFile := filepath.Join(projectName, ".gitignore")
    file, err := os.Create(gitignoreFile)
    if err != nil {
        return fmt.Errorf("It couldnt create the .gitignore file")
    }
    defer file.Close()

    content := `# Byte-compiled / optimized / DLL files
**/__pycache__/
*.py[cod]
*$py.class

# Virtual environment
venv/
env/

# Distribution / packaging
build/
dist/
*.egg-info/

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
