package prune

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	kapi "k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/client/unversioned/fake"
	ktc "k8s.io/kubernetes/pkg/client/unversioned/testclient"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/sets"

	buildapi "github.com/openshift/origin/pkg/build/api"
	"github.com/openshift/origin/pkg/client/testclient"
	deployapi "github.com/openshift/origin/pkg/deploy/api"
	imageapi "github.com/openshift/origin/pkg/image/api"
)

type fakeRegistryPinger struct {
	err      error
	requests []string
}

func (f *fakeRegistryPinger) ping(registry string) error {
	f.requests = append(f.requests, registry)
	return f.err
}

func imageList(images ...imageapi.Image) imageapi.ImageList {
	return imageapi.ImageList{
		Items: images,
	}
}

const (
	layer1 = "tarsum.dev+sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	layer2 = "tarsum.dev+sha256:b194de3772ebbcdc8f244f663669799ac1cb141834b7cb8b69100285d357a2b0"
	layer3 = "tarsum.dev+sha256:c937c4bb1c1a21cc6d94340812262c6472092028972ae69b551b1a70d4276171"
	layer4 = "tarsum.dev+sha256:2aaacc362ac6be2b9e9ae8c6029f6f616bb50aec63746521858e47841b90fabd"
	layer5 = "tarsum.dev+sha256:e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
)

func agedImage(id, ref string, ageInMinutes int64) imageapi.Image {
	image := imageWithLayers(id, ref, layer1, layer2, layer3, layer4, layer5)

	if ageInMinutes >= 0 {
		image.CreationTimestamp = unversioned.NewTime(unversioned.Now().Add(time.Duration(-1*ageInMinutes) * time.Minute))
	}

	return image
}

func sizedImage(id, ref string, size int64) imageapi.Image {
	image := imageWithLayers(id, ref, layer1, layer2, layer3, layer4, layer5)
	image.CreationTimestamp = unversioned.NewTime(unversioned.Now().Add(time.Duration(-1) * time.Minute))
	image.DockerImageMetadata.Size = size

	return image
}

func image(id, ref string) imageapi.Image {
	return agedImage(id, ref, -1)
}

func imageWithLayers(id, ref string, layers ...string) imageapi.Image {
	image := imageapi.Image{
		ObjectMeta: kapi.ObjectMeta{
			Name: id,
			Annotations: map[string]string{
				imageapi.ManagedByOpenShiftAnnotation: "true",
			},
		},
		DockerImageReference: ref,
	}

	image.DockerImageLayers = []imageapi.ImageLayer{}
	for _, layer := range layers {
		image.DockerImageLayers = append(image.DockerImageLayers, imageapi.ImageLayer{Name: layer})
	}

	return image
}

func unmanagedImage(id, ref string, hasAnnotations bool, annotation, value string) imageapi.Image {
	image := imageWithLayers(id, ref)
	if !hasAnnotations {
		image.Annotations = nil
	} else {
		delete(image.Annotations, imageapi.ManagedByOpenShiftAnnotation)
		image.Annotations[annotation] = value
	}
	return image
}

func podList(pods ...kapi.Pod) kapi.PodList {
	return kapi.PodList{
		Items: pods,
	}
}

func pod(namespace, name string, phase kapi.PodPhase, containerImages ...string) kapi.Pod {
	return agedPod(namespace, name, phase, -1, containerImages...)
}

func agedPod(namespace, name string, phase kapi.PodPhase, ageInMinutes int64, containerImages ...string) kapi.Pod {
	pod := kapi.Pod{
		ObjectMeta: kapi.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: podSpec(containerImages...),
		Status: kapi.PodStatus{
			Phase: phase,
		},
	}

	if ageInMinutes >= 0 {
		pod.CreationTimestamp = unversioned.NewTime(unversioned.Now().Add(time.Duration(-1*ageInMinutes) * time.Minute))
	}

	return pod
}

func podSpec(containerImages ...string) kapi.PodSpec {
	spec := kapi.PodSpec{
		Containers: []kapi.Container{},
	}
	for _, image := range containerImages {
		container := kapi.Container{
			Image: image,
		}
		spec.Containers = append(spec.Containers, container)
	}
	return spec
}

func streamList(streams ...imageapi.ImageStream) imageapi.ImageStreamList {
	return imageapi.ImageStreamList{
		Items: streams,
	}
}

func stream(registry, namespace, name string, tags map[string]imageapi.TagEventList) imageapi.ImageStream {
	return agedStream(registry, namespace, name, -1, tags)
}

