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

type ApplicationsResourceStatus struct {
	Group           string                    `json:"group,omitempty"`
	Version         string                    `json:"version,omitempty"`
	Kind            string                    `json:"kind,omitempty"`
	Namespace       string                    `json:"namespace,omitempty"`
	Name            string                    `json:"name,omitempty"`
	Status          string                    `json:"status,omitempty"`
	Health          *ApplicationsHealthStatus `json:"health,omitempty"`
	Hook            bool                      `json:"hook,omitempty"`
	RequiresPruning bool                      `json:"requiresPruning,omitempty"`
}