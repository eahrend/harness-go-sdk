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

type OneofOptions struct {
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
	ParserForType *ParserOneofOptions `json:"parserForType,omitempty"`
	SerializedSize int32 `json:"serializedSize,omitempty"`
	DefaultInstanceForType *OneofOptions `json:"defaultInstanceForType,omitempty"`
	UninterpretedOptionList []UninterpretedOption `json:"uninterpretedOptionList,omitempty"`
	UninterpretedOptionCount int32 `json:"uninterpretedOptionCount,omitempty"`
	UninterpretedOptionOrBuilderList []UninterpretedOptionOrBuilder `json:"uninterpretedOptionOrBuilderList,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	AllFieldsRaw map[string]interface{} `json:"allFieldsRaw,omitempty"`
	MemoizedSerializedSize int32 `json:"memoizedSerializedSize,omitempty"`
}
