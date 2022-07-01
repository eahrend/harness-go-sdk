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

// This is the NTLM configuration details defined in Harness.
type NtlmConfig struct {
	Type_ string `json:"type"`
	// This is the NTLM domain name.
	Domain string `json:"domain"`
	// This is the NTLM user name.
	Username string `json:"username"`
	// This is the NTLM either to use SSL/https .
	UseSSL bool `json:"useSSL,omitempty"`
	// This is the NTLM either to skip certificate checks .
	SkipCertChecks bool `json:"skipCertChecks,omitempty"`
	// This is the NTLM powershell runs without loading profile .
	UseNoProfile bool `json:"useNoProfile,omitempty"`
	Password string `json:"password"`
}
