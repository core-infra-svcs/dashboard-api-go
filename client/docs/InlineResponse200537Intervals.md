# InlineResponse200537Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **string** | The start time of the query range with iso8601 format | [optional] 
**EndTs** | Pointer to **string** | The end time of the query range  with iso8601 format | [optional] 
**Overall** | Pointer to [**InlineResponse200537Overall**](InlineResponse200537Overall.md) |  | [optional] 
**ByCore** | Pointer to [**[]InlineResponse200537ByCore**](InlineResponse200537ByCore.md) | The CPU usage per core of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200537Intervals

`func NewInlineResponse200537Intervals() *InlineResponse200537Intervals`

NewInlineResponse200537Intervals instantiates a new InlineResponse200537Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200537IntervalsWithDefaults

`func NewInlineResponse200537IntervalsWithDefaults() *InlineResponse200537Intervals`

NewInlineResponse200537IntervalsWithDefaults instantiates a new InlineResponse200537Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200537Intervals) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200537Intervals) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200537Intervals) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200537Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200537Intervals) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200537Intervals) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200537Intervals) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200537Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOverall

`func (o *InlineResponse200537Intervals) GetOverall() InlineResponse200537Overall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200537Intervals) GetOverallOk() (*InlineResponse200537Overall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200537Intervals) SetOverall(v InlineResponse200537Overall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200537Intervals) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByCore

`func (o *InlineResponse200537Intervals) GetByCore() []InlineResponse200537ByCore`

GetByCore returns the ByCore field if non-nil, zero value otherwise.

### GetByCoreOk

`func (o *InlineResponse200537Intervals) GetByCoreOk() (*[]InlineResponse200537ByCore, bool)`

GetByCoreOk returns a tuple with the ByCore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCore

`func (o *InlineResponse200537Intervals) SetByCore(v []InlineResponse200537ByCore)`

SetByCore sets ByCore field to given value.

### HasByCore

`func (o *InlineResponse200537Intervals) HasByCore() bool`

HasByCore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


