package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init <project-name>",
	Aliases: []string{"initialize", "initialise", "create"},
	Short:   "Initialize a new Go server project",
	Long: `
The 'init' command of 'lets-go' helps developers quickly scaffold a new Go server project with essential packages pre-configured.

When provided with a project path like 'github.com/username/my-awesome-project', the command:

1. Creates a new directory named 'my-awesome-project' on your file system.
2. Initializes this directory as a new Go module with the specified path.
3. Adds Go Fiber to the project's dependencies, allowing developers to quickly build performant web applications.
4. Installs SQLC as a global Go tool, offering a suite of functionalities for handling SQL in Go projects.

The goal is to provide a fast and streamlined process for setting up new Go backend projects.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fullPath := args[0]
		projectName := filepath.Base(fullPath) // Extract the last part of the path as directory name

		// Step 1: Create project directory
		err := os.Mkdir(projectName, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		// Step 2: Execute `go mod init <project-name>`
		execCmd(filepath.Join(projectName, "."), "go", "mod", "init", fullPath)

		// Step 3: Get Go Fiber
		execCmd(filepath.Join(projectName, "."), "go", "get", "github.com/gofiber/fiber/v2")

		// Step 4: Install SQLC
		execCmd("", "go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")
	},
}

func execCmd(dir string, command string, args ...string) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	fmt.Println(string(output))
}

func init() {
	rootCmd.AddCommand(initCmd)
}
