#!/bin/bash

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default values
REPO="zeroxsolutions/beginning"
VERSION="latest"
OS=""
ARCH=""
INSTALL_DIR="/usr/local/bin"

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to detect OS and architecture
detect_platform() {
    print_status "Detecting platform..."
    
    case "$(uname -s)" in
        Linux*)     OS="linux";;
        Darwin*)    OS="darwin";;
        CYGWIN*|MINGW32*|MSYS*|MINGW*) OS="windows";;
        *)          print_error "Unsupported operating system"; exit 1;;
    esac
    
    case "$(uname -m)" in
        x86_64)    ARCH="amd64";;
        aarch64)   ARCH="arm64";;
        arm64)     ARCH="arm64";;
        *)         print_error "Unsupported architecture"; exit 1;;
    esac
    
    print_success "Detected: $OS-$ARCH"
}

# Function to download and install
install_cli() {
    local tag="$VERSION"
    local registry="ghcr.io"
    local url="https://$registry/$REPO:$tag"
    
    print_status "Installing beginning CLI version $VERSION for $OS-$ARCH..."
    
    # Create temporary directory
    local temp_dir=$(mktemp -d)
    cd "$temp_dir"
    
    # Download using oras
    if command -v oras &> /dev/null; then
        print_status "Using oras to download from GitHub Container Registry..."
        oras pull "$url" --output .
    else
        print_warning "oras not found, trying alternative download method..."
        # Alternative: download from GitHub releases
        local release_url="https://github.com/$REPO/releases/download/$VERSION"
        if [ "$OS" = "windows" ]; then
            local file="beginning-$OS-$ARCH.zip"
            curl -L -o "$file" "$release_url/$file"
            unzip "$file"
        else
            local file="beginning-$OS-$ARCH.tar.gz"
            curl -L -o "$file" "$release_url/$file"
            tar -xzf "$file"
        fi
    fi
    
    # Find the binary
    local binary=""
    if [ "$OS" = "windows" ]; then
        binary="beginning-$OS-$ARCH.exe"
    else
        binary="beginning-$OS-$ARCH"
    fi
    
    if [ ! -f "$binary" ]; then
        print_error "Binary not found after download"
        exit 1
    fi
    
    # Make binary executable
    chmod +x "$binary"
    
    # Install to system
    if [ "$EUID" -eq 0 ]; then
        # Running as root
        cp "$binary" "$INSTALL_DIR/beginning"
        print_success "Installed to $INSTALL_DIR/beginning"
    else
        # Not running as root, try to install to user directory
        local user_bin="$HOME/.local/bin"
        mkdir -p "$user_bin"
        cp "$binary" "$user_bin/beginning"
        
        # Add to PATH if not already there
        if [[ ":$PATH:" != *":$user_bin:"* ]]; then
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.bashrc"
            echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.zshrc"
            print_warning "Added $user_bin to PATH in shell config files"
            print_warning "Please restart your terminal or run: source ~/.bashrc (or ~/.zshrc)"
        fi
        
        print_success "Installed to $user_bin/beginning"
    fi
    
    # Clean up
    cd /
    rm -rf "$temp_dir"
    
    print_success "Installation completed successfully!"
}

# Function to show help
show_help() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -v, --version VERSION    Specify version to install (default: latest)"
    echo "  -d, --directory DIR      Specify installation directory (default: /usr/local/bin)"
    echo "  -h, --help               Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0                       # Install latest version"
    echo "  $0 -v v0.0.1            # Install specific version"
    echo "  $0 -d ~/bin             # Install to custom directory"
    echo ""
    echo "Note: This script requires oras CLI tool for optimal experience."
    echo "Install oras: https://oras.land/cli/"
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -v|--version)
            VERSION="$2"
            shift 2
            ;;
        -d|--directory)
            INSTALL_DIR="$2"
            shift 2
            ;;
        -h|--help)
            show_help
            exit 0
            ;;
        *)
            print_error "Unknown option: $1"
            show_help
            exit 1
            ;;
    esac
done

# Main installation process
main() {
    print_status "Beginning CLI installation script"
    print_status "Repository: $REPO"
    print_status "Version: $VERSION"
    
    detect_platform
    install_cli
    
    print_success "You can now use 'beginning --help' to see available commands"
}

# Run main function
main "$@" 