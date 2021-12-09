# Go API client for authz

This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./authz"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.harness.io/gateway/authz/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AclApi* | [**GetAccessControlList**](docs/AclApi.md#getaccesscontrollist) | **Post** /acl | Check for permission on resource(s) for a principal
*DefaultApi* | [**EvaluateCustomFeatureRestriction**](docs/DefaultApi.md#evaluatecustomfeaturerestriction) | **Put** /enforcement/client/custom/{featureRestrictionName} | 
*DefaultApi* | [**GetFeatureUsage**](docs/DefaultApi.md#getfeatureusage) | **Put** /enforcement/client/usage/{featureRestrictionName} | 
*PermissionsApi* | [**GetPermissionList**](docs/PermissionsApi.md#getpermissionlist) | **Get** /permissions | Get all permissions in a scope or all permissions in the system.
*PermissionsApi* | [**GetPermissionResourceTypesList**](docs/PermissionsApi.md#getpermissionresourcetypeslist) | **Get** /permissions/resourcetypes | Get all resource types for permissions in a scope or in the system.
*RoleAssignmentsApi* | [**DeleteRoleAssignment**](docs/RoleAssignmentsApi.md#deleteroleassignment) | **Delete** /roleassignments/{identifier} | Delete an existing role assignment by identifier
*RoleAssignmentsApi* | [**GetFilteredRoleAssignmentList**](docs/RoleAssignmentsApi.md#getfilteredroleassignmentlist) | **Post** /roleassignments/filter | List role assignments in the scope according to the given filter
*RoleAssignmentsApi* | [**GetRoleAssignmentAggregateList**](docs/RoleAssignmentsApi.md#getroleassignmentaggregatelist) | **Post** /roleassignments/aggregate | List role assignments in the scope according to the given filter with added metadata
*RoleAssignmentsApi* | [**GetRoleAssignmentList**](docs/RoleAssignmentsApi.md#getroleassignmentlist) | **Get** /roleassignments | List role assignments in the given scope
*RoleAssignmentsApi* | [**PostRoleAssignment**](docs/RoleAssignmentsApi.md#postroleassignment) | **Post** /roleassignments | Create role assignment in the given scope
*RoleAssignmentsApi* | [**PostRoleAssignments**](docs/RoleAssignmentsApi.md#postroleassignments) | **Post** /roleassignments/multi | Create multiple role assignments in a scope. Returns all successfully created role assignments. Ignores failures and duplicates.
*RoleAssignmentsApi* | [**PutRoleAssignment**](docs/RoleAssignmentsApi.md#putroleassignment) | **Put** /roleassignments/{identifier} | Update existing role assignment by identifier and scope. Only changing the disabled/enabled state is allowed.
*RoleAssignmentsApi* | [**ValidateRoleAssignment**](docs/RoleAssignmentsApi.md#validateroleassignment) | **Post** /roleassignments/validate | Check whether a proposed role assignment is valid.
*RolesApi* | [**DeleteRole**](docs/RolesApi.md#deleterole) | **Delete** /roles/{identifier} | Delete a Custom Role in a scope
*RolesApi* | [**GetRole**](docs/RolesApi.md#getrole) | **Get** /roles/{identifier} | Get a Role by identifier
*RolesApi* | [**GetRoleList**](docs/RolesApi.md#getrolelist) | **Get** /roles | List roles in the given scope
*RolesApi* | [**PostRole**](docs/RolesApi.md#postrole) | **Post** /roles | Create a Custom Role in a scope
*RolesApi* | [**PutRole**](docs/RolesApi.md#putrole) | **Put** /roles/{identifier} | Update a Custom Role by identifier

## Documentation For Models

 - [AccessCheckRequest](docs/AccessCheckRequest.md)
 - [AccessCheckResponse](docs/AccessCheckResponse.md)
 - [AccessControl](docs/AccessControl.md)
 - [AccessDeniedError](docs/AccessDeniedError.md)
 - [AvailabilityRestrictionMetadataDto](docs/AvailabilityRestrictionMetadataDto.md)
 - [CustomRestrictionEvaluationDto](docs/CustomRestrictionEvaluationDto.md)
 - [CustomRestrictionMetadataDto](docs/CustomRestrictionMetadataDto.md)
 - [DurationRestrictionMetadataDto](docs/DurationRestrictionMetadataDto.md)
 - [ErrorMetadataDto](docs/ErrorMetadataDto.md)
 - [Failure](docs/Failure.md)
 - [FeatureRestrictionUsageDto](docs/FeatureRestrictionUsageDto.md)
 - [ModelError](docs/ModelError.md)
 - [PageResponseRoleAssignmentResponse](docs/PageResponseRoleAssignmentResponse.md)
 - [PageResponseRoleResponse](docs/PageResponseRoleResponse.md)
 - [Permission](docs/Permission.md)
 - [PermissionCheck](docs/PermissionCheck.md)
 - [PermissionResponse](docs/PermissionResponse.md)
 - [Principal](docs/Principal.md)
 - [RateLimitRestrictionMetadataDto](docs/RateLimitRestrictionMetadataDto.md)
 - [ResourceGroup](docs/ResourceGroup.md)
 - [ResourceScope](docs/ResourceScope.md)
 - [ResponseDtoAccessCheckResponse](docs/ResponseDtoAccessCheckResponse.md)
 - [ResponseDtoBoolean](docs/ResponseDtoBoolean.md)
 - [ResponseDtoFeatureRestrictionUsageDto](docs/ResponseDtoFeatureRestrictionUsageDto.md)
 - [ResponseDtoListPermissionResponse](docs/ResponseDtoListPermissionResponse.md)
 - [ResponseDtoListRoleAssignmentResponse](docs/ResponseDtoListRoleAssignmentResponse.md)
 - [ResponseDtoPageResponseRoleAssignmentResponse](docs/ResponseDtoPageResponseRoleAssignmentResponse.md)
 - [ResponseDtoPageResponseRoleResponse](docs/ResponseDtoPageResponseRoleResponse.md)
 - [ResponseDtoRoleAssignmentAggregateResponse](docs/ResponseDtoRoleAssignmentAggregateResponse.md)
 - [ResponseDtoRoleAssignmentResponse](docs/ResponseDtoRoleAssignmentResponse.md)
 - [ResponseDtoRoleAssignmentValidationResponse](docs/ResponseDtoRoleAssignmentValidationResponse.md)
 - [ResponseDtoRoleResponse](docs/ResponseDtoRoleResponse.md)
 - [ResponseDtoSetString](docs/ResponseDtoSetString.md)
 - [ResponseMessage](docs/ResponseMessage.md)
 - [ResponseMessageException](docs/ResponseMessageException.md)
 - [ResponseMessageExceptionStackTrace](docs/ResponseMessageExceptionStackTrace.md)
 - [ResponseMessageExceptionSuppressed](docs/ResponseMessageExceptionSuppressed.md)
 - [RestrictionMetadataDto](docs/RestrictionMetadataDto.md)
 - [Role](docs/Role.md)
 - [RoleAssignment](docs/RoleAssignment.md)
 - [RoleAssignmentAggregateResponse](docs/RoleAssignmentAggregateResponse.md)
 - [RoleAssignmentCreateRequest](docs/RoleAssignmentCreateRequest.md)
 - [RoleAssignmentFilter](docs/RoleAssignmentFilter.md)
 - [RoleAssignmentResponse](docs/RoleAssignmentResponse.md)
 - [RoleAssignmentValidationRequest](docs/RoleAssignmentValidationRequest.md)
 - [RoleAssignmentValidationResponse](docs/RoleAssignmentValidationResponse.md)
 - [RoleResponse](docs/RoleResponse.md)
 - [SampleErrorMetadataDto](docs/SampleErrorMetadataDto.md)
 - [Scope](docs/Scope.md)
 - [SortOrder](docs/SortOrder.md)
 - [StaticLimitRestrictionMetadataDto](docs/StaticLimitRestrictionMetadataDto.md)
 - [TimeUnit](docs/TimeUnit.md)
 - [ValidationError](docs/ValidationError.md)
 - [ValidationResult](docs/ValidationResult.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

contact@harness.io