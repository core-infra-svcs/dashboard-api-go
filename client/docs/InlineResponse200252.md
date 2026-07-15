# InlineResponse200252

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID outage schedule is enabled. | [optional] 
**Ranges** | Pointer to [**[]InlineResponse200252Ranges**](InlineResponse200252Ranges.md) | List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. | [optional] 
**RangesInSeconds** | Pointer to [**[]InlineResponse200252RangesInSeconds**](InlineResponse200252RangesInSeconds.md) | List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. | [optional] 

## Methods

### NewInlineResponse200252

`func NewInlineResponse200252() *InlineResponse200252`

NewInlineResponse200252 instantiates a new InlineResponse200252 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200252WithDefaults

`func NewInlineResponse200252WithDefaults() *InlineResponse200252`

NewInlineResponse200252WithDefaults instantiates a new InlineResponse200252 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200252) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200252) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200252) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200252) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *InlineResponse200252) GetRanges() []InlineResponse200252Ranges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *InlineResponse200252) GetRangesOk() (*[]InlineResponse200252Ranges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *InlineResponse200252) SetRanges(v []InlineResponse200252Ranges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *InlineResponse200252) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRangesInSeconds

`func (o *InlineResponse200252) GetRangesInSeconds() []InlineResponse200252RangesInSeconds`

GetRangesInSeconds returns the RangesInSeconds field if non-nil, zero value otherwise.

### GetRangesInSecondsOk

`func (o *InlineResponse200252) GetRangesInSecondsOk() (*[]InlineResponse200252RangesInSeconds, bool)`

GetRangesInSecondsOk returns a tuple with the RangesInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesInSeconds

`func (o *InlineResponse200252) SetRangesInSeconds(v []InlineResponse200252RangesInSeconds)`

SetRangesInSeconds sets RangesInSeconds field to given value.

### HasRangesInSeconds

`func (o *InlineResponse200252) HasRangesInSeconds() bool`

HasRangesInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


