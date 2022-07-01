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

type RoutingData struct {
	Instance *InstanceBasedRoutingData `json:"instance,omitempty"`
	Ports []PortConfig `json:"ports,omitempty"`
	Lb string `json:"lb,omitempty"`
	K8s *RoutingDataK8s `json:"k8s,omitempty"`
	CustomDomainProviders *interface{} `json:"custom_domain_providers,omitempty"`
	ContainerSvc *ContainerSvc `json:"container_svc,omitempty"`
	Database *RdsDatabase `json:"database,omitempty"`
}
