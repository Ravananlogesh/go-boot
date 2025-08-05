package cmd

import (
	"fmt"

	"github.com/Ravananlogesh/go-boot/internal/scaffolder"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go project",
	Run: func(cmd *cobra.Command, args []string) {
		scaffolder.GenerateProject("my-app")
		fmt.Println("Project initialized!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
