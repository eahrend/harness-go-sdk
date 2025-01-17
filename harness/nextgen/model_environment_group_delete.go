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

// This is the view of Environment Group Delete Response defined in Harness
type EnvironmentGroupDelete struct {
	// Value true, when the Entity is deleted
	Deleted bool `json:"deleted,omitempty"`
	// Account Identifier for the Entity.
	AccountId string `json:"accountId,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Identifier for the Entity.
	Identifier string `json:"identifier,omitempty"`
}
