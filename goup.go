package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Config struct {
	Editor string
}

func readConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configFile := filepath.Join(homeDir, ".config", "goup")

	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	config := &Config{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "editor=") {
			config.Editor = strings.TrimPrefix(line, "editor=")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if config.Editor == "" {
		return nil, fmt.Errorf("editor not specified in config")
	}

	return config, nil
}

func initializeGoProject(moduleName string) error {

	err := os.Mkdir(moduleName, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %v", moduleName, err)
	}

	err = os.Chdir(moduleName)
	if err != nil {
		return fmt.Errorf("failed to change directory to %s: %v", moduleName, err)
	}

	err = exec.Command("go", "mod", "init", moduleName).Run()
	if err != nil {
		return fmt.Errorf("failed to initialize go module: %v", err)
	}

	mainFileContents := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
`
	err = os.WriteFile("main.go", []byte(mainFileContents), 0644)
	if err != nil {
		return fmt.Errorf("failed to create main.go: %v", err)
	}

	err = exec.Command("git", "init", ".").Run()
	if err != nil {
		fmt.Printf("Failed to initialize git repository: %v\n", err)
	} else {
		fmt.Println("Git repository initialized.")
	}

	return nil
}

func main() {

	editorFlag := flag.String("e", "", "Override the default editor")
	editorLongFlag := flag.String("editor", "", "Override the default editor")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("Usage: goup [options] <module-name>")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	moduleName := flag.Arg(0)

	config, err := readConfig()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	// Check if an editor override was provided via command-line flags
	if *editorFlag != "" {
		config.Editor = *editorFlag
	} else if *editorLongFlag != "" {
		config.Editor = *editorLongFlag
	}

	// Ensure the editor is specified
	if config.Editor == "" {
		fmt.Println("No editor specified. Please set an editor in the config file or use the -e or -editor flag.")
		os.Exit(1)
	}

	err = initializeGoProject(moduleName)
	if err != nil {
		fmt.Printf("Error initializing project: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(config.Editor, "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error opening editor: %v\n", err)
		os.Exit(1)
	}
}
