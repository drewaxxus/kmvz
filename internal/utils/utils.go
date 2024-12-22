package utils

import (
    "os/exec"
)

// RunCommand executes a command and returns the output
func RunCommand(name string, args ...string) (string, error) {
    cmd := exec.Command(name, args...)
    output, err := cmd.CombinedOutput()
    return string(output), err
}
