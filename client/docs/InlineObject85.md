# InlineObject85

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the custom performance class | 
**MaxLatency** | Pointer to **NullableInt32** | Maximum latency in milliseconds | [optional] 
**MaxJitter** | Pointer to **NullableInt32** | Maximum jitter in milliseconds | [optional] 
**MaxLossPercentage** | Pointer to **NullableInt32** | Maximum percentage of packet loss | [optional] 

## Methods

### NewInlineObject85

`func NewInlineObject85(name string, ) *InlineObject85`

NewInlineObject85 instantiates a new InlineObject85 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject85WithDefaults

`func NewInlineObject85WithDefaults() *InlineObject85`

NewInlineObject85WithDefaults instantiates a new InlineObject85 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject85) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject85) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject85) SetName(v string)`

SetName sets Name field to given value.


### GetMaxLatency

`func (o *InlineObject85) GetMaxLatency() int32`

GetMaxLatency returns the MaxLatency field if non-nil, zero value otherwise.

### GetMaxLatencyOk

`func (o *InlineObject85) GetMaxLatencyOk() (*int32, bool)`

GetMaxLatencyOk returns a tuple with the MaxLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLatency

`func (o *InlineObject85) SetMaxLatency(v int32)`

SetMaxLatency sets MaxLatency field to given value.

### HasMaxLatency

`func (o *InlineObject85) HasMaxLatency() bool`

HasMaxLatency returns a boolean if a field has been set.

### SetMaxLatencyNil

`func (o *InlineObject85) SetMaxLatencyNil(b bool)`

 SetMaxLatencyNil sets the value for MaxLatency to be an explicit nil

### UnsetMaxLatency
`func (o *InlineObject85) UnsetMaxLatency()`

UnsetMaxLatency ensures that no value is present for MaxLatency, not even an explicit nil
### GetMaxJitter

`func (o *InlineObject85) GetMaxJitter() int32`

GetMaxJitter returns the MaxJitter field if non-nil, zero value otherwise.

### GetMaxJitterOk

`func (o *InlineObject85) GetMaxJitterOk() (*int32, bool)`

GetMaxJitterOk returns a tuple with the MaxJitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJitter

`func (o *InlineObject85) SetMaxJitter(v int32)`

SetMaxJitter sets MaxJitter field to given value.

### HasMaxJitter

`func (o *InlineObject85) HasMaxJitter() bool`

HasMaxJitter returns a boolean if a field has been set.

### SetMaxJitterNil

`func (o *InlineObject85) SetMaxJitterNil(b bool)`

 SetMaxJitterNil sets the value for MaxJitter to be an explicit nil

### UnsetMaxJitter
`func (o *InlineObject85) UnsetMaxJitter()`

UnsetMaxJitter ensures that no value is present for MaxJitter, not even an explicit nil
### GetMaxLossPercentage

`func (o *InlineObject85) GetMaxLossPercentage() int32`

GetMaxLossPercentage returns the MaxLossPercentage field if non-nil, zero value otherwise.

### GetMaxLossPercentageOk

`func (o *InlineObject85) GetMaxLossPercentageOk() (*int32, bool)`

GetMaxLossPercentageOk returns a tuple with the MaxLossPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLossPercentage

`func (o *InlineObject85) SetMaxLossPercentage(v int32)`

SetMaxLossPercentage sets MaxLossPercentage field to given value.

### HasMaxLossPercentage

`func (o *InlineObject85) HasMaxLossPercentage() bool`

HasMaxLossPercentage returns a boolean if a field has been set.

### SetMaxLossPercentageNil

`func (o *InlineObject85) SetMaxLossPercentageNil(b bool)`

 SetMaxLossPercentageNil sets the value for MaxLossPercentage to be an explicit nil

### UnsetMaxLossPercentage
`func (o *InlineObject85) UnsetMaxLossPercentage()`

UnsetMaxLossPercentage ensures that no value is present for MaxLossPercentage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


