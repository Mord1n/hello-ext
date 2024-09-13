package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// HelloCmd represents the hello command
var HelloCmd = &cobra.Command{
	Use:   "hello-ext",
	Short: "Prints 'hello from hello-ext'",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from hello-ext")
	},
}
