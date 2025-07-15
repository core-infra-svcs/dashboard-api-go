# InlineResponse200361Series

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ts** | Pointer to **time.Time** | Timestamp of the cpu load measurement | [optional] 
**CpuLoad5** | Pointer to **int32** | The 5 minutes cpu load average of the device | [optional] 

## Methods

### NewInlineResponse200361Series

`func NewInlineResponse200361Series() *InlineResponse200361Series`

NewInlineResponse200361Series instantiates a new InlineResponse200361Series object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200361SeriesWithDefaults

`func NewInlineResponse200361SeriesWithDefaults() *InlineResponse200361Series`

NewInlineResponse200361SeriesWithDefaults instantiates a new InlineResponse200361Series object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTs

`func (o *InlineResponse200361Series) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *InlineResponse200361Series) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *InlineResponse200361Series) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *InlineResponse200361Series) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetCpuLoad5

`func (o *InlineResponse200361Series) GetCpuLoad5() int32`

GetCpuLoad5 returns the CpuLoad5 field if non-nil, zero value otherwise.

### GetCpuLoad5Ok

`func (o *InlineResponse200361Series) GetCpuLoad5Ok() (*int32, bool)`

GetCpuLoad5Ok returns a tuple with the CpuLoad5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad5

`func (o *InlineResponse200361Series) SetCpuLoad5(v int32)`

SetCpuLoad5 sets CpuLoad5 field to given value.

### HasCpuLoad5

`func (o *InlineResponse200361Series) HasCpuLoad5() bool`

HasCpuLoad5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


