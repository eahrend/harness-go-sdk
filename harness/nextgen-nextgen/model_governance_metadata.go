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

type GovernanceMetadata struct {
	UnknownFields *UnknownFieldSet `json:"unknownFields,omitempty"`
	Action string `json:"action,omitempty"`
	Entity string `json:"entity,omitempty"`
	Created int64 `json:"created,omitempty"`
	Status string `json:"status,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
	Initialized bool `json:"initialized,omitempty"`
	ParserForType *ParserGovernanceMetadata `json:"parserForType,omitempty"`
	SerializedSize int32 `json:"serializedSize,omitempty"`
	DefaultInstanceForType *GovernanceMetadata `json:"defaultInstanceForType,omitempty"`
	TypeBytes *ByteString `json:"typeBytes,omitempty"`
	MessageBytes *ByteString `json:"messageBytes,omitempty"`
	EntityBytes *ByteString `json:"entityBytes,omitempty"`
	ActionBytes *ByteString `json:"actionBytes,omitempty"`
	StatusBytes *ByteString `json:"statusBytes,omitempty"`
	AccountId string `json:"accountId,omitempty"`
	AccountIdBytes *ByteString `json:"accountIdBytes,omitempty"`
	OrgId string `json:"orgId,omitempty"`
	OrgIdBytes *ByteString `json:"orgIdBytes,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
	ProjectIdBytes *ByteString `json:"projectIdBytes,omitempty"`
	Deny bool `json:"deny,omitempty"`
	IdBytes *ByteString `json:"idBytes,omitempty"`
	DetailsList []PolicySetMetadata `json:"detailsList,omitempty"`
	DetailsOrBuilderList []PolicySetMetadataOrBuilder `json:"detailsOrBuilderList,omitempty"`
	DetailsCount int32 `json:"detailsCount,omitempty"`
	Message string `json:"message,omitempty"`
	Id string `json:"id,omitempty"`
	Type_ string `json:"type,omitempty"`
	AllFields map[string]interface{} `json:"allFields,omitempty"`
	DescriptorForType *Descriptor `json:"descriptorForType,omitempty"`
	InitializationErrorString string `json:"initializationErrorString,omitempty"`
	MemoizedSerializedSize int32 `json:"memoizedSerializedSize,omitempty"`
}
