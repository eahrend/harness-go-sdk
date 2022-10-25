/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
type V1EventSeries struct {
	Count            int32        `json:"count,omitempty"`
	LastObservedTime *V1MicroTime `json:"lastObservedTime,omitempty"`
}
