# Leap CLI

## Usage

## Installation

Run `leap --help` to see available options.

Also available [here](docs/leap.md)

### Automatic installation with script

with `curl`:

```sh
curl -s https://raw.githubusercontent.com/tensorleap/leap-cli/master/install.sh | bash
```

with `wget`:

```sh
wget -q -O - https://raw.githubusercontent.com/tensorleap/leap-cli/master/install.sh | bash
```

## Connect an AI assistant (MCP)

Right after installing leap, before you have a Tensorleap server, run:

```sh
leap mcp config
```

It asks which assistant (Claude Code, Cursor, VS Code/Copilot, Codex, Windsurf, Claude Desktop),
registers the Tensorleap MCP server and installs every
[Tensorleap skill](https://github.com/tensorleap/skills) the assistant supports. Skills already
installed are left alone; each keeps itself up to date. It ends with the next thing to ask the
assistant, and leap suggests each step after it once, when it becomes possible: install
Tensorleap (after `leap mcp config`), write the integration for your model (after
`leap server install` or `leap auth login`), then the analysis report (after `leap push --eval`).
Run `leap mcp` in a terminal to check the connection once you are signed in.

### Manually

1. Download the binary from the [releases page](https://github.com/tensorleap/leap-cli/releases) that matches your OS and machine architecture.
2. Rename the binary to `leap`
3. Save it `/usr/local/bin` (or anywhere else in your `$PATH`)

### Windows installation

To install on Windows, go to the [releases page](https://github.com/tensorleap/leap-cli/releases)  and download the Windows installer. Please note that the server installation is not supported on Windows; you can only use this installation to connect to an existing server.

### Shell Completions

In order to enabled shell completions, add `source <(leap completion zsh)` to your `~/.zshrc` file

## Building locally

1. Clone this repo
2. Run `go build .` (the built file will be named `leap-cli`)

you can also run `go run .` to compile and run the source

### Updating docs

```sh
make docgen
```

### Uninstall

```sh
sudo rm $(which leap)
```

### Development requirements

> brew install go

> brew install golangci-lint