func agedStream(registry, namespace, name string, ageInMinutes int64, tags map[string]imageapi.TagEventList) imageapi.ImageStream {
	stream := imageapi.ImageStream{
		ObjectMeta: kapi.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Status: imageapi.ImageStreamStatus{
			DockerImageRepository: fmt.Sprintf("%s/%s/%s", registry, namespace, name),
			Tags: tags,
		},
	}

	if ageInMinutes >= 0 {
		stream.CreationTimestamp = unversioned.NewTime(unversioned.Now().Add(time.Duration(-1*ageInMinutes) * time.Minute))
	}

	return stream
}

func streamPtr(registry, namespace, name string, tags map[string]imageapi.TagEventList) *imageapi.ImageStream {
	s := stream(registry, namespace, name, tags)
	return &s
}

func tags(list ...namedTagEventList) map[string]imageapi.TagEventList {
	m := make(map[string]imageapi.TagEventList, len(list))
	for _, tag := range list {
		m[tag.name] = tag.events
	}
	return m
}

type namedTagEventList struct {
	name   string
	events imageapi.TagEventList
}

func tag(name string, events ...imageapi.TagEvent) namedTagEventList {
	return namedTagEventList{
		name: name,
		events: imageapi.TagEventList{
			Items: events,
		},
	}
}

func tagEvent(id, ref string) imageapi.TagEvent {
	return imageapi.TagEvent{
		Image:                id,
		DockerImageReference: ref,
	}
}

func rcList(rcs ...kapi.ReplicationController) kapi.ReplicationControllerList {
	return kapi.ReplicationControllerList{
		Items: rcs,
	}
}

func rc(namespace, name string, containerImages ...string) kapi.ReplicationController {
	return kapi.ReplicationController{
		ObjectMeta: kapi.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: kapi.ReplicationControllerSpec{
			Template: &kapi.PodTemplateSpec{
				Spec: podSpec(containerImages...),
			},
		},
	}
}

func dcList(dcs ...deployapi.DeploymentConfig) deployapi.DeploymentConfigList {
	return deployapi.DeploymentConfigList{
		Items: dcs,
	}
}

func dc(namespace, name string, containerImages ...string) deployapi.DeploymentConfig {
	return deployapi.DeploymentConfig{
		ObjectMeta: kapi.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: deployapi.DeploymentConfigSpec{
			Template: &kapi.PodTemplateSpec{
				Spec: podSpec(containerImages...),
			},
		},
	}
}

func bcList(bcs ...buildapi.BuildConfig) buildapi.BuildConfigList {
	return buildapi.BuildConfigList{
		Items: bcs,
	}
}

func bc(namespace, name, strategyType, fromKind, fromNamespace, fromName string) buildapi.BuildConfig {
	return buildapi.BuildConfig{
		ObjectMeta: kapi.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: buildapi.BuildConfigSpec{
			CommonSpec: commonSpec(strategyType, fromKind, fromNamespace, fromName),
		},
	}
}

func buildList(builds ...buildapi.Build) buildapi.BuildList {
	return buildapi.BuildList{
		Items: builds,
	}
}

func build(namespace, name, strategyType, fromKind, fromNamespace, fromName string) buildapi.Build {
	return buildapi.Build{
		ObjectMeta: kapi.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
		Spec: buildapi.BuildSpec{
			CommonSpec: commonSpec(strategyType, fromKind, fromNamespace, fromName),
		},
	}
}

func limitList(limits ...int64) []*kapi.LimitRange {
	list := make([]*kapi.LimitRange, len(limits))
	for _, limit := range limits {
		quantity := resource.NewQuantity(limit, resource.BinarySI)
		list = append(list, &kapi.LimitRange{
			Spec: kapi.LimitRangeSpec{
				Limits: []kapi.LimitRangeItem{
					{
						Type: imageapi.LimitTypeImage,
						Max: kapi.ResourceList{
							kapi.ResourceStorage: *quantity,
						},
					},
				},
			},
		})
	}
	return list
}

func commonSpec(strategyType, fromKind, fromNamespace, fromName string) buildapi.CommonSpec {
	spec := buildapi.CommonSpec{
		Strategy: buildapi.BuildStrategy{},
	}
	switch strategyType {
	case "source":
		spec.Strategy.SourceStrategy = &buildapi.SourceBuildStrategy{
			From: kapi.ObjectReference{
				Kind:      fromKind,
				Namespace: fromNamespace,
				Name:      fromName,
			},
		}
	case "docker":
		spec.Strategy.DockerStrategy = &buildapi.DockerBuildStrategy{
			From: &kapi.ObjectReference{
				Kind:      fromKind,
				Namespace: fromNamespace,
				Name:      fromName,
			},
		}
	case "custom":
		spec.Strategy.CustomStrategy = &buildapi.CustomBuildStrategy{
			From: kapi.ObjectReference{
				Kind:      fromKind,
				Namespace: fromNamespace,
				Name:      fromName,
			},
		}
	}

	return spec
}

