package main

import (
	"github.com/spf13/cobra"
	"hello-ext/cmd" // Import your commands
)

func main() {
	var rootCmd = &cobra.Command{Use: "hello-ext"}
	rootCmd.AddCommand(cmd.HelloCmd)
	rootCmd.Execute()
}
