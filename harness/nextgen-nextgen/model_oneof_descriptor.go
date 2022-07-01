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

type OneofDescriptor struct {
	Index int32 `json:"index,omitempty"`
	Proto *OneofDescriptorProto `json:"proto,omitempty"`
	FullName string `json:"fullName,omitempty"`
	File *FileDescriptor `json:"file,omitempty"`
	ContainingType *Descriptor `json:"containingType,omitempty"`
	FieldCount int32 `json:"fieldCount,omitempty"`
	Fields []FieldDescriptor `json:"fields,omitempty"`
	Options *OneofOptions `json:"options,omitempty"`
	Name string `json:"name,omitempty"`
}
