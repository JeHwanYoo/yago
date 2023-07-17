package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"yago/src/yago/injectable"
)

var output string

var Root = &cobra.Command{
	Use:   "yago",
	Short: "Yago is a YAML based programming language",
	Long:  `Yago is a declarative language that uses YAML syntax. It takes YAML files as input, converting them into an Abstract Syntax Tree (AST), then generates corresponding Go code. This tool serves as the compiler for Yago.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Please provide a yaml file as argument")
		}

		if output == "" {
			log.Fatal("Please provide an output file with the -o option")
		}

		filePath := args[0]
		writer := cmd.Context().Value("writer").(injectable.Writer)

		yamlFile, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("yamlFile.Get err #%v ", err)
		}

		parser := cmd.Context().Value("parser").(injectable.Parser)
		ast, err2 := parser.Parse(yamlFile)
		if err2 != nil {
			log.Fatalf("Failed to parse YAML: %v", err)
		}

		generator := cmd.Context().Value("generator").(injectable.Generator)
		goCode, err3 := generator.Generate(ast)
		if err3 != nil {
			log.Fatalf("Failed to generate Go code: %v", err)
		}

		err = writer.WriteFile(output, []byte(goCode), 0644)
		if err != nil {
			log.Fatalf("Failed to write to output file: %v", err)
		}
	},
}

func Init() {
	Root.PersistentFlags().StringVarP(&output, "output", "o", "", "output file (must be specified with -o option)")
}