type fakeImageDeleter struct {
	invocations sets.String
	err         error
}

var _ ImageDeleter = &fakeImageDeleter{}

func (p *fakeImageDeleter) DeleteImage(image *imageapi.Image) error {
	p.invocations.Insert(image.Name)
	return p.err
}

type fakeImageStreamDeleter struct {
	invocations sets.String
	err         error
}

var _ ImageStreamDeleter = &fakeImageStreamDeleter{}

func (p *fakeImageStreamDeleter) DeleteImageStream(stream *imageapi.ImageStream, image *imageapi.Image, updatedTags []string) (*imageapi.ImageStream, error) {
	p.invocations.Insert(fmt.Sprintf("%s/%s|%s", stream.Namespace, stream.Name, image.Name))
	return stream, p.err
}

type fakeBlobDeleter struct {
	invocations sets.String
	err         error
}

var _ BlobDeleter = &fakeBlobDeleter{}

func (p *fakeBlobDeleter) DeleteBlob(registryClient *http.Client, registryURL, blob string) error {
	p.invocations.Insert(fmt.Sprintf("%s|%s", registryURL, blob))
	return p.err
}

type fakeLayerDeleter struct {
	invocations sets.String
	err         error
}

var _ LayerDeleter = &fakeLayerDeleter{}

func (p *fakeLayerDeleter) DeleteLayer(registryClient *http.Client, registryURL, repo, layer string) error {
	p.invocations.Insert(fmt.Sprintf("%s|%s|%s", registryURL, repo, layer))
	return p.err
}

type fakeManifestDeleter struct {
	invocations sets.String
	err         error
}

var _ ManifestDeleter = &fakeManifestDeleter{}

func (p *fakeManifestDeleter) DeleteManifest(registryClient *http.Client, registryURL, repo, manifest string) error {
	p.invocations.Insert(fmt.Sprintf("%s|%s|%s", registryURL, repo, manifest))
	return p.err
}

var logLevel = flag.Int("loglevel", 0, "")
var testCase = flag.String("testcase", "", "")

