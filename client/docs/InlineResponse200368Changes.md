# InlineResponse200368Changes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTs** | Pointer to **time.Time** | The start time(UTC seconds) of the wireless LAN controller connectivity status change | [optional] 
**EndTs** | Pointer to **string** | The end time(UTC seconds) of the wireless LAN controller connectivity status change. This attribute is set to be null by default if there&#39;s no need to assign. | [optional] 
**Status** | Pointer to **string** | The wireless LAN controller connectivity status | [optional] 

## Methods

### NewInlineResponse200368Changes

`func NewInlineResponse200368Changes() *InlineResponse200368Changes`

NewInlineResponse200368Changes instantiates a new InlineResponse200368Changes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200368ChangesWithDefaults

`func NewInlineResponse200368ChangesWithDefaults() *InlineResponse200368Changes`

NewInlineResponse200368ChangesWithDefaults instantiates a new InlineResponse200368Changes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTs

`func (o *InlineResponse200368Changes) GetStartTs() time.Time`

GetStartTs returns the StartTs field if non-nil, zero value otherwise.

### GetStartTsOk

`func (o *InlineResponse200368Changes) GetStartTsOk() (*time.Time, bool)`

GetStartTsOk returns a tuple with the StartTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTs

`func (o *InlineResponse200368Changes) SetStartTs(v time.Time)`

SetStartTs sets StartTs field to given value.

### HasStartTs

`func (o *InlineResponse200368Changes) HasStartTs() bool`

HasStartTs returns a boolean if a field has been set.

### GetEndTs

`func (o *InlineResponse200368Changes) GetEndTs() string`

GetEndTs returns the EndTs field if non-nil, zero value otherwise.

### GetEndTsOk

`func (o *InlineResponse200368Changes) GetEndTsOk() (*string, bool)`

GetEndTsOk returns a tuple with the EndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTs

`func (o *InlineResponse200368Changes) SetEndTs(v string)`

SetEndTs sets EndTs field to given value.

### HasEndTs

`func (o *InlineResponse200368Changes) HasEndTs() bool`

HasEndTs returns a boolean if a field has been set.

### GetStatus

`func (o *InlineResponse200368Changes) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InlineResponse200368Changes) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InlineResponse200368Changes) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InlineResponse200368Changes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


