package k3d

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	cliutil "github.com/k3d-io/k3d/v5/cmd/util"
	k3dCluster "github.com/k3d-io/k3d/v5/pkg/client"
	"github.com/k3d-io/k3d/v5/pkg/config"
	k3dConfTypes "github.com/k3d-io/k3d/v5/pkg/config/types"
	conf "github.com/k3d-io/k3d/v5/pkg/config/v1alpha5"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
	k3d "github.com/k3d-io/k3d/v5/pkg/types"
	"github.com/k3d-io/k3d/v5/version"
	"github.com/tensorleap/leap-cli/pkg/log"
)

type Cluster = k3d.Cluster

const CLUSTER_NAME = "tensorleap"

var (
	K3sVersion          = version.K3sVersion
	K3sImage            = fmt.Sprintf("%s:%s", k3d.DefaultK3sImageRepo, K3sVersion)
	K3sGpuVersion       = "v1.26.4-k3s1"
	K3sGpuVersionSuffix = "cuda-11.8.0-ubuntu-22.04"
	K3sGpuImage         = fmt.Sprintf("us-central1-docker.pkg.dev/tensorleap/main/k3s:%s-%s", K3sGpuVersion, K3sGpuVersionSuffix)
)

func GetCluster(ctx context.Context) (*Cluster, error) {
	clusters, err := k3dCluster.ClusterList(ctx, runtimes.SelectedRuntime)
	if err != nil {
		return nil, err
	}

	for _, cluster := range clusters {
		if cluster.Name == CLUSTER_NAME {
			return cluster, nil
		}
	}

	return nil, nil
}

func CreateCluster(ctx context.Context, port uint, volumes []string, useGpu bool) error {
	log.SendCloudReport("info", "Creating cluster", "Running", &map[string]interface{}{"useGpu": useGpu, "port": port})
	clusterConfig := createClusterConfig(ctx, port, volumes, useGpu)

	if _, err := k3dCluster.ClusterGet(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster); err == nil {
		log.Println("Found existing tensorleap cluster!")
		log.SendCloudReport("info", "Cluster already exists", "Running", &map[string]interface{}{"useGpu": useGpu, "port": port})
		return nil
	}

	if err := k3dCluster.ClusterRun(ctx, runtimes.SelectedRuntime, clusterConfig); err != nil {
		log.Println(err)
		log.Println("Failed to create cluster >>> Rolling Back")
		log.SendCloudReport("error", "Failed creating cluster", "Failed",
			&map[string]interface{}{"selectedRuntime": runtimes.SelectedRuntime, "error": err.Error()})
		if err := k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, k3d.ClusterDeleteOpts{SkipRegistryCheck: true}); err != nil {
			log.Println(err)
			log.Fatalln("Cluster creation FAILED, also FAILED to rollback changes!")
			log.SendCloudReport("error", "Failed rolling back cluster changes", "Failed",
				&map[string]interface{}{"error": err.Error()})
		}
		log.SendCloudReport("error", "Successfully rolled back cluster changes", "Failed", nil)
		log.Fatalln("Cluster creation FAILED, all changes have been rolled back!")
	}
	log.Printf("Cluster '%s' created successfully!\n", clusterConfig.Cluster.Name)
	log.SendCloudReport("info", "Created cluster successfully", "Running", nil)

	if _, err := k3dCluster.KubeconfigGetWrite(ctx, runtimes.SelectedRuntime, &clusterConfig.Cluster, "", &k3dCluster.WriteKubeConfigOptions{
		UpdateExisting:       true,
		OverwriteExisting:    false,
		UpdateCurrentContext: true,
	}); err != nil {
		log.Println(err)
	}

	return nil
}

func IsGpuCluster(cluster *Cluster) bool {
	return len(cluster.Nodes) > 0 && strings.Contains(cluster.Nodes[0].Image, "cuda")
}