func TestImagePruning(t *testing.T) {
	flag.Lookup("v").Value.Set(fmt.Sprint(*logLevel))
	registryURL := "registry"

	tests := map[string]struct {
		pruneOverSizeLimit     *bool
		registryURLs           []string
		images                 imageapi.ImageList
		pods                   kapi.PodList
		streams                imageapi.ImageStreamList
		rcs                    kapi.ReplicationControllerList
		bcs                    buildapi.BuildConfigList
		builds                 buildapi.BuildList
		dcs                    deployapi.DeploymentConfigList
		limits                 map[string][]*kapi.LimitRange
		expectedImageDeletions []string
		expectedStreamUpdates  []string
		expectedLayerDeletions []string
		expectedBlobDeletions  []string
	}{
		"1 pod - phase pending - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods:   podList(pod("foo", "pod1", kapi.PodPending, registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"3 pods - last phase pending - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodSucceeded, registryURL+"/foo/bar@id"),
				pod("foo", "pod2", kapi.PodSucceeded, registryURL+"/foo/bar@id"),
				pod("foo", "pod3", kapi.PodPending, registryURL+"/foo/bar@id"),
			),
			expectedImageDeletions: []string{},
		},
		"1 pod - phase running - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods:   podList(pod("foo", "pod1", kapi.PodRunning, registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"3 pods - last phase running - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodSucceeded, registryURL+"/foo/bar@id"),
				pod("foo", "pod2", kapi.PodSucceeded, registryURL+"/foo/bar@id"),
				pod("foo", "pod3", kapi.PodRunning, registryURL+"/foo/bar@id"),
			),
			expectedImageDeletions: []string{},
		},
		"pod phase succeeded - prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods:   podList(pod("foo", "pod1", kapi.PodSucceeded, registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{"id"},
			expectedBlobDeletions: []string{
				registryURL + "|" + layer1,
				registryURL + "|" + layer2,
				registryURL + "|" + layer3,
				registryURL + "|" + layer4,
				registryURL + "|" + layer5,
			},
		},
		"pod phase succeeded, pod less than min pruning age - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods:   podList(agedPod("foo", "pod1", kapi.PodSucceeded, 5, registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"pod phase succeeded, image less than min pruning age - don't prune": {
			images: imageList(agedImage("id", registryURL+"/foo/bar@id", 5)),
			pods:   podList(pod("foo", "pod1", kapi.PodSucceeded, registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"pod phase failed - prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodFailed, registryURL+"/foo/bar@id"),
				pod("foo", "pod2", kapi.PodFailed, registryURL+"/foo/bar@id"),
				pod("foo", "pod3", kapi.PodFailed, registryURL+"/foo/bar@id"),
			),
			expectedImageDeletions: []string{"id"},
			expectedBlobDeletions: []string{
				registryURL + "|" + layer1,
				registryURL + "|" + layer2,
				registryURL + "|" + layer3,
				registryURL + "|" + layer4,
				registryURL + "|" + layer5,
			},
		},
		"pod phase unknown - prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodUnknown, registryURL+"/foo/bar@id"),
				pod("foo", "pod2", kapi.PodUnknown, registryURL+"/foo/bar@id"),
				pod("foo", "pod3", kapi.PodUnknown, registryURL+"/foo/bar@id"),
			),
			expectedImageDeletions: []string{"id"},
			expectedBlobDeletions: []string{
				registryURL + "|" + layer1,
				registryURL + "|" + layer2,
				registryURL + "|" + layer3,
				registryURL + "|" + layer4,
				registryURL + "|" + layer5,
			},
		},
		"pod container image not parsable": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodRunning, "a/b/c/d/e"),
			),
			expectedImageDeletions: []string{"id"},
			expectedBlobDeletions: []string{
				registryURL + "|" + layer1,
				registryURL + "|" + layer2,
				registryURL + "|" + layer3,
				registryURL + "|" + layer4,
				registryURL + "|" + layer5,
			},
		},
		"pod container image doesn't have an id": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodRunning, "foo/bar:latest"),
			),
			expectedImageDeletions: []string{"id"},
			expectedBlobDeletions: []string{
				registryURL + "|" + layer1,
				registryURL + "|" + layer2,
				registryURL + "|" + layer3,
				registryURL + "|" + layer4,
				registryURL + "|" + layer5,
			},
		},
		"pod refers to image not in graph": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			pods: podList(
				pod("foo", "pod1", kapi.PodRunning, registryURL+"/foo/bar@otherid"),
			),
			expectedImageDeletions: []string{"id"},
			expectedBlobDeletions: []string{
				registryURL + "|" + layer1,
				registryURL + "|" + layer2,
				registryURL + "|" + layer3,
				registryURL + "|" + layer4,
				registryURL + "|" + layer5,
			},
		},
		"referenced by rc - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			rcs:    rcList(rc("foo", "rc1", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by dc - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			dcs:    dcList(dc("foo", "rc1", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by bc - sti - ImageStreamImage - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			bcs:    bcList(bc("foo", "bc1", "source", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by bc - docker - ImageStreamImage - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			bcs:    bcList(bc("foo", "bc1", "docker", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by bc - custom - ImageStreamImage - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			bcs:    bcList(bc("foo", "bc1", "custom", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by bc - sti - DockerImage - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			bcs:    bcList(bc("foo", "bc1", "source", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by bc - docker - DockerImage - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			bcs:    bcList(bc("foo", "bc1", "docker", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by bc - custom - DockerImage - don't prune": {
			images: imageList(image("id", registryURL+"/foo/bar@id")),
			bcs:    bcList(bc("foo", "bc1", "custom", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by build - sti - ImageStreamImage - don't prune": {
			images:                 imageList(image("id", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "source", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by build - docker - ImageStreamImage - don't prune": {
			images:                 imageList(image("id", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "docker", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by build - custom - ImageStreamImage - don't prune": {
			images:                 imageList(image("id", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "custom", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by build - sti - DockerImage - don't prune": {
			images:                 imageList(image("id", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "source", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by build - docker - DockerImage - don't prune": {
			images:                 imageList(image("id", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "docker", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"referenced by build - custom - DockerImage - don't prune": {
			images:                 imageList(image("id", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "custom", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			expectedImageDeletions: []string{},
		},
		"image stream - keep most recent n images": {
			images: imageList(
				unmanagedImage("id", "otherregistry/foo/bar@id", false, "", ""),
				image("id2", registryURL+"/foo/bar@id2"),
				image("id3", registryURL+"/foo/bar@id3"),
				image("id4", registryURL+"/foo/bar@id4"),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id", "otherregistry/foo/bar@id"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
						tagEvent("id3", registryURL+"/foo/bar@id3"),
						tagEvent("id4", registryURL+"/foo/bar@id4"),
					),
				)),
			),
			expectedImageDeletions: []string{"id4"},
			expectedStreamUpdates:  []string{"foo/bar|id4"},
		},
		"image stream - same manifest listed multiple times in tag history": {
			images: imageList(
				image("id1", registryURL+"/foo/bar@id1"),
				image("id2", registryURL+"/foo/bar@id2"),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id1", registryURL+"/foo/bar@id1"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
						tagEvent("id1", registryURL+"/foo/bar@id1"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
					),
				)),
			),
		},
		"image stream age less than min pruning age - don't prune": {
			images: imageList(
				image("id", registryURL+"/foo/bar@id"),
				image("id2", registryURL+"/foo/bar@id2"),
				image("id3", registryURL+"/foo/bar@id3"),
				image("id4", registryURL+"/foo/bar@id4"),
			),
			streams: streamList(
				agedStream(registryURL, "foo", "bar", 5, tags(
					tag("latest",
						tagEvent("id", registryURL+"/foo/bar@id"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
						tagEvent("id3", registryURL+"/foo/bar@id3"),
						tagEvent("id4", registryURL+"/foo/bar@id4"),
					),
				)),
			),
			expectedImageDeletions: []string{},
			expectedStreamUpdates:  []string{},
		},
		"multiple resources pointing to image - don't prune": {
			images: imageList(
				image("id", registryURL+"/foo/bar@id"),
				image("id2", registryURL+"/foo/bar@id2"),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id", registryURL+"/foo/bar@id"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
					),
				)),
			),
			rcs:                    rcList(rc("foo", "rc1", registryURL+"/foo/bar@id2")),
			pods:                   podList(pod("foo", "pod1", kapi.PodRunning, registryURL+"/foo/bar@id2")),
			dcs:                    dcList(dc("foo", "rc1", registryURL+"/foo/bar@id")),
			bcs:                    bcList(bc("foo", "bc1", "source", "DockerImage", "foo", registryURL+"/foo/bar@id")),
			builds:                 buildList(build("foo", "build1", "custom", "ImageStreamImage", "foo", "bar@id")),
			expectedImageDeletions: []string{},
			expectedStreamUpdates:  []string{},
		},
		"image with nil annotations": {
			images: imageList(
				unmanagedImage("id", "someregistry/foo/bar@id", false, "", ""),
			),
			expectedImageDeletions: []string{},
			expectedStreamUpdates:  []string{},
		},
		"image missing managed annotation": {
			images: imageList(
				unmanagedImage("id", "someregistry/foo/bar@id", true, "foo", "bar"),
			),
			expectedImageDeletions: []string{},
			expectedStreamUpdates:  []string{},
		},
		"image with managed annotation != true": {
			images: imageList(
				unmanagedImage("id", "someregistry/foo/bar@id", true, imageapi.ManagedByOpenShiftAnnotation, "false"),
				unmanagedImage("id", "someregistry/foo/bar@id", true, imageapi.ManagedByOpenShiftAnnotation, "0"),
				unmanagedImage("id", "someregistry/foo/bar@id", true, imageapi.ManagedByOpenShiftAnnotation, "1"),
				unmanagedImage("id", "someregistry/foo/bar@id", true, imageapi.ManagedByOpenShiftAnnotation, "True"),
				unmanagedImage("id", "someregistry/foo/bar@id", true, imageapi.ManagedByOpenShiftAnnotation, "yes"),
				unmanagedImage("id", "someregistry/foo/bar@id", true, imageapi.ManagedByOpenShiftAnnotation, "Yes"),
			),
			expectedImageDeletions: []string{},
			expectedStreamUpdates:  []string{},
		},
		"image with layers": {
			images: imageList(
				imageWithLayers("id1", registryURL+"/foo/bar@id1", "layer1", "layer2", "layer3", "layer4"),
				imageWithLayers("id2", registryURL+"/foo/bar@id2", "layer1", "layer2", "layer3", "layer4"),
				imageWithLayers("id3", registryURL+"/foo/bar@id3", "layer1", "layer2", "layer3", "layer4"),
				imageWithLayers("id4", registryURL+"/foo/bar@id4", "layer5", "layer6", "layer7", "layer8"),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id1", registryURL+"/foo/bar@id1"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
						tagEvent("id3", registryURL+"/foo/bar@id3"),
						tagEvent("id4", registryURL+"/foo/bar@id4"),
					),
				)),
			),
			expectedImageDeletions: []string{"id4"},
			expectedStreamUpdates:  []string{"foo/bar|id4"},
			expectedLayerDeletions: []string{
				registryURL + "|foo/bar|layer5",
				registryURL + "|foo/bar|layer6",
				registryURL + "|foo/bar|layer7",
				registryURL + "|foo/bar|layer8",
			},
			expectedBlobDeletions: []string{
				registryURL + "|layer5",
				registryURL + "|layer6",
				registryURL + "|layer7",
				registryURL + "|layer8",
			},
		},
		"image exceeding limits": {
			pruneOverSizeLimit: newBool(true),
			images: imageList(
				unmanagedImage("id", "otherregistry/foo/bar@id", false, "", ""),
				sizedImage("id2", registryURL+"/foo/bar@id2", 100),
				sizedImage("id3", registryURL+"/foo/bar@id3", 200),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id", "otherregistry/foo/bar@id"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
						tagEvent("id3", registryURL+"/foo/bar@id3"),
					),
				)),
			),
			limits: map[string][]*kapi.LimitRange{
				"foo": limitList(100, 200),
			},
			expectedImageDeletions: []string{"id3"},
			expectedStreamUpdates:  []string{"foo/bar|id3"},
		},
		"multiple images in different namespaces exceeding different limits": {
			pruneOverSizeLimit: newBool(true),
			images: imageList(
				sizedImage("id1", registryURL+"/foo/bar@id1", 100),
				sizedImage("id2", registryURL+"/foo/bar@id2", 200),
				sizedImage("id3", registryURL+"/bar/foo@id3", 500),
				sizedImage("id4", registryURL+"/bar/foo@id4", 600),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id1", registryURL+"/foo/bar@id1"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
					),
				)),
				stream(registryURL, "bar", "foo", tags(
					tag("latest",
						tagEvent("id3", registryURL+"/bar/foo@id3"),
						tagEvent("id4", registryURL+"/bar/foo@id4"),
					),
				)),
			),
			limits: map[string][]*kapi.LimitRange{
				"foo": limitList(150),
				"bar": limitList(550),
			},
			expectedImageDeletions: []string{"id2", "id4"},
			expectedStreamUpdates:  []string{"foo/bar|id2", "bar/foo|id4"},
		},
		"image within allowed limits": {
			pruneOverSizeLimit: newBool(true),
			images: imageList(
				unmanagedImage("id", "otherregistry/foo/bar@id", false, "", ""),
				sizedImage("id2", registryURL+"/foo/bar@id2", 100),
				sizedImage("id3", registryURL+"/foo/bar@id3", 200),
			),
			streams: streamList(
				stream(registryURL, "foo", "bar", tags(
					tag("latest",
						tagEvent("id", "otherregistry/foo/bar@id"),
						tagEvent("id2", registryURL+"/foo/bar@id2"),
						tagEvent("id3", registryURL+"/foo/bar@id3"),
					),
				)),
			),
			limits: map[string][]*kapi.LimitRange{
				"foo": limitList(300),
			},
			expectedImageDeletions: []string{},
			expectedStreamUpdates:  []string{},
		},
	}

	for name, test := range tests {
		tcFilter := flag.Lookup("testcase").Value.String()
		if len(tcFilter) > 0 && name != tcFilter {
			continue
		}

		options := PrunerOptions{
			Images:      &test.images,
			Streams:     &test.streams,
			Pods:        &test.pods,
			RCs:         &test.rcs,
			BCs:         &test.bcs,
			Builds:      &test.builds,
			DCs:         &test.dcs,
			LimitRanges: test.limits,
		}
		if test.pruneOverSizeLimit != nil {
			options.PruneOverSizeLimit = test.pruneOverSizeLimit
		} else {
			keepYoungerThan := 60 * time.Minute
			keepTagRevisions := 3
			options.KeepYoungerThan = &keepYoungerThan
			options.KeepTagRevisions = &keepTagRevisions
		}
		p := NewPruner(options)
		p.(*pruner).registryPinger = &fakeRegistryPinger{}

		imageDeleter := &fakeImageDeleter{invocations: sets.NewString()}
		streamDeleter := &fakeImageStreamDeleter{invocations: sets.NewString()}
		layerDeleter := &fakeLayerDeleter{invocations: sets.NewString()}
		blobDeleter := &fakeBlobDeleter{invocations: sets.NewString()}
		manifestDeleter := &fakeManifestDeleter{invocations: sets.NewString()}

		p.Prune(imageDeleter, streamDeleter, layerDeleter, blobDeleter, manifestDeleter)

		expectedImageDeletions := sets.NewString(test.expectedImageDeletions...)
		if !reflect.DeepEqual(expectedImageDeletions, imageDeleter.invocations) {
			t.Errorf("%s: expected image deletions %q, got %q", name, expectedImageDeletions.List(), imageDeleter.invocations.List())
		}

		expectedStreamUpdates := sets.NewString(test.expectedStreamUpdates...)
		if !reflect.DeepEqual(expectedStreamUpdates, streamDeleter.invocations) {
			t.Errorf("%s: expected stream updates %q, got %q", name, expectedStreamUpdates.List(), streamDeleter.invocations.List())
		}

		expectedLayerDeletions := sets.NewString(test.expectedLayerDeletions...)
		if !reflect.DeepEqual(expectedLayerDeletions, layerDeleter.invocations) {
			t.Errorf("%s: expected layer deletions %q, got %q", name, expectedLayerDeletions.List(), layerDeleter.invocations.List())
		}

		expectedBlobDeletions := sets.NewString(test.expectedBlobDeletions...)
		if !reflect.DeepEqual(expectedBlobDeletions, blobDeleter.invocations) {
			t.Errorf("%s: expected blob deletions %q, got %q", name, expectedBlobDeletions.List(), blobDeleter.invocations.List())
		}
	}
}

