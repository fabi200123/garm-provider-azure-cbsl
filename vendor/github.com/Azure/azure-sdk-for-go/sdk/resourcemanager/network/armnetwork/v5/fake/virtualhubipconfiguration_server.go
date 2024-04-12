//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualHubIPConfigurationServer is a fake server for instances of the armnetwork.VirtualHubIPConfigurationClient type.
type VirtualHubIPConfigurationServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualHubIPConfigurationClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, parameters armnetwork.HubIPConfiguration, options *armnetwork.VirtualHubIPConfigurationClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubIPConfigurationClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualHubIPConfigurationClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *armnetwork.VirtualHubIPConfigurationClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualHubIPConfigurationClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualHubIPConfigurationClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualHubName string, ipConfigName string, options *armnetwork.VirtualHubIPConfigurationClientGetOptions) (resp azfake.Responder[armnetwork.VirtualHubIPConfigurationClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualHubIPConfigurationClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, virtualHubName string, options *armnetwork.VirtualHubIPConfigurationClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualHubIPConfigurationClientListResponse])
}

// NewVirtualHubIPConfigurationServerTransport creates a new instance of VirtualHubIPConfigurationServerTransport with the provided implementation.
// The returned VirtualHubIPConfigurationServerTransport instance is connected to an instance of armnetwork.VirtualHubIPConfigurationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualHubIPConfigurationServerTransport(srv *VirtualHubIPConfigurationServer) *VirtualHubIPConfigurationServerTransport {
	return &VirtualHubIPConfigurationServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.VirtualHubIPConfigurationClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.VirtualHubIPConfigurationClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.VirtualHubIPConfigurationClientListResponse]](),
	}
}

// VirtualHubIPConfigurationServerTransport connects instances of armnetwork.VirtualHubIPConfigurationClient to instances of VirtualHubIPConfigurationServer.
// Don't use this type directly, use NewVirtualHubIPConfigurationServerTransport instead.
type VirtualHubIPConfigurationServerTransport struct {
	srv                 *VirtualHubIPConfigurationServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.VirtualHubIPConfigurationClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.VirtualHubIPConfigurationClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.VirtualHubIPConfigurationClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualHubIPConfigurationServerTransport.
func (v *VirtualHubIPConfigurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualHubIPConfigurationClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualHubIPConfigurationClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualHubIPConfigurationClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualHubIPConfigurationClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualHubIPConfigurationServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipConfigurations/(?P<ipConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.HubIPConfiguration](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		ipConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, virtualHubNameParam, ipConfigNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualHubIPConfigurationServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipConfigurations/(?P<ipConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		ipConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualHubNameParam, ipConfigNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualHubIPConfigurationServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipConfigurations/(?P<ipConfigName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
	if err != nil {
		return nil, err
	}
	ipConfigNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ipConfigName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, virtualHubNameParam, ipConfigNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HubIPConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualHubIPConfigurationServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualHubs/(?P<virtualHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualHubName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, virtualHubNameParam, nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualHubIPConfigurationClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}
