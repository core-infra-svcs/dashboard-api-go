# InlineResponse200378Intervals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **string** | The start time interval snapshots of the query range with iso8601 format | [optional] 
**EndTs** | Pointer to **string** | The end time interval snapshots of the query range with iso8601 format | [optional] 
**Overall** | Pointer to [**InlineResponse200378Overall**](InlineResponse200378Overall.md) |  | [optional] 
**ByInterface** | Pointer to [**[]InlineResponse200378ByInterface**](InlineResponse200378ByInterface.md) | The usage data on the interfaces of the wireless LAN controller | [optional] 

## Methods

### NewInlineResponse200378Intervals

`func NewInlineResponse200378Intervals() *InlineResponse200378Intervals`

NewInlineResponse200378Intervals instantiates a new InlineResponse200378Intervals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200378IntervalsWithDefaults

`func NewInlineResponse200378IntervalsWithDefaults() *InlineResponse200378Intervals`

NewInlineResponse200378IntervalsWithDefaults instantiates a new InlineResponse200378Intervals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200378Intervals) GetStartTs() string`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200378Intervals) GetStartTsOk() (*string, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200378Intervals) SetStartTs(v string)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200378Intervals) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200378Intervals) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200378Intervals) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200378Intervals) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200378Intervals) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetOverall

`func (o *InlineResponse200378Intervals) GetOverall() InlineResponse200378Overall`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *InlineResponse200378Intervals) GetOverallOk() (*InlineResponse200378Overall, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *InlineResponse200378Intervals) SetOverall(v InlineResponse200378Overall)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *InlineResponse200378Intervals) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByInterface

`func (o *InlineResponse200378Intervals) GetByInterface() []InlineResponse200378ByInterface`

GetByInterface returns the ByInterface field if non-nil, zero value otherwise.

### GetByInterfaceOk

`func (o *InlineResponse200378Intervals) GetByInterfaceOk() (*[]InlineResponse200378ByInterface, bool)`

GetByInterfaceOk returns a tuple with the ByInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByInterface

`func (o *InlineResponse200378Intervals) SetByInterface(v []InlineResponse200378ByInterface)`

SetByInterface sets ByInterface field to given value.

### HasByInterface

`func (o *InlineResponse200378Intervals) HasByInterface() bool`

HasByInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


