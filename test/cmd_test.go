package test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"yago/src/yago/cmd"
	"yago/test/mock"
)

func TestRootCommand(t *testing.T) {
	testCases := []struct {
		name          string
		args          []string
		writer        *mock.Writer
		parser        *mock.Parser
		generator     *mock.Generator
		expectedError string
		expectedOut   string
	}{
		{
			name:          "no yaml file provided",
			args:          []string{"-o", "output.go"},
			writer:        &mock.Writer{},
			parser:        &mock.Parser{},
			generator:     &mock.Generator{},
			expectedError: "please provide a yaml file as argument",
		},
		{
			name:          "no output file provided",
			args:          []string{"program/hello-world.yaml"},
			writer:        &mock.Writer{},
			parser:        &mock.Parser{},
			generator:     &mock.Generator{},
			expectedError: "please provide an output file with the -o option",
		},
		{
			name:          "invalid yaml file",
			args:          []string{"program/invalid.yaml", "-o", "output.go"},
			writer:        &mock.Writer{},
			parser:        &mock.Parser{},
			generator:     &mock.Generator{},
			expectedError: "failed to read YAML file: program/invalid.yaml",
		},
		{
			name:   "parse yaml file error",
			args:   []string{"program/hello-world.yaml", "-o", "output.go"},
			writer: &mock.Writer{},
			parser: &mock.Parser{
				ParseFunc: func(data []byte) (interface{}, error) {
					return nil, errors.New("parse error")
				},
			},
			generator:     &mock.Generator{},
			expectedError: "failed to parse YAML: parse error",
		},
		{
			name: "write file error",
			args: []string{"program/hello-world.yaml", "-o", "output.go"},
			writer: &mock.Writer{
				WriteFileFunc: func(name string, data []byte, perm os.FileMode) error {
					return errors.New("write error")
				},
			},
			parser:        &mock.Parser{},
			generator:     &mock.Generator{},
			expectedError: "failed to write to output file: write error",
		},
		{
			name:      "success",
			args:      []string{"program/hello-world.yaml", "-o", "output.go"},
			writer:    &mock.Writer{},
			parser:    &mock.Parser{},
			generator: &mock.Generator{},
			expectedOut: `package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}`,
		},
	}

	cmd.Init()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd.Reset()

			ctx := context.WithValue(context.Background(), "writer", tc.writer)
			ctx = context.WithValue(ctx, "parser", tc.parser)
			ctx = context.WithValue(ctx, "generator", tc.generator)

			cmdRoot := cmd.Root
			cmdRoot.SetContext(ctx)

			cmdRoot.SetArgs(tc.args)

			err := cmdRoot.Execute()

			if tc.expectedError != "" {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), tc.expectedError)
			} else {
				assert.NoError(t, err)

				got := string(tc.writer.GetWrittenData())
				assert.Equal(t, tc.expectedOut, got)
			}
		})
	}
}
