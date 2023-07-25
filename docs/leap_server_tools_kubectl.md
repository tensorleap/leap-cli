## leap server tools kubectl

kubectl controls the Kubernetes cluster manager

### Synopsis

kubectl controls the Kubernetes cluster manager.

 Find more information at: https://kubernetes.io/docs/reference/kubectl/

```
leap server tools kubectl [flags]
```

### Options

```
      --as string                      Username to impersonate for the operation. User could be a regular user or a service account in a namespace.
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --as-uid string                  UID to impersonate for the operation.
      --cache-dir string               Default cache directory (default "/Users/moshekabala/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
      --context string                 The name of the kubeconfig context to use
      --disable-compression            If true, opt-out of response compression for all requests to the server
  -h, --help                           help for kubectl
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
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

### Options inherited from parent commands

```
      --apiKey string   Tensorleap Api key
      --apiUrl string   Tensorleap api url
      --config string   config file (default is $HOME/.config/tensorleap/config.yaml)
```

### SEE ALSO

* [leap server tools](leap_server_tools.md)	 - 3rd party tools to help manage local environment
* [leap server tools kubectl annotate](leap_server_tools_kubectl_annotate.md)	 - Update the annotations on a resource
* [leap server tools kubectl api-resources](leap_server_tools_kubectl_api-resources.md)	 - Print the supported API resources on the server
* [leap server tools kubectl api-versions](leap_server_tools_kubectl_api-versions.md)	 - Print the supported API versions on the server, in the form of "group/version"
* [leap server tools kubectl apply](leap_server_tools_kubectl_apply.md)	 - Apply a configuration to a resource by file name or stdin
* [leap server tools kubectl attach](leap_server_tools_kubectl_attach.md)	 - Attach to a running container
* [leap server tools kubectl auth](leap_server_tools_kubectl_auth.md)	 - Inspect authorization
* [leap server tools kubectl autoscale](leap_server_tools_kubectl_autoscale.md)	 - Auto-scale a deployment, replica set, stateful set, or replication controller
* [leap server tools kubectl certificate](leap_server_tools_kubectl_certificate.md)	 - Modify certificate resources.
* [leap server tools kubectl cluster-info](leap_server_tools_kubectl_cluster-info.md)	 - Display cluster information
* [leap server tools kubectl completion](leap_server_tools_kubectl_completion.md)	 - Output shell completion code for the specified shell (bash, zsh, fish, or powershell)
* [leap server tools kubectl config](leap_server_tools_kubectl_config.md)	 - Modify kubeconfig files
* [leap server tools kubectl cordon](leap_server_tools_kubectl_cordon.md)	 - Mark node as unschedulable
* [leap server tools kubectl cp](leap_server_tools_kubectl_cp.md)	 - Copy files and directories to and from containers
* [leap server tools kubectl create](leap_server_tools_kubectl_create.md)	 - Create a resource from a file or from stdin
* [leap server tools kubectl debug](leap_server_tools_kubectl_debug.md)	 - Create debugging sessions for troubleshooting workloads and nodes
* [leap server tools kubectl delete](leap_server_tools_kubectl_delete.md)	 - Delete resources by file names, stdin, resources and names, or by resources and label selector
* [leap server tools kubectl describe](leap_server_tools_kubectl_describe.md)	 - Show details of a specific resource or group of resources
* [leap server tools kubectl diff](leap_server_tools_kubectl_diff.md)	 - Diff the live version against a would-be applied version
* [leap server tools kubectl drain](leap_server_tools_kubectl_drain.md)	 - Drain node in preparation for maintenance
* [leap server tools kubectl edit](leap_server_tools_kubectl_edit.md)	 - Edit a resource on the server
* [leap server tools kubectl events](leap_server_tools_kubectl_events.md)	 - List events
* [leap server tools kubectl exec](leap_server_tools_kubectl_exec.md)	 - Execute a command in a container
* [leap server tools kubectl explain](leap_server_tools_kubectl_explain.md)	 - Get documentation for a resource
* [leap server tools kubectl expose](leap_server_tools_kubectl_expose.md)	 - Take a replication controller, service, deployment or pod and expose it as a new Kubernetes service
* [leap server tools kubectl get](leap_server_tools_kubectl_get.md)	 - Display one or many resources
* [leap server tools kubectl kustomize](leap_server_tools_kubectl_kustomize.md)	 - Build a kustomization target from a directory or URL
* [leap server tools kubectl label](leap_server_tools_kubectl_label.md)	 - Update the labels on a resource
* [leap server tools kubectl logs](leap_server_tools_kubectl_logs.md)	 - Print the logs for a container in a pod
* [leap server tools kubectl options](leap_server_tools_kubectl_options.md)	 - Print the list of flags inherited by all commands
* [leap server tools kubectl patch](leap_server_tools_kubectl_patch.md)	 - Update fields of a resource
* [leap server tools kubectl plugin](leap_server_tools_kubectl_plugin.md)	 - Provides utilities for interacting with plugins
* [leap server tools kubectl port-forward](leap_server_tools_kubectl_port-forward.md)	 - Forward one or more local ports to a pod
* [leap server tools kubectl proxy](leap_server_tools_kubectl_proxy.md)	 - Run a proxy to the Kubernetes API server
* [leap server tools kubectl replace](leap_server_tools_kubectl_replace.md)	 - Replace a resource by file name or stdin
* [leap server tools kubectl rollout](leap_server_tools_kubectl_rollout.md)	 - Manage the rollout of a resource
* [leap server tools kubectl run](leap_server_tools_kubectl_run.md)	 - Run a particular image on the cluster
* [leap server tools kubectl scale](leap_server_tools_kubectl_scale.md)	 - Set a new size for a deployment, replica set, or replication controller
* [leap server tools kubectl set](leap_server_tools_kubectl_set.md)	 - Set specific features on objects
* [leap server tools kubectl taint](leap_server_tools_kubectl_taint.md)	 - Update the taints on one or more nodes
* [leap server tools kubectl top](leap_server_tools_kubectl_top.md)	 - Display resource (CPU/memory) usage
* [leap server tools kubectl uncordon](leap_server_tools_kubectl_uncordon.md)	 - Mark node as schedulable
* [leap server tools kubectl version](leap_server_tools_kubectl_version.md)	 - Print the client and server version information
* [leap server tools kubectl wait](leap_server_tools_kubectl_wait.md)	 - Experimental: Wait for a specific condition on one or many resources

###### Auto generated by spf13/cobra on 25-Jul-2023
