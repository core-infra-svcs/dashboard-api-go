# InlineResponse200467Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **string** | The start time interval snapshots of the query range with iso8601 format | [optional] 
**EndTs** | Pointer to **string** | The end time interval snapshots of the query range with iso8601 format | [optional] 
**Overall** | Pointer to [**InlineResponse200467Overall**](InlineResponse200467Overall.md) |  | [optional] 
**ByInterface** | Pointer to [**[]InlineResponse200467ByInterface**](InlineResponse200467ByInterface.md) | The usage data on the interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200467Intervals

`func NewInlineResponse200467Intervals() *InlineResponse200467Intervals`

NewInlineResponse200467Intervals instantiates a new InlineResponse200467Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200467IntervalsWithDefaults

`func NewInlineResponse200467IntervalsWithDefaults() *InlineResponse200467Intervals`

NewInlineResponse200467IntervalsWithDefaults instantiates a new InlineResponse200467Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200467Intervals) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200467Intervals) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200467Intervals) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200467Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200467Intervals) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200467Intervals) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200467Intervals) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200467Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOverall

`func (o *InlineResponse200467Intervals) GetOverall() InlineResponse200467Overall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200467Intervals) GetOverallOk() (*InlineResponse200467Overall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200467Intervals) SetOverall(v InlineResponse200467Overall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200467Intervals) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByInterface

`func (o *InlineResponse200467Intervals) GetByInterface() []InlineResponse200467ByInterface`

GetByInterface returns the ByInterface field if non-nil, zero value otherwise.

### GetByInterfaceOk

`func (o *InlineResponse200467Intervals) GetByInterfaceOk() (*[]InlineResponse200467ByInterface, bool)`

GetByInterfaceOk returns a tuple with the ByInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByInterface

`func (o *InlineResponse200467Intervals) SetByInterface(v []InlineResponse200467ByInterface)`

SetByInterface sets ByInterface field to given value.

### HasByInterface

`func (o *InlineResponse200467Intervals) HasByInterface() bool`

HasByInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


