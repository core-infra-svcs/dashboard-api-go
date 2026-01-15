# InlineObject196

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on this SSID. | [optional] 
**Rules** | Pointer to [**[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules**](NetworksNetworkIdGroupPoliciesBonjourForwardingRules.md) | List of bonjour forwarding rules. | [optional] 
**Exception** | Pointer to [**InlineResponse200209Exception**](InlineResponse200209Exception.md) |  | [optional] 

## Methods

### NewInlineObject196

`func NewInlineObject196() *InlineObject196`

NewInlineObject196 instantiates a new InlineObject196 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject196WithDefaults

`func NewInlineObject196WithDefaults() *InlineObject196`

NewInlineObject196WithDefaults instantiates a new InlineObject196 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject196) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject196) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject196) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject196) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject196) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject196) GetRulesOk() (*[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject196) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject196) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetException

`func (o *InlineObject196) GetException() InlineResponse200209Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineObject196) GetExceptionOk() (*InlineResponse200209Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineObject196) SetException(v InlineResponse200209Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineObject196) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


