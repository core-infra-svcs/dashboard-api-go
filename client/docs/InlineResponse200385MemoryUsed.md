# InlineResponse200385MemoryUsed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Minimum** | Pointer to **int32** | Minimum memory in kB used on the device over the interval | [optional] 
**Maximum** | Pointer to **int32** | Maximum memory in kB used on the device over the interval | [optional] 
**Median** | Pointer to **int32** | Median memory in kB used on the device over the interval rounded up to nearest integer | [optional] 
**Percentages** | Pointer to [**InlineResponse200385MemoryUsedPercentages**](InlineResponse200385MemoryUsedPercentages.md) |  | [optional] 

## Methods

### NewInlineResponse200385MemoryUsed

`func NewInlineResponse200385MemoryUsed() *InlineResponse200385MemoryUsed`

NewInlineResponse200385MemoryUsed instantiates a new InlineResponse200385MemoryUsed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200385MemoryUsedWithDefaults

`func NewInlineResponse200385MemoryUsedWithDefaults() *InlineResponse200385MemoryUsed`

NewInlineResponse200385MemoryUsedWithDefaults instantiates a new InlineResponse200385MemoryUsed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinimum

`func (o *InlineResponse200385MemoryUsed) GetMinimum() int32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *InlineResponse200385MemoryUsed) GetMinimumOk() (*int32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *InlineResponse200385MemoryUsed) SetMinimum(v int32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *InlineResponse200385MemoryUsed) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetMaximum

`func (o *InlineResponse200385MemoryUsed) GetMaximum() int32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *InlineResponse200385MemoryUsed) GetMaximumOk() (*int32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *InlineResponse200385MemoryUsed) SetMaximum(v int32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *InlineResponse200385MemoryUsed) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMedian

`func (o *InlineResponse200385MemoryUsed) GetMedian() int32`

GetMedian returns the Median field if non-nil, zero value otherwise.

### GetMedianOk

`func (o *InlineResponse200385MemoryUsed) GetMedianOk() (*int32, bool)`

GetMedianOk returns a tuple with the Median field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian

`func (o *InlineResponse200385MemoryUsed) SetMedian(v int32)`

SetMedian sets Median field to given value.

### HasMedian

`func (o *InlineResponse200385MemoryUsed) HasMedian() bool`

HasMedian returns a boolean if a field has been set.

### GetPercentages

`func (o *InlineResponse200385MemoryUsed) GetPercentages() InlineResponse200385MemoryUsedPercentages`

GetPercentages returns the Percentages field if non-nil, zero value otherwise.

### GetPercentagesOk

`func (o *InlineResponse200385MemoryUsed) GetPercentagesOk() (*InlineResponse200385MemoryUsedPercentages, bool)`

GetPercentagesOk returns a tuple with the Percentages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentages

`func (o *InlineResponse200385MemoryUsed) SetPercentages(v InlineResponse200385MemoryUsedPercentages)`

SetPercentages sets Percentages field to given value.

### HasPercentages

`func (o *InlineResponse200385MemoryUsed) HasPercentages() bool`

HasPercentages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


