# goup

goup is a command-line tool to quickly create a new Go environment, initialize a Go module, and open the project in your preferred editor.

## Features

Initializes a new Go module in a specified directory.
Creates a basic main.go file with a "Hello, World!" example.
Opens the project in your preferred editor.
Allows overriding the default editor via command-line options.
Installation

## Build

```bash
make build
```

## Install

```bash
make install
```

## Configuration

Create a configuration file at `~/.config/goup` to set your default editor:

```bash
editor=nvim
```

Replace `nvim` with the command for your preferred editor (e.g., `vim`, `nano`, `subl`, `code`).

## Usage

Run `goup` with a module name to create a new Go project:

```bash
goup <module-name>
```

### Command-Line Options

`-e` or `-editor`: Override the default editor specified in the configuration file.

Examples:

```bash
goup my-module                # Uses the editor specified in ~/.config/goup
goup -e nvim my-module         # Opens with Vim editor
goup -editor nano my-module   # Opens with Nano editor
```

### License

This project is licensed under the MIT License.
