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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// GalleryApplicationsServer is a fake server for instances of the armcompute.GalleryApplicationsClient type.
type GalleryApplicationsServer struct {
	// BeginCreateOrUpdate is the fake for method GalleryApplicationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication armcompute.GalleryApplication, options *armcompute.GalleryApplicationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryApplicationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GalleryApplicationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *armcompute.GalleryApplicationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.GalleryApplicationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GalleryApplicationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *armcompute.GalleryApplicationsClientGetOptions) (resp azfake.Responder[armcompute.GalleryApplicationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByGalleryPager is the fake for method GalleryApplicationsClient.NewListByGalleryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByGalleryPager func(resourceGroupName string, galleryName string, options *armcompute.GalleryApplicationsClientListByGalleryOptions) (resp azfake.PagerResponder[armcompute.GalleryApplicationsClientListByGalleryResponse])

	// BeginUpdate is the fake for method GalleryApplicationsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplication armcompute.GalleryApplicationUpdate, options *armcompute.GalleryApplicationsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryApplicationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGalleryApplicationsServerTransport creates a new instance of GalleryApplicationsServerTransport with the provided implementation.
// The returned GalleryApplicationsServerTransport instance is connected to an instance of armcompute.GalleryApplicationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGalleryApplicationsServerTransport(srv *GalleryApplicationsServer) *GalleryApplicationsServerTransport {
	return &GalleryApplicationsServerTransport{
		srv:                   srv,
		beginCreateOrUpdate:   newTracker[azfake.PollerResponder[armcompute.GalleryApplicationsClientCreateOrUpdateResponse]](),
		beginDelete:           newTracker[azfake.PollerResponder[armcompute.GalleryApplicationsClientDeleteResponse]](),
		newListByGalleryPager: newTracker[azfake.PagerResponder[armcompute.GalleryApplicationsClientListByGalleryResponse]](),
		beginUpdate:           newTracker[azfake.PollerResponder[armcompute.GalleryApplicationsClientUpdateResponse]](),
	}
}

// GalleryApplicationsServerTransport connects instances of armcompute.GalleryApplicationsClient to instances of GalleryApplicationsServer.
// Don't use this type directly, use NewGalleryApplicationsServerTransport instead.
type GalleryApplicationsServerTransport struct {
	srv                   *GalleryApplicationsServer
	beginCreateOrUpdate   *tracker[azfake.PollerResponder[armcompute.GalleryApplicationsClientCreateOrUpdateResponse]]
	beginDelete           *tracker[azfake.PollerResponder[armcompute.GalleryApplicationsClientDeleteResponse]]
	newListByGalleryPager *tracker[azfake.PagerResponder[armcompute.GalleryApplicationsClientListByGalleryResponse]]
	beginUpdate           *tracker[azfake.PollerResponder[armcompute.GalleryApplicationsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for GalleryApplicationsServerTransport.
func (g *GalleryApplicationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "GalleryApplicationsClient.BeginCreateOrUpdate":
		resp, err = g.dispatchBeginCreateOrUpdate(req)
	case "GalleryApplicationsClient.BeginDelete":
		resp, err = g.dispatchBeginDelete(req)
	case "GalleryApplicationsClient.Get":
		resp, err = g.dispatchGet(req)
	case "GalleryApplicationsClient.NewListByGalleryPager":
		resp, err = g.dispatchNewListByGalleryPager(req)
	case "GalleryApplicationsClient.BeginUpdate":
		resp, err = g.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *GalleryApplicationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryApplication](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, galleryNameParam, galleryApplicationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GalleryApplicationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := g.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), resourceGroupNameParam, galleryNameParam, galleryApplicationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		g.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		g.beginDelete.remove(req)
	}

	return resp, nil
}

func (g *GalleryApplicationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
	if err != nil {
		return nil, err
	}
	galleryApplicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, galleryNameParam, galleryApplicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GalleryApplication, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GalleryApplicationsServerTransport) dispatchNewListByGalleryPager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByGalleryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByGalleryPager not implemented")}
	}
	newListByGalleryPager := g.newListByGalleryPager.get(req)
	if newListByGalleryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByGalleryPager(resourceGroupNameParam, galleryNameParam, nil)
		newListByGalleryPager = &resp
		g.newListByGalleryPager.add(req, newListByGalleryPager)
		server.PagerResponderInjectNextLinks(newListByGalleryPager, req, func(page *armcompute.GalleryApplicationsClientListByGalleryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByGalleryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByGalleryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByGalleryPager) {
		g.newListByGalleryPager.remove(req)
	}
	return resp, nil
}

func (g *GalleryApplicationsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := g.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<galleryApplicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryApplicationUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryApplicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryApplicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), resourceGroupNameParam, galleryNameParam, galleryApplicationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		g.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		g.beginUpdate.remove(req)
	}

	return resp, nil
}
