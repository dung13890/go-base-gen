# go-base-gen

![workflow status](https://github.com/dung13890/go-base-gen/actions/workflows/go-ci.yml/badge.svg)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

A command-line tool to quickly generate Go projects using clean architecture principles.

## What It Does

- 🚀 **Generate Project**: Creates a complete Go project structure based on [clean architecture](https://github.com/dung13890/go-clean-architecture)
- 📦 **Generate Domain**: Adds new domain modules to your existing project

## Installation

```bash
go install github.com/dung13890/go-base-gen@latest
```

## Quick Start

### 1. Create a New Project

```bash
go-base-gen project --pkg github.com/yourusername/yourproject --path yourproject
```

**Parameters:**
- `--pkg`: Your project's Go module path (required)
- `--path`: Directory where the project will be created (optional, defaults to current directory)

### 2. Set Up Your Project

```bash
cd yourproject
go mod tidy              # Download dependencies
cp .env.example .env     # Create environment configuration
go run cmd/migrate/main.go  # Set up database
```

### 3. Add a New Domain

```bash
go-base-gen domain --dn product --pkg github.com/yourusername/yourproject
```

**Parameters:**
- `--dn`: Domain name (e.g., user, product, order)
- `--pkg`: Your project's Go module path (same as project creation)
- `--path`: Project directory (optional)

### 4. Run Your Project

```bash
make dev
```

## Complete Example

```bash
# Step 1: Generate project
go-base-gen project --pkg github.com/johndoe/shopapi --path shopapi

# Step 2: Navigate to project
cd shopapi

# Step 3: Install dependencies
go mod tidy

# Step 4: Configure environment
cp .env.example .env
# Edit .env with your database credentials

# Step 5: Initialize database
go run cmd/migrate/main.go

# Step 6: Create domains
go-base-gen domain --dn product --pkg github.com/johndoe/shopapi
go-base-gen domain --dn customer --pkg github.com/johndoe/shopapi

# Step 7: Start development server
make dev
```

## Commands Reference

```bash
# Show help
go-base-gen --help

# Generate project
go-base-gen project --pkg <module-path> [--path <directory>]

# Generate domain
go-base-gen domain --dn <domain-name> --pkg <module-path> [--path <directory>]

# Show version
go-base-gen --version
```

## Project Structure

After generation, your project will follow clean architecture with:
- **API Layer**: HTTP handlers and routing
- **Use Cases**: Business logic
- **Repositories**: Data access layer
- **Entities**: Domain models

Learn more about the architecture in the [go-clean-architecture repository](https://github.com/dung13890/go-clean-architecture).

## Support

If you find this tool helpful, consider:

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/dung13890)

## License

MIT License - see [LICENSE](https://opensource.org/licenses/MIT) for details.
