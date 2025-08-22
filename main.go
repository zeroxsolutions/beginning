package main

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Values struct {
	ModuleName string `yaml:"ModuleName"`
	RepoName   string `yaml:"RepoName"`
	GoVersion  string `yaml:"GoVersion"`
}

//go:embed template
var templateFS embed.FS

var (
	valuesFile   string
	moduleName   string
	repoName     string
	goVersion    string
	outputDir    string
	templateType string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "beginning",
		Short: "A powerful Go project scaffolder and generator",
		Long: `Beginning is a comprehensive CLI tool designed to scaffold Go projects from predefined templates.

It supports multiple project types including microservices, libraries, APIs, and more. 
The tool automatically generates project structure, dependencies, and configuration files
based on best practices and your specific requirements.

Features:
• Multiple template types (service, library, etc.)
• Flexible output directory configuration
• Template customization via values.yaml or CLI flags
• Automatic dependency management
• Post-generation setup scripts
• Cross-platform compatibility

Examples:
  beginning list                           # Show available template types
  beginning create -t service -r myapi    # Create a service project
  beginning create -t library -r mylib    # Create a library project
  beginning create --help                 # Show detailed help`,
	}

	var scaffoldCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new Go project from templates",
		Long: `Create a new Go project using one of the available templates.

This command will:
1. Generate the complete project structure
2. Apply template variables (module name, repo name, Go version)
3. Create all necessary files and directories
4. Run post-generation setup scripts (if available)
5. Initialize Go modules and dependencies

Template Types:
• service: Full-featured microservice with API, database, swagger docs
• library: Simple Go library with basic structure
• (more types can be added to template/ directory)

Examples:
  beginning create -t service -r myapi -m github.com/company/myapi
  beginning create -t library -r myutils -o /path/to/output
  beginning create -v custom-values.yaml`,
		Run: runScaffold,
	}

	scaffoldCmd.Flags().StringVarP(&valuesFile, "values", "v", "values.yaml", "Path to values.yaml configuration file (optional if using CLI flags)")
	scaffoldCmd.Flags().StringVarP(&moduleName, "module", "m", "", "Go module name (e.g., github.com/company/project)")
	scaffoldCmd.Flags().StringVarP(&repoName, "repo", "r", "", "Repository/project name (used for directory naming)")
	scaffoldCmd.Flags().StringVarP(&goVersion, "go-version", "g", "", "Go version to use (defaults to 1.21 if not specified)")
	scaffoldCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory path (defaults to ./{repo-name})")
	scaffoldCmd.Flags().StringVarP(&templateType, "type", "t", "service", "Template type to use (service, library, etc.)")

	rootCmd.AddCommand(scaffoldCmd)

	// Add list command to show available template types
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List available template types",
		Long: `Display all available template types that can be used with the create command.

This command scans the template directory and shows you what project types
are available for scaffolding. Each template type represents a different
project structure and configuration.

Examples:
  beginning list                    # Show all available templates
  beginning list --help            # Show detailed help`,
		Run: listTemplates,
	}
	rootCmd.AddCommand(listCmd)

	rootCmd.Execute()
}

func listTemplates(cmd *cobra.Command, args []string) {
	fmt.Println("Available template types:")

	err := fs.WalkDir(templateFS, "template", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if path == "template" {
			return nil
		}

		// Get the first level directory (template type)
		parts := strings.Split(path, "/")
		if len(parts) == 2 && d.IsDir() {
			fmt.Printf("  - %s\n", parts[1])
		}

		return nil
	})

	if err != nil {
		fmt.Printf("❌ Error listing templates: %v\n", err)
		os.Exit(1)
	}
}

