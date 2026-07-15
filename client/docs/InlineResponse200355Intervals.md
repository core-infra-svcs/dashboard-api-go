# InlineResponse200355Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | Timestamp for the beginning of the historical snapshot, exclusive. | [optional] 
**EndTs** | Pointer to **time.Time** | Timestamp for the end of the historical snapshot, inclusive. | [optional] 
**Memory** | Pointer to [**InlineResponse200355Memory**](InlineResponse200355Memory.md) |  | [optional] 

## Methods

### NewInlineResponse200355Intervals

`func NewInlineResponse200355Intervals() *InlineResponse200355Intervals`

NewInlineResponse200355Intervals instantiates a new InlineResponse200355Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200355IntervalsWithDefaults

`func NewInlineResponse200355IntervalsWithDefaults() *InlineResponse200355Intervals`

NewInlineResponse200355IntervalsWithDefaults instantiates a new InlineResponse200355Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200355Intervals) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200355Intervals) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200355Intervals) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200355Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200355Intervals) GetEndTs() time.Time`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200355Intervals) GetEndTsOk() (*time.Time, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200355Intervals) SetEndTs(v time.Time)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200355Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetMemory

`func (o *InlineResponse200355Intervals) GetMemory() InlineResponse200355Memory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *InlineResponse200355Intervals) GetMemoryOk() (*InlineResponse200355Memory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *InlineResponse200355Intervals) SetMemory(v InlineResponse200355Memory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *InlineResponse200355Intervals) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


