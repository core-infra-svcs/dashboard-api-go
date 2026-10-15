# OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusServers**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusServers.md) | Organization RADIUS servers, ordered by priority | [optional] 
**Accounting** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusAccounting**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusAccounting.md) |  | [optional] 
**Policies** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusPolicies**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusPolicies.md) |  | [optional] 
**Proxy** | Pointer to [**InlineResponse200517SsidRadiusProxy**](InlineResponse200517SsidRadiusProxy.md) |  | [optional] 
**Testing** | Pointer to [**InlineResponse200517SsidRadiusTesting**](InlineResponse200517SsidRadiusTesting.md) |  | [optional] 
**Identifiers** | Pointer to [**InlineResponse200517SsidRadiusIdentifiers**](InlineResponse200517SsidRadiusIdentifiers.md) |  | [optional] 
**Das** | Pointer to [**InlineResponse200517SsidRadiusDas**](InlineResponse200517SsidRadiusDas.md) |  | [optional] 
**EapTimers** | Pointer to [**OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusEapTimers**](OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusEapTimers.md) |  | [optional] 

## Methods

### NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius

`func NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius() *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius`

NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius instantiates a new OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusWithDefaults

`func NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusWithDefaults() *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius`

NewOrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusWithDefaults instantiates a new OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetServers() []OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetServersOk() (*[]OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetServers(v []OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounting

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetAccounting() OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusAccounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetAccountingOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusAccounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetAccounting(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusAccounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetPolicies

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetPolicies() OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetPoliciesOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetPolicies(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProxy

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetProxy() InlineResponse200517SsidRadiusProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetProxyOk() (*InlineResponse200517SsidRadiusProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetProxy(v InlineResponse200517SsidRadiusProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTesting

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetTesting() InlineResponse200517SsidRadiusTesting`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetTestingOk() (*InlineResponse200517SsidRadiusTesting, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetTesting(v InlineResponse200517SsidRadiusTesting)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetIdentifiers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetIdentifiers() InlineResponse200517SsidRadiusIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetIdentifiersOk() (*InlineResponse200517SsidRadiusIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetIdentifiers(v InlineResponse200517SsidRadiusIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDas

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetDas() InlineResponse200517SsidRadiusDas`

GetDas returns the Das field if non-nil, zero value otherwise.

### GetDasOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetDasOk() (*InlineResponse200517SsidRadiusDas, bool)`

GetDasOk returns a tuple with the Das field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDas

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetDas(v InlineResponse200517SsidRadiusDas)`

SetDas sets Das field to given value.

### HasDas

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasDas() bool`

HasDas returns a boolean if a field has been set.

### GetEapTimers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetEapTimers() OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusEapTimers`

GetEapTimers returns the EapTimers field if non-nil, zero value otherwise.

### GetEapTimersOk

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) GetEapTimersOk() (*OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusEapTimers, bool)`

GetEapTimersOk returns a tuple with the EapTimers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapTimers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) SetEapTimers(v OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadiusEapTimers)`

SetEapTimers sets EapTimers field to given value.

### HasEapTimers

`func (o *OrganizationsOrganizationIdWirelessSsidsProfilesSsidRadius) HasEapTimers() bool`

HasEapTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


