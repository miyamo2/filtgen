/*
Copyright © 2024 - miyamo2 <miyamo2@outlook.com>
*/
package cmd

import (
	"github.com/miyamo2/filtgen/internal"
	"github.com/spf13/cobra"
)

var (
	source string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate iterators based on the file specified by the source flag.",
	Long:  `generate iterators based on the file specified by the source flag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		return internal.Generate(ctx, source)
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&source, "source", "s", "", "source file path")
	err := generateCmd.MarkFlagRequired("source")
	if err != nil {
		panic(err)
	}
}
