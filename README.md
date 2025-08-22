# Beginning - Go Project Scaffolder

A powerful CLI tool designed to scaffold Go projects from predefined templates with best practices and modern architecture patterns.

## 🚀 Features

- **Multiple Template Types**: Support for services, libraries, APIs, and more
- **Flexible Configuration**: Use values.yaml or CLI flags for customization
- **Smart Defaults**: Sensible defaults with easy override options
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Post-Generation Setup**: Automatic dependency management and initialization
- **Extensible**: Easy to add new template types
- **Go 1.24+ Ready**: Built with latest Go features and optimizations

## 📦 Installation

> **Note**: This repository uses the `master` branch. Make sure to use `master` instead of `main` in URLs.

### Quick Install (Recommended)
```bash
# Install latest version
go install github.com/zeroxsolutions/beginning@latest

# Install specific version
go install github.com/zeroxsolutions/beginning@v0.0.1
```

**That's it!** 🎉 

The CLI tool will be automatically built for your platform and installed to your Go bin directory. All templates are embedded in the binary, so no additional downloads are needed.

### Alternative: From Source
```bash
# Clone and build from source
git clone https://github.com/zeroxsolutions/beginning.git
cd beginning
go install .
```

### From Source
```bash
git clone <repository-url>
cd beginning
go build -o beginning main.go
```

### Make Executable Available Globally
```bash
# Move to a directory in your PATH
sudo mv beginning /usr/local/bin/

# Or add to your local bin
mkdir -p ~/bin
mv beginning ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
```

## 🎯 Quick Start

### List Available Templates
```bash
beginning list
```

### Create a Service Project
```bash
beginning create -t service -r myapi -m github.com/company/myapi
```

### Create a Library Project
```bash
beginning create -t library -r myutils -m github.com/company/myutils
```

### Create Project in Specific Directory
```bash
beginning create -t service -r myproject -o /path/to/output
```

## 📚 Usage

### Root Command
```bash
beginning --help
```

### Create Command
```bash
beginning create --help
```

### List Command
```bash
beginning list --help
```

## 🔧 Configuration

### CLI Flags

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--type` | `-t` | Template type (service, library, etc.) | `service` |
| `--repo` | `-r` | Repository/project name | Required |
| `--module` | `-m` | Go module name | Required |
| `--output` | `-o` | Output directory path | `./{repo-name}` |
| `--go-version` | `-g` | Go version to use | `1.24` |
| `--values` | `-v` | Path to values.yaml file | `values.yaml` |

### Values File (Optional)

Create a `values.yaml` file to store default values:

```yaml
ModuleName: github.com/company/project
RepoName: myproject
GoVersion: 1.24
```

## 🏗️ Template Types

### Service Template
Full-featured microservice with:
- REST API structure
- Database integration
- Swagger documentation
- Dependency injection (Wire)
- Configuration management
- Testing setup

### Library Template
Simple Go library with:
- Basic module structure
- README documentation
- Go module configuration

### Adding New Templates
1. Create a new directory in `template/`
2. Add your template files
3. Use `.tmpl` extension for files that need variable substitution
4. CLI will automatically detect new template types

## 📁 Project Structure

```
template/
├── service/          # Microservice template
│   ├── cmd/         # Application entry points
│   ├── internal/    # Private application code
│   ├── config/      # Configuration files
│   ├── bin/         # Build and setup scripts
│   └── ...
└── library/          # Library template
    ├── go.mod.tmpl
    └── README.md.tmpl
```

## 🔍 Examples

### Basic Service Creation
```bash
beginning create -t service -r user-service -m github.com/company/user-service
```

### Library with Custom Output
```bash
beginning create -t library -r utils -m github.com/company/utils -o ~/Projects/
```

### Using Values File
```bash
# Create values.yaml
echo "ModuleName: github.com/company/api
RepoName: api-service
GoVersion: 1.24" > values.yaml

# Create project
beginning create -t service
```

### Override Values File
```bash
beginning create -t service -r custom-name -m github.com/company/custom
```

## 🚀 Releases

### Versioning
We use semantic versioning (SemVer) for releases:
- Format: `vX.Y.Z` (e.g., `v0.0.1`, `v1.0.0`)
- Latest version: `github.com/zeroxsolutions/beginning@latest`
- Specific version: `github.com/zeroxsolutions/beginning@v0.0.1`

### Release Process
1. Create and push a new tag: `git tag v0.0.1 && git push origin v0.0.1`
2. GitHub Actions automatically verifies and tests the Go module
3. Go modules are automatically published for `go install`
4. All templates are embedded in the binary - no separate downloads needed

### How It Works
```bash
# When you run this:
go install github.com/zeroxsolutions/beginning@latest

# Go automatically:
# 1. Downloads the source code
# 2. Builds the binary for your platform (Linux/macOS/Windows)
# 3. Embeds all templates into the binary
# 4. Installs to your Go bin directory
# 5. Makes it available as 'beginning' command
```

### Why This Approach?
- 🚀 **Automatic**: No need to choose the right binary for your platform
- 🔧 **Native**: Built specifically for your OS and architecture
- 📦 **Complete**: All templates are embedded, no missing files
- 🎯 **Simple**: One command, everything works

## 🛠️ Development

### Building
```bash
# Basic build
go build -o beginning main.go

# Build with Go 1.24+ optimizations
go build -ldflags="-s -w -X main.version=$(git describe --tags --always --dirty)" -o beginning main.go

# Build for specific platform (for testing)
GOOS=linux GOARCH=amd64 go build -o beginning-linux-amd64 main.go
GOOS=darwin GOARCH=arm64 go build -o beginning-darwin-arm64 main.go
```

### Local Development
```bash
# Install locally for development
go install .

# Run directly
go run main.go list
go run main.go create -t service -r myproject -m github.com/company/myproject
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...
```

### Adding Dependencies
```bash
go get github.com/spf13/cobra
go mod tidy
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

- **Issues**: Create an issue on GitHub
- **Documentation**: Check this README and help commands
- **Examples**: See the examples section above

## 🔧 Troubleshooting

### Installation Issues
- **404 Error**: Make sure you're using the correct branch (`master`, not `main`)
- **Permission Denied**: Use `sudo` or install to user directory with `~/bin`
- **Command Not Found**: Add the installation directory to your PATH

### Common Problems
- **Wrong Branch**: Repository uses `master` branch, not `main`
- **Path Issues**: Ensure the binary is in your PATH
- **Version Mismatch**: Check Go version compatibility (Go 1.24+ recommended)

---

**Happy Scaffolding! 🎉**
