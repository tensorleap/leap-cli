## leap server install

Installs tensorleap on the local machine, running in a docker container

### Synopsis

Installs tensorleap on the local machine, running in a docker container

```
leap server install [flags]
```

### Options

```
      --data-volume string   Data Volume maps the user's local directory to the container's directory, enabling access to code integration for training and evaluation (default "/Users/moshekabala/tensorleap/data:/Users/moshekabala/tensorleap/data")
      --gpu                  Enable GPU usage for training and evaluating
  -h, --help                 help for install
  -p, --port uint            Port to be used for tensorleap installation (default 4589)
      --registry-port uint   Port to be used for docker registry (default 5699)
```

### Options inherited from parent commands

```
      --apiKey string   Tensorleap Api key
      --apiUrl string   Tensorleap api url
      --config string   config file (default is $HOME/.config/tensorleap/config.yaml)
```

### SEE ALSO

* [leap server](leap_server.md)	 - Manage local server installation of Tensorleap

###### Auto generated by spf13/cobra on 25-Jul-2023