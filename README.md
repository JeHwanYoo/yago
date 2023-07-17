# Yago

> Currently under development.

## Describe

Yago is a declarative language that uses YAML syntax. It takes YAML files as input, converting them into an Abstract
Syntax Tree (AST), then generates corresponding Go code. This tool serves as the compiler for Yago.

This project is being developed in compliance with TDD as a principle.

## Architecture

<a href="https://ibb.co/sPb9MHq"><img src="https://i.ibb.co/JdmjMQH/yago-drawio.png" alt="yago-drawio" border="0"></a>

## Synopsis

```
Usage:
  yago <yaml-file> [flags]

Flags:
  -h, --help            help for yago
  -o, --output string   output file (must be specified with -o option)
```
