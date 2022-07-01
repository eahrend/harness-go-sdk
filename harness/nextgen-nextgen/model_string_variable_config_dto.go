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

type StringVariableConfigDto struct {
	// Type of Value of the Variable.
	ValueType string `json:"valueType"`
	Value *interface{} `json:"value,omitempty"`
	// Fixed Value of the Variable.
	FixedValue string `json:"fixedValue,omitempty"`
	// Default Value of the Variable.
	DefaultValue string `json:"defaultValue,omitempty"`
	// Set of Values allowed for the Variable.
	AllowedValues []string `json:"allowedValues,omitempty"`
	Regex string `json:"regex,omitempty"`
}
