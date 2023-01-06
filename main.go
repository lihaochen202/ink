package main

import (
	"github.com/lihaochen202/ink/cmd/npm"
	"github.com/lihaochen202/ink/cmd/prettier"
	"github.com/lihaochen202/ink/cmd/tailwind"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "ink",
	}
	rootCmd.AddCommand(npm.RootCmd)
	rootCmd.AddCommand(prettier.RootCmd)
	rootCmd.AddCommand(tailwind.RootCmd)
	rootCmd.Execute()
}
