package main

import (
	"fmt"
	"os"
	"toy-project/svc/cmd/serve"

	"github.com/spf13/cobra"
)

var verbose bool

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "toy-project",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			os.Exit(2)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(serve.Cmd)
	return rootCmd
}

func main() {

	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	fmt.Println("app starded successfully !!")
}
