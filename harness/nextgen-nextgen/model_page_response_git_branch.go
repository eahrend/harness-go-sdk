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

// This contains details of all the branches of given repo
type PageResponseGitBranch struct {
	TotalPages int64 `json:"totalPages,omitempty"`
	TotalItems int64 `json:"totalItems,omitempty"`
	PageItemCount int64 `json:"pageItemCount,omitempty"`
	PageSize int64 `json:"pageSize,omitempty"`
	Content []GitBranch `json:"content,omitempty"`
	PageIndex int64 `json:"pageIndex,omitempty"`
	Empty bool `json:"empty,omitempty"`
}
