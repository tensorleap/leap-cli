## leap server tools kubectl config

Modify kubeconfig files

### Synopsis

Modify kubeconfig files using subcommands like "kubectl config set current-context my-context"

 The loading order follows these rules:

  1.  If the --kubeconfig flag is set, then only that file is loaded. The flag may only be set once and no merging takes place.
  2.  If $KUBECONFIG environment variable is set, then it is used as a list of paths (normal path delimiting rules for your system). These paths are merged. When a value is modified, it is modified in the file that defines the stanza. When a value is created, it is created in the first file that exists. If no files in the chain exist, then it creates the last file in the list.
  3.  Otherwise, ${HOME}/.kube/config is used and no merging takes place.

```
leap server tools kubectl config SUBCOMMAND
```

### Options

```
  -h, --help                help for config
      --kubeconfig string   use a particular kubeconfig file
```

### Options inherited from parent commands

```
      --apiKey string                  Tensorleap Api key
      --apiUrl string                  Tensorleap api url
      --as string                      Username to impersonate for the operation. User could be a regular user or a service account in a namespace.
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --as-uid string                  UID to impersonate for the operation.
      --cache-dir string               Default cache directory (default "/Users/moshekabala/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --config string                  config file (default is $HOME/.config/tensorleap/config.yaml)
      --context string                 The name of the kubeconfig context to use
      --disable-compression            If true, opt-out of response compression for all requests to the server
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --match-server-version           Require server version to match client version
  -n, --namespace string               If present, the namespace scope for this CLI request
      --password string                Password for basic authentication to the API server
      --profile string                 Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex) (default "none")
      --profile-output string          Name of the file to write the profile to (default "profile.pprof")
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
      --user string                    The name of the kubeconfig user to use
      --username string                Username for basic authentication to the API server
      --warnings-as-errors             Treat warnings received from the server as errors and exit with a non-zero exit code
```

### SEE ALSO

* [leap server tools kubectl](leap_server_tools_kubectl.md)	 - kubectl controls the Kubernetes cluster manager
* [leap server tools kubectl config current-context](leap_server_tools_kubectl_config_current-context.md)	 - Display the current-context
* [leap server tools kubectl config delete-cluster](leap_server_tools_kubectl_config_delete-cluster.md)	 - Delete the specified cluster from the kubeconfig
* [leap server tools kubectl config delete-context](leap_server_tools_kubectl_config_delete-context.md)	 - Delete the specified context from the kubeconfig
* [leap server tools kubectl config delete-user](leap_server_tools_kubectl_config_delete-user.md)	 - Delete the specified user from the kubeconfig
* [leap server tools kubectl config get-clusters](leap_server_tools_kubectl_config_get-clusters.md)	 - Display clusters defined in the kubeconfig
* [leap server tools kubectl config get-contexts](leap_server_tools_kubectl_config_get-contexts.md)	 - Describe one or many contexts
* [leap server tools kubectl config get-users](leap_server_tools_kubectl_config_get-users.md)	 - Display users defined in the kubeconfig
* [leap server tools kubectl config rename-context](leap_server_tools_kubectl_config_rename-context.md)	 - Rename a context from the kubeconfig file
* [leap server tools kubectl config set](leap_server_tools_kubectl_config_set.md)	 - Set an individual value in a kubeconfig file
* [leap server tools kubectl config set-cluster](leap_server_tools_kubectl_config_set-cluster.md)	 - Set a cluster entry in kubeconfig
* [leap server tools kubectl config set-context](leap_server_tools_kubectl_config_set-context.md)	 - Set a context entry in kubeconfig
* [leap server tools kubectl config set-credentials](leap_server_tools_kubectl_config_set-credentials.md)	 - Set a user entry in kubeconfig
* [leap server tools kubectl config unset](leap_server_tools_kubectl_config_unset.md)	 - Unset an individual value in a kubeconfig file
* [leap server tools kubectl config use-context](leap_server_tools_kubectl_config_use-context.md)	 - Set the current-context in a kubeconfig file
* [leap server tools kubectl config view](leap_server_tools_kubectl_config_view.md)	 - Display merged kubeconfig settings or a specified kubeconfig file

###### Auto generated by spf13/cobra on 7-Aug-2023
