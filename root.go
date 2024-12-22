package root

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/yourusername/kmvz/internal/aur"
    "github.com/yourusername/kmvz/internal/utils"
)

// NewRootCmd creates the root command
func NewRootCmd(aurClient *aur.Client) *cobra.Command {
    rootCmd := &cobra.Command{
        Use:   "kmvz",
        Short: "know my vucking zhit",
    }

    // Add subcommands
    rootCmd.AddCommand(
        installCmd(aurClient),
        uninstallCmd(aurClient),
        updateCmd(aurClient),
        searchCmd(aurClient),
        helpCmd(),
    )

    return rootCmd
}

// installCmd creates the install command
func installCmd(aurClient *aur.Client) *cobra.Command {
    return &cobra.Command{
        Use:   "install <package>",
        Short: "Install a package from AUR",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            pkgName := args[0]
            if err := aurClient.InstallPackage(pkgName); err != nil {
                fmt.Println("Error installing package:", err)
                return
            }
            fmt.Println("Package installed successfully.")
        },
    }
}

// uninstallCmd creates the uninstall command
func uninstallCmd(aurClient *aur.Client) *cobra.Command {
    return &cobra.Command{
        Use:   "uninstall <package>",
        Short: "Uninstall a package from AUR",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            pkgName := args[0]
            if err := aurClient.UninstallPackage(pkgName); err != nil {
                fmt.Println("Error uninstalling package:", err)
                return
            }
            fmt.Println("Package uninstalled successfully.")
        },
    }
}

// updateCmd creates the update command
func updateCmd(aurClient *aur.Client) *cobra.Command {
    return &cobra.Command{
        Use:   "update",
        Short: "Update all packages",
        Run: func(cmd *cobra.Command, args []string) {
            if err := aurClient.UpdatePackages(); err != nil {
                fmt.Println("Error updating packages:", err)
                return
            }
            fmt.Println("Packages updated successfully.")
        },
    }
}

// searchCmd creates the search command
func searchCmd(aurClient *aur.Client) *cobra.Command {
    return &cobra.Command{
        Use:   "search <query>",
        Short: "Search for a package in AUR",
        Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            query := args[0]
            results, err := aurClient.SearchPackage(query)
            if err != nil {
                fmt.Println("Error searching for package:", err)
                return
            }
            for _, pkg := range results {
                fmt.Println(pkg.Name, "-", pkg.Description)
            }
        },
    }
}

// helpCmd creates the help command
func helpCmd() *cobra.Command {
    return &cobra.Command{
        Use:   "help",
        Short: "Show help information",
        Run: func(cmd *cobra.Command, args []string) {
            cmd.Help() // Use the built-in Cobra help command
        },
    }
}