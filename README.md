# Beginning - Go Project Scaffolder

A powerful CLI tool designed to scaffold Go projects from predefined templates with best practices and modern architecture patterns.

## 🚀 Features

- **Multiple Template Types**: Support for services, libraries, APIs, and more
- **Flexible Configuration**: Use values.yaml or CLI flags for customization
- **Smart Defaults**: Sensible defaults with easy override options
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **Post-Generation Setup**: Automatic dependency management and initialization
- **Extensible**: Easy to add new template types

## 📦 Installation

### Quick Install (Recommended)
```bash
# Install latest version
curl -fsSL https://raw.githubusercontent.com/zeroxsolutions/beginning/main/install.sh | bash

# Install specific version
curl -fsSL https://raw.githubusercontent.com/zeroxsolutions/beginning/main/install.sh | bash -s -- -v v0.0.1
```

### From GitHub Container Registry
```bash
# Install oras CLI tool first
# macOS
brew install oras

# Linux
curl -LO https://github.com/oras-project/oras/releases/latest/download/oras_linux_amd64.tar.gz
tar -xzf oras_linux_amd64.tar.gz
sudo mv oras /usr/local/bin/

# Download and install from GitHub Container Registry
oras pull ghcr.io/zeroxsolutions/beginning:latest --output .
chmod +x beginning-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m)
sudo mv beginning-$(uname -s | tr '[:upper:]' '[:lower:]')-$(uname -m) /usr/local/bin/beginning
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
| `--go-version` | `-g` | Go version to use | `1.21` |
| `--values` | `-v` | Path to values.yaml file | `values.yaml` |

### Values File (Optional)

Create a `values.yaml` file to store default values:

```yaml
ModuleName: github.com/company/project
RepoName: myproject
GoVersion: 1.21
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
GoVersion: 1.21" > values.yaml

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
- Latest version: `ghcr.io/zeroxsolutions/beginning:latest`
- Specific version: `ghcr.io/zeroxsolutions/beginning:v0.0.1`

### Release Process
1. Create and push a new tag: `git tag v0.0.1 && git push origin v0.0.1`
2. GitHub Actions automatically builds and publishes to GitHub Container Registry
3. Binary files are available for all supported platforms
4. GitHub Release is created with downloadable assets

## 🛠️ Development

### Building
```bash
go build -o beginning main.go
```

### Testing
```bash
go test ./...
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

---

**Happy Scaffolding! 🎉**
