package cmd

import (
	"github.com/spf13/cobra"
)

var output string

var Root = &cobra.Command{
	Use:   "yago",
	Short: "Yago is a YAML based programming language",
	Long:  `Yago is a declarative language that uses YAML syntax. It takes YAML files as input, converting them into an Abstract Syntax Tree (AST), then generates corresponding Go code. This tool serves as the compiler for Yago.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Init() {
	Root.PersistentFlags().StringVarP(&output, "output", "o", "", "output file (must be specified with -o option)")
}
