# Beginning - Go Project Scaffolder

A powerful CLI tool designed to scaffold Go projects from predefined templates with automatic auto-completion support.

## 🚀 Features

- **Multiple Template Types**: Service, library, API, and more
- **Automatic Auto-completion**: Works out of the box after installation
- **Flexible Configuration**: Via CLI flags or values.yaml
- **Cross-platform**: Supports bash, zsh, fish, and PowerShell
- **Smart Detection**: Automatically detects your shell and installs completion
- **Dynamic Completion**: Real-time suggestions for templates, Go versions, and commands

## 📦 Installation

### Method 1: Go Install (Recommended)
```bash
go install github.com/zeroxsolutions/beginning@latest
```

**Note:** When installing via `go install`, the binary will be placed in `$GOPATH/bin` or `$HOME/go/bin`. 
Make sure this directory is in your `$PATH` for the `beginning` command to work globally.

### Method 2: Build from Source
```bash
git clone https://github.com/zeroxsolutions/beginning.git
cd beginning
go build -o beginning main.go
```

## 🎯 Auto-completion Setup

### 🎉 Automatic Setup (Recommended)
After installing, simply run:
```bash
beginning install-completion
```

This command will:
1. 🔍 Automatically detect your shell (zsh, bash, fish, PowerShell)
2. 📝 Generate the appropriate completion script
3. 📁 Install it to the correct directory
4. ⚙️ Configure your shell configuration files
5. ✅ Verify the installation

**Supported Shells & Installation Paths:**
- **zsh**: `~/.zsh/completions/_beginning` + updates `~/.zshrc`
- **bash**: `~/.local/share/bash-completion/completions/beginning`
- **fish**: `~/.config/fish/completions/beginning.fish`
- **PowerShell**: Generates script for manual profile addition

### 🔧 Manual Setup
If you prefer manual setup or need custom configuration:
```bash
# For zsh
beginning completion zsh > ~/.zsh/completions/_beginning
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
echo 'autoload -U compinit && compinit' >> ~/.zshrc

# For bash
beginning completion bash > ~/.local/share/bash-completion/completions/beginning
echo 'source ~/.local/share/bash-completion/completions/beginning' >> ~/.bashrc

# For fish
beginning completion fish > ~/.config/fish/completions/beginning.fish

# For PowerShell
beginning completion powershell > beginning-completion.ps1
# Then add to your PowerShell profile
```

### 🚀 Quick Start with Auto-completion
```bash
# 1. Install beginning
go install github.com/zeroxsolutions/beginning@latest

# 2. Setup auto-completion (one-time setup)
beginning install-completion

# 3. Restart your shell or source config
source ~/.zshrc  # for zsh
# or restart terminal

# 4. Test auto-completion
beginning [TAB]  # Should show all commands
beginning create -t [TAB]  # Should show template types
```

## 🎮 Usage

### Basic Commands
```bash
# List available templates
beginning list

# Create a new project
beginning create -t service -r myapi -m github.com/company/myapi

# Show help
beginning --help
beginning create --help
```

### 🎯 Auto-completion Examples
```bash
# Commands (press TAB after 'beginning')
beginning [TAB]
# → create, list, completion, install-completion, help

# Flags (press TAB after '-')
beginning create -[TAB]
# → -t, --type, -r, --repo, -m, --module, -g, --go-version, -o, --output, -v, --values

# Template types (press TAB after '-t')
beginning create -t [TAB]
# → service, library

# Go versions (press TAB after '-g')
beginning create -g [TAB]
# → 1.24, 1.25, 1.26, 1.27, 1.28, 1.29, 1.30

# Repository names (dynamic completion)
beginning create -r [TAB]
# → Shows suggestions based on current directory
```

### Project Creation
```bash
# Create a microservice
beginning create -t service -r myapi -m github.com/company/myapi -g 1.25

# Create a library
beginning create -t library -r myutils -m github.com/company/myutils

# Use custom values file
beginning create -v custom-values.yaml
```

## 📋 Available Templates

### Service Template
Full-featured microservice with:
- API endpoints
- Database integration
- Swagger documentation
- Dependency injection (Wire)
- Testing setup
- Docker configuration

### Library Template
Simple Go library with:
- Basic structure
- Testing setup
- Documentation
- Examples

## ⚙️ Configuration

### CLI Flags
- `-t, --type`: Template type (service, library, etc.)
- `-r, --repo`: Repository/project name
- `-m, --module`: Go module name
- `-g, --go-version`: Go version (default: 1.24)
- `-o, --output`: Output directory
- `-v, --values`: Path to values.yaml file

