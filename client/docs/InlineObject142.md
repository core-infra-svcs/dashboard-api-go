# InlineObject142

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, the SSID outage schedule is enabled. | [optional] 
**Ranges** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberSchedulesRanges**](NetworksNetworkIdWirelessSsidsNumberSchedulesRanges.md) | List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence. | [optional] 
**RangesInSeconds** | Pointer to [**[]NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds**](NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds.md) | List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence. | [optional] 

## Methods

### NewInlineObject142

`func NewInlineObject142() *InlineObject142`

NewInlineObject142 instantiates a new InlineObject142 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject142WithDefaults

`func NewInlineObject142WithDefaults() *InlineObject142`

NewInlineObject142WithDefaults instantiates a new InlineObject142 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineObject142) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineObject142) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineObject142) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineObject142) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRanges

`func (o *InlineObject142) GetRanges() []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *InlineObject142) GetRangesOk() (*[]NetworksNetworkIdWirelessSsidsNumberSchedulesRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *InlineObject142) SetRanges(v []NetworksNetworkIdWirelessSsidsNumberSchedulesRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *InlineObject142) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetRangesInSeconds

`func (o *InlineObject142) GetRangesInSeconds() []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds`

GetRangesInSeconds returns the RangesInSeconds field if non-nil, zero value otherwise.

### GetRangesInSecondsOk

`func (o *InlineObject142) GetRangesInSecondsOk() (*[]NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds, bool)`

GetRangesInSecondsOk returns a tuple with the RangesInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangesInSeconds

`func (o *InlineObject142) SetRangesInSeconds(v []NetworksNetworkIdWirelessSsidsNumberSchedulesRangesInSeconds)`

SetRangesInSeconds sets RangesInSeconds field to given value.

### HasRangesInSeconds

`func (o *InlineObject142) HasRangesInSeconds() bool`

HasRangesInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


