/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type PrometheusMetricDefinition struct {
	Identifier string `json:"identifier"`
	MetricName string `json:"metricName"`
	RiskProfile *RiskProfile `json:"riskProfile,omitempty"`
	Analysis *AnalysisDto `json:"analysis,omitempty"`
	Sli *Slidto `json:"sli,omitempty"`
	Query string `json:"query,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	ServiceInstanceFieldName string `json:"serviceInstanceFieldName,omitempty"`
	PrometheusMetric string `json:"prometheusMetric,omitempty"`
	ServiceFilter []PrometheusFilter `json:"serviceFilter,omitempty"`
	EnvFilter []PrometheusFilter `json:"envFilter,omitempty"`
	AdditionalFilters []PrometheusFilter `json:"additionalFilters,omitempty"`
	Aggregation string `json:"aggregation,omitempty"`
	IsManualQuery bool `json:"isManualQuery,omitempty"`
}
