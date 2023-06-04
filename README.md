# Tensorleap CLI

## Installation

with `curl`:
```sh
curl -s https://raw.githubusercontent.com/tensorleap/cli-go/master/install.sh | bash
```

with `wget`:
```sh
wget -q -O - https://raw.githubusercontent.com/tensorleap/cli-go/master/install.sh | bash
```

### Completions
In order to enabled shell completions, add `source <(tensorleap completion zsh)` to your `~/.zshrc` file

## Building locally
1. Clone this repo
2. Run `go build .` (the built file will be named `cli-go`)

you can also run `go run .` to compile and run the source
