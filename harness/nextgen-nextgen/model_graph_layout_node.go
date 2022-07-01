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

// This is the view of the Graph for execution of the Pipeline.
type GraphLayoutNode struct {
	NodeType string `json:"nodeType,omitempty"`
	NodeGroup string `json:"nodeGroup,omitempty"`
	NodeIdentifier string `json:"nodeIdentifier,omitempty"`
	Name string `json:"name,omitempty"`
	NodeUuid string `json:"nodeUuid,omitempty"`
	// This is the Execution Status of the entity
	Status string `json:"status,omitempty"`
	Module string `json:"module,omitempty"`
	ModuleInfo map[string]map[string]interface{} `json:"moduleInfo,omitempty"`
	StartTs int64 `json:"startTs,omitempty"`
	EndTs int64 `json:"endTs,omitempty"`
	EdgeLayoutList *EdgeLayoutList `json:"edgeLayoutList,omitempty"`
	SkipInfo *SkipInfo `json:"skipInfo,omitempty"`
	NodeRunInfo *NodeRunInfo `json:"nodeRunInfo,omitempty"`
	BarrierFound bool `json:"barrierFound,omitempty"`
	FailureInfo *ExecutionErrorInfo `json:"failureInfo,omitempty"`
	FailureInfoDTO *FailureInfoDto `json:"failureInfoDTO,omitempty"`
	StepDetails map[string]map[string]interface{} `json:"stepDetails,omitempty"`
	Hidden bool `json:"hidden,omitempty"`
	NodeExecutionId string `json:"nodeExecutionId,omitempty"`
	StrategyMetadata *StrategyMetadata `json:"strategyMetadata,omitempty"`
}
