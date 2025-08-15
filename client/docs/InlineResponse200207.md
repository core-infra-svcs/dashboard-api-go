# InlineResponse200207

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID outage schedule is enabled. | [optional] 
**Ranges** | Pointer to [**[]InlineResponse200207Ranges**](InlineResponse200207Ranges.md) | List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. | [optional] 
**RangesInSeconds** | Pointer to [**[]InlineResponse200207RangesInSeconds**](InlineResponse200207RangesInSeconds.md) | List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. | [optional] 

## Methods

### NewInlineResponse200207

`func NewInlineResponse200207() *InlineResponse200207`

NewInlineResponse200207 instantiates a new InlineResponse200207 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200207WithDefaults

`func NewInlineResponse200207WithDefaults() *InlineResponse200207`

NewInlineResponse200207WithDefaults instantiates a new InlineResponse200207 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200207) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200207) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200207) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200207) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *InlineResponse200207) GetRanges() []InlineResponse200207Ranges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *InlineResponse200207) GetRangesOk() (*[]InlineResponse200207Ranges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *InlineResponse200207) SetRanges(v []InlineResponse200207Ranges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *InlineResponse200207) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRangesInSeconds

`func (o *InlineResponse200207) GetRangesInSeconds() []InlineResponse200207RangesInSeconds`

GetRangesInSeconds returns the RangesInSeconds field if non-nil, zero value otherwise.

### GetRangesInSecondsOk

`func (o *InlineResponse200207) GetRangesInSecondsOk() (*[]InlineResponse200207RangesInSeconds, bool)`

GetRangesInSecondsOk returns a tuple with the RangesInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesInSeconds

`func (o *InlineResponse200207) SetRangesInSeconds(v []InlineResponse200207RangesInSeconds)`

SetRangesInSeconds sets RangesInSeconds field to given value.

### HasRangesInSeconds

`func (o *InlineResponse200207) HasRangesInSeconds() bool`

HasRangesInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


