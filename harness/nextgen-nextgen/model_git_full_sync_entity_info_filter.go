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

// Filter details for Git Full Sync.
type GitFullSyncEntityInfoFilter struct {
	// List of entity Types to filter on the entities.
	EntityTypes []string `json:"entityTypes,omitempty"`
	// Sync Status of the Entity that may be QUEUED, SUCCESS or FAILED.
	SyncStatus string `json:"syncStatus,omitempty"`
}
