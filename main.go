package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

type Values struct {
	ModuleName string `yaml:"ModuleName"`
	RepoName   string `yaml:"RepoName"`
	GoVersion  string `yaml:"GoVersion"`
}

func main() {
	values := loadValues()

	projectDir := fmt.Sprintf("./%s", values.RepoName)

	fmt.Println("Scaffolding project in:", projectDir)
	err := os.Mkdir(projectDir, 0755)
	check(err)

	err = filepath.Walk("_template", func(path string, info os.FileInfo, err error) error {
		check(err)

		relPath, _ := filepath.Rel("_template", path)
		tmplPath := templatePath(relPath, values)

		targetPath := filepath.Join(projectDir, tmplPath)

		if info.IsDir() {
			os.MkdirAll(targetPath, 0755)
			return nil
		}

		if filepath.Ext(path) == ".tmpl" {
			renderTemplate(path, targetPath[:len(targetPath)-5], values)
		} else {
			copyFile(path, targetPath)
		}
		return nil
	})
	check(err)

	fmt.Println("Project scaffolded:", projectDir)

	os.Chdir(projectDir)

	runCommand("chmod +x bin/*")
	runCommand("./bin/swagger.sh")
	runCommand("go mod tidy")
	runCommand("./bin/wire.sh")
}

func runCommand(cmdStr string) {
	fmt.Println("⚙️  Running:", cmdStr)
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}

func loadValues() Values {
	f, err := os.Open("values.yaml")
	check(err)
	defer f.Close()

	var values Values
	err = yaml.NewDecoder(f).Decode(&values)
	check(err)
	return values
}

func renderTemplate(tmplPath, outputPath string, values Values) {
	tmpl, err := template.ParseFiles(tmplPath)
	check(err)

	out, err := os.Create(outputPath)
	check(err)
	defer out.Close()

	err = tmpl.Execute(out, values)
	check(err)
}

func copyFile(src, dst string) {
	in, err := os.Open(src)
	check(err)
	defer in.Close()

	out, err := os.Create(dst)
	check(err)
	defer out.Close()

	_, err = io.Copy(out, in)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func templatePath(path string, data Values) string {
	tmpl, err := template.New("path").Parse(path)
	check(err)

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	check(err)

	return buf.String()
}
