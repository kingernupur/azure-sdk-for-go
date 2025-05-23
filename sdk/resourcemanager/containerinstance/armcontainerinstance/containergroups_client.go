//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerinstance

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

// ContainerGroupsClient contains the methods for the ContainerGroups group.
// Don't use this type directly, use NewContainerGroupsClient() instead.
type ContainerGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContainerGroupsClient creates a new instance of ContainerGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContainerGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContainerGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContainerGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update container groups with specified configurations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - containerGroup - The properties of the container group to be created or updated.
//   - options - ContainerGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the ContainerGroupsClient.BeginCreateOrUpdate
//     method.
func (client *ContainerGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup ContainerGroup, options *ContainerGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ContainerGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, containerGroupName, containerGroup, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerGroupsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create or update container groups with specified configurations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *ContainerGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup ContainerGroup, options *ContainerGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerGroupsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, containerGroupName, containerGroup, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContainerGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup ContainerGroup, options *ContainerGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, containerGroup); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the specified container group in the specified subscription and resource group. The operation does
// not delete other resources provided by the user, such as volumes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - options - ContainerGroupsClientBeginDeleteOptions contains the optional parameters for the ContainerGroupsClient.BeginDelete
//     method.
func (client *ContainerGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginDeleteOptions) (*runtime.Poller[ContainerGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, containerGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerGroupsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the specified container group in the specified subscription and resource group. The operation does not
// delete other resources provided by the user, such as volumes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *ContainerGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerGroupsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, containerGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContainerGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the properties of the specified container group in the specified subscription and resource group. The operation
// returns the properties of each container group including containers, image registry
// credentials, restart policy, IP address type, OS type, state, and volumes.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - options - ContainerGroupsClientGetOptions contains the optional parameters for the ContainerGroupsClient.Get method.
func (client *ContainerGroupsClient) Get(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientGetOptions) (ContainerGroupsClientGetResponse, error) {
	var err error
	const operationName = "ContainerGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, containerGroupName, options)
	if err != nil {
		return ContainerGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContainerGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContainerGroupsClient) getHandleResponse(resp *http.Response) (ContainerGroupsClientGetResponse, error) {
	result := ContainerGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroup); err != nil {
		return ContainerGroupsClientGetResponse{}, err
	}
	return result, nil
}

// GetOutboundNetworkDependenciesEndpoints - Gets all the network dependencies for this container group to allow complete
// control of network setting and configuration. For container groups, this will always be an empty list.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - options - ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsOptions contains the optional parameters for the
//     ContainerGroupsClient.GetOutboundNetworkDependenciesEndpoints method.
func (client *ContainerGroupsClient) GetOutboundNetworkDependenciesEndpoints(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsOptions) (ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse, error) {
	var err error
	const operationName = "ContainerGroupsClient.GetOutboundNetworkDependenciesEndpoints"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getOutboundNetworkDependenciesEndpointsCreateRequest(ctx, resourceGroupName, containerGroupName, options)
	if err != nil {
		return ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	resp, err := client.getOutboundNetworkDependenciesEndpointsHandleResponse(httpResp)
	return resp, err
}

// getOutboundNetworkDependenciesEndpointsCreateRequest creates the GetOutboundNetworkDependenciesEndpoints request.
func (client *ContainerGroupsClient) getOutboundNetworkDependenciesEndpointsCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/outboundNetworkDependenciesEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOutboundNetworkDependenciesEndpointsHandleResponse handles the GetOutboundNetworkDependenciesEndpoints response.
func (client *ContainerGroupsClient) getOutboundNetworkDependenciesEndpointsHandleResponse(resp *http.Response) (ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse, error) {
	result := ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StringArray); err != nil {
		return ContainerGroupsClientGetOutboundNetworkDependenciesEndpointsResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of container groups in the specified subscription. This operation returns properties of each
// container group including containers, image registry credentials, restart policy, IP address
// type, OS type, state, and volumes.
//
// Generated from API version 2024-05-01-preview
//   - options - ContainerGroupsClientListOptions contains the optional parameters for the ContainerGroupsClient.NewListPager
//     method.
func (client *ContainerGroupsClient) NewListPager(options *ContainerGroupsClientListOptions) *runtime.Pager[ContainerGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerGroupsClientListResponse]{
		More: func(page ContainerGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerGroupsClientListResponse) (ContainerGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ContainerGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ContainerGroupsClient) listCreateRequest(ctx context.Context, options *ContainerGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ContainerInstance/containerGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContainerGroupsClient) listHandleResponse(resp *http.Response) (ContainerGroupsClientListResponse, error) {
	result := ContainerGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupListResult); err != nil {
		return ContainerGroupsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of container groups in a specified subscription and resource group. This operation
// returns properties of each container group including containers, image registry credentials, restart
// policy, IP address type, OS type, state, and volumes.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ContainerGroupsClientListByResourceGroupOptions contains the optional parameters for the ContainerGroupsClient.NewListByResourceGroupPager
//     method.
func (client *ContainerGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *ContainerGroupsClientListByResourceGroupOptions) *runtime.Pager[ContainerGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContainerGroupsClientListByResourceGroupResponse]{
		More: func(page ContainerGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContainerGroupsClientListByResourceGroupResponse) (ContainerGroupsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContainerGroupsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ContainerGroupsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ContainerGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ContainerGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups"
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
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ContainerGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ContainerGroupsClientListByResourceGroupResponse, error) {
	result := ContainerGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroupListResult); err != nil {
		return ContainerGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// BeginRestart - Restarts all containers in a container group in place. If container image has updates, new image will be
// downloaded.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - options - ContainerGroupsClientBeginRestartOptions contains the optional parameters for the ContainerGroupsClient.BeginRestart
//     method.
func (client *ContainerGroupsClient) BeginRestart(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginRestartOptions) (*runtime.Poller[ContainerGroupsClientRestartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.restart(ctx, resourceGroupName, containerGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerGroupsClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerGroupsClientRestartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Restart - Restarts all containers in a container group in place. If container image has updates, new image will be downloaded.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *ContainerGroupsClient) restart(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginRestartOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerGroupsClient.BeginRestart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.restartCreateRequest(ctx, resourceGroupName, containerGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// restartCreateRequest creates the Restart request.
func (client *ContainerGroupsClient) restartCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginStart - Starts all containers in a container group. Compute resources will be allocated and billing will start.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - options - ContainerGroupsClientBeginStartOptions contains the optional parameters for the ContainerGroupsClient.BeginStart
//     method.
func (client *ContainerGroupsClient) BeginStart(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginStartOptions) (*runtime.Poller[ContainerGroupsClientStartResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.start(ctx, resourceGroupName, containerGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ContainerGroupsClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ContainerGroupsClientStartResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Start - Starts all containers in a container group. Compute resources will be allocated and billing will start.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *ContainerGroupsClient) start(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginStartOptions) (*http.Response, error) {
	var err error
	const operationName = "ContainerGroupsClient.BeginStart"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startCreateRequest(ctx, resourceGroupName, containerGroupName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// startCreateRequest creates the Start request.
func (client *ContainerGroupsClient) startCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientBeginStartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/start"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Stop - Stops all containers in a container group. Compute resources will be deallocated and billing will stop.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - options - ContainerGroupsClientStopOptions contains the optional parameters for the ContainerGroupsClient.Stop method.
func (client *ContainerGroupsClient) Stop(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientStopOptions) (ContainerGroupsClientStopResponse, error) {
	var err error
	const operationName = "ContainerGroupsClient.Stop"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopCreateRequest(ctx, resourceGroupName, containerGroupName, options)
	if err != nil {
		return ContainerGroupsClientStopResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerGroupsClientStopResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ContainerGroupsClientStopResponse{}, err
	}
	return ContainerGroupsClientStopResponse{}, nil
}

// stopCreateRequest creates the Stop request.
func (client *ContainerGroupsClient) stopCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, options *ContainerGroupsClientStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/stop"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Updates container group tags with specified values.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - containerGroupName - The name of the container group.
//   - resource - The container group resource with just the tags to be updated.
//   - options - ContainerGroupsClientUpdateOptions contains the optional parameters for the ContainerGroupsClient.Update method.
func (client *ContainerGroupsClient) Update(ctx context.Context, resourceGroupName string, containerGroupName string, resource Resource, options *ContainerGroupsClientUpdateOptions) (ContainerGroupsClientUpdateResponse, error) {
	var err error
	const operationName = "ContainerGroupsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, containerGroupName, resource, options)
	if err != nil {
		return ContainerGroupsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContainerGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContainerGroupsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ContainerGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, containerGroupName string, resource Resource, options *ContainerGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if containerGroupName == "" {
		return nil, errors.New("parameter containerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerGroupName}", url.PathEscape(containerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ContainerGroupsClient) updateHandleResponse(resp *http.Response) (ContainerGroupsClientUpdateResponse, error) {
	result := ContainerGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContainerGroup); err != nil {
		return ContainerGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
