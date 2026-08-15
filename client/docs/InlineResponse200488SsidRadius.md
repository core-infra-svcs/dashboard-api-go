# InlineResponse200488SsidRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]InlineResponse200488SsidRadiusServers**](InlineResponse200488SsidRadiusServers.md) | RADIUS servers for the SSID | [optional] 
**Accounting** | Pointer to [**InlineResponse200488SsidRadiusAccounting**](InlineResponse200488SsidRadiusAccounting.md) |  | [optional] 
**Policies** | Pointer to [**InlineResponse200488SsidRadiusPolicies**](InlineResponse200488SsidRadiusPolicies.md) |  | [optional] 
**Proxy** | Pointer to [**InlineResponse200488SsidRadiusProxy**](InlineResponse200488SsidRadiusProxy.md) |  | [optional] 
**Testing** | Pointer to [**InlineResponse200488SsidRadiusTesting**](InlineResponse200488SsidRadiusTesting.md) |  | [optional] 
**Identifiers** | Pointer to [**InlineResponse200488SsidRadiusIdentifiers**](InlineResponse200488SsidRadiusIdentifiers.md) |  | [optional] 
**Das** | Pointer to [**InlineResponse200488SsidRadiusDas**](InlineResponse200488SsidRadiusDas.md) |  | [optional] 
**EapTimers** | Pointer to [**InlineResponse200488SsidRadiusEapTimers**](InlineResponse200488SsidRadiusEapTimers.md) |  | [optional] 

## Methods

### NewInlineResponse200488SsidRadius

`func NewInlineResponse200488SsidRadius() *InlineResponse200488SsidRadius`

NewInlineResponse200488SsidRadius instantiates a new InlineResponse200488SsidRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200488SsidRadiusWithDefaults

`func NewInlineResponse200488SsidRadiusWithDefaults() *InlineResponse200488SsidRadius`

NewInlineResponse200488SsidRadiusWithDefaults instantiates a new InlineResponse200488SsidRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *InlineResponse200488SsidRadius) GetServers() []InlineResponse200488SsidRadiusServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200488SsidRadius) GetServersOk() (*[]InlineResponse200488SsidRadiusServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200488SsidRadius) SetServers(v []InlineResponse200488SsidRadiusServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *InlineResponse200488SsidRadius) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounting

`func (o *InlineResponse200488SsidRadius) GetAccounting() InlineResponse200488SsidRadiusAccounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *InlineResponse200488SsidRadius) GetAccountingOk() (*InlineResponse200488SsidRadiusAccounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *InlineResponse200488SsidRadius) SetAccounting(v InlineResponse200488SsidRadiusAccounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *InlineResponse200488SsidRadius) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200488SsidRadius) GetPolicies() InlineResponse200488SsidRadiusPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200488SsidRadius) GetPoliciesOk() (*InlineResponse200488SsidRadiusPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200488SsidRadius) SetPolicies(v InlineResponse200488SsidRadiusPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200488SsidRadius) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProxy

`func (o *InlineResponse200488SsidRadius) GetProxy() InlineResponse200488SsidRadiusProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *InlineResponse200488SsidRadius) GetProxyOk() (*InlineResponse200488SsidRadiusProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *InlineResponse200488SsidRadius) SetProxy(v InlineResponse200488SsidRadiusProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *InlineResponse200488SsidRadius) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTesting

`func (o *InlineResponse200488SsidRadius) GetTesting() InlineResponse200488SsidRadiusTesting`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *InlineResponse200488SsidRadius) GetTestingOk() (*InlineResponse200488SsidRadiusTesting, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *InlineResponse200488SsidRadius) SetTesting(v InlineResponse200488SsidRadiusTesting)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *InlineResponse200488SsidRadius) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetIdentifiers

`func (o *InlineResponse200488SsidRadius) GetIdentifiers() InlineResponse200488SsidRadiusIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *InlineResponse200488SsidRadius) GetIdentifiersOk() (*InlineResponse200488SsidRadiusIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *InlineResponse200488SsidRadius) SetIdentifiers(v InlineResponse200488SsidRadiusIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *InlineResponse200488SsidRadius) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDas

`func (o *InlineResponse200488SsidRadius) GetDas() InlineResponse200488SsidRadiusDas`

GetDas returns the Das field if non-nil, zero value otherwise.

### GetDasOk

`func (o *InlineResponse200488SsidRadius) GetDasOk() (*InlineResponse200488SsidRadiusDas, bool)`

GetDasOk returns a tuple with the Das field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDas

`func (o *InlineResponse200488SsidRadius) SetDas(v InlineResponse200488SsidRadiusDas)`

SetDas sets Das field to given value.

### HasDas

`func (o *InlineResponse200488SsidRadius) HasDas() bool`

HasDas returns a boolean if a field has been set.

### GetEapTimers

`func (o *InlineResponse200488SsidRadius) GetEapTimers() InlineResponse200488SsidRadiusEapTimers`

GetEapTimers returns the EapTimers field if non-nil, zero value otherwise.

### GetEapTimersOk

`func (o *InlineResponse200488SsidRadius) GetEapTimersOk() (*InlineResponse200488SsidRadiusEapTimers, bool)`

GetEapTimersOk returns a tuple with the EapTimers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapTimers

`func (o *InlineResponse200488SsidRadius) SetEapTimers(v InlineResponse200488SsidRadiusEapTimers)`

SetEapTimers sets EapTimers field to given value.

### HasEapTimers

`func (o *InlineResponse200488SsidRadius) HasEapTimers() bool`

HasEapTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


