
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

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type PipelineExecutionApiService service
/*
PipelineExecutionApiService Execute a Pipeline with Input Set References
Execute a Pipeline with Input Set References
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param accountIdentifier Account Identifier for the Entity.
 * @param orgIdentifier Organization Identifier for the Entity.
 * @param projectIdentifier Project Identifier for the Entity.
 * @param moduleType Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI
 * @param identifier Pipeline identifier for the entity. Identifier of the Pipeline to be executed
 * @param optional nil or *PipelineExecutionApiPostPipelineExecuteWithInputSetListOpts - Optional Parameters:
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
     * @param "UseFQNIfError" (optional.Bool) - 
@return ResponseDtoPlanExecutionResponse
*/

type PipelineExecutionApiPostPipelineExecuteWithInputSetListOpts struct {
    Branch optional.String
    RepoIdentifier optional.String
    GetDefaultFromOtherRepo optional.Bool
    UseFQNIfError optional.Bool
}

func (a *PipelineExecutionApiService) PostPipelineExecuteWithInputSetList(ctx context.Context, body MergeInputSetRequest, accountIdentifier string, orgIdentifier string, projectIdentifier string, moduleType string, identifier string, localVarOptionals *PipelineExecutionApiPostPipelineExecuteWithInputSetListOpts) (ResponseDtoPlanExecutionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoPlanExecutionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pipeline/api/pipeline/execute/{identifier}/inputSetList"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", fmt.Sprintf("%v", identifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	localVarQueryParams.Add("orgIdentifier", parameterToString(orgIdentifier, ""))
	localVarQueryParams.Add("projectIdentifier", parameterToString(projectIdentifier, ""))
	localVarQueryParams.Add("moduleType", parameterToString(moduleType, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseFQNIfError.IsSet() {
		localVarQueryParams.Add("useFQNIfError", parameterToString(localVarOptionals.UseFQNIfError.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPlanExecutionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
PipelineExecutionApiService Execute a Pipeline with Runtime Input YAML
Execute a Pipeline with Runtime Input YAML
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountIdentifier Account Identifier for the Entity.
 * @param orgIdentifier Organization Identifier for the Entity.
 * @param projectIdentifier Project Identifier for the Entity.
 * @param moduleType Module type for the entity. If its from deployments,type will be CD , if its from build type will be CI
 * @param identifier Pipeline identifier for the entity. Identifier of the Pipeline to be executed
 * @param optional nil or *PipelineExecutionApiPostPipelineExecuteWithInputSetYamlOpts - Optional Parameters:
     * @param "Body" (optional.Interface of string) -  Enter Runtime Input YAML if the Pipeline contains Runtime Inputs. Please refer to https://ngdocs.harness.io/article/f6yobn7iq0 and https://ngdocs.harness.io/article/1eishcolt3 to see how to generate Runtime Input YAML for a Pipeline.
     * @param "Branch" (optional.String) -  Name of the branch.
     * @param "RepoIdentifier" (optional.String) -  Git Sync Config Id.
     * @param "GetDefaultFromOtherRepo" (optional.Bool) -  if true, return all the default entities
     * @param "UseFQNIfError" (optional.Bool) - 
@return ResponseDtoPlanExecutionResponse
*/

type PipelineExecutionApiPostPipelineExecuteWithInputSetYamlOpts struct {
    Body optional.Interface
    Branch optional.String
    RepoIdentifier optional.String
    GetDefaultFromOtherRepo optional.Bool
    UseFQNIfError optional.Bool
}

func (a *PipelineExecutionApiService) PostPipelineExecuteWithInputSetYaml(ctx context.Context, accountIdentifier string, orgIdentifier string, projectIdentifier string, moduleType string, identifier string, localVarOptionals *PipelineExecutionApiPostPipelineExecuteWithInputSetYamlOpts) (ResponseDtoPlanExecutionResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoPlanExecutionResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pipeline/api/pipeline/execute/{identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"identifier"+"}", fmt.Sprintf("%v", identifier), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	localVarQueryParams.Add("orgIdentifier", parameterToString(orgIdentifier, ""))
	localVarQueryParams.Add("projectIdentifier", parameterToString(projectIdentifier, ""))
	localVarQueryParams.Add("moduleType", parameterToString(moduleType, ""))
	if localVarOptionals != nil && localVarOptionals.Branch.IsSet() {
		localVarQueryParams.Add("branch", parameterToString(localVarOptionals.Branch.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RepoIdentifier.IsSet() {
		localVarQueryParams.Add("repoIdentifier", parameterToString(localVarOptionals.RepoIdentifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.GetDefaultFromOtherRepo.IsSet() {
		localVarQueryParams.Add("getDefaultFromOtherRepo", parameterToString(localVarOptionals.GetDefaultFromOtherRepo.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.UseFQNIfError.IsSet() {
		localVarQueryParams.Add("useFQNIfError", parameterToString(localVarOptionals.UseFQNIfError.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/yaml"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		
		localVarOptionalBody:= localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoPlanExecutionResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
PipelineExecutionApiService Execute an Interrupt
Executes an Interrupt on a Given Execution
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param accountIdentifier Account Identifier for the Entity.
 * @param orgIdentifier Organization Identifier for the Entity.
 * @param projectIdentifier Project Identifier for the Entity.
 * @param interruptType The Interrupt type needed to be applied to the execution. Choose a value from the enum list.
 * @param planExecutionId The Pipeline Execution Id on which the Interrupt needs to be applied.
@return ResponseDtoInterruptResponse
*/
func (a *PipelineExecutionApiService) PutHandleInterrupt(ctx context.Context, accountIdentifier string, orgIdentifier string, projectIdentifier string, interruptType string, planExecutionId string) (ResponseDtoInterruptResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue ResponseDtoInterruptResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pipeline/api/pipeline/execute/interrupt/{planExecutionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"planExecutionId"+"}", fmt.Sprintf("%v", planExecutionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("accountIdentifier", parameterToString(accountIdentifier, ""))
	localVarQueryParams.Add("orgIdentifier", parameterToString(orgIdentifier, ""))
	localVarQueryParams.Add("projectIdentifier", parameterToString(projectIdentifier, ""))
	localVarQueryParams.Add("interruptType", parameterToString(interruptType, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/yaml"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["x-api-key"] = key
			
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v ResponseDtoInterruptResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
