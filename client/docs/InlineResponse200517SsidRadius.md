# InlineResponse200517SsidRadius

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]InlineResponse200517SsidRadiusServers**](InlineResponse200517SsidRadiusServers.md) | RADIUS servers for the SSID | [optional] 
**Accounting** | Pointer to [**InlineResponse200517SsidRadiusAccounting**](InlineResponse200517SsidRadiusAccounting.md) |  | [optional] 
**Policies** | Pointer to [**InlineResponse200517SsidRadiusPolicies**](InlineResponse200517SsidRadiusPolicies.md) |  | [optional] 
**Proxy** | Pointer to [**InlineResponse200517SsidRadiusProxy**](InlineResponse200517SsidRadiusProxy.md) |  | [optional] 
**Testing** | Pointer to [**InlineResponse200517SsidRadiusTesting**](InlineResponse200517SsidRadiusTesting.md) |  | [optional] 
**Identifiers** | Pointer to [**InlineResponse200517SsidRadiusIdentifiers**](InlineResponse200517SsidRadiusIdentifiers.md) |  | [optional] 
**Das** | Pointer to [**InlineResponse200517SsidRadiusDas**](InlineResponse200517SsidRadiusDas.md) |  | [optional] 
**EapTimers** | Pointer to [**InlineResponse200517SsidRadiusEapTimers**](InlineResponse200517SsidRadiusEapTimers.md) |  | [optional] 

## Methods

### NewInlineResponse200517SsidRadius

`func NewInlineResponse200517SsidRadius() *InlineResponse200517SsidRadius`

NewInlineResponse200517SsidRadius instantiates a new InlineResponse200517SsidRadius object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200517SsidRadiusWithDefaults

`func NewInlineResponse200517SsidRadiusWithDefaults() *InlineResponse200517SsidRadius`

NewInlineResponse200517SsidRadiusWithDefaults instantiates a new InlineResponse200517SsidRadius object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *InlineResponse200517SsidRadius) GetServers() []InlineResponse200517SsidRadiusServers`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *InlineResponse200517SsidRadius) GetServersOk() (*[]InlineResponse200517SsidRadiusServers, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *InlineResponse200517SsidRadius) SetServers(v []InlineResponse200517SsidRadiusServers)`

SetServers sets Servers field to given value.

### HasServers

`func (o *InlineResponse200517SsidRadius) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetAccounting

`func (o *InlineResponse200517SsidRadius) GetAccounting() InlineResponse200517SsidRadiusAccounting`

GetAccounting returns the Accounting field if non-nil, zero value otherwise.

### GetAccountingOk

`func (o *InlineResponse200517SsidRadius) GetAccountingOk() (*InlineResponse200517SsidRadiusAccounting, bool)`

GetAccountingOk returns a tuple with the Accounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounting

`func (o *InlineResponse200517SsidRadius) SetAccounting(v InlineResponse200517SsidRadiusAccounting)`

SetAccounting sets Accounting field to given value.

### HasAccounting

`func (o *InlineResponse200517SsidRadius) HasAccounting() bool`

HasAccounting returns a boolean if a field has been set.

### GetPolicies

`func (o *InlineResponse200517SsidRadius) GetPolicies() InlineResponse200517SsidRadiusPolicies`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *InlineResponse200517SsidRadius) GetPoliciesOk() (*InlineResponse200517SsidRadiusPolicies, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *InlineResponse200517SsidRadius) SetPolicies(v InlineResponse200517SsidRadiusPolicies)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *InlineResponse200517SsidRadius) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetProxy

`func (o *InlineResponse200517SsidRadius) GetProxy() InlineResponse200517SsidRadiusProxy`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *InlineResponse200517SsidRadius) GetProxyOk() (*InlineResponse200517SsidRadiusProxy, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *InlineResponse200517SsidRadius) SetProxy(v InlineResponse200517SsidRadiusProxy)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *InlineResponse200517SsidRadius) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetTesting

`func (o *InlineResponse200517SsidRadius) GetTesting() InlineResponse200517SsidRadiusTesting`

GetTesting returns the Testing field if non-nil, zero value otherwise.

### GetTestingOk

`func (o *InlineResponse200517SsidRadius) GetTestingOk() (*InlineResponse200517SsidRadiusTesting, bool)`

GetTestingOk returns a tuple with the Testing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTesting

`func (o *InlineResponse200517SsidRadius) SetTesting(v InlineResponse200517SsidRadiusTesting)`

SetTesting sets Testing field to given value.

### HasTesting

`func (o *InlineResponse200517SsidRadius) HasTesting() bool`

HasTesting returns a boolean if a field has been set.

### GetIdentifiers

`func (o *InlineResponse200517SsidRadius) GetIdentifiers() InlineResponse200517SsidRadiusIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *InlineResponse200517SsidRadius) GetIdentifiersOk() (*InlineResponse200517SsidRadiusIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *InlineResponse200517SsidRadius) SetIdentifiers(v InlineResponse200517SsidRadiusIdentifiers)`

SetIdentifiers sets Identifiers field to given value.

### HasIdentifiers

`func (o *InlineResponse200517SsidRadius) HasIdentifiers() bool`

HasIdentifiers returns a boolean if a field has been set.

### GetDas

`func (o *InlineResponse200517SsidRadius) GetDas() InlineResponse200517SsidRadiusDas`

GetDas returns the Das field if non-nil, zero value otherwise.

### GetDasOk

`func (o *InlineResponse200517SsidRadius) GetDasOk() (*InlineResponse200517SsidRadiusDas, bool)`

GetDasOk returns a tuple with the Das field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDas

`func (o *InlineResponse200517SsidRadius) SetDas(v InlineResponse200517SsidRadiusDas)`

SetDas sets Das field to given value.

### HasDas

`func (o *InlineResponse200517SsidRadius) HasDas() bool`

HasDas returns a boolean if a field has been set.

### GetEapTimers

`func (o *InlineResponse200517SsidRadius) GetEapTimers() InlineResponse200517SsidRadiusEapTimers`

GetEapTimers returns the EapTimers field if non-nil, zero value otherwise.

### GetEapTimersOk

`func (o *InlineResponse200517SsidRadius) GetEapTimersOk() (*InlineResponse200517SsidRadiusEapTimers, bool)`

GetEapTimersOk returns a tuple with the EapTimers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapTimers

`func (o *InlineResponse200517SsidRadius) SetEapTimers(v InlineResponse200517SsidRadiusEapTimers)`

SetEapTimers sets EapTimers field to given value.

### HasEapTimers

`func (o *InlineResponse200517SsidRadius) HasEapTimers() bool`

HasEapTimers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


