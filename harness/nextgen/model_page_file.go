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

import (
	"os"
)

type PageFile struct {
	TotalElements    int64      `json:"totalElements,omitempty"`
	TotalPages       int32      `json:"totalPages,omitempty"`
	Sort             *Sort      `json:"sort,omitempty"`
	First            bool       `json:"first,omitempty"`
	NumberOfElements int32      `json:"numberOfElements,omitempty"`
	Last             bool       `json:"last,omitempty"`
	Pageable         *Pageable  `json:"pageable,omitempty"`
	Number           int32      `json:"number,omitempty"`
	Size             int32      `json:"size,omitempty"`
	Content          []*os.File `json:"content,omitempty"`
	Empty            bool       `json:"empty,omitempty"`
}