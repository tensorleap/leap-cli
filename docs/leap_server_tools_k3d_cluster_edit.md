## leap server tools k3d cluster edit

[EXPERIMENTAL] Edit cluster(s).

### Synopsis

[EXPERIMENTAL] Edit cluster(s).

```
leap server tools k3d cluster edit CLUSTER [flags]
```

### Options

```
  -h, --help                                                               help for edit
      --port-add [HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]   [EXPERIMENTAL] Map ports from the node containers (via the serverlb) to the host (Format: [HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER])
                                                                            - Example: `k3d node edit k3d-mycluster-serverlb --port-add 8080:80`
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

###### Auto generated by spf13/cobra on 7-Aug-2023
