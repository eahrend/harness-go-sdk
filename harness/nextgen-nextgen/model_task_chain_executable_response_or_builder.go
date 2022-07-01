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

type TaskChainExecutableResponseOrBuilder struct {
	TaskName string `json:"taskName,omitempty"`
	TaskId string `json:"taskId,omitempty"`
	ChainEnd bool `json:"chainEnd,omitempty"`
	PassThroughData *ByteString `json:"passThroughData,omitempty"`
	TaskIdBytes *ByteString `json:"taskIdBytes,omitempty"`
	LogKeysList []string `json:"logKeysList,omitempty"`
	LogKeysCount int32 `json:"logKeysCount,omitempty"`
	UnitsList []string `json:"unitsList,omitempty"`
	UnitsCount int32 `json:"unitsCount,omitempty"`
	TaskNameBytes *ByteString `json:"taskNameBytes,omitempty"`
	TaskCategory string `json:"taskCategory,omitempty"`
	TaskCategoryValue int32 `json:"taskCategoryValue,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	DefaultInstanceForType *Message `json:"defaultInstanceForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
}
