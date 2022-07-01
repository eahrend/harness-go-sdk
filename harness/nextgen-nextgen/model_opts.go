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

type Opts struct {
	PreservePrivateIp bool `json:"preserve_private_ip,omitempty"`
	DeleteCloudResources bool `json:"delete_cloud_resources,omitempty"`
	AlwaysUsePrivateIp bool `json:"always_use_private_ip,omitempty"`
	AccessDetails *interface{} `json:"access_details,omitempty"`
	HideProgressPage bool `json:"hide_progress_page,omitempty"`
}
