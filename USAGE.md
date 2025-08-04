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
10. [Server Management Commands](#server-management-commands)
11. [CLI Management Commands](#cli-management-commands)
12. [Platform Support & Availability](#platform-support--availability)

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

## Server Management Commands

### `leap server` - Local Server Management
**Platform Support**: Unix/Linux/macOS only (Windows not supported)

#### `leap server`
**Description**: Manage local server installation of TensorLeap

**Usage**:
```bash
leap server [flags]
leap server --info
```

**Options**:
- `-i, --info`: Show server info

#### `leap server install`
**Description**: Server installation command
**Note**: This command's flags are dynamically set by the helm-charts package

**Usage**:
```bash
leap server install [flags]
```

#### `leap server upgrade`
**Description**: Server upgrade command
**Note**: This command's flags are dynamically set by the helm-charts package

**Usage**:
```bash
leap server upgrade [flags]
```

#### `leap server reinstall`
**Description**: Reinstall TensorLeap
**Note**: This command's flags are dynamically set by the helm-charts package

**Usage**:
```bash
leap server reinstall [flags]
```

#### `leap server uninstall`
**Description**: Remove local TensorLeap installation

**Usage**:
```bash
leap server uninstall [flags]
leap server uninstall --purge
```

**Options**:
- `--purge`: Remove all data and cached files

#### `leap server info`
**Description**: Server installation information

**Usage**:
```bash
leap server info
```

#### `leap server check`
**Description**: Troubleshoot local installation issues

**Usage**:
```bash
leap server check
```

**Note**: Run this if your local installation stopped working.

### Server Tools - Third-party Management Tools

#### `leap server tools` - Third-party Tools Hub
**Description**: 3rd party tools to help manage local environment

**Usage**:
```bash
leap server tools [subcommand]
```

### K3D Integration - `leap server tools k3d`

Complete k3d (k3s in Docker) integration for Kubernetes cluster management.

#### `leap server tools k3d`
**Description**: https://k3d.io/ -> Run k3s in Docker!

**Usage**:
```bash
leap server tools k3d [flags]
```

**Options**:
- `--timestamps`: Enable log timestamps
- `--trace`: Enable super verbose output (trace logging)
- `--verbose`: Enable verbose output (debug logging)
- `--version`: Show k3d and default k3s version

#### K3D Cluster Management

**`leap server tools k3d cluster`** - Manage cluster(s)

Sub-commands:
- `leap server tools k3d cluster create` - Create a new k3s cluster
- `leap server tools k3d cluster delete` - Delete cluster(s)
- `leap server tools k3d cluster edit` - Edit cluster(s)
- `leap server tools k3d cluster list` - List cluster(s)
- `leap server tools k3d cluster start` - Start existing k3s cluster(s)
- `leap server tools k3d cluster stop` - Stop existing k3s cluster(s)

#### K3D Node Management

**`leap server tools k3d node`** - Manage node(s)

Sub-commands:
- `leap server tools k3d node create` - Create a new k3s node
- `leap server tools k3d node delete` - Delete node(s)
- `leap server tools k3d node edit` - Edit node(s)
- `leap server tools k3d node list` - List node(s)
- `leap server tools k3d node start` - Start node(s)
- `leap server tools k3d node stop` - Stop node(s)

#### K3D Registry Management

**`leap server tools k3d registry`** - Manage registry/registries  

Sub-commands:
- `leap server tools k3d registry create` - Create a new registry
- `leap server tools k3d registry delete` - Delete registry/registries
- `leap server tools k3d registry list` - List registries

#### K3D Configuration Management

**`leap server tools k3d config`** - Work with config file(s)

Sub-commands:
- `leap server tools k3d config init` - Create a new k3d config file
- `leap server tools k3d config migrate` - Migrate old config file format to new format

#### K3D Image Management

**`leap server tools k3d image`** - Handle container images

Sub-commands:
- `leap server tools k3d image import` - Import image(s) from docker into k3d cluster(s)

#### K3D Kubeconfig Management

**`leap server tools k3d kubeconfig`** - Manage kubeconfig(s)

Sub-commands:
- `leap server tools k3d kubeconfig get` - Get kubeconfig(s)
- `leap server tools k3d kubeconfig merge` - Merge kubeconfig(s)

#### K3D Additional Commands

- `leap server tools k3d completion` - Generate completion scripts for [bash, zsh, fish, powershell | psh]
- `leap server tools k3d version` - Show k3d and default k3s version
  - `leap server tools k3d version list` - List available k3s versions

### Kubectl Integration - `leap server tools kubectl`

Complete kubectl integration for Kubernetes cluster management.

#### `leap server tools kubectl`
**Description**: kubectl controls the Kubernetes cluster manager

**Usage**:
```bash
leap server tools kubectl [flags]
```

**Standard kubectl Options**:
- `--as string`: Username to impersonate for the operation
- `--as-group stringArray`: Group to impersonate for the operation (repeatable)
- `--as-uid string`: UID to impersonate for the operation
- `--cache-dir string`: Default cache directory
- `--certificate-authority string`: Path to a cert file for the certificate authority
- `--client-certificate string`: Path to a client certificate file for TLS
- `--client-key string`: Path to a client key file for TLS
- `--cluster string`: The name of the kubeconfig cluster to use
- `--context string`: The name of the kubeconfig context to use
- `--disable-compression`: Opt-out of response compression for all requests
- `--insecure-skip-tls-verify`: Skip server certificate validity check
- `--kubeconfig string`: Path to the kubeconfig file
- `--match-server-version`: Require server version to match client version
- `-n, --namespace string`: Namespace scope for this CLI request
- `--password string`: Password for basic authentication
- `--profile string`: Profile to capture [none|cpu|heap|goroutine|threadcreate|block|mutex]
- `--profile-output string`: File to write the profile to
- `--request-timeout string`: Timeout for server requests
- `-s, --server string`: Address and port of the Kubernetes API server
- `--tls-server-name string`: Server name for server certificate validation
- `--token string`: Bearer token for authentication
- `--user string`: Name of the kubeconfig user to use
- `--username string`: Username for basic authentication
- `--warnings-as-errors`: Treat warnings as errors

#### Kubectl Resource Management Commands

**Core Resource Operations**:
- `leap server tools kubectl get` - Display one or many resources
- `leap server tools kubectl describe` - Show details of a specific resource
- `leap server tools kubectl create` - Create a resource from a file or stdin
- `leap server tools kubectl apply` - Apply a configuration to a resource
- `leap server tools kubectl delete` - Delete resources
- `leap server tools kubectl edit` - Edit a resource on the server
- `leap server tools kubectl replace` - Replace a resource
- `leap server tools kubectl patch` - Update fields of a resource

**Resource Creation Commands**:
- `leap server tools kubectl create clusterrole` - Create a cluster role
- `leap server tools kubectl create clusterrolebinding` - Create a cluster role binding
- `leap server tools kubectl create configmap` - Create a configmap
- `leap server tools kubectl create cronjob` - Create a cronjob
- `leap server tools kubectl create deployment` - Create a deployment
- `leap server tools kubectl create ingress` - Create an ingress
- `leap server tools kubectl create job` - Create a job
- `leap server tools kubectl create namespace` - Create a namespace
- `leap server tools kubectl create poddisruptionbudget` - Create a pod disruption budget
- `leap server tools kubectl create priorityclass` - Create a priority class
- `leap server tools kubectl create quota` - Create a quota
- `leap server tools kubectl create role` - Create a role
- `leap server tools kubectl create rolebinding` - Create a role binding
- `leap server tools kubectl create secret` - Create a secret
  - `leap server tools kubectl create secret docker-registry` - Create a docker-registry secret
  - `leap server tools kubectl create secret generic` - Create a generic secret
  - `leap server tools kubectl create secret tls` - Create a TLS secret
- `leap server tools kubectl create service` - Create a service
  - `leap server tools kubectl create service clusterip` - Create a ClusterIP service
  - `leap server tools kubectl create service externalname` - Create an ExternalName service  
  - `leap server tools kubectl create service loadbalancer` - Create a LoadBalancer service
  - `leap server tools kubectl create service nodeport` - Create a NodePort service
- `leap server tools kubectl create serviceaccount` - Create a service account
- `leap server tools kubectl create token` - Create a token

**Workload Management**:
- `leap server tools kubectl run` - Run a particular image on the cluster
- `leap server tools kubectl expose` - Expose a resource as a new Kubernetes service
- `leap server tools kubectl scale` - Set a new size for a deployment, replica set, or replication controller
- `leap server tools kubectl autoscale` - Auto-scale a deployment, replica set, stateful set, or replication controller

**Rollout Management**:
- `leap server tools kubectl rollout` - Manage the rollout of a resource
  - `leap server tools kubectl rollout history` - View rollout history
  - `leap server tools kubectl rollout pause` - Mark the provided resource as paused
  - `leap server tools kubectl rollout restart` - Restart a resource
  - `leap server tools kubectl rollout resume` - Resume a paused resource
  - `leap server tools kubectl rollout status` - Show the status of the rollout
  - `leap server tools kubectl rollout undo` - Undo a previous rollout

**Configuration Management**:
- `leap server tools kubectl config` - Modify kubeconfig files
  - `leap server tools kubectl config current-context` - Display current-context
  - `leap server tools kubectl config delete-cluster` - Delete the specified cluster
  - `leap server tools kubectl config delete-context` - Delete the specified context
  - `leap server tools kubectl config delete-user` - Delete the specified user
  - `leap server tools kubectl config get-clusters` - Display clusters
  - `leap server tools kubectl config get-contexts` - Display list of contexts
  - `leap server tools kubectl config get-users` - Display users
  - `leap server tools kubectl config rename-context` - Rename a context
  - `leap server tools kubectl config set-cluster` - Set a cluster entry
  - `leap server tools kubectl config set-context` - Set a context entry
  - `leap server tools kubectl config set-credentials` - Set a user entry
  - `leap server tools kubectl config set` - Set an individual value
  - `leap server tools kubectl config unset` - Unset an individual value
  - `leap server tools kubectl config use-context` - Set the current-context
  - `leap server tools kubectl config view` - Display merged kubeconfig settings

**Resource Modification**:
- `leap server tools kubectl annotate` - Update the annotations on a resource
- `leap server tools kubectl label` - Update the labels on a resource
- `leap server tools kubectl set` - Set specific features on objects
  - `leap server tools kubectl set env` - Update environment variables
  - `leap server tools kubectl set image` - Update image of a pod template
  - `leap server tools kubectl set resources` - Update resource requests/limits
  - `leap server tools kubectl set selector` - Set the selector on a resource
  - `leap server tools kubectl set serviceaccount` - Update ServiceAccount of a resource
  - `leap server tools kubectl set subject` - Update User, Group or ServiceAccount

**Node Management**:
- `leap server tools kubectl cordon` - Mark node as unschedulable
- `leap server tools kubectl uncordon` - Mark node as schedulable
- `leap server tools kubectl drain` - Drain node in preparation for maintenance
- `leap server tools kubectl taint` - Update the taints on one or more nodes

**Debugging & Monitoring**:
- `leap server tools kubectl logs` - Print the logs for a container in a pod
- `leap server tools kubectl exec` - Execute a command in a container
- `leap server tools kubectl attach` - Attach to a running container
- `leap server tools kubectl debug` - Create debugging sessions for troubleshooting workloads and nodes
- `leap server tools kubectl port-forward` - Forward one or more local ports to a pod
- `leap server tools kubectl proxy` - Run a proxy to the Kubernetes API server
- `leap server tools kubectl cp` - Copy files and directories to and from containers

**Resource Information**:
- `leap server tools kubectl explain` - Get documentation for a resource
- `leap server tools kubectl api-resources` - Print the supported API resources on the server
- `leap server tools kubectl api-versions` - Print the supported API versions on the server
- `leap server tools kubectl version` - Print the client and server version information
- `leap server tools kubectl cluster-info` - Display cluster information
  - `leap server tools kubectl cluster-info dump` - Dump cluster information

**Resource Monitoring**:
- `leap server tools kubectl top` - Display resource (CPU/memory) usage
  - `leap server tools kubectl top node` - Display resource usage of nodes
  - `leap server tools kubectl top pod` - Display resource usage of pods
- `leap server tools kubectl events` - List events
- `leap server tools kubectl wait` - Wait for a specific condition on one or many resources

**Authentication & Authorization**:
- `leap server tools kubectl auth` - Inspect authorization
  - `leap server tools kubectl auth can-i` - Check whether an action is allowed
  - `leap server tools kubectl auth reconcile` - Reconcile rules for RBAC objects
  - `leap server tools kubectl auth whoami` - Display information about the current user

**Certificate Management**:
- `leap server tools kubectl certificate` - Modify certificate resources
  - `leap server tools kubectl certificate approve` - Approve a certificate signing request
  - `leap server tools kubectl certificate deny` - Deny a certificate signing request

**Advanced Operations**:
- `leap server tools kubectl diff` - Diff the live version against a would-be applied version
- `leap server tools kubectl kustomize` - Build a kustomization target from a directory or URL
- `leap server tools kubectl plugin` - Provides utilities for interacting with plugins
  - `leap server tools kubectl plugin list` - List all available plugin files

**Apply Sub-commands**:
- `leap server tools kubectl apply edit-last-applied` - Edit latest last-applied-configuration annotations
- `leap server tools kubectl apply set-last-applied` - Set the last-applied-configuration annotation
- `leap server tools kubectl apply view-last-applied` - View latest last-applied-configuration annotations

**Utility Commands**:
- `leap server tools kubectl completion` - Output shell completion code for the specified shell (bash, zsh, fish, or powershell)
- `leap server tools kubectl options` - Print the list of flags inherited by all commands

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

#### Server Commands  
**Platform Support**:
- **Unix/Linux/macOS**: Full functionality
- **Windows**: Not supported - displays "This command is not supported on this platform"

**Affected Commands**:
- `leap server install`
- `leap server upgrade`  
- `leap server reinstall`
- `leap server uninstall`
- `leap server info`
- `leap server check`
- All `leap server tools` commands

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