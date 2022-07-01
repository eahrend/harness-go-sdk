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

type AuthenticationsettingsSamlmetadatauploadBody1 struct {
	File *interface{} `json:"file,omitempty"`
	FileMetadata *FormDataContentDisposition `json:"fileMetadata,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	GroupMembershipAttr string `json:"groupMembershipAttr,omitempty"`
	AuthorizationEnabled bool `json:"authorizationEnabled,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
	EntityIdentifier string `json:"entityIdentifier,omitempty"`
	SamlProviderType string `json:"samlProviderType,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
}
