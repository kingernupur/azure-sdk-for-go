// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridkubernetes/armhybridkubernetes/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectedClusterServer is a fake server for instances of the armhybridkubernetes.ConnectedClusterClient type.
type ConnectedClusterServer struct {
	// BeginCreateOrReplace is the fake for method ConnectedClusterClient.BeginCreateOrReplace
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrReplace func(ctx context.Context, resourceGroupName string, clusterName string, connectedCluster armhybridkubernetes.ConnectedCluster, options *armhybridkubernetes.ConnectedClusterClientBeginCreateOrReplaceOptions) (resp azfake.PollerResponder[armhybridkubernetes.ConnectedClusterClientCreateOrReplaceResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ConnectedClusterClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, options *armhybridkubernetes.ConnectedClusterClientBeginDeleteOptions) (resp azfake.PollerResponder[armhybridkubernetes.ConnectedClusterClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectedClusterClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, options *armhybridkubernetes.ConnectedClusterClientGetOptions) (resp azfake.Responder[armhybridkubernetes.ConnectedClusterClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ConnectedClusterClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armhybridkubernetes.ConnectedClusterClientListByResourceGroupOptions) (resp azfake.PagerResponder[armhybridkubernetes.ConnectedClusterClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ConnectedClusterClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armhybridkubernetes.ConnectedClusterClientListBySubscriptionOptions) (resp azfake.PagerResponder[armhybridkubernetes.ConnectedClusterClientListBySubscriptionResponse])

	// ListClusterUserCredential is the fake for method ConnectedClusterClient.ListClusterUserCredential
	// HTTP status codes to indicate success: http.StatusOK
	ListClusterUserCredential func(ctx context.Context, resourceGroupName string, clusterName string, properties armhybridkubernetes.ListClusterUserCredentialProperties, options *armhybridkubernetes.ConnectedClusterClientListClusterUserCredentialOptions) (resp azfake.Responder[armhybridkubernetes.ConnectedClusterClientListClusterUserCredentialResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method ConnectedClusterClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, clusterName string, connectedClusterPatch armhybridkubernetes.ConnectedClusterPatch, options *armhybridkubernetes.ConnectedClusterClientUpdateOptions) (resp azfake.Responder[armhybridkubernetes.ConnectedClusterClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewConnectedClusterServerTransport creates a new instance of ConnectedClusterServerTransport with the provided implementation.
// The returned ConnectedClusterServerTransport instance is connected to an instance of armhybridkubernetes.ConnectedClusterClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectedClusterServerTransport(srv *ConnectedClusterServer) *ConnectedClusterServerTransport {
	return &ConnectedClusterServerTransport{
		srv:                         srv,
		beginCreateOrReplace:        newTracker[azfake.PollerResponder[armhybridkubernetes.ConnectedClusterClientCreateOrReplaceResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armhybridkubernetes.ConnectedClusterClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armhybridkubernetes.ConnectedClusterClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armhybridkubernetes.ConnectedClusterClientListBySubscriptionResponse]](),
	}
}

// ConnectedClusterServerTransport connects instances of armhybridkubernetes.ConnectedClusterClient to instances of ConnectedClusterServer.
// Don't use this type directly, use NewConnectedClusterServerTransport instead.
type ConnectedClusterServerTransport struct {
	srv                         *ConnectedClusterServer
	beginCreateOrReplace        *tracker[azfake.PollerResponder[armhybridkubernetes.ConnectedClusterClientCreateOrReplaceResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armhybridkubernetes.ConnectedClusterClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armhybridkubernetes.ConnectedClusterClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armhybridkubernetes.ConnectedClusterClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for ConnectedClusterServerTransport.
func (c *ConnectedClusterServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ConnectedClusterServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if connectedClusterServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = connectedClusterServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ConnectedClusterClient.BeginCreateOrReplace":
				res.resp, res.err = c.dispatchBeginCreateOrReplace(req)
			case "ConnectedClusterClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ConnectedClusterClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ConnectedClusterClient.NewListByResourceGroupPager":
				res.resp, res.err = c.dispatchNewListByResourceGroupPager(req)
			case "ConnectedClusterClient.NewListBySubscriptionPager":
				res.resp, res.err = c.dispatchNewListBySubscriptionPager(req)
			case "ConnectedClusterClient.ListClusterUserCredential":
				res.resp, res.err = c.dispatchListClusterUserCredential(req)
			case "ConnectedClusterClient.Update":
				res.resp, res.err = c.dispatchUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (c *ConnectedClusterServerTransport) dispatchBeginCreateOrReplace(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrReplace == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrReplace not implemented")}
	}
	beginCreateOrReplace := c.beginCreateOrReplace.get(req)
	if beginCreateOrReplace == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridkubernetes.ConnectedCluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrReplace(req.Context(), resourceGroupNameParam, clusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrReplace = &respr
		c.beginCreateOrReplace.add(req, beginCreateOrReplace)
	}

	resp, err := server.PollerResponderNext(beginCreateOrReplace, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrReplace.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrReplace) {
		c.beginCreateOrReplace.remove(req)
	}

	return resp, nil
}

func (c *ConnectedClusterServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ConnectedClusterServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectedCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectedClusterServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armhybridkubernetes.ConnectedClusterClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *ConnectedClusterServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armhybridkubernetes.ConnectedClusterClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *ConnectedClusterServerTransport) dispatchListClusterUserCredential(req *http.Request) (*http.Response, error) {
	if c.srv.ListClusterUserCredential == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListClusterUserCredential not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listClusterUserCredential`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridkubernetes.ListClusterUserCredentialProperties](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.ListClusterUserCredential(req.Context(), resourceGroupNameParam, clusterNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CredentialResults, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectedClusterServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Kubernetes/connectedClusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhybridkubernetes.ConnectedClusterPatch](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, clusterNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConnectedCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ConnectedClusterServerTransport
var connectedClusterServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
