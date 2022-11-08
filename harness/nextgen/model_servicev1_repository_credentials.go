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

type Servicev1RepositoryCredentials struct {
	// Account Identifier for the Entity.
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Agent identifier for entity.
	AgentIdentifier string `json:"agentIdentifier,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	RepoCreds *HrepocredsRepoCreds `json:"repoCreds,omitempty"`
	CreatedAt *V1Time `json:"createdAt,omitempty"`
	LastModifiedAt *V1Time `json:"lastModifiedAt,omitempty"`
	Stale bool `json:"stale,omitempty"`
}
