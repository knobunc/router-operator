package v1

import (
	v1 "github.com/openshift/api/image/v1"
	"github.com/openshift/client-go/image/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ImageV1Interface interface {
	RESTClient() rest.Interface
	ImagesGetter
	ImageSignaturesGetter
	ImageStreamsGetter
	ImageStreamImagesGetter
	ImageStreamImportsGetter
	ImageStreamMappingsGetter
	ImageStreamTagsGetter
}

// ImageV1Client is used to interact with features provided by the image.openshift.io group.
type ImageV1Client struct {
	restClient rest.Interface
}

func (c *ImageV1Client) Images() ImageInterface {
	return newImages(c)
}

func (c *ImageV1Client) ImageSignatures() ImageSignatureInterface {
	return newImageSignatures(c)
}

func (c *ImageV1Client) ImageStreams(namespace string) ImageStreamInterface {
	return newImageStreams(c, namespace)
}

func (c *ImageV1Client) ImageStreamImages(namespace string) ImageStreamImageInterface {
	return newImageStreamImages(c, namespace)
}

func (c *ImageV1Client) ImageStreamImports(namespace string) ImageStreamImportInterface {
	return newImageStreamImports(c, namespace)
}

func (c *ImageV1Client) ImageStreamMappings(namespace string) ImageStreamMappingInterface {
	return newImageStreamMappings(c, namespace)
}

func (c *ImageV1Client) ImageStreamTags(namespace string) ImageStreamTagInterface {
	return newImageStreamTags(c, namespace)
}

// NewForConfig creates a new ImageV1Client for the given config.
func NewForConfig(c *rest.Config) (*ImageV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ImageV1Client{client}, nil
}

// NewForConfigOrDie creates a new ImageV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ImageV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ImageV1Client for the given RESTClient.
func New(c rest.Interface) *ImageV1Client {
	return &ImageV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ImageV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