### Values File (values.yaml)
```yaml
ModuleName: github.com/company/project
RepoName: myproject
GoVersion: 1.25
```

## 🔧 Development

### Building
```bash
go build -o beginning main.go
```

### Testing
```bash
go test ./...
```

### Adding New Templates
1. Create a new directory in `template/`
2. Add your template files
3. Use `.tmpl` extension for files that need variable substitution
4. Add any post-generation scripts in `bin/`

## 🌟 Auto-completion Features

### 🚀 Global Installation Support
The auto-completion system works seamlessly whether you:
- **Build from source**: `go build -o beginning main.go`
- **Install globally**: `go install github.com/zeroxsolutions/beginning@latest`

The tool automatically detects the executable path and generates completion scripts correctly.

### 🧠 Smart Detection
- Automatically detects your shell
- Installs to the correct directories
- Updates shell configuration files
- Handles different shell environments

### 🎯 Dynamic Values
- Template types loaded from available templates
- Go versions with latest releases
- Command and flag suggestions
- Context-aware completions

### 🔄 Cross-shell Support
- **zsh**: Full completion with descriptions and advanced features
- **bash**: Full completion with flag suggestions and command completion
- **fish**: Full completion with command descriptions and advanced features
- **PowerShell**: Full completion with profile integration and cross-platform support

### 🚀 Advanced Completion Features
- **Command Chaining**: Complete subcommands and flags
- **Flag Validation**: Only show valid flags for each command
- **Value Suggestions**: Smart suggestions for template types and Go versions
- **Error Handling**: Graceful fallback if completion fails

## 🛠️ Troubleshooting

### Auto-completion Not Working?
```bash
# 1. Check if completion is installed
beginning install-completion --force

# 2. Verify shell configuration
# For zsh: check ~/.zshrc contains completion setup
# For bash: check ~/.bashrc contains source command

# 3. Restart your shell or source config
source ~/.zshrc  # for zsh
source ~/.bashrc # for bash

# 4. Test completion
beginning [TAB]
```

### Common Issues
- **Completion not loading**: Restart shell or source configuration file
- **Wrong shell detected**: Use `--force` flag to reinstall
- **Permission denied**: Check directory permissions for completion files
- **Fish shell issues**: Ensure fish completions directory exists
- **Go install PATH issues**: Ensure `$GOPATH/bin` or `$HOME/go/bin` is in your `$PATH`

### Debug Mode
```bash
# Enable debug output for completion
beginning completion zsh --debug

# Check completion script location
ls -la ~/.zsh/completions/_beginning  # for zsh
ls -la ~/.local/share/bash-completion/completions/beginning  # for bash
```

## 📝 Examples

### Complete Workflow
```bash
# 1. Install the tool
go install github.com/zeroxsolutions/beginning@latest

# 2. Setup auto-completion (one-time)
beginning install-completion

# 3. Create a project with auto-completion
beginning create -t [TAB] -r [TAB] -m [TAB]

# 4. Enjoy auto-completion forever! 🎉
```

### Custom Values
```bash
# Create values.yaml
cat > values.yaml << EOF
ModuleName: github.com/mycompany/myapi
RepoName: myapi
GoVersion: 1.26
EOF

# Use custom values
beginning create -v values.yaml
```

### Advanced Auto-completion Usage
```bash
# Complete command with description
beginning [TAB]
# → create: Create a new Go project from templates
# → list: List available template types
# → completion: Generate completion script for specified shell

# Complete flags with descriptions
beginning create --help
beginning create -[TAB]
# → Shows all available flags with descriptions

# Complete template types dynamically
beginning create -t [TAB]
# → Shows available templates from template/ directory
```

## 🔍 Command Reference

### `beginning install-completion`
Automatically installs completion scripts for your shell.

**Flags:**
- `--force, -f`: Force reinstall even if already installed

**Examples:**
```bash
beginning install-completion           # Auto-detect and install
beginning install-completion --force  # Force reinstall
```

### `beginning completion`
Generates completion scripts for different shells.

**Arguments:**
- `bash`: Generate bash completion script
- `zsh`: Generate zsh completion script
- `fish`: Generate fish completion script
- `powershell`: Generate PowerShell completion script

**Examples:**
```bash
beginning completion zsh > ~/.zsh/completions/_beginning
beginning completion bash > ~/.local/share/bash-completion/completions/beginning
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

MIT License - see LICENSE file for details

## 🆘 Support

- **Issues**: GitHub Issues
- **Documentation**: This README
- **Examples**: See `template/` directory
- **Auto-completion**: Run `beginning install-completion --help`

---

**Happy Scaffolding! 🚀**

*Auto-completion makes everything better! 🎯*
