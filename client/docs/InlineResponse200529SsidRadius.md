# InlineResponse200529SsidRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]InlineResponse200529SsidRadiusServers**](InlineResponse200529SsidRadiusServers.md) | RADIUS servers for the SSID | [optional] 
**Accounting** | Pointer to [**InlineResponse200529SsidRadiusAccounting**](InlineResponse200529SsidRadiusAccounting.md) |  | [optional] 
**Policies** | Pointer to [**InlineResponse200529SsidRadiusPolicies**](InlineResponse200529SsidRadiusPolicies.md) |  | [optional] 
**Proxy** | Pointer to [**InlineResponse200529SsidRadiusProxy**](InlineResponse200529SsidRadiusProxy.md) |  | [optional] 
**Testing** | Pointer to [**InlineResponse200529SsidRadiusTesting**](InlineResponse200529SsidRadiusTesting.md) |  | [optional] 
**Identifiers** | Pointer to [**InlineResponse200529SsidRadiusIdentifiers**](InlineResponse200529SsidRadiusIdentifiers.md) |  | [optional] 
**Das** | Pointer to [**InlineResponse200529SsidRadiusDas**](InlineResponse200529SsidRadiusDas.md) |  | [optional] 
**EapTimers** | Pointer to [**InlineResponse200529SsidRadiusEapTimers**](InlineResponse200529SsidRadiusEapTimers.md) |  | [optional] 

## Methods

### NewInlineResponse200529SsidRadius

`func NewInlineResponse200529SsidRadius() *InlineResponse200529SsidRadius`

NewInlineResponse200529SsidRadius instantiates a new InlineResponse200529SsidRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200529SsidRadiusWithDefaults

`func NewInlineResponse200529SsidRadiusWithDefaults() *InlineResponse200529SsidRadius`

NewInlineResponse200529SsidRadiusWithDefaults instantiates a new InlineResponse200529SsidRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *InlineResponse200529SsidRadius) GetServers() []InlineResponse200529SsidRadiusServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200529SsidRadius) GetServersOk() (*[]InlineResponse200529SsidRadiusServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200529SsidRadius) SetServers(v []InlineResponse200529SsidRadiusServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *InlineResponse200529SsidRadius) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounting

`func (o *InlineResponse200529SsidRadius) GetAccounting() InlineResponse200529SsidRadiusAccounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *InlineResponse200529SsidRadius) GetAccountingOk() (*InlineResponse200529SsidRadiusAccounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *InlineResponse200529SsidRadius) SetAccounting(v InlineResponse200529SsidRadiusAccounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *InlineResponse200529SsidRadius) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200529SsidRadius) GetPolicies() InlineResponse200529SsidRadiusPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200529SsidRadius) GetPoliciesOk() (*InlineResponse200529SsidRadiusPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200529SsidRadius) SetPolicies(v InlineResponse200529SsidRadiusPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200529SsidRadius) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProxy

`func (o *InlineResponse200529SsidRadius) GetProxy() InlineResponse200529SsidRadiusProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *InlineResponse200529SsidRadius) GetProxyOk() (*InlineResponse200529SsidRadiusProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *InlineResponse200529SsidRadius) SetProxy(v InlineResponse200529SsidRadiusProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *InlineResponse200529SsidRadius) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTesting

`func (o *InlineResponse200529SsidRadius) GetTesting() InlineResponse200529SsidRadiusTesting`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *InlineResponse200529SsidRadius) GetTestingOk() (*InlineResponse200529SsidRadiusTesting, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *InlineResponse200529SsidRadius) SetTesting(v InlineResponse200529SsidRadiusTesting)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *InlineResponse200529SsidRadius) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetIdentifiers

`func (o *InlineResponse200529SsidRadius) GetIdentifiers() InlineResponse200529SsidRadiusIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *InlineResponse200529SsidRadius) GetIdentifiersOk() (*InlineResponse200529SsidRadiusIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *InlineResponse200529SsidRadius) SetIdentifiers(v InlineResponse200529SsidRadiusIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *InlineResponse200529SsidRadius) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDas

`func (o *InlineResponse200529SsidRadius) GetDas() InlineResponse200529SsidRadiusDas`

GetDas returns the Das field if non-nil, zero value otherwise.

### GetDasOk

`func (o *InlineResponse200529SsidRadius) GetDasOk() (*InlineResponse200529SsidRadiusDas, bool)`

GetDasOk returns a tuple with the Das field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDas

`func (o *InlineResponse200529SsidRadius) SetDas(v InlineResponse200529SsidRadiusDas)`

SetDas sets Das field to given value.

### HasDas

`func (o *InlineResponse200529SsidRadius) HasDas() bool`

HasDas returns a boolean if a field has been set.

### GetEapTimers

`func (o *InlineResponse200529SsidRadius) GetEapTimers() InlineResponse200529SsidRadiusEapTimers`

GetEapTimers returns the EapTimers field if non-nil, zero value otherwise.

### GetEapTimersOk

`func (o *InlineResponse200529SsidRadius) GetEapTimersOk() (*InlineResponse200529SsidRadiusEapTimers, bool)`

GetEapTimersOk returns a tuple with the EapTimers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapTimers

`func (o *InlineResponse200529SsidRadius) SetEapTimers(v InlineResponse200529SsidRadiusEapTimers)`

SetEapTimers sets EapTimers field to given value.

### HasEapTimers

`func (o *InlineResponse200529SsidRadius) HasEapTimers() bool`

HasEapTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


