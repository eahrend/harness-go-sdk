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

// This contains the Audit Event filter information. This is used to filter Audit Events depending on the information provided.
type AuditFilterProperties struct {
	// List of Resource Scopes
	Scopes []AuditResourceScope `json:"scopes,omitempty"`
	// List of Resources
	Resources []AuditResource `json:"resources,omitempty"`
	// List of Module Types
	Modules []string `json:"modules,omitempty"`
	// List of Actions
	Actions []string `json:"actions,omitempty"`
	// List of Environments
	Environments []Environment `json:"environments,omitempty"`
	// List of Principals
	Principals []AuditPrincipal `json:"principals,omitempty"`
	// Pre-defined Filter
	StaticFilter string `json:"staticFilter,omitempty"`
	// Used to specify a start time for retrieving Audit events that occurred at or after the time indicated.
	StartTime int64 `json:"startTime,omitempty"`
	// Used to specify the end time for retrieving Audit events that occurred at or before the time indicated.
	EndTime int64 `json:"endTime,omitempty"`
	// Filter tags as a key-value pair.
	Tags map[string]string `json:"tags,omitempty"`
	// This specifies the corresponding Entity of the filter.
	FilterType string `json:"filterType,omitempty"`
}
