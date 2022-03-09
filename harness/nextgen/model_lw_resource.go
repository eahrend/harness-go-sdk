/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type LwResource struct {
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Region           string   `json:"region,omitempty"`
	AvailabilityZone string   `json:"availability_zone,omitempty"`
	Status           string   `json:"status,omitempty"`
	Type_            string   `json:"type,omitempty"`
	LaunchTime       string   `json:"launch_time,omitempty"`
	Ipv4             []string `json:"ipv4,omitempty"`
	PrivateIpv4      []string `json:"private_ipv4,omitempty"`
	// tag key as attribute key and tag value as attribute value
	Tags           *interface{} `json:"tags,omitempty"`
	ResourceType   string       `json:"resource_type,omitempty"`
	ProviderName   string       `json:"provider_name,omitempty"`
	IsSpot         bool         `json:"is_spot,omitempty"`
	Platform       string       `json:"platform,omitempty"`
	CloudAccountId float64      `json:"cloud_account_id,omitempty"`
	Metadata       *interface{} `json:"metadata,omitempty"`
	ProviderType   string       `json:"provider_type,omitempty"`
}