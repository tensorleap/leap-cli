# Tensorleap CLI

## Usage

## Installation
Run `tensorleap --help` to see available options.

Also available [here](docs/tensorleap.md)

### Automatic installation with script

with `curl`:
```sh
curl -s https://raw.githubusercontent.com/tensorleap/cli-go/master/install.sh | bash
```

with `wget`:
```sh
wget -q -O - https://raw.githubusercontent.com/tensorleap/cli-go/master/install.sh | bash
```

### Manually
1. Download the binary from the [releases page](https://github.com/tensorleap/cli-go/releases) that matches your OS and machine architecture.
2. Rename the binary to `tensorleap`
3. Save it `/usr/local/bin` (or anywhere else in your `$PATH`)

### Shell Completions
In order to enabled shell completions, add `source <(tensorleap completion zsh)` to your `~/.zshrc` file

## Building locally
1. Clone this repo
2. Run `go build .` (the built file will be named `cli-go`)

you can also run `go run .` to compile and run the source

### Updating docs 

```sh
make docgen
```

### Uninstall
```sh
sudo rm $(which tensorleap)
```

### Development requirements

> brew install go
> brew install golangci-lint