func runScaffold(cmd *cobra.Command, args []string) {
	// Validate template type exists
	if !templateTypeExists(templateType) {
		fmt.Printf("❌ Template type '%s' not found!\n", templateType)
		fmt.Println("Use 'beginning list' to see available template types")
		os.Exit(1)
	}

	values := loadValues()

	// Determine output directory
	if outputDir == "" {
		outputDir = fmt.Sprintf("./%s", values.RepoName)
	}

	// Convert to absolute path if relative
	if !filepath.IsAbs(outputDir) {
		absPath, err := filepath.Abs(outputDir)
		if err != nil {
			fmt.Printf("❌ Error resolving output path: %v\n", err)
			os.Exit(1)
		}
		outputDir = absPath
	}

	fmt.Printf("Scaffolding %s project in: %s\n", templateType, outputDir)
	check(os.MkdirAll(outputDir, 0755))

	// Use the specific template type
	templatePath := fmt.Sprintf("template/%s", templateType)

	err := fs.WalkDir(templateFS, templatePath, func(path string, d fs.DirEntry, err error) error {
		check(err)
		if path == templatePath {
			return nil
		}

		relPath, _ := filepath.Rel(templatePath, path)
		tmplPath, err := templatePathFunc(relPath, values)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(outputDir, tmplPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		data, err := templateFS.ReadFile(path)
		check(err)

		if filepath.Ext(path) == ".tmpl" {
			// Special handling for gitignore.tmpl -> .gitignore
			if strings.HasSuffix(path, "gitignore.tmpl") {
				gitignorePath := filepath.Join(filepath.Dir(targetPath), ".gitignore")
				return renderTemplateBytes(data, gitignorePath, values)
			}
			// Regular template files: remove .tmpl extension
			return renderTemplateBytes(data, targetPath[:len(targetPath)-5], values)
		} else {
			return os.WriteFile(targetPath, data, 0644)
		}
	})
	check(err)

	fmt.Printf("✅ %s project scaffolded: %s\n", strings.Title(templateType), outputDir)

	// Change to output directory for running commands
	originalDir, _ := os.Getwd()
	check(os.Chdir(outputDir))

	// Run post-scaffold commands (only if they exist)
	if fileExists("bin/swagger.sh") {
		runCommand("chmod +x bin/*")
		runCommand("./bin/swagger.sh")
	}

	if fileExists("go.mod") {
		runCommand("go mod tidy")
	}

	if fileExists("bin/wire.sh") {
		runCommand("./bin/wire.sh")
	}

	// Return to original directory
	check(os.Chdir(originalDir))
}

func runCommand(cmdStr string) {
	fmt.Println("⚙️  Running:", cmdStr)
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	check(cmd.Run())
}

func loadValues() Values {
	values := Values{}

	// Try to load from values.yaml if it exists
	if _, err := os.Stat(valuesFile); err == nil {
		f, err := os.Open(valuesFile)
		if err == nil {
			defer f.Close()
			if yaml.NewDecoder(f).Decode(&values) == nil {
				fmt.Printf("📁 Loaded values from %s\n", valuesFile)
			}
		}
	}

	// Override with CLI flags if provided
	if moduleName != "" {
		values.ModuleName = moduleName
	}
	if repoName != "" {
		values.RepoName = repoName
	}
	if goVersion != "" {
		values.GoVersion = goVersion
	}

	// Validate required values
	if values.ModuleName == "" {
		fmt.Println("❌ Module name is required. Use -m flag or provide in values.yaml")
		os.Exit(1)
	}
	if values.RepoName == "" {
		fmt.Println("❌ Repository name is required. Use -r flag or provide in values.yaml")
		os.Exit(1)
	}
	if values.GoVersion == "" {
		values.GoVersion = "1.24" // Default Go version
		fmt.Printf("ℹ️  Using default Go version: %s\n", values.GoVersion)
	}

	// Validate minimum Go version
	if !isValidGoVersion(values.GoVersion) {
		fmt.Printf("❌ Go version %s is below minimum required version 1.24\n", values.GoVersion)
		os.Exit(1)
	}

	return values
}

func renderTemplateBytes(content []byte, outputPath string, values Values) error {
	tmpl, err := template.New("file").Parse(string(content))
	check(err)

	out, err := os.Create(outputPath)
	check(err)
	defer out.Close()

	return tmpl.Execute(out, values)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func templatePathFunc(path string, data Values) (string, error) {
	tmpl, err := template.New("path").Parse(path)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func templateTypeExists(templateType string) bool {
	entries, err := templateFS.ReadDir("template")
	if err != nil {
		return false
	}

	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == templateType {
			return true
		}
	}
	return false
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isValidGoVersion(version string) bool {
	// Parse version string (e.g., "1.24", "1.25", "1.24.1")
	parts := strings.Split(version, ".")
	if len(parts) < 2 {
		return false
	}

	major, err := parseVersionPart(parts[0])
	if err != nil {
		return false
	}

	minor, err := parseVersionPart(parts[1])
	if err != nil {
		return false
	}

	// Check if version is >= 1.24
	return major > 1 || (major == 1 && minor >= 24)
}

func parseVersionPart(part string) (int, error) {
	// Remove any non-numeric suffix
	cleanPart := strings.TrimRightFunc(part, func(r rune) bool {
		return r < '0' || r > '9'
	})

	var result int
	_, err := fmt.Sscanf(cleanPart, "%d", &result)
	return result, err
}
