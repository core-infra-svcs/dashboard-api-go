# InlineObject203

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on this SSID. | [optional] 
**Rules** | Pointer to [**[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules**](NetworksNetworkIdGroupPoliciesBonjourForwardingRules.md) | List of bonjour forwarding rules. | [optional] 
**Exception** | Pointer to [**InlineResponse200236Exception**](InlineResponse200236Exception.md) |  | [optional] 

## Methods

### NewInlineObject203

`func NewInlineObject203() *InlineObject203`

NewInlineObject203 instantiates a new InlineObject203 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject203WithDefaults

`func NewInlineObject203WithDefaults() *InlineObject203`

NewInlineObject203WithDefaults instantiates a new InlineObject203 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject203) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject203) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject203) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject203) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRules

`func (o *InlineObject203) GetRules() []NetworksNetworkIdGroupPoliciesBonjourForwardingRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject203) GetRulesOk() (*[]NetworksNetworkIdGroupPoliciesBonjourForwardingRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject203) SetRules(v []NetworksNetworkIdGroupPoliciesBonjourForwardingRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject203) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetException

`func (o *InlineObject203) GetException() InlineResponse200236Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineObject203) GetExceptionOk() (*InlineResponse200236Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineObject203) SetException(v InlineResponse200236Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineObject203) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


