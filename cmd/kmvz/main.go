package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/drewaxxus/kmvz/cmd/kmvz/root"
    "github.com/drewaxxus/kmvz/internal/aur"
    "github.com/drewaxxus/kmvz/pkg/config"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        fmt.Println("Error loading configuration:", err)
        os.Exit(1)
    }

    // Initialize AUR client
    aurClient := aur.NewClient(cfg.AurAPIURL)

    // Initialize root command
    rootCmd := root.NewRootCmd(aurClient)

    // Execute the command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
