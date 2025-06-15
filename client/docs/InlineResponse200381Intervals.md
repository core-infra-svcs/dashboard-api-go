# InlineResponse200381Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **string** | The start time of the query range with iso8601 format | [optional] 
**EndTs** | Pointer to **string** | The end time of the query range  with iso8601 format | [optional] 
**Overall** | Pointer to [**InlineResponse200381Overall**](InlineResponse200381Overall.md) |  | [optional] 
**ByCore** | Pointer to [**[]InlineResponse200381ByCore**](InlineResponse200381ByCore.md) | The CPU usage per core of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200381Intervals

`func NewInlineResponse200381Intervals() *InlineResponse200381Intervals`

NewInlineResponse200381Intervals instantiates a new InlineResponse200381Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200381IntervalsWithDefaults

`func NewInlineResponse200381IntervalsWithDefaults() *InlineResponse200381Intervals`

NewInlineResponse200381IntervalsWithDefaults instantiates a new InlineResponse200381Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200381Intervals) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200381Intervals) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200381Intervals) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200381Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200381Intervals) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200381Intervals) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200381Intervals) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200381Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOverall

`func (o *InlineResponse200381Intervals) GetOverall() InlineResponse200381Overall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200381Intervals) GetOverallOk() (*InlineResponse200381Overall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200381Intervals) SetOverall(v InlineResponse200381Overall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200381Intervals) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByCore

`func (o *InlineResponse200381Intervals) GetByCore() []InlineResponse200381ByCore`

GetByCore returns the ByCore field if non-nil, zero value otherwise.

### GetByCoreOk

`func (o *InlineResponse200381Intervals) GetByCoreOk() (*[]InlineResponse200381ByCore, bool)`

GetByCoreOk returns a tuple with the ByCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCore

`func (o *InlineResponse200381Intervals) SetByCore(v []InlineResponse200381ByCore)`

SetByCore sets ByCore field to given value.

### HasByCore

`func (o *InlineResponse200381Intervals) HasByCore() bool`

HasByCore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


