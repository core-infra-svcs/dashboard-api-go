# InlineObject201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID outage schedule is enabled. | [optional] 
**Ranges** | Pointer to [**[]InlineResponse200205Ranges**](InlineResponse200205Ranges.md) | List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. | [optional] 
**RangesInSeconds** | Pointer to [**[]InlineResponse200205RangesInSeconds**](InlineResponse200205RangesInSeconds.md) | List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. | [optional] 

## Methods

### NewInlineObject201

`func NewInlineObject201() *InlineObject201`

NewInlineObject201 instantiates a new InlineObject201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject201WithDefaults

`func NewInlineObject201WithDefaults() *InlineObject201`

NewInlineObject201WithDefaults instantiates a new InlineObject201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject201) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject201) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject201) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject201) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *InlineObject201) GetRanges() []InlineResponse200205Ranges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *InlineObject201) GetRangesOk() (*[]InlineResponse200205Ranges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *InlineObject201) SetRanges(v []InlineResponse200205Ranges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *InlineObject201) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRangesInSeconds

`func (o *InlineObject201) GetRangesInSeconds() []InlineResponse200205RangesInSeconds`

GetRangesInSeconds returns the RangesInSeconds field if non-nil, zero value otherwise.

### GetRangesInSecondsOk

`func (o *InlineObject201) GetRangesInSecondsOk() (*[]InlineResponse200205RangesInSeconds, bool)`

GetRangesInSecondsOk returns a tuple with the RangesInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesInSeconds

`func (o *InlineObject201) SetRangesInSeconds(v []InlineResponse200205RangesInSeconds)`

SetRangesInSeconds sets RangesInSeconds field to given value.

### HasRangesInSeconds

`func (o *InlineObject201) HasRangesInSeconds() bool`

HasRangesInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


