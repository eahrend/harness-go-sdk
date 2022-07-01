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

// This contains details of the Bitbucket credentials used via HTTP connections
type BitbucketHttpCredentials struct {
	Type_ string `json:"type"`
	Spec *BitbucketHttpCredentialsSpec `json:"spec"`
}
