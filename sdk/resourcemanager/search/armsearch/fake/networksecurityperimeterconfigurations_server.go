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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
	"net/http"
	"net/url"
	"regexp"
)

// NetworkSecurityPerimeterConfigurationsServer is a fake server for instances of the armsearch.NetworkSecurityPerimeterConfigurationsClient type.
type NetworkSecurityPerimeterConfigurationsServer struct {
	// Get is the fake for method NetworkSecurityPerimeterConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, options *armsearch.NetworkSecurityPerimeterConfigurationsClientGetOptions) (resp azfake.Responder[armsearch.NetworkSecurityPerimeterConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method NetworkSecurityPerimeterConfigurationsClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, searchServiceName string, options *armsearch.NetworkSecurityPerimeterConfigurationsClientListByServiceOptions) (resp azfake.PagerResponder[armsearch.NetworkSecurityPerimeterConfigurationsClientListByServiceResponse])

	// BeginReconcile is the fake for method NetworkSecurityPerimeterConfigurationsClient.BeginReconcile
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginReconcile func(ctx context.Context, resourceGroupName string, searchServiceName string, nspConfigName string, options *armsearch.NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (resp azfake.PollerResponder[armsearch.NetworkSecurityPerimeterConfigurationsClientReconcileResponse], errResp azfake.ErrorResponder)
}

// NewNetworkSecurityPerimeterConfigurationsServerTransport creates a new instance of NetworkSecurityPerimeterConfigurationsServerTransport with the provided implementation.
// The returned NetworkSecurityPerimeterConfigurationsServerTransport instance is connected to an instance of armsearch.NetworkSecurityPerimeterConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNetworkSecurityPerimeterConfigurationsServerTransport(srv *NetworkSecurityPerimeterConfigurationsServer) *NetworkSecurityPerimeterConfigurationsServerTransport {
	return &NetworkSecurityPerimeterConfigurationsServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armsearch.NetworkSecurityPerimeterConfigurationsClientListByServiceResponse]](),
		beginReconcile:        newTracker[azfake.PollerResponder[armsearch.NetworkSecurityPerimeterConfigurationsClientReconcileResponse]](),
	}
}

// NetworkSecurityPerimeterConfigurationsServerTransport connects instances of armsearch.NetworkSecurityPerimeterConfigurationsClient to instances of NetworkSecurityPerimeterConfigurationsServer.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsServerTransport instead.
type NetworkSecurityPerimeterConfigurationsServerTransport struct {
	srv                   *NetworkSecurityPerimeterConfigurationsServer
	newListByServicePager *tracker[azfake.PagerResponder[armsearch.NetworkSecurityPerimeterConfigurationsClientListByServiceResponse]]
	beginReconcile        *tracker[azfake.PollerResponder[armsearch.NetworkSecurityPerimeterConfigurationsClientReconcileResponse]]
}

// Do implements the policy.Transporter interface for NetworkSecurityPerimeterConfigurationsServerTransport.
func (n *NetworkSecurityPerimeterConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if networkSecurityPerimeterConfigurationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = networkSecurityPerimeterConfigurationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NetworkSecurityPerimeterConfigurationsClient.Get":
				res.resp, res.err = n.dispatchGet(req)
			case "NetworkSecurityPerimeterConfigurationsClient.NewListByServicePager":
				res.resp, res.err = n.dispatchNewListByServicePager(req)
			case "NetworkSecurityPerimeterConfigurationsClient.BeginReconcile":
				res.resp, res.err = n.dispatchBeginReconcile(req)
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

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<nspConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
	if err != nil {
		return nil, err
	}
	nspConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("nspConfigName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, searchServiceNameParam, nspConfigNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSecurityPerimeterConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := n.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByServicePager(resourceGroupNameParam, searchServiceNameParam, nil)
		newListByServicePager = &resp
		n.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armsearch.NetworkSecurityPerimeterConfigurationsClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		n.newListByServicePager.remove(req)
	}
	return resp, nil
}

func (n *NetworkSecurityPerimeterConfigurationsServerTransport) dispatchBeginReconcile(req *http.Request) (*http.Response, error) {
	if n.srv.BeginReconcile == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReconcile not implemented")}
	}
	beginReconcile := n.beginReconcile.get(req)
	if beginReconcile == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Search/searchServices/(?P<searchServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<nspConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reconcile`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		searchServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("searchServiceName")])
		if err != nil {
			return nil, err
		}
		nspConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("nspConfigName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginReconcile(req.Context(), resourceGroupNameParam, searchServiceNameParam, nspConfigNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReconcile = &respr
		n.beginReconcile.add(req, beginReconcile)
	}

	resp, err := server.PollerResponderNext(beginReconcile, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginReconcile.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReconcile) {
		n.beginReconcile.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to NetworkSecurityPerimeterConfigurationsServerTransport
var networkSecurityPerimeterConfigurationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
