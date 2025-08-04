# Leap CLI - Comprehensive Usage Documentation

Leap CLI is a comprehensive command-line interface for managing TensorLeap deep learning model debugging and analysis workflows. This document provides complete usage information for every command, flag, and option available in the CLI.

## Table of Contents

1. [Installation & Configuration](#installation--configuration)
2. [Global Options](#global-options)
3. [Core Commands](#core-commands)
4. [Authentication Commands](#authentication-commands)
5. [Project Management Commands](#project-management-commands)
6. [Code Integration Commands](#code-integration-commands)
7. [Model Management Commands](#model-management-commands)
8. [Secrets Management Commands](#secrets-management-commands)
9. [Hub Commands](#hub-commands)
10. [CLI Management Commands](#cli-management-commands)
11. [Platform Support & Availability](#platform-support--availability)

---

## Installation & Configuration

### Basic Usage
```bash
leap [command] [subcommand] [flags]
```

### Configuration File
Default configuration location: `$HOME/.config/tensorleap/config.yaml`

---

## Global Options

These flags are available on all commands:

| Flag | Description | Default |
|------|-------------|---------|
| `--config string` | Config file path | `$HOME/.config/tensorleap/config.yaml` |
| `--apiKey string` | TensorLeap API key | - |
| `--apiUrl string` | TensorLeap API URL | - |
| `-h, --help` | Help for any command | - |

---

## Core Commands

### Root Command: `leap`

**Description**: Leap - Deepbug your models!

**Usage**: 
```bash
leap [flags]
leap --version    # Show version information
```

**Options**:
- `--version`: Show version information

### `leap completion` - Shell Completion

**Description**: Generate the autocompletion script for leap for the specified shell

**Usage**:
```bash
leap completion [command]
```

**Available Commands**:
- `leap completion bash` - Generate the autocompletion script for bash
- `leap completion fish` - Generate the autocompletion script for fish  
- `leap completion powershell` - Generate the autocompletion script for powershell
- `leap completion zsh` - Generate the autocompletion script for zsh

---

## Authentication Commands

### `leap auth` - Authentication Management

#### `leap auth login [environment url]`
**Description**: Login using API key or username/password

**Usage**:
```bash
leap auth login [environment_url] [flags]
leap auth login https://api.tensorleap.ai --api-key YOUR_API_KEY
leap auth login https://api.tensorleap.ai --username john --password secret
```

**Options**:
- `-n, --name string`: Name of the environment to login to (defaults to parsing the API URL)
- `-k, --api-key string`: API key to login with
- `-u, --username string`: Username for username/password authentication
- `-p, --password string`: Password for username/password authentication

#### `leap auth logout`
**Description**: Remove API key from the machine

**Usage**:
```bash
leap auth logout [flags]
leap auth logout --name production
```

**Options**:
- `-n, --name string`: Name of the environment to logout from (defaults to current environment)

#### `leap auth whoami`
**Description**: Get information about the authenticated user

**Usage**:
```bash
leap auth whoami
```

#### `leap auth select [environment name]`
**Description**: Select an environment to use

**Usage**:
```bash
leap auth select [environment_name]
leap auth select production
```

---

## Project Management Commands

### `leap projects` (aliases: `project`, `pro`) - Project Management

#### `leap projects create [projectName]`
**Description**: Create new project

**Usage**:
```bash
leap projects create [projectName] [flags]
leap projects create "My ML Project" --description "Computer vision model"
```

**Options**:
- `-d, --description string`: Project description

#### `leap projects list` (alias: `ls`)
**Description**: List projects

**Usage**:
```bash
leap projects list
leap projects ls
```

#### `leap projects delete`
**Description**: Delete project

**Usage**:
```bash
leap projects delete [flags]
leap projects delete --projectId PROJECT_ID --skipConfirm
```

**Options**:
- `--projectId string`: Project ID to delete
- `-y, --skipConfirm`: Skip deletion confirmation

#### `leap projects info`
**Description**: Show project and code integration information

**Usage**:
```bash
leap projects info
```

#### `leap projects init`
**Description**: Create initial project environment files in the current directory

**Usage**:
```bash
leap projects init [flags]
leap projects init --projectId PROJECT_ID --codeId CODE_ID
```

**Options**:
- `--projectId string`: Project ID
- `--codeId string`: Code integration ID to bind
- `--secretId string`: Secret manager ID for new code integration
- `--pythonVersion string`: Python version for the code integration
- `--branch string`: Branch of the code integration to bind

#### `leap projects copy [env:source-name] [env:target-name]` (alias: `cp`)
**Description**: Copy a project

**Usage**:
```bash
leap projects copy [env:source-name] [env:target-name] [flags]
leap projects copy prod:my-project dev:my-project-copy --no-cache
```

**Options**:
- `-n, --no-cache`: Do not use project exported cache
- `--exclude-calc-files`: Do not include calculated files in export (decreases export size)

#### `leap projects export [project name]`
**Description**: Export project to file

**Usage**:
```bash
leap projects export [project_name] [flags]
leap projects export "My Project" --output-dir ./exports
```

**Options**:
- `-o, --output-dir string`: Output directory
- `--no-cache`: Do not use previously cached export
- `--exclude-calc-files`: Do not include calculated files in export (decreases export size)

#### `leap projects import [project name]`
**Description**: Import project from hub

**Usage**:
```bash
leap projects import [project_name] [flags]
leap projects import "External Project" --file ./project.tar.gz
```

**Options**:
- `-f, --file string`: Import project from file instead of hub

#### `leap projects publish [project names]` *(Hub Required)*
**Description**: Publish project to the hub
**Availability**: Only when `LEAP_HUB_ENABLED=true`

**Usage**:
```bash
leap projects publish [project_names] [flags]
leap projects publish "My Project" --override
leap projects publish --all
```

**Options**:
- `-o, --override`: Override existing project
- `--use-export-cache`: Use project exported cache (default true unless override flag is set)
- `--exclude-calc-files`: Do not include calculated files in export (decreases export size)
- `-a, --all`: Publish all public projects
- `-s, --use-signed-url`: Use signed URL to publish project (requires appropriate permissions)

#### `leap projects push <modelPath>`
**Description**: Push new version into a project with its model and code integration

**Usage**:
```bash
leap projects push <modelPath> [flags]
leap projects push ./model.h5 --model-name "v1.2" --type H5_TF2
```

**Options**:
- `-m, --model-name string`: Model version name
- `--code-message string`: Code version message
- `--type string`: Model file type [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]
- `--model-branch string`: Name of the model branch (optional)
- `--code-branch string`: Name of the code branch (optional)
- `--secretId string`: Secret ID
- `-f, --force`: Force push code integration
- `--transform-input`: Transpose input data to channel-last format
- `--no-wait`: Do not wait for push to complete
- `--leap-mapping string`: Path to external leap mapping file

#### `leap projects set-secret`
**Description**: Set secret on local leap config

**Usage**:
```bash
leap projects set-secret [flags]
leap projects set-secret --secret-id SECRET_ID
```

**Options**:
- `--secret-id string`: Secret ID

---

## Code Integration Commands

### `leap code` (alias: `code-integration`) - Code Integration Management

#### `leap code create [code-integration-name]`
**Description**: Create a new code integration

**Usage**:
```bash
leap code create [code-integration-name]
leap code create "my-model-code"
```

#### `leap code list` (alias: `ls`)
**Description**: List available code integrations

**Usage**:
```bash
leap code list
leap code ls
```

#### `leap code delete`
**Description**: Delete code integration

**Usage**:
```bash
leap code delete [flags]
leap code delete --codeId CODE_ID --skipConfirm
```

**Options**:
- `--codeId string`: Code Integration ID to delete
- `-y, --skipConfirm`: Skip deletion confirmation

#### `leap code info`
**Description**: Show code integration information

**Usage**:
```bash
leap code info
```

#### `leap code init`
**Description**: Create initial code integration files in the current directory

**Usage**:
```bash
leap code init [flags]
leap code init --new "my-code" --secretId SECRET_ID
leap code init --codeId EXISTING_CODE_ID
```

**Options**:
- `--codeId string`: Code integration ID of existing dataset (mutually exclusive with --new)
- `--new string`: Name for new database (mutually exclusive with --codeId)
- `--secretId string`: Secret manager ID for new code integration
- `-b, --branch string`: Branch for code integration
- `-p, --pythonVersion string`: Python version for code integration

#### `leap code list-versions`
**Description**: List code integration versions

**Usage**:
```bash
leap code list-versions [flags]
leap code list-versions --codeId CODE_ID
```

**Options**:
- `--codeId string`: Code integration ID to get versions from

#### `leap code pull [dataset-id] [branch]`
**Description**: Pull dataset into a new directory

**Usage**:
```bash
leap code pull [dataset-id] [branch] [flags]
leap code pull DATASET_ID main --mapping-only
```

**Options**:
- `-f, --mapping-only`: Pull mapping only

#### `leap code push`
**Description**: Push code integration

**Usage**:
```bash
leap code push [flags]
leap code push --branch main --message "Updated model architecture"
```

**Options**:
- `--secretId string`: Secret ID
- `-b, --branch string`: Branch
- `-m, --message string`: Commit message
- `--no-wait`: Do not wait for code parsing
- `-f, --force`: Force push code integration
- `-p, --python-version string`: Python version
- `--leap-mapping string`: Path to external leap mapping file

---

## Model Management Commands

### `leap models` - Model Management

#### `leap models import <modelPath>`
**Description**: Import a model into TensorLeap

**Usage**:
```bash
leap models import <modelPath> [flags]
leap models import ./model.onnx --projectId PROJECT_ID --type ONNX
```

**Options**:
- `--projectId string`: Project ID the model will be imported to
- `-m, --message string`: Version message
- `--type string`: Model file type [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]
- `--model-branch string`: Name of the model branch (optional)
- `--code-branch string`: Name of the code integration branch (optional)
- `--codeId string`: Code integration ID (will use the last valid dataset version)
- `--transform-input`: Transpose input data to channel-last format
- `--no-wait`: Do not wait for push to complete

---

## Secrets Management Commands

### `leap secrets` - Secrets Management

#### `leap secrets create [name] [secretKeyPath]`
**Description**: Create a new secret

**Usage**:
```bash
leap secrets create [name] [secretKeyPath] [flags]
leap secrets create "aws-creds" "./aws-key.json"
leap secrets create "api-token" --secret-key-content "my-secret-content"
```

**Options**:
- `-k, --secret-key-content string`: Secret key content

#### `leap secrets list` (alias: `ls`)
**Description**: List secrets

**Usage**:
```bash
leap secrets list
leap secrets ls
```

#### `leap secrets delete [name]`
**Description**: Delete a secret

**Usage**:
```bash
leap secrets delete [name]
leap secrets delete "old-credentials"
```

#### `leap secrets set`
**Description**: Set secret on local leap config

**Usage**:
```bash
leap secrets set [flags]
leap secrets set --secret-id SECRET_ID
```

**Options**:
- `--secret-id string`: Secret ID

---

## Hub Commands

### `leap hub` - Hub Management
**Availability**: Only when `LEAP_HUB_ENABLED=true` environment variable is set

#### `leap hub list` (alias: `ls`)
**Description**: List projects from the hub

**Usage**:
```bash
leap hub list
leap hub ls
```

#### `leap hub download [project name]`
**Description**: Download project from the hub

**Usage**:
```bash
leap hub download [project_name] [flags]
leap hub download "Public Model" --version "1.2.0"
```

**Options**:
- `-v, --version string`: Version of the project to download (defaults to latest)

#### `leap hub publish <project tar file path>`
**Description**: Publish project to the hub

**Usage**:
```bash
leap hub publish <project_tar_file_path> [flags]
leap hub publish ./my-project.tar.gz --override
```

**Options**:
- `-o, --override`: Override existing project

#### `leap hub delete [project name]`
**Description**: Delete project from the hub

**Usage**:
```bash
leap hub delete [project_name] [flags]
leap hub delete "My Project" --version "1.0.0"
leap hub delete "My Project" --all
```

**Options**:
- `-v, --version string`: Version of the project to delete (defaults to latest)
- `-a, --all`: Delete all versions of the project

---

## CLI Management Commands

### `leap cli` - CLI Management

#### `leap cli`
**Description**: Manage leap cli

**Usage**:
```bash
leap cli [flags]
leap cli --version
```

**Options**:
- `--version`: Show version information

#### `leap cli upgrade`
**Description**: Helps to upgrade the leap cli

**Usage**:
```bash
leap cli upgrade [flags]
leap cli upgrade --script
```

**Options**:
- `-s, --script`: Print the install script to stdout

---

## Platform Support & Availability

### Conditional Features

#### Hub Commands
**Requirement**: `LEAP_HUB_ENABLED=true` environment variable must be set
**Affected Commands**:
- All `leap hub` commands
- `leap projects publish`


### Common Flag Patterns

#### Confirmation Flags
- `-y, --skipConfirm`: Skip confirmation prompts (used in delete operations)

#### Async Operation Flags  
- `--no-wait`: Don't wait for operations to complete
- `-f, --force`: Force operations

#### Version & Branch Flags
- `-m, --message`: Commit/version messages
- `-b, --branch`: Branch specifications
- `--model-branch`: Model-specific branch
- `--code-branch`: Code-specific branch

#### Output & Directory Flags
- `-o, --output-dir`: Output directories  
- `-o, --override`: Override existing resources

#### Authentication Flags
- `--secretId`: Secret manager references
- `--codeId`: Code integration references
- `--projectId`: Project references

#### Model-Specific Flags
- `--type`: Model file type [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]
- `--transform-input`: Transpose input data to channel-last format

#### Caching & Performance Flags
- `--no-cache`: Disable caching
- `--use-export-cache`: Use export cache
- `--exclude-calc-files`: Exclude calculated files to reduce size

### Command Aliases

- `leap projects` → `leap project`, `leap pro`
- `leap code` → `leap code-integration`  
- `leap projects list` → `leap projects ls`
- `leap projects copy` → `leap projects cp`
- `leap code list` → `leap code ls`
- `leap secrets list` → `leap secrets ls`
- `leap hub list` → `leap hub ls`

---

## Environment Variables

| Variable | Description | Impact |
|----------|-------------|---------|
| `LEAP_HUB_ENABLED=true` | Enables hub functionality | Makes `leap hub` and `leap projects publish` available |

---

## Configuration Context

The leap CLI operates in the context of:
1. **Global Configuration**: Authentication, API endpoints
2. **Workspace Configuration**: Project-specific settings, code integration bindings
3. **Environment Selection**: Multi-environment support for different TensorLeap instances

Most commands automatically detect and use the appropriate context based on the current directory and global configuration.

---

*This documentation covers the complete leap CLI as of the current codebase analysis. For the most up-to-date information, run `leap [command] --help` for any specific command.*