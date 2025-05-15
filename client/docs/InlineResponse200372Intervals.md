# InlineResponse200372Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **string** | The start time of the query range with iso8601 format | [optional] 
**EndTs** | Pointer to **string** | The end time of the query range  with iso8601 format | [optional] 
**Overall** | Pointer to [**InlineResponse200372Overall**](InlineResponse200372Overall.md) |  | [optional] 
**ByCore** | Pointer to [**[]InlineResponse200372ByCore**](InlineResponse200372ByCore.md) | The CPU usage per core of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200372Intervals

`func NewInlineResponse200372Intervals() *InlineResponse200372Intervals`

NewInlineResponse200372Intervals instantiates a new InlineResponse200372Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200372IntervalsWithDefaults

`func NewInlineResponse200372IntervalsWithDefaults() *InlineResponse200372Intervals`

NewInlineResponse200372IntervalsWithDefaults instantiates a new InlineResponse200372Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200372Intervals) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200372Intervals) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200372Intervals) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200372Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200372Intervals) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200372Intervals) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200372Intervals) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200372Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOverall

`func (o *InlineResponse200372Intervals) GetOverall() InlineResponse200372Overall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200372Intervals) GetOverallOk() (*InlineResponse200372Overall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200372Intervals) SetOverall(v InlineResponse200372Overall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200372Intervals) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByCore

`func (o *InlineResponse200372Intervals) GetByCore() []InlineResponse200372ByCore`

GetByCore returns the ByCore field if non-nil, zero value otherwise.

### GetByCoreOk

`func (o *InlineResponse200372Intervals) GetByCoreOk() (*[]InlineResponse200372ByCore, bool)`

GetByCoreOk returns a tuple with the ByCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCore

`func (o *InlineResponse200372Intervals) SetByCore(v []InlineResponse200372ByCore)`

SetByCore sets ByCore field to given value.

### HasByCore

`func (o *InlineResponse200372Intervals) HasByCore() bool`

HasByCore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


