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

type VariableDto struct {
	// Identifier of the Variable.
	Identifier string `json:"identifier"`
	// Name of the Variable.
	Name string `json:"name"`
	// Description of the entity
	Description string `json:"description,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Type of the Variable.
	Type_ string             `json:"type"`
	Spec  *VariableConfigDto `json:"spec"`
}