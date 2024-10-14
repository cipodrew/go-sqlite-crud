package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a task manager",
	Long:  `command line todo application`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// 	fmt.Printf("Hello World!")
	// },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
