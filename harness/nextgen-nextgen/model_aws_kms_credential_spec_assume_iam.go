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

// Returns the Delegate Selectors used by this AWS KMS Secret Manager Connector.
type AwsKmsCredentialSpecAssumeIam struct {
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors []string `json:"delegateSelectors"`
}
