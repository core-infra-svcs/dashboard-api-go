# InlineObject216

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on this SSID. | [optional] 
**Rules** | Pointer to [**[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules**](NetworksNetworkIdGroupPoliciesBonjourForwardingRules.md) | List of bonjour forwarding rules. | [optional] 
**Exception** | Pointer to [**InlineResponse200246Exception**](InlineResponse200246Exception.md) |  | [optional] 

## Methods

### NewInlineObject216

`func NewInlineObject216() *InlineObject216`

NewInlineObject216 instantiates a new InlineObject216 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject216WithDefaults

`func NewInlineObject216WithDefaults() *InlineObject216`

NewInlineObject216WithDefaults instantiates a new InlineObject216 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject216) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject216) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject216) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject216) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject216) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject216) GetRulesOk() (*[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject216) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject216) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetException

`func (o *InlineObject216) GetException() InlineResponse200246Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineObject216) GetExceptionOk() (*InlineResponse200246Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineObject216) SetException(v InlineResponse200246Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineObject216) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


