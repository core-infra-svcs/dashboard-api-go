# InlineObject143

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BroadcastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**MulticastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**UnknownUnicastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 

## Methods

### NewInlineObject143

`func NewInlineObject143() *InlineObject143`

NewInlineObject143 instantiates a new InlineObject143 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject143WithDefaults

`func NewInlineObject143WithDefaults() *InlineObject143`

NewInlineObject143WithDefaults instantiates a new InlineObject143 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcastThreshold

`func (o *InlineObject143) GetBroadcastThreshold() int32`

GetBroadcastThreshold returns the BroadcastThreshold field if non-nil, zero value otherwise.

### GetBroadcastThresholdOk

`func (o *InlineObject143) GetBroadcastThresholdOk() (*int32, bool)`

GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastThreshold

`func (o *InlineObject143) SetBroadcastThreshold(v int32)`

SetBroadcastThreshold sets BroadcastThreshold field to given value.

### HasBroadcastThreshold

`func (o *InlineObject143) HasBroadcastThreshold() bool`

HasBroadcastThreshold returns a boolean if a field has been set.

### GetMulticastThreshold

`func (o *InlineObject143) GetMulticastThreshold() int32`

GetMulticastThreshold returns the MulticastThreshold field if non-nil, zero value otherwise.

### GetMulticastThresholdOk

`func (o *InlineObject143) GetMulticastThresholdOk() (*int32, bool)`

GetMulticastThresholdOk returns a tuple with the MulticastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastThreshold

`func (o *InlineObject143) SetMulticastThreshold(v int32)`

SetMulticastThreshold sets MulticastThreshold field to given value.

### HasMulticastThreshold

`func (o *InlineObject143) HasMulticastThreshold() bool`

HasMulticastThreshold returns a boolean if a field has been set.

### GetUnknownUnicastThreshold

`func (o *InlineObject143) GetUnknownUnicastThreshold() int32`

GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field if non-nil, zero value otherwise.

### GetUnknownUnicastThresholdOk

`func (o *InlineObject143) GetUnknownUnicastThresholdOk() (*int32, bool)`

GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicastThreshold

`func (o *InlineObject143) SetUnknownUnicastThreshold(v int32)`

SetUnknownUnicastThreshold sets UnknownUnicastThreshold field to given value.

### HasUnknownUnicastThreshold

`func (o *InlineObject143) HasUnknownUnicastThreshold() bool`

HasUnknownUnicastThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


