package v2

import (
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// Artifact describes a registry artifact.
// This structure provides `application/vnd.oci.artifact.manifest.v1+json` mediatype when marshalled to JSON.
type Artifact struct {
	specs.Versioned

	// MediaType is the media type of the object this schema refers to.
	MediaType string `json:"mediaType"`

	// ArtifactType is the artifact type of the object this schema refers to.
	ArtifactType string `json:"artifactType"`

	// Config references the index configuration.
	Config v1.Descriptor `json:"config"`

	// Blobs is a collection of blobs referenced by this manifest.
	Blobs []v1.Descriptor `json:"blobs"`

	// Manifests is a collection of manifests this artifact is linked to.
	Manifests []v1.Descriptor `json:"manifests"`

	// Annotations contains arbitrary metadata for the artifact manifest.
	Annotations map[string]string `json:"annotations,omitempty"`
}
