package cmd_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"yago/src/yago/cmd"
	"yago/test/mock"
)

func TestRootCommand(t *testing.T) {
	ctx := context.WithValue(context.Background(), "writer", &mock.Writer{})
	ctx = context.WithValue(ctx, "parser", mock.Parser{})
	ctx = context.WithValue(ctx, "generator", mock.Generator{})

	cmd.Init()

	cmdRoot := cmd.Root
	cmdRoot.SetContext(ctx)

	cmdRoot.SetArgs([]string{"program/hello-world.yaml", "-o", "output.go"})

	assert.NoError(t, cmdRoot.Execute())

	expected := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}`

	writer := ctx.Value("writer").(*mock.Writer)
	got := string(writer.GetWrittenData())

	assert.Equal(t, expected, got)
}
