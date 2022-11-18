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

type ApplicationsApplicationSummary struct {
	// ExternalURLs holds all external URLs of application child resources.
	ExternalURLs []string `json:"externalURLs,omitempty"`
	// Images holds all images of application child resources.
	Images []string `json:"images,omitempty"`
}