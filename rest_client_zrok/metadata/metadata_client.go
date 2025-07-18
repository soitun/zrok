// Code generated by go-swagger; DO NOT EDIT.

package metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metadata API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metadata API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ClientVersionCheck(params *ClientVersionCheckParams, opts ...ClientOption) (*ClientVersionCheckOK, error)

	Configuration(params *ConfigurationParams, opts ...ClientOption) (*ConfigurationOK, error)

	GetAccountDetail(params *GetAccountDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountDetailOK, error)

	GetAccountMetrics(params *GetAccountMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountMetricsOK, error)

	GetEnvironmentDetail(params *GetEnvironmentDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentDetailOK, error)

	GetEnvironmentMetrics(params *GetEnvironmentMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentMetricsOK, error)

	GetFrontendDetail(params *GetFrontendDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFrontendDetailOK, error)

	GetShareDetail(params *GetShareDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetShareDetailOK, error)

	GetShareMetrics(params *GetShareMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetShareMetricsOK, error)

	GetSparklines(params *GetSparklinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSparklinesOK, error)

	ListMemberships(params *ListMembershipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListMembershipsOK, error)

	ListOrgMembers(params *ListOrgMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrgMembersOK, error)

	ListPublicFrontendsForAccount(params *ListPublicFrontendsForAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPublicFrontendsForAccountOK, error)

	OrgAccountOverview(params *OrgAccountOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrgAccountOverviewOK, error)

	Overview(params *OverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OverviewOK, error)

	Version(params *VersionParams, opts ...ClientOption) (*VersionOK, error)

	VersionInventory(params *VersionInventoryParams, opts ...ClientOption) (*VersionInventoryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ClientVersionCheck client version check API
*/
func (a *Client) ClientVersionCheck(params *ClientVersionCheckParams, opts ...ClientOption) (*ClientVersionCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClientVersionCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "clientVersionCheck",
		Method:             "POST",
		PathPattern:        "/clientVersionCheck",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClientVersionCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClientVersionCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clientVersionCheck: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Configuration configuration API
*/
func (a *Client) Configuration(params *ConfigurationParams, opts ...ClientOption) (*ConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigurationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "configuration",
		Method:             "GET",
		PathPattern:        "/configuration",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConfigurationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for configuration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccountDetail get account detail API
*/
func (a *Client) GetAccountDetail(params *GetAccountDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccountDetail",
		Method:             "GET",
		PathPattern:        "/detail/account",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccountDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccountDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAccountMetrics get account metrics API
*/
func (a *Client) GetAccountMetrics(params *GetAccountMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAccountMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAccountMetrics",
		Method:             "GET",
		PathPattern:        "/metrics/account",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAccountMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAccountMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAccountMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironmentDetail get environment detail API
*/
func (a *Client) GetEnvironmentDetail(params *GetEnvironmentDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironmentDetail",
		Method:             "GET",
		PathPattern:        "/detail/environment/{envZId}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEnvironmentDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetEnvironmentMetrics get environment metrics API
*/
func (a *Client) GetEnvironmentMetrics(params *GetEnvironmentMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetEnvironmentMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEnvironmentMetrics",
		Method:             "GET",
		PathPattern:        "/metrics/environment/{envId}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetEnvironmentMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetEnvironmentMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEnvironmentMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetFrontendDetail get frontend detail API
*/
func (a *Client) GetFrontendDetail(params *GetFrontendDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFrontendDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFrontendDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFrontendDetail",
		Method:             "GET",
		PathPattern:        "/detail/frontend/{frontendId}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFrontendDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFrontendDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getFrontendDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetShareDetail get share detail API
*/
func (a *Client) GetShareDetail(params *GetShareDetailParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetShareDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShareDetailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getShareDetail",
		Method:             "GET",
		PathPattern:        "/detail/share/{shareToken}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetShareDetailReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetShareDetailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getShareDetail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetShareMetrics get share metrics API
*/
func (a *Client) GetShareMetrics(params *GetShareMetricsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetShareMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShareMetricsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getShareMetrics",
		Method:             "GET",
		PathPattern:        "/metrics/share/{shareToken}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetShareMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetShareMetricsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getShareMetrics: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetSparklines get sparklines API
*/
func (a *Client) GetSparklines(params *GetSparklinesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSparklinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSparklinesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSparklines",
		Method:             "POST",
		PathPattern:        "/sparklines",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSparklinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSparklinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getSparklines: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListMemberships list memberships API
*/
func (a *Client) ListMemberships(params *ListMembershipsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListMembershipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListMembershipsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listMemberships",
		Method:             "GET",
		PathPattern:        "/memberships",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListMembershipsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListMembershipsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listMemberships: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListOrgMembers list org members API
*/
func (a *Client) ListOrgMembers(params *ListOrgMembersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListOrgMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListOrgMembersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listOrgMembers",
		Method:             "GET",
		PathPattern:        "/members/{organizationToken}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListOrgMembersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListOrgMembersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listOrgMembers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListPublicFrontendsForAccount list public frontends for account API
*/
func (a *Client) ListPublicFrontendsForAccount(params *ListPublicFrontendsForAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListPublicFrontendsForAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPublicFrontendsForAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listPublicFrontendsForAccount",
		Method:             "GET",
		PathPattern:        "/overview/public-frontends",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPublicFrontendsForAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPublicFrontendsForAccountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPublicFrontendsForAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
OrgAccountOverview org account overview API
*/
func (a *Client) OrgAccountOverview(params *OrgAccountOverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OrgAccountOverviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOrgAccountOverviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "orgAccountOverview",
		Method:             "GET",
		PathPattern:        "/overview/{organizationToken}/{accountEmail}",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OrgAccountOverviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OrgAccountOverviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for orgAccountOverview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Overview overview API
*/
func (a *Client) Overview(params *OverviewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*OverviewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOverviewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "overview",
		Method:             "GET",
		PathPattern:        "/overview",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OverviewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OverviewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for overview: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Version version API
*/
func (a *Client) Version(params *VersionParams, opts ...ClientOption) (*VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "version",
		Method:             "GET",
		PathPattern:        "/version",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for version: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VersionInventory version inventory API
*/
func (a *Client) VersionInventory(params *VersionInventoryParams, opts ...ClientOption) (*VersionInventoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionInventoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "versionInventory",
		Method:             "GET",
		PathPattern:        "/versions",
		ProducesMediaTypes: []string{"application/zrok.v1+json"},
		ConsumesMediaTypes: []string{"application/zrok.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &VersionInventoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VersionInventoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for versionInventory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
