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

// This contains the Service Account details and its metadata.
type ServiceAccountAggregate struct {
	ServiceAccount *ServiceAccount `json:"serviceAccount"`
	// This is the time at which Service Account was created.
	CreatedAt int64 `json:"createdAt"`
	// This is the time at which Service Account was last modified.
	LastModifiedAt int64 `json:"lastModifiedAt"`
	// This is the total number of tokens in a Service Account.
	TokensCount int32 `json:"tokensCount,omitempty"`
	// This is the list of Role Assignments for the Service Account.
	RoleAssignmentsMetadataDTO []RoleAssignmentMetadata `json:"roleAssignmentsMetadataDTO,omitempty"`
}
