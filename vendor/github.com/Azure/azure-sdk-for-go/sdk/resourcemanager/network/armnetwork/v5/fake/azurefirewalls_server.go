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

// AzureFirewallsServer is a fake server for instances of the armnetwork.AzureFirewallsClient type.
type AzureFirewallsServer struct {
	// BeginCreateOrUpdate is the fake for method AzureFirewallsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters armnetwork.AzureFirewall, options *armnetwork.AzureFirewallsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AzureFirewallsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, azureFirewallName string, options *armnetwork.AzureFirewallsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AzureFirewallsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureFirewallName string, options *armnetwork.AzureFirewallsClientGetOptions) (resp azfake.Responder[armnetwork.AzureFirewallsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AzureFirewallsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnetwork.AzureFirewallsClientListOptions) (resp azfake.PagerResponder[armnetwork.AzureFirewallsClientListResponse])

	// NewListAllPager is the fake for method AzureFirewallsClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnetwork.AzureFirewallsClientListAllOptions) (resp azfake.PagerResponder[armnetwork.AzureFirewallsClientListAllResponse])

	// BeginListLearnedPrefixes is the fake for method AzureFirewallsClient.BeginListLearnedPrefixes
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListLearnedPrefixes func(ctx context.Context, resourceGroupName string, azureFirewallName string, options *armnetwork.AzureFirewallsClientBeginListLearnedPrefixesOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientListLearnedPrefixesResponse], errResp azfake.ErrorResponder)

	// BeginPacketCapture is the fake for method AzureFirewallsClient.BeginPacketCapture
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginPacketCapture func(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters armnetwork.FirewallPacketCaptureParameters, options *armnetwork.AzureFirewallsClientBeginPacketCaptureOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientPacketCaptureResponse], errResp azfake.ErrorResponder)

	// BeginUpdateTags is the fake for method AzureFirewallsClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, azureFirewallName string, parameters armnetwork.TagsObject, options *armnetwork.AzureFirewallsClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armnetwork.AzureFirewallsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewAzureFirewallsServerTransport creates a new instance of AzureFirewallsServerTransport with the provided implementation.
// The returned AzureFirewallsServerTransport instance is connected to an instance of armnetwork.AzureFirewallsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureFirewallsServerTransport(srv *AzureFirewallsServer) *AzureFirewallsServerTransport {
	return &AzureFirewallsServerTransport{
		srv:                      srv,
		beginCreateOrUpdate:      newTracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientCreateOrUpdateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientDeleteResponse]](),
		newListPager:             newTracker[azfake.PagerResponder[armnetwork.AzureFirewallsClientListResponse]](),
		newListAllPager:          newTracker[azfake.PagerResponder[armnetwork.AzureFirewallsClientListAllResponse]](),
		beginListLearnedPrefixes: newTracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientListLearnedPrefixesResponse]](),
		beginPacketCapture:       newTracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientPacketCaptureResponse]](),
		beginUpdateTags:          newTracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientUpdateTagsResponse]](),
	}
}

// AzureFirewallsServerTransport connects instances of armnetwork.AzureFirewallsClient to instances of AzureFirewallsServer.
// Don't use this type directly, use NewAzureFirewallsServerTransport instead.
type AzureFirewallsServerTransport struct {
	srv                      *AzureFirewallsServer
	beginCreateOrUpdate      *tracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientCreateOrUpdateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientDeleteResponse]]
	newListPager             *tracker[azfake.PagerResponder[armnetwork.AzureFirewallsClientListResponse]]
	newListAllPager          *tracker[azfake.PagerResponder[armnetwork.AzureFirewallsClientListAllResponse]]
	beginListLearnedPrefixes *tracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientListLearnedPrefixesResponse]]
	beginPacketCapture       *tracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientPacketCaptureResponse]]
	beginUpdateTags          *tracker[azfake.PollerResponder[armnetwork.AzureFirewallsClientUpdateTagsResponse]]
}

// Do implements the policy.Transporter interface for AzureFirewallsServerTransport.
func (a *AzureFirewallsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureFirewallsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AzureFirewallsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AzureFirewallsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureFirewallsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AzureFirewallsClient.NewListAllPager":
		resp, err = a.dispatchNewListAllPager(req)
	case "AzureFirewallsClient.BeginListLearnedPrefixes":
		resp, err = a.dispatchBeginListLearnedPrefixes(req)
	case "AzureFirewallsClient.BeginPacketCapture":
		resp, err = a.dispatchBeginPacketCapture(req)
	case "AzureFirewallsClient.BeginUpdateTags":
		resp, err = a.dispatchBeginUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.AzureFirewall](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, azureFirewallNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, azureFirewallNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureFirewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, azureFirewallNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureFirewall, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.AzureFirewallsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := a.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListAllPager(nil)
		newListAllPager = &resp
		a.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armnetwork.AzureFirewallsClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		a.newListAllPager.remove(req)
	}
	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginListLearnedPrefixes(req *http.Request) (*http.Response, error) {
	if a.srv.BeginListLearnedPrefixes == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListLearnedPrefixes not implemented")}
	}
	beginListLearnedPrefixes := a.beginListLearnedPrefixes.get(req)
	if beginListLearnedPrefixes == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/learnedIPPrefixes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginListLearnedPrefixes(req.Context(), resourceGroupNameParam, azureFirewallNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListLearnedPrefixes = &respr
		a.beginListLearnedPrefixes.add(req, beginListLearnedPrefixes)
	}

	resp, err := server.PollerResponderNext(beginListLearnedPrefixes, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginListLearnedPrefixes.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListLearnedPrefixes) {
		a.beginListLearnedPrefixes.remove(req)
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginPacketCapture(req *http.Request) (*http.Response, error) {
	if a.srv.BeginPacketCapture == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPacketCapture not implemented")}
	}
	beginPacketCapture := a.beginPacketCapture.get(req)
	if beginPacketCapture == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/packetCapture`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.FirewallPacketCaptureParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginPacketCapture(req.Context(), resourceGroupNameParam, azureFirewallNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPacketCapture = &respr
		a.beginPacketCapture.add(req, beginPacketCapture)
	}

	resp, err := server.PollerResponderNext(beginPacketCapture, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		a.beginPacketCapture.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPacketCapture) {
		a.beginPacketCapture.remove(req)
	}

	return resp, nil
}

func (a *AzureFirewallsServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	beginUpdateTags := a.beginUpdateTags.get(req)
	if beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/azureFirewalls/(?P<azureFirewallName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureFirewallNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureFirewallName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdateTags(req.Context(), resourceGroupNameParam, azureFirewallNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateTags = &respr
		a.beginUpdateTags.add(req, beginUpdateTags)
	}

	resp, err := server.PollerResponderNext(beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdateTags.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateTags) {
		a.beginUpdateTags.remove(req)
	}

	return resp, nil
}