/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type DeploymentStatsSummary struct {
	TotalCount int64 `json:"totalCount,omitempty"`
	TotalCountChangeRate float64 `json:"totalCountChangeRate,omitempty"`
	FailureRate float64 `json:"failureRate,omitempty"`
	FailureRateChangeRate float64 `json:"failureRateChangeRate,omitempty"`
	DeploymentRate float64 `json:"deploymentRate,omitempty"`
	DeploymentRateChangeRate float64 `json:"deploymentRateChangeRate,omitempty"`
	TimeBasedDeploymentInfoList []TimeBasedDeploymentInfo `json:"timeBasedDeploymentInfoList,omitempty"`
}