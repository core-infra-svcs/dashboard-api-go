# InlineResponse200493SsidRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]InlineResponse200493SsidRadiusServers**](InlineResponse200493SsidRadiusServers.md) | RADIUS servers for the SSID | [optional] 
**Accounting** | Pointer to [**InlineResponse200493SsidRadiusAccounting**](InlineResponse200493SsidRadiusAccounting.md) |  | [optional] 
**Policies** | Pointer to [**InlineResponse200493SsidRadiusPolicies**](InlineResponse200493SsidRadiusPolicies.md) |  | [optional] 
**Proxy** | Pointer to [**InlineResponse200493SsidRadiusProxy**](InlineResponse200493SsidRadiusProxy.md) |  | [optional] 
**Testing** | Pointer to [**InlineResponse200493SsidRadiusTesting**](InlineResponse200493SsidRadiusTesting.md) |  | [optional] 
**Identifiers** | Pointer to [**InlineResponse200493SsidRadiusIdentifiers**](InlineResponse200493SsidRadiusIdentifiers.md) |  | [optional] 
**Das** | Pointer to [**InlineResponse200493SsidRadiusDas**](InlineResponse200493SsidRadiusDas.md) |  | [optional] 
**EapTimers** | Pointer to [**InlineResponse200493SsidRadiusEapTimers**](InlineResponse200493SsidRadiusEapTimers.md) |  | [optional] 

## Methods

### NewInlineResponse200493SsidRadius

`func NewInlineResponse200493SsidRadius() *InlineResponse200493SsidRadius`

NewInlineResponse200493SsidRadius instantiates a new InlineResponse200493SsidRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200493SsidRadiusWithDefaults

`func NewInlineResponse200493SsidRadiusWithDefaults() *InlineResponse200493SsidRadius`

NewInlineResponse200493SsidRadiusWithDefaults instantiates a new InlineResponse200493SsidRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *InlineResponse200493SsidRadius) GetServers() []InlineResponse200493SsidRadiusServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200493SsidRadius) GetServersOk() (*[]InlineResponse200493SsidRadiusServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200493SsidRadius) SetServers(v []InlineResponse200493SsidRadiusServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *InlineResponse200493SsidRadius) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounting

`func (o *InlineResponse200493SsidRadius) GetAccounting() InlineResponse200493SsidRadiusAccounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *InlineResponse200493SsidRadius) GetAccountingOk() (*InlineResponse200493SsidRadiusAccounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *InlineResponse200493SsidRadius) SetAccounting(v InlineResponse200493SsidRadiusAccounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *InlineResponse200493SsidRadius) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200493SsidRadius) GetPolicies() InlineResponse200493SsidRadiusPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200493SsidRadius) GetPoliciesOk() (*InlineResponse200493SsidRadiusPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200493SsidRadius) SetPolicies(v InlineResponse200493SsidRadiusPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200493SsidRadius) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProxy

`func (o *InlineResponse200493SsidRadius) GetProxy() InlineResponse200493SsidRadiusProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *InlineResponse200493SsidRadius) GetProxyOk() (*InlineResponse200493SsidRadiusProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *InlineResponse200493SsidRadius) SetProxy(v InlineResponse200493SsidRadiusProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *InlineResponse200493SsidRadius) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTesting

`func (o *InlineResponse200493SsidRadius) GetTesting() InlineResponse200493SsidRadiusTesting`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *InlineResponse200493SsidRadius) GetTestingOk() (*InlineResponse200493SsidRadiusTesting, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *InlineResponse200493SsidRadius) SetTesting(v InlineResponse200493SsidRadiusTesting)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *InlineResponse200493SsidRadius) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetIdentifiers

`func (o *InlineResponse200493SsidRadius) GetIdentifiers() InlineResponse200493SsidRadiusIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *InlineResponse200493SsidRadius) GetIdentifiersOk() (*InlineResponse200493SsidRadiusIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *InlineResponse200493SsidRadius) SetIdentifiers(v InlineResponse200493SsidRadiusIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *InlineResponse200493SsidRadius) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDas

`func (o *InlineResponse200493SsidRadius) GetDas() InlineResponse200493SsidRadiusDas`

GetDas returns the Das field if non-nil, zero value otherwise.

### GetDasOk

`func (o *InlineResponse200493SsidRadius) GetDasOk() (*InlineResponse200493SsidRadiusDas, bool)`

GetDasOk returns a tuple with the Das field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDas

`func (o *InlineResponse200493SsidRadius) SetDas(v InlineResponse200493SsidRadiusDas)`

SetDas sets Das field to given value.

### HasDas

`func (o *InlineResponse200493SsidRadius) HasDas() bool`

HasDas returns a boolean if a field has been set.

### GetEapTimers

`func (o *InlineResponse200493SsidRadius) GetEapTimers() InlineResponse200493SsidRadiusEapTimers`

GetEapTimers returns the EapTimers field if non-nil, zero value otherwise.

### GetEapTimersOk

`func (o *InlineResponse200493SsidRadius) GetEapTimersOk() (*InlineResponse200493SsidRadiusEapTimers, bool)`

GetEapTimersOk returns a tuple with the EapTimers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapTimers

`func (o *InlineResponse200493SsidRadius) SetEapTimers(v InlineResponse200493SsidRadiusEapTimers)`

SetEapTimers sets EapTimers field to given value.

### HasEapTimers

`func (o *InlineResponse200493SsidRadius) HasEapTimers() bool`

HasEapTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


