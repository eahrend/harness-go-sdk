/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type ProtectionDefPullReq struct {
	Approvals    *ProtectionDefApprovals    `json:"approvals,omitempty"`
	Comments     *ProtectionDefComments     `json:"comments,omitempty"`
	Merge        *ProtectionDefMerge        `json:"merge,omitempty"`
	StatusChecks *ProtectionDefStatusChecks `json:"status_checks,omitempty"`
}