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

type FieldDescriptorProtoOrBuilder struct {
	Options *FieldOptions `json:"options,omitempty"`
	Label string `json:"label,omitempty"`
	Number int32 `json:"number,omitempty"`
	OptionsOrBuilder *FieldOptionsOrBuilder `json:"optionsOrBuilder,omitempty"`
	NameBytes *ByteString `json:"nameBytes,omitempty"`
	JsonName string `json:"jsonName,omitempty"`
	OneofIndex int32 `json:"oneofIndex,omitempty"`
	Extendee string `json:"extendee,omitempty"`
	TypeNameBytes *ByteString `json:"typeNameBytes,omitempty"`
	ExtendeeBytes *ByteString `json:"extendeeBytes,omitempty"`
	DefaultValueBytes *ByteString `json:"defaultValueBytes,omitempty"`
	JsonNameBytes *ByteString `json:"jsonNameBytes,omitempty"`
	Name string `json:"name,omitempty"`
	TypeName string `json:"typeName,omitempty"`
	Type_ string `json:"type,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	DefaultInstanceForType *Message `json:"defaultInstanceForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
}
