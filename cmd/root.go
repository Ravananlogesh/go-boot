package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-boot",
	Short: "Go Boot-like project generator for Go",
	Long:  `go-boot is a CLI tool to quickly bootstrap production-ready Go applications with best practices.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}
