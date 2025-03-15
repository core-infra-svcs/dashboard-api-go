# InlineObject163

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BroadcastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**MulticastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**UnknownUnicastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**TreatTheseTrafficTypesAsOneThreshold** | Pointer to **[]string** | Grouped traffic types | [optional] 

## Methods

### NewInlineObject163

`func NewInlineObject163() *InlineObject163`

NewInlineObject163 instantiates a new InlineObject163 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject163WithDefaults

`func NewInlineObject163WithDefaults() *InlineObject163`

NewInlineObject163WithDefaults instantiates a new InlineObject163 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcastThreshold

`func (o *InlineObject163) GetBroadcastThreshold() int32`

GetBroadcastThreshold returns the BroadcastThreshold field if non-nil, zero value otherwise.

### GetBroadcastThresholdOk

`func (o *InlineObject163) GetBroadcastThresholdOk() (*int32, bool)`

GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastThreshold

`func (o *InlineObject163) SetBroadcastThreshold(v int32)`

SetBroadcastThreshold sets BroadcastThreshold field to given value.

### HasBroadcastThreshold

`func (o *InlineObject163) HasBroadcastThreshold() bool`

HasBroadcastThreshold returns a boolean if a field has been set.

### GetMulticastThreshold

`func (o *InlineObject163) GetMulticastThreshold() int32`

GetMulticastThreshold returns the MulticastThreshold field if non-nil, zero value otherwise.

### GetMulticastThresholdOk

`func (o *InlineObject163) GetMulticastThresholdOk() (*int32, bool)`

GetMulticastThresholdOk returns a tuple with the MulticastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastThreshold

`func (o *InlineObject163) SetMulticastThreshold(v int32)`

SetMulticastThreshold sets MulticastThreshold field to given value.

### HasMulticastThreshold

`func (o *InlineObject163) HasMulticastThreshold() bool`

HasMulticastThreshold returns a boolean if a field has been set.

### GetUnknownUnicastThreshold

`func (o *InlineObject163) GetUnknownUnicastThreshold() int32`

GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field if non-nil, zero value otherwise.

### GetUnknownUnicastThresholdOk

`func (o *InlineObject163) GetUnknownUnicastThresholdOk() (*int32, bool)`

GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicastThreshold

`func (o *InlineObject163) SetUnknownUnicastThreshold(v int32)`

SetUnknownUnicastThreshold sets UnknownUnicastThreshold field to given value.

### HasUnknownUnicastThreshold

`func (o *InlineObject163) HasUnknownUnicastThreshold() bool`

HasUnknownUnicastThreshold returns a boolean if a field has been set.

### GetTreatTheseTrafficTypesAsOneThreshold

`func (o *InlineObject163) GetTreatTheseTrafficTypesAsOneThreshold() []string`

GetTreatTheseTrafficTypesAsOneThreshold returns the TreatTheseTrafficTypesAsOneThreshold field if non-nil, zero value otherwise.

### GetTreatTheseTrafficTypesAsOneThresholdOk

`func (o *InlineObject163) GetTreatTheseTrafficTypesAsOneThresholdOk() (*[]string, bool)`

GetTreatTheseTrafficTypesAsOneThresholdOk returns a tuple with the TreatTheseTrafficTypesAsOneThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreatTheseTrafficTypesAsOneThreshold

`func (o *InlineObject163) SetTreatTheseTrafficTypesAsOneThreshold(v []string)`

SetTreatTheseTrafficTypesAsOneThreshold sets TreatTheseTrafficTypesAsOneThreshold field to given value.

### HasTreatTheseTrafficTypesAsOneThreshold

`func (o *InlineObject163) HasTreatTheseTrafficTypesAsOneThreshold() bool`

HasTreatTheseTrafficTypesAsOneThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


