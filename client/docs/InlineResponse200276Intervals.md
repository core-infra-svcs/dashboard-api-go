# InlineResponse200276Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Timestamp for the beginning of the historical snapshot, exclusive. | [optional] 
**EndTs** | Pointer to **time.Time** | Timestamp for the end of the historical snapshot, inclusive. | [optional] 
**Memory** | Pointer to [**InlineResponse200276Memory**](InlineResponse200276Memory.md) |  | [optional] 

## Methods

### NewInlineResponse200276Intervals

`func NewInlineResponse200276Intervals() *InlineResponse200276Intervals`

NewInlineResponse200276Intervals instantiates a new InlineResponse200276Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200276IntervalsWithDefaults

`func NewInlineResponse200276IntervalsWithDefaults() *InlineResponse200276Intervals`

NewInlineResponse200276IntervalsWithDefaults instantiates a new InlineResponse200276Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200276Intervals) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200276Intervals) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200276Intervals) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200276Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200276Intervals) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200276Intervals) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200276Intervals) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200276Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetMemory

`func (o *InlineResponse200276Intervals) GetMemory() InlineResponse200276Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InlineResponse200276Intervals) GetMemoryOk() (*InlineResponse200276Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InlineResponse200276Intervals) SetMemory(v InlineResponse200276Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InlineResponse200276Intervals) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


