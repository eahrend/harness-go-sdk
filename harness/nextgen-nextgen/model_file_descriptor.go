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

type FileDescriptor struct {
	Proto *FileDescriptorProto `json:"proto,omitempty"`
	MessageTypes []Descriptor `json:"messageTypes,omitempty"`
	EnumTypes []EnumDescriptor `json:"enumTypes,omitempty"`
	Services []ServiceDescriptor `json:"services,omitempty"`
	Extensions []FieldDescriptor `json:"extensions,omitempty"`
	Dependencies []FileDescriptor `json:"dependencies,omitempty"`
	PublicDependencies []FileDescriptor `json:"publicDependencies,omitempty"`
	Options *FileOptions `json:"options,omitempty"`
	FullName string `json:"fullName,omitempty"`
	Syntax string `json:"syntax,omitempty"`
	Name string `json:"name,omitempty"`
	Package_ string `json:"package,omitempty"`
	File *FileDescriptor `json:"file,omitempty"`
}
