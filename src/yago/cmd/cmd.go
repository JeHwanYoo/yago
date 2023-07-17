package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"yago/src/yago/injectable"
)

var output string

var Root = &cobra.Command{
	Use:   "yago",
	Short: "Yago is a YAML based programming language",
	Long:  `Yago is a declarative language that uses YAML syntax. It takes YAML files as input, converting them into an Abstract Syntax Tree (AST), then generates corresponding Go code. This tool serves as the compiler for Yago.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("please provide a yaml file as argument")
		}

		if output == "" {
			return fmt.Errorf("please provide an output file with the -o option")
		}

		filePath := args[0]
		writer := cmd.Context().Value("writer").(injectable.Writer)

		yamlFile, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read YAML file: %s", filePath)
		}

		parser := cmd.Context().Value("parser").(injectable.Parser)
		ast, err2 := parser.Parse(yamlFile)
		if err2 != nil {
			return fmt.Errorf("failed to parse YAML: %v", err2)
		}

		generator := cmd.Context().Value("generator").(injectable.Generator)
		goCode, err3 := generator.Generate(ast)
		if err3 != nil {
			return fmt.Errorf("failed to generate Go code: %v", err3)
		}

		err4 := writer.WriteFile(output, []byte(goCode), 0644)
		if err4 != nil {
			return fmt.Errorf("failed to write to output file: %v", err4)
		}

		return nil
	},
}

func Init() {
	Root.PersistentFlags().StringVarP(&output, "output", "o", "", "output file (must be specified with -o option)")
}

func Reset() {
	output = ""
}