func TestImageDeleter(t *testing.T) {
	flag.Lookup("v").Value.Set(fmt.Sprint(*logLevel))

	tests := map[string]struct {
		imageDeletionError error
	}{
		"no error": {},
		"delete error": {
			imageDeletionError: fmt.Errorf("foo"),
		},
	}

	for name, test := range tests {
		imageClient := testclient.Fake{}
		imageClient.AddReactor("delete", "images", func(action ktc.Action) (handled bool, ret runtime.Object, err error) {
			return true, nil, test.imageDeletionError
		})
		imageDeleter := NewImageDeleter(imageClient.Images())
		err := imageDeleter.DeleteImage(&imageapi.Image{ObjectMeta: kapi.ObjectMeta{Name: "id2"}})
		if test.imageDeletionError != nil {
			if e, a := test.imageDeletionError, err; e != a {
				t.Errorf("%s: err: expected %v, got %v", name, e, a)
			}
			continue
		}

		if e, a := 1, len(imageClient.Actions()); e != a {
			t.Errorf("%s: expected %d actions, got %d: %#v", name, e, a, imageClient.Actions())
			continue
		}

		if !imageClient.Actions()[0].Matches("delete", "images") {
			t.Errorf("%s: expected action %s, got %v", name, "delete-images", imageClient.Actions()[0])
		}
	}
}

