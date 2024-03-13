/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type TypesCommitFilesResponse struct {
	CommitId       string                `json:"commit_id,omitempty"`
	DryRunRules    bool                  `json:"dry_run_rules,omitempty"`
	RuleViolations []TypesRuleViolations `json:"rule_violations,omitempty"`
}
