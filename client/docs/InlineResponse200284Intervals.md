# InlineResponse200284Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Timestamp for the beginning of the historical snapshot, exclusive. | [optional] 
**EndTs** | Pointer to **time.Time** | Timestamp for the end of the historical snapshot, inclusive. | [optional] 
**Memory** | Pointer to [**InlineResponse200284Memory**](InlineResponse200284Memory.md) |  | [optional] 

## Methods

### NewInlineResponse200284Intervals

`func NewInlineResponse200284Intervals() *InlineResponse200284Intervals`

NewInlineResponse200284Intervals instantiates a new InlineResponse200284Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200284IntervalsWithDefaults

`func NewInlineResponse200284IntervalsWithDefaults() *InlineResponse200284Intervals`

NewInlineResponse200284IntervalsWithDefaults instantiates a new InlineResponse200284Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200284Intervals) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200284Intervals) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200284Intervals) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200284Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200284Intervals) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200284Intervals) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200284Intervals) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200284Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetMemory

`func (o *InlineResponse200284Intervals) GetMemory() InlineResponse200284Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InlineResponse200284Intervals) GetMemoryOk() (*InlineResponse200284Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InlineResponse200284Intervals) SetMemory(v InlineResponse200284Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InlineResponse200284Intervals) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