func TestLayerDeleter(t *testing.T) {
	flag.Lookup("v").Value.Set(fmt.Sprint(*logLevel))

	var actions []string
	client := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		actions = append(actions, req.Method+":"+req.URL.String())
		return &http.Response{StatusCode: http.StatusServiceUnavailable, Body: ioutil.NopCloser(bytes.NewReader([]byte{}))}, nil
	})
	layerDeleter := NewLayerDeleter()
	layerDeleter.DeleteLayer(client, "registry1", "repo", "layer1")

	if !reflect.DeepEqual(actions, []string{"DELETE:https://registry1/v2/repo/blobs/layer1",
		"DELETE:http://registry1/v2/repo/blobs/layer1"}) {
		t.Errorf("Unexpected actions %v", actions)
	}
}

func TestNotFoundLayerDeleter(t *testing.T) {
	flag.Lookup("v").Value.Set(fmt.Sprint(*logLevel))

	var actions []string
	client := fake.CreateHTTPClient(func(req *http.Request) (*http.Response, error) {
		actions = append(actions, req.Method+":"+req.URL.String())
		return &http.Response{StatusCode: http.StatusNotFound, Body: ioutil.NopCloser(bytes.NewReader([]byte{}))}, nil
	})
	layerDeleter := NewLayerDeleter()
	layerDeleter.DeleteLayer(client, "registry1", "repo", "layer1")

	if !reflect.DeepEqual(actions, []string{"DELETE:https://registry1/v2/repo/blobs/layer1"}) {
		t.Errorf("Unexpected actions %v", actions)
	}
}

