package k3d

import (
	"fmt"
	"testing"

	"github.com/tensorleap/leap-cli/pkg/server/manifest"
)

func TestCreateMirrorFromManifest(t *testing.T) {
	mfs := manifest.InstallationManifest{
		Images: manifest.ManifestImages{

			ServerImages: []string{
				"docker.elastic.co/elasticsearch/elasticsearch:7.10.2",
				"docker.io/bitnami/mongodb:5.0.9-debian-11-r3",
				"docker.io/library/mongo:5.0.11",
				"docker.io/library/rabbitmq:3.9.22",
				"quay.io/minio/minio:RELEASE.2021-12-20T22-07-16Z",
				"registry.k8s.io/ingress-nginx/controller:v1.8.0",
				"registry.k8s.io/ingress-nginx/kube-webhook-certgen:v20230407",
				"us-central1-docker.pkg.dev/tensorleap/main/engine:master-40246f22-stable",
				"us-central1-docker.pkg.dev/tensorleap/main/node-server:master-cc43e60b-stable",
				"us-central1-docker.pkg.dev/tensorleap/main/web-ui:master-6ea417b8-stable",
			},
			K3sImages: []string{
				"docker.io/rancher/klipper-helm:v0.7.7-build20230403",
				"docker.io/rancher/klipper-lb:v0.4.3",
				"docker.io/rancher/local-path-provisioner:v0.0.24",
				"docker.io/rancher/mirrored-coredns-coredns:1.10.1",
				"docker.io/rancher/mirrored-library-busybox:1.34.1",
				"docker.io/rancher/mirrored-library-traefik:2.9.4",
				"docker.io/rancher/mirrored-metrics-server:v0.6.2",
				"docker.io/rancher/mirrored-pause:3.6",
			},
			K3sGpuImages: []string{
				"docker.io/rancher/klipper-helm:v0.7.7-build20230403",
				"docker.io/rancher/klipper-lb:v0.4.3",
				"docker.io/rancher/local-path-provisioner:v0.0.24",
				"docker.io/rancher/mirrored-coredns-coredns:1.10.1",
				"docker.io/rancher/mirrored-library-busybox:1.34.1",
				"docker.io/rancher/mirrored-library-traefik:2.9.4",
				"docker.io/rancher/mirrored-metrics-server:v0.6.2",
				"docker.io/rancher/mirrored-pause:3.6",
			},
		},
	}

	mirrorConfig, err := CreateMirrorFromManifest(&mfs, "some url")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(mirrorConfig)
}
