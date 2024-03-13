/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This is the ClusterBatchRequest entity defined in Harness
type ClusterBatchRequest struct {
	// organization identifier of the cluster
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// project identifier of the cluster
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// environment identifier of the cluster
	EnvRef string `json:"envRef"`
	// link all clusters
	LinkAllClusters bool `json:"linkAllClusters,omitempty"`
	// unlink all clusters
	UnlinkAllClusters bool `json:"unlinkAllClusters,omitempty"`
	// search term if applicable. only valid if linking all clusters
	SearchTerm string `json:"searchTerm,omitempty"`
	// list of cluster identifiers and names
	Clusters []ClusterBasicDto `json:"clusters,omitempty"`
}
