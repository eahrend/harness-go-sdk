/*
  - Harness NextGen Software Delivery Platform API Reference
    *
  - This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of

the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject:
<security-definitions> -->

	*
	* API version: 3.0
	* Contact: contact@harness.io
	* Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
*/
package nextgen

type ServiceOverridesSpec struct {
	Variables           []map[string]interface{} `json:"variables,omitempty"`
	Manifests           []map[string]interface{} `json:"manifests,omitempty"`
	ConfigFiles         []map[string]interface{} `json:"config_files,omitempty"`
	ApplicationSettings []map[string]interface{} `json:"application_settings,omitempty"`
	ConnectionStrings   []map[string]interface{} `json:"connection_strings,omitempty"`
}