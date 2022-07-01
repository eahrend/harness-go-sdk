/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This is the view of the SecretManagerMetadataRequest entity defined in Harness
type Invite struct {
	// This specifies the type of encryption used by the Secret Manager to encrypt Secrets.
	EncryptionType string `json:"encryptionType"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Identifier of the SecretManager metadata.
	Identifier string `json:"identifier"`
	Spec *SecretManagerMetadataRequestSpecDto `json:"spec"`
}
