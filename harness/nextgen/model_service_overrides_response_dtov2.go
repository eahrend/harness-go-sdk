/*
  - Harness NextGen Software Delivery Platform API Reference
    *
  - This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of

the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject:
<security-definitions> -->

	*
	* API version: 3.0
	* Contact: contact@harness.io
	* Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
*/
package nextgen

type ServiceOverridesResponseDtov2 struct {
	Identifier        string `json:"identifier,omitempty"`
	AccountId         string `json:"accountId,omitempty"`
	OrgIdentifier     string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	EnvironmentRef    string `json:"environmentRef,omitempty"`
	ServiceRef        string `json:"serviceRef,omitempty"`
	InfraIdentifier   string `json:"infraIdentifier,omitempty"`
	ClusterIdentifier string `json:"clusterIdentifier,omitempty"`
	Type_             string `json:"type,omitempty"`
	Spec              string `json:"spec,omitempty"`
	NewlyCreated      bool   `json:"newlyCreated,omitempty"`
}
