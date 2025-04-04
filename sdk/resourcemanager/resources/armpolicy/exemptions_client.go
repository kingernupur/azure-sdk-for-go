// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExemptionsClient contains the methods for the PolicyExemptions group.
// Don't use this type directly, use NewExemptionsClient() instead.
type ExemptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExemptionsClient creates a new instance of ExemptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExemptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExemptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExemptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - This operation creates or updates a policy exemption with the given scope and name. Policy exemptions
// apply to all resources contained within their scope. For example, when you create a policy
// exemption at resource group scope for a policy assignment at the same or above level, the exemption exempts to all applicable
// resources in the resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - scope - The scope of the policy exemption. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format:
//     '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}',
//     or resource (format:
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
//   - policyExemptionName - The name of the policy exemption to delete.
//   - parameters - Parameters for the policy exemption.
//   - options - ExemptionsClientCreateOrUpdateOptions contains the optional parameters for the ExemptionsClient.CreateOrUpdate
//     method.
func (client *ExemptionsClient) CreateOrUpdate(ctx context.Context, scope string, policyExemptionName string, parameters Exemption, options *ExemptionsClientCreateOrUpdateOptions) (ExemptionsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ExemptionsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, scope, policyExemptionName, parameters, options)
	if err != nil {
		return ExemptionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExemptionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ExemptionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExemptionsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, policyExemptionName string, parameters Exemption, _ *ExemptionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ExemptionsClient) createOrUpdateHandleResponse(resp *http.Response) (ExemptionsClientCreateOrUpdateResponse, error) {
	result := ExemptionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Exemption); err != nil {
		return ExemptionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - This operation deletes a policy exemption, given its name and the scope it was created in. The scope of a policy
// exemption is the part of its ID preceding
// '/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - scope - The scope of the policy exemption. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format:
//     '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}',
//     or resource (format:
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
//   - policyExemptionName - The name of the policy exemption to delete.
//   - options - ExemptionsClientDeleteOptions contains the optional parameters for the ExemptionsClient.Delete method.
func (client *ExemptionsClient) Delete(ctx context.Context, scope string, policyExemptionName string, options *ExemptionsClientDeleteOptions) (ExemptionsClientDeleteResponse, error) {
	var err error
	const operationName = "ExemptionsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, scope, policyExemptionName, options)
	if err != nil {
		return ExemptionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExemptionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ExemptionsClientDeleteResponse{}, err
	}
	return ExemptionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExemptionsClient) deleteCreateRequest(ctx context.Context, scope string, policyExemptionName string, _ *ExemptionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - This operation retrieves a single policy exemption, given its name and the scope it was created at.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - scope - The scope of the policy exemption. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format:
//     '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}',
//     or resource (format:
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
//   - policyExemptionName - The name of the policy exemption to delete.
//   - options - ExemptionsClientGetOptions contains the optional parameters for the ExemptionsClient.Get method.
func (client *ExemptionsClient) Get(ctx context.Context, scope string, policyExemptionName string, options *ExemptionsClientGetOptions) (ExemptionsClientGetResponse, error) {
	var err error
	const operationName = "ExemptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, policyExemptionName, options)
	if err != nil {
		return ExemptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExemptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExemptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ExemptionsClient) getCreateRequest(ctx context.Context, scope string, policyExemptionName string, _ *ExemptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExemptionsClient) getHandleResponse(resp *http.Response) (ExemptionsClientGetResponse, error) {
	result := ExemptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Exemption); err != nil {
		return ExemptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - This operation retrieves the list of all policy exemptions associated with the given subscription that match
// the optional given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()',
// 'excludeExpired()' or 'policyAssignmentId eq '{value}”. If $filter is not provided, the unfiltered list includes all policy
// exemptions associated with the subscription, including those that apply
// directly or from management groups that contain the given subscription, as well as any applied to objects contained within
// the subscription.
//
// Generated from API version 2022-07-01-preview
//   - options - ExemptionsClientListOptions contains the optional parameters for the ExemptionsClient.NewListPager method.
func (client *ExemptionsClient) NewListPager(options *ExemptionsClientListOptions) *runtime.Pager[ExemptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExemptionsClientListResponse]{
		More: func(page ExemptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExemptionsClientListResponse) (ExemptionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExemptionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ExemptionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExemptionsClient) listCreateRequest(ctx context.Context, options *ExemptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policyExemptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExemptionsClient) listHandleResponse(resp *http.Response) (ExemptionsClientListResponse, error) {
	result := ExemptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExemptionListResult); err != nil {
		return ExemptionsClientListResponse{}, err
	}
	return result, nil
}

// NewListForManagementGroupPager - This operation retrieves the list of all policy exemptions applicable to the management
// group that match the given $filter. Valid values for $filter are: 'atScope()', 'atExactScope()',
// 'excludeExpired()' or 'policyAssignmentId eq '{value}”. If $filter=atScope() is provided, the returned list includes all
// policy exemptions that are assigned to the management group or the management
// group's ancestors.
//
// Generated from API version 2022-07-01-preview
//   - managementGroupID - The ID of the management group.
//   - options - ExemptionsClientListForManagementGroupOptions contains the optional parameters for the ExemptionsClient.NewListForManagementGroupPager
//     method.
func (client *ExemptionsClient) NewListForManagementGroupPager(managementGroupID string, options *ExemptionsClientListForManagementGroupOptions) *runtime.Pager[ExemptionsClientListForManagementGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExemptionsClientListForManagementGroupResponse]{
		More: func(page ExemptionsClientListForManagementGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExemptionsClientListForManagementGroupResponse) (ExemptionsClientListForManagementGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExemptionsClient.NewListForManagementGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForManagementGroupCreateRequest(ctx, managementGroupID, options)
			}, nil)
			if err != nil {
				return ExemptionsClientListForManagementGroupResponse{}, err
			}
			return client.listForManagementGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForManagementGroupCreateRequest creates the ListForManagementGroup request.
func (client *ExemptionsClient) listForManagementGroupCreateRequest(ctx context.Context, managementGroupID string, options *ExemptionsClientListForManagementGroupOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/managementGroups/{managementGroupId}/providers/Microsoft.Authorization/policyExemptions"
	if managementGroupID == "" {
		return nil, errors.New("parameter managementGroupID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managementGroupId}", url.PathEscape(managementGroupID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForManagementGroupHandleResponse handles the ListForManagementGroup response.
func (client *ExemptionsClient) listForManagementGroupHandleResponse(resp *http.Response) (ExemptionsClientListForManagementGroupResponse, error) {
	result := ExemptionsClientListForManagementGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExemptionListResult); err != nil {
		return ExemptionsClientListForManagementGroupResponse{}, err
	}
	return result, nil
}

// NewListForResourcePager - This operation retrieves the list of all policy exemptions associated with the specified resource
// in the given resource group and subscription that match the optional given $filter. Valid values for
// $filter are: 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}”. If $filter is not
// provided, the unfiltered list includes all policy exemptions associated with the
// resource, including those that apply directly or from all containing scopes, as well as any applied to resources contained
// within the resource. Three parameters plus the resource name are used to
// identify a specific resource. If the resource is not part of a parent resource (the more common case), the parent resource
// path should not be provided (or provided as ”). For example a web app could
// be specified as ({resourceProviderNamespace} == 'Microsoft.Web', {parentResourcePath} == ”, {resourceType} == 'sites',
// {resourceName} == 'MyWebApp'). If the resource is part of a parent resource,
// then all parameters should be provided. For example a virtual machine DNS name could be specified as ({resourceProviderNamespace}
// == 'Microsoft.Compute', {parentResourcePath} ==
// 'virtualMachines/MyVirtualMachine', {resourceType} == 'domainNames', {resourceName} == 'MyComputerName'). A convenient
// alternative to providing the namespace and type name separately is to provide
// both in the {resourceType} parameter, format: ({resourceProviderNamespace} == ”, {parentResourcePath} == ”, {resourceType}
// == 'Microsoft.Web/sites', {resourceName} == 'MyWebApp').
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group containing the resource.
//   - resourceProviderNamespace - The namespace of the resource provider. For example, the namespace of a virtual machine is
//     Microsoft.Compute (from Microsoft.Compute/virtualMachines)
//   - parentResourcePath - The parent resource path. Use empty string if there is none.
//   - resourceType - The resource type name. For example the type name of a web app is 'sites' (from Microsoft.Web/sites).
//   - resourceName - The name of the resource.
//   - options - ExemptionsClientListForResourceOptions contains the optional parameters for the ExemptionsClient.NewListForResourcePager
//     method.
func (client *ExemptionsClient) NewListForResourcePager(resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *ExemptionsClientListForResourceOptions) *runtime.Pager[ExemptionsClientListForResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExemptionsClientListForResourceResponse]{
		More: func(page ExemptionsClientListForResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExemptionsClientListForResourceResponse) (ExemptionsClientListForResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExemptionsClient.NewListForResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, options)
			}, nil)
			if err != nil {
				return ExemptionsClientListForResourceResponse{}, err
			}
			return client.listForResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *ExemptionsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, options *ExemptionsClientListForResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/policyExemptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", url.PathEscape(resourceProviderNamespace))
	urlPath = strings.ReplaceAll(urlPath, "{parentResourcePath}", parentResourcePath)
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *ExemptionsClient) listForResourceHandleResponse(resp *http.Response) (ExemptionsClientListForResourceResponse, error) {
	result := ExemptionsClientListForResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExemptionListResult); err != nil {
		return ExemptionsClientListForResourceResponse{}, err
	}
	return result, nil
}

// NewListForResourceGroupPager - This operation retrieves the list of all policy exemptions associated with the given resource
// group in the given subscription that match the optional given $filter. Valid values for $filter are:
// 'atScope()', 'atExactScope()', 'excludeExpired()' or 'policyAssignmentId eq '{value}”. If $filter is not provided, the
// unfiltered list includes all policy exemptions associated with the resource
// group, including those that apply directly or apply from containing scopes, as well as any applied to resources contained
// within the resource group.
//
// Generated from API version 2022-07-01-preview
//   - resourceGroupName - The name of the resource group containing the resource.
//   - options - ExemptionsClientListForResourceGroupOptions contains the optional parameters for the ExemptionsClient.NewListForResourceGroupPager
//     method.
func (client *ExemptionsClient) NewListForResourceGroupPager(resourceGroupName string, options *ExemptionsClientListForResourceGroupOptions) *runtime.Pager[ExemptionsClientListForResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExemptionsClientListForResourceGroupResponse]{
		More: func(page ExemptionsClientListForResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExemptionsClientListForResourceGroupResponse) (ExemptionsClientListForResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExemptionsClient.NewListForResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ExemptionsClientListForResourceGroupResponse{}, err
			}
			return client.listForResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *ExemptionsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExemptionsClientListForResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/policyExemptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *ExemptionsClient) listForResourceGroupHandleResponse(resp *http.Response) (ExemptionsClientListForResourceGroupResponse, error) {
	result := ExemptionsClientListForResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExemptionListResult); err != nil {
		return ExemptionsClientListForResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - This operation updates a policy exemption with the given scope and name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01-preview
//   - scope - The scope of the policy exemption. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'),
//     subscription (format:
//     '/subscriptions/{subscriptionId}'), resource group (format: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}',
//     or resource (format:
//     '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/[{parentResourcePath}/]{resourceType}/{resourceName}'
//   - policyExemptionName - The name of the policy exemption to delete.
//   - parameters - Parameters for policy exemption patch request.
//   - options - ExemptionsClientUpdateOptions contains the optional parameters for the ExemptionsClient.Update method.
func (client *ExemptionsClient) Update(ctx context.Context, scope string, policyExemptionName string, parameters ExemptionUpdate, options *ExemptionsClientUpdateOptions) (ExemptionsClientUpdateResponse, error) {
	var err error
	const operationName = "ExemptionsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, scope, policyExemptionName, parameters, options)
	if err != nil {
		return ExemptionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExemptionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExemptionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ExemptionsClient) updateCreateRequest(ctx context.Context, scope string, policyExemptionName string, parameters ExemptionUpdate, _ *ExemptionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/policyExemptions/{policyExemptionName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if policyExemptionName == "" {
		return nil, errors.New("parameter policyExemptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{policyExemptionName}", url.PathEscape(policyExemptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ExemptionsClient) updateHandleResponse(resp *http.Response) (ExemptionsClientUpdateResponse, error) {
	result := ExemptionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Exemption); err != nil {
		return ExemptionsClientUpdateResponse{}, err
	}
	return result, nil
}
