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

type PagePipelineExecutionSummary struct {
	TotalPages int32 `json:"totalPages,omitempty"`
	TotalElements int64 `json:"totalElements,omitempty"`
	First bool `json:"first,omitempty"`
	Last bool `json:"last,omitempty"`
	Sort *Sort `json:"sort,omitempty"`
	Number int32 `json:"number,omitempty"`
	NumberOfElements int32 `json:"numberOfElements,omitempty"`
	Pageable *Pageable `json:"pageable,omitempty"`
	Size int32 `json:"size,omitempty"`
	Content []PipelineExecutionSummary `json:"content,omitempty"`
	Empty bool `json:"empty,omitempty"`
}
