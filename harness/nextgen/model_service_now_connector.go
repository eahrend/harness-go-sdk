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

// ServiceNow Connector details.
type ServiceNowConnector struct {
	ConnectorType     string                    `json:"connectorType"`
	ServiceNowUrl     string                    `json:"serviceNowUrl"`
	Username          string                    `json:"username,omitempty"`
	UsernameRef       string                    `json:"usernameRef,omitempty"`
	PasswordRef       string                    `json:"passwordRef,omitempty"`
	DelegateSelectors []string                  `json:"delegateSelectors,omitempty"`
	Auth              *ServiceNowAuthentication `json:"auth"`
}
