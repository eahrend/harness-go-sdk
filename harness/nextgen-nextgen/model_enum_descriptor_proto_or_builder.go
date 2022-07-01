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

type EnumDescriptorProtoOrBuilder struct {
	Options *EnumOptions `json:"options,omitempty"`
	ReservedRangeList []EnumReservedRange `json:"reservedRangeList,omitempty"`
	ReservedNameList []string `json:"reservedNameList,omitempty"`
	OptionsOrBuilder *EnumOptionsOrBuilder `json:"optionsOrBuilder,omitempty"`
	NameBytes *ByteString `json:"nameBytes,omitempty"`
	ReservedRangeOrBuilderList []EnumReservedRangeOrBuilder `json:"reservedRangeOrBuilderList,omitempty"`
	ReservedRangeCount int32 `json:"reservedRangeCount,omitempty"`
	ReservedNameCount int32 `json:"reservedNameCount,omitempty"`
	ValueList []EnumValueDescriptorProto `json:"valueList,omitempty"`
	ValueOrBuilderList []EnumValueDescriptorProtoOrBuilder `json:"valueOrBuilderList,omitempty"`
	ValueCount int32 `json:"valueCount,omitempty"`
	Name string `json:"name,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	DefaultInstanceForType *Message `json:"defaultInstanceForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
}
