/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type CreateRequestBody struct {
	// The harness connector used for authenticating on the git provider
	GitConnectorRef string `json:"git_connector_ref,omitempty"`
	// The path to the file in the git repo
	GitPath string `json:"git_path,omitempty"`
	// The git repo the policy resides in
	GitRepo string `json:"git_repo,omitempty"`
	// Identifier for the policy
	Identifier string `json:"identifier"`
	// Name of the policy
	Name string `json:"name"`
	// Rego that defines the policy policy
	Rego string `json:"rego"`
}