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

// This is the view of the Pipeline Executions for a particular Date
type PipelineExecution struct {
	Date int64 `json:"date,omitempty"`
	Count *PipelineCount `json:"count,omitempty"`
}
