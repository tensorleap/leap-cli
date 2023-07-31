## leap server tools k3d cluster start

Start existing k3d cluster(s)

### Synopsis

Start existing k3d cluster(s)

```
leap server tools k3d cluster start [NAME [NAME...] | --all] [flags]
```

### Options

```
  -a, --all                Start all existing clusters
  -h, --help               help for start
      --timeout duration   Maximum waiting time for '--wait' before canceling/returning.
      --wait               Wait for the server(s) (and loadbalancer) to be ready before returning. (default true)
```

### Options inherited from parent commands

```
      --apiKey string   Tensorleap Api key
      --apiUrl string   Tensorleap api url
      --config string   config file (default is $HOME/.config/tensorleap/config.yaml)
      --timestamps      Enable Log timestamps
      --trace           Enable super verbose output (trace logging)
      --verbose         Enable verbose output (debug logging)
```

### SEE ALSO

* [leap server tools k3d cluster](leap_server_tools_k3d_cluster.md)	 - Manage cluster(s)

###### Auto generated by spf13/cobra on 25-Jul-2023