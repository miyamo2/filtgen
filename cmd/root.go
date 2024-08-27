/*
Copyright © 2024 - miyamo2 <miyamo2@outlook.com>
*/
package cmd

import (
	"fmt"
	"github.com/miyamo2/filtgen/internal"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "filtgen",
	Short: "filtgen is the CLI tool for generating iterators that selects items by their field values.",
	Long:  `filtgen is the CLI tool for generating iterators that selects items by their field values.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", internal.Version)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
