/*
 * Governance Policy Management API
 *
 * Read and manage OPA Governance policies, policy sets and evaluations
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package policymgmt

type LinkedPolicy2 struct {
	// Harness account ID associated with this policy
	AccountId string `json:"account_id"`
	// Time the policy was created
	Created int64 `json:"created"`
	// identifier of the policy
	Identifier string `json:"identifier"`
	// Name of the policy
	Name string `json:"name"`
	// Harness organization ID associated with this policy
	OrgId string `json:"org_id"`
	// Harness project ID associated with this policy
	ProjectId string `json:"project_id"`
	// Rego that defines the policy
	Rego string `json:"rego"`
	// The severity of this policy in this context
	Severity string `json:"severity"`
	// Time the policy was last updated
	Updated int64 `json:"updated"`
}