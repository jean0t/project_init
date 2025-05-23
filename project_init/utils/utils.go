package utils

import (
    "os/exec"
    "strings"
)

type Config struct {
    Name string
    Language string
    License string
    GitIgnore bool
}


func GitAvailable() bool {
    _, err := exec.LookPath("git")
    return err == nil
}


func GetGitUsername() string {
    if !GitAvailable() {
        return "<name>"
    }

    var cmd *exec.Cmd = exec.Command("git", "config", "--global", "user.name")
    output, err := cmd.Output()
    if err != nil {
        return "<name>"
    }

    var name string = strings.TrimSpace(string(output))
    if name == "" {
        return "<name>"
    }

    return name
}

