# InlineResponse200388Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **string** | The start time interval snapshots of the query range with iso8601 format | [optional] 
**EndTs** | Pointer to **string** | The end time interval snapshots of the query range with iso8601 format | [optional] 
**Overall** | Pointer to [**InlineResponse200388Overall**](InlineResponse200388Overall.md) |  | [optional] 
**ByInterface** | Pointer to [**[]InlineResponse200388ByInterface**](InlineResponse200388ByInterface.md) | The usage data on the interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200388Intervals

`func NewInlineResponse200388Intervals() *InlineResponse200388Intervals`

NewInlineResponse200388Intervals instantiates a new InlineResponse200388Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200388IntervalsWithDefaults

`func NewInlineResponse200388IntervalsWithDefaults() *InlineResponse200388Intervals`

NewInlineResponse200388IntervalsWithDefaults instantiates a new InlineResponse200388Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200388Intervals) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200388Intervals) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200388Intervals) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200388Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200388Intervals) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200388Intervals) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200388Intervals) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200388Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOverall

`func (o *InlineResponse200388Intervals) GetOverall() InlineResponse200388Overall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200388Intervals) GetOverallOk() (*InlineResponse200388Overall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200388Intervals) SetOverall(v InlineResponse200388Overall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200388Intervals) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByInterface

`func (o *InlineResponse200388Intervals) GetByInterface() []InlineResponse200388ByInterface`

GetByInterface returns the ByInterface field if non-nil, zero value otherwise.

### GetByInterfaceOk

`func (o *InlineResponse200388Intervals) GetByInterfaceOk() (*[]InlineResponse200388ByInterface, bool)`

GetByInterfaceOk returns a tuple with the ByInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByInterface

`func (o *InlineResponse200388Intervals) SetByInterface(v []InlineResponse200388ByInterface)`

SetByInterface sets ByInterface field to given value.

### HasByInterface

`func (o *InlineResponse200388Intervals) HasByInterface() bool`

HasByInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


