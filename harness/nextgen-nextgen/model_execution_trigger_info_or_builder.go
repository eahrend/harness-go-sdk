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

type ExecutionTriggerInfoOrBuilder struct {
	TriggerTypeValue int32 `json:"triggerTypeValue,omitempty"`
	TriggerType string `json:"triggerType,omitempty"`
	TriggeredBy *TriggeredBy `json:"triggeredBy,omitempty"`
	TriggeredByOrBuilder *TriggeredByOrBuilder `json:"triggeredByOrBuilder,omitempty"`
	IsRerun bool `json:"isRerun,omitempty"`
	RerunInfo *RerunInfo `json:"rerunInfo,omitempty"`
	RerunInfoOrBuilder *RerunInfoOrBuilder `json:"rerunInfoOrBuilder,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	DefaultInstanceForType *Message `json:"defaultInstanceForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
}