func createClusterConfig(ctx context.Context, port uint, volumes []string, useGpu bool) *conf.ClusterConfig {
	freePort, err := cliutil.GetFreePort()
	if err != nil {
		log.Fatalln(err)
	}

	image := K3sImage
	if useGpu {
		image = K3sGpuImage
	}

	simpleK3dConfig := conf.SimpleConfig{
		TypeMeta: k3dConfTypes.TypeMeta{
			Kind:       "Simple",
			APIVersion: "k3d.io/v1alpha5",
		},
		ObjectMeta: k3dConfTypes.ObjectMeta{
			Name: CLUSTER_NAME,
		},
		Servers: 1,
		ExposeAPI: conf.SimpleExposureOpts{
			HostPort: strconv.Itoa(freePort),
		},
		Image:   image,
		Volumes: make([]conf.VolumeWithNodeFilters, len(volumes)),
		Ports: []conf.PortWithNodeFilters{{
			Port:        fmt.Sprintf("%v:80", port),
			NodeFilters: []string{"server:*:direct"},
		}},
		Env: []conf.EnvVarWithNodeFilters{
			{
				EnvVar:      fmt.Sprintf("all_proxy=%s", os.Getenv("all_proxy")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("ALL_PROXY=%s", os.Getenv("ALL_PROXY")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("http_proxy=%s", os.Getenv("http_proxy")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("HTTP_PROXY=%s", os.Getenv("HTTP_PROXY")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("https_proxy=%s", os.Getenv("https_proxy")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("HTTPS_PROXY=%s", os.Getenv("HTTPS_PROXY")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("no_proxy=%s", os.Getenv("no_proxy")),
				NodeFilters: []string{"server:*"},
			},
			{
				EnvVar:      fmt.Sprintf("NO_PROXY=%s", os.Getenv("NO_PROXY")),
				NodeFilters: []string{"server:*"},
			},
		},
		Registries: conf.SimpleConfigRegistries{
			Use: []string{"tensorleap-registry"},
			Config: `
mirrors:
  docker.io:
    endpoint:
      - http://k3d-tensorleap-registry:5000
  k8s.gcr.io:
    endpoint:
      - http://k3d-tensorleap-registry:5000
  gcr.io:
    endpoint:
      - http://k3d-tensorleap-registry:5000
  docker.elastic.co:
    endpoint:
      - http://k3d-tensorleap-registry:5000
  quay.io:
    endpoint:
      - http://k3d-tensorleap-registry:5000
  us-central1-docker.pkg.dev:
    endpoint:
      - http://k3d-tensorleap-registry:5000
          `,
		},
		Options: conf.SimpleConfigOptions{
			K3dOptions: conf.SimpleConfigOptionsK3d{
				Wait:                true,
				DisableLoadbalancer: true,
			},
			K3sOptions: conf.SimpleConfigOptionsK3s{
				ExtraArgs: []conf.K3sArgWithNodeFilters{{
					Arg:         "--disable=traefik",
					NodeFilters: []string{"server:*"},
				}},
			},
			KubeconfigOptions: conf.SimpleConfigOptionsKubeconfig{
				UpdateDefaultKubeconfig: true,
				SwitchCurrentContext:    true,
			},
		},
	}
	if useGpu {
		simpleK3dConfig.Options.Runtime.GPURequest = "all"
	}

	for i, v := range volumes {
		simpleK3dConfig.Volumes[i] = conf.VolumeWithNodeFilters{
			Volume:      v,
			NodeFilters: []string{"server:*"},
		}
	}

	k3dClusterConfig, err := config.TransformSimpleToClusterConfig(ctx, runtimes.SelectedRuntime, simpleK3dConfig)
	if err != nil {
		log.Fatalln(err)
	}

	k3dClusterConfig, err = config.ProcessClusterConfig(*k3dClusterConfig)
	if err != nil {
		log.Fatalln(err)
	}

	return k3dClusterConfig
}

func UninstallCluster(ctx context.Context) error {
	cluster, err := GetCluster(ctx)
	if err != nil {
		log.SendCloudReport("error", "Failed getting cluster", "Failed", &map[string]interface{}{"error": err.Error()})
		return err
	}
	if cluster == nil {
		log.SendCloudReport("info", "Cluster not found", "Running", nil)
		log.Info("Cluster 'tensorleap' not found")
		return nil
	}
	log.Info("Deleting cluster 'tensorleap'")
	opt := k3d.ClusterDeleteOpts{
		SkipRegistryCheck: true,
	}
	err = k3dCluster.ClusterDelete(ctx, runtimes.SelectedRuntime, cluster, opt)
	if err != nil {
		log.SendCloudReport("error", "Failed deleting cluster", "Failed",
			&map[string]interface{}{"opt": opt, "cluster": cluster, "runtime": runtimes.SelectedRuntime, "error": err.Error()})
	}
	return err
}