func TestRegistryPruning(t *testing.T) {
	flag.Lookup("v").Value.Set(fmt.Sprint(*logLevel))

	tests := map[string]struct {
		images                    imageapi.ImageList
		streams                   imageapi.ImageStreamList
		expectedLayerDeletions    sets.String
		expectedBlobDeletions     sets.String
		expectedManifestDeletions sets.String
		pingErr                   error
	}{
		"layers unique to id1 pruned": {
			images: imageList(
				imageWithLayers("id1", "registry1/foo/bar@id1", "layer1", "layer2", "layer3", "layer4"),
				imageWithLayers("id2", "registry1/foo/bar@id2", "layer3", "layer4", "layer5", "layer6"),
			),
			streams: streamList(
				stream("registry1", "foo", "bar", tags(
					tag("latest",
						tagEvent("id2", "registry1/foo/bar@id2"),
						tagEvent("id1", "registry1/foo/bar@id1"),
					),
				)),
				stream("registry1", "foo", "other", tags(
					tag("latest",
						tagEvent("id2", "registry1/foo/other@id2"),
					),
				)),
			),
			expectedLayerDeletions: sets.NewString(
				"registry1|foo/bar|layer1",
				"registry1|foo/bar|layer2",
			),
			expectedBlobDeletions: sets.NewString(
				"registry1|layer1",
				"registry1|layer2",
			),
			expectedManifestDeletions: sets.NewString(
				"registry1|foo/bar|id1",
			),
		},
		"no pruning when no images are pruned": {
			images: imageList(
				imageWithLayers("id1", "registry1/foo/bar@id1", "layer1", "layer2", "layer3", "layer4"),
			),
			streams: streamList(
				stream("registry1", "foo", "bar", tags(
					tag("latest",
						tagEvent("id1", "registry1/foo/bar@id1"),
					),
				)),
			),
			expectedLayerDeletions:    sets.NewString(),
			expectedBlobDeletions:     sets.NewString(),
			expectedManifestDeletions: sets.NewString(),
		},
		"blobs pruned when streams have already been deleted": {
			images: imageList(
				imageWithLayers("id1", "registry1/foo/bar@id1", "layer1", "layer2", "layer3", "layer4"),
				imageWithLayers("id2", "registry1/foo/bar@id2", "layer3", "layer4", "layer5", "layer6"),
			),
			expectedLayerDeletions: sets.NewString(),
			expectedBlobDeletions: sets.NewString(
				"registry1|layer1",
				"registry1|layer2",
				"registry1|layer3",
				"registry1|layer4",
				"registry1|layer5",
				"registry1|layer6",
			),
			expectedManifestDeletions: sets.NewString(),
		},
		"ping error": {
			images: imageList(
				imageWithLayers("id1", "registry1/foo/bar@id1", "layer1", "layer2", "layer3", "layer4"),
				imageWithLayers("id2", "registry1/foo/bar@id2", "layer3", "layer4", "layer5", "layer6"),
			),
			streams: streamList(
				stream("registry1", "foo", "bar", tags(
					tag("latest",
						tagEvent("id2", "registry1/foo/bar@id2"),
						tagEvent("id1", "registry1/foo/bar@id1"),
					),
				)),
				stream("registry1", "foo", "other", tags(
					tag("latest",
						tagEvent("id2", "registry1/foo/other@id2"),
					),
				)),
			),
			expectedLayerDeletions:    sets.NewString(),
			expectedBlobDeletions:     sets.NewString(),
			expectedManifestDeletions: sets.NewString(),
			pingErr:                   errors.New("foo"),
		},
	}

	for name, test := range tests {
		tcFilter := flag.Lookup("testcase").Value.String()
		if len(tcFilter) > 0 && name != tcFilter {
			continue
		}

		t.Logf("Running test case %s", name)

		keepYoungerThan := 60 * time.Minute
		keepTagRevisions := 1
		options := PrunerOptions{
			KeepYoungerThan:  &keepYoungerThan,
			KeepTagRevisions: &keepTagRevisions,
			Images:           &test.images,
			Streams:          &test.streams,
			Pods:             &kapi.PodList{},
			RCs:              &kapi.ReplicationControllerList{},
			BCs:              &buildapi.BuildConfigList{},
			Builds:           &buildapi.BuildList{},
			DCs:              &deployapi.DeploymentConfigList{},
		}
		p := NewPruner(options)
		p.(*pruner).registryPinger = &fakeRegistryPinger{err: test.pingErr}

		imageDeleter := &fakeImageDeleter{invocations: sets.NewString()}
		streamDeleter := &fakeImageStreamDeleter{invocations: sets.NewString()}
		layerDeleter := &fakeLayerDeleter{invocations: sets.NewString()}
		blobDeleter := &fakeBlobDeleter{invocations: sets.NewString()}
		manifestDeleter := &fakeManifestDeleter{invocations: sets.NewString()}

		p.Prune(imageDeleter, streamDeleter, layerDeleter, blobDeleter, manifestDeleter)

		if !reflect.DeepEqual(test.expectedLayerDeletions, layerDeleter.invocations) {
			t.Errorf("%s: expected layer deletions %#v, got %#v", name, test.expectedLayerDeletions, layerDeleter.invocations)
		}
		if !reflect.DeepEqual(test.expectedBlobDeletions, blobDeleter.invocations) {
			t.Errorf("%s: expected blob deletions %#v, got %#v", name, test.expectedBlobDeletions, blobDeleter.invocations)
		}
		if !reflect.DeepEqual(test.expectedManifestDeletions, manifestDeleter.invocations) {
			t.Errorf("%s: expected manifest deletions %#v, got %#v", name, test.expectedManifestDeletions, manifestDeleter.invocations)
		}
	}
}

func newBool(a bool) *bool {
	r := new(bool)
	*r = a
	return r
}
