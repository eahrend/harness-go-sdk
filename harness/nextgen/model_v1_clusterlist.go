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

type V1Clusterlist struct {
	Content       []Servicev1Cluster `json:"content,omitempty"`
	TotalPages    int32              `json:"totalPages,omitempty"`
	TotalItems    int32              `json:"totalItems,omitempty"`
	PageItemCount int32              `json:"pageItemCount,omitempty"`
	PageSize      int32              `json:"pageSize,omitempty"`
	PageIndex     int32              `json:"pageIndex,omitempty"`
	Empty         bool               `json:"empty,omitempty"`
}
