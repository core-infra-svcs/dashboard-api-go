# InlineResponse200244TwoFourGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPower** | Pointer to **int32** | Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30. | [optional] 
**MinPower** | Pointer to **int32** | Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5. | [optional] 
**MinBitrate** | Pointer to **float32** | Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of &#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. Defaults to 11. | [optional] 
**ValidAutoChannels** | Pointer to **[]int32** | Sets valid auto channels for 2.4Ghz band. Can be one of &#39;1&#39;, &#39;6&#39; or &#39;11&#39;. Defaults to [1, 6, 11]. | [optional] 
**AxEnabled** | Pointer to **bool** | Determines whether IEEE 802.11ax (Wi-Fi 6) on the 2.4 GHz band is enabled. Can be either true or false. If false, we highly recommend disabling band steering. Deprecated; use dot11ax.enabled instead. Defaults to true. | [optional] 
**Dot11ax** | Pointer to [**InlineResponse200244TwoFourGhzSettingsDot11ax**](InlineResponse200244TwoFourGhzSettingsDot11ax.md) |  | [optional] 
**Rxsop** | Pointer to **NullableInt32** | The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default. | [optional] 

## Methods

### NewInlineResponse200244TwoFourGhzSettings

`func NewInlineResponse200244TwoFourGhzSettings() *InlineResponse200244TwoFourGhzSettings`

NewInlineResponse200244TwoFourGhzSettings instantiates a new InlineResponse200244TwoFourGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200244TwoFourGhzSettingsWithDefaults

`func NewInlineResponse200244TwoFourGhzSettingsWithDefaults() *InlineResponse200244TwoFourGhzSettings`

NewInlineResponse200244TwoFourGhzSettingsWithDefaults instantiates a new InlineResponse200244TwoFourGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPower

`func (o *InlineResponse200244TwoFourGhzSettings) GetMaxPower() int32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetMaxPowerOk() (*int32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *InlineResponse200244TwoFourGhzSettings) SetMaxPower(v int32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *InlineResponse200244TwoFourGhzSettings) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *InlineResponse200244TwoFourGhzSettings) GetMinPower() int32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetMinPowerOk() (*int32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *InlineResponse200244TwoFourGhzSettings) SetMinPower(v int32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *InlineResponse200244TwoFourGhzSettings) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse200244TwoFourGhzSettings) GetMinBitrate() float32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetMinBitrateOk() (*float32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse200244TwoFourGhzSettings) SetMinBitrate(v float32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineResponse200244TwoFourGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetValidAutoChannels

`func (o *InlineResponse200244TwoFourGhzSettings) GetValidAutoChannels() []int32`

GetValidAutoChannels returns the ValidAutoChannels field if non-nil, zero value otherwise.

### GetValidAutoChannelsOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetValidAutoChannelsOk() (*[]int32, bool)`

GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAutoChannels

`func (o *InlineResponse200244TwoFourGhzSettings) SetValidAutoChannels(v []int32)`

SetValidAutoChannels sets ValidAutoChannels field to given value.

### HasValidAutoChannels

`func (o *InlineResponse200244TwoFourGhzSettings) HasValidAutoChannels() bool`

HasValidAutoChannels returns a boolean if a field has been set.

### GetAxEnabled

`func (o *InlineResponse200244TwoFourGhzSettings) GetAxEnabled() bool`

GetAxEnabled returns the AxEnabled field if non-nil, zero value otherwise.

### GetAxEnabledOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetAxEnabledOk() (*bool, bool)`

GetAxEnabledOk returns a tuple with the AxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxEnabled

`func (o *InlineResponse200244TwoFourGhzSettings) SetAxEnabled(v bool)`

SetAxEnabled sets AxEnabled field to given value.

### HasAxEnabled

`func (o *InlineResponse200244TwoFourGhzSettings) HasAxEnabled() bool`

HasAxEnabled returns a boolean if a field has been set.

### GetDot11ax

`func (o *InlineResponse200244TwoFourGhzSettings) GetDot11ax() InlineResponse200244TwoFourGhzSettingsDot11ax`

GetDot11ax returns the Dot11ax field if non-nil, zero value otherwise.

### GetDot11axOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetDot11axOk() (*InlineResponse200244TwoFourGhzSettingsDot11ax, bool)`

GetDot11axOk returns a tuple with the Dot11ax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot11ax

`func (o *InlineResponse200244TwoFourGhzSettings) SetDot11ax(v InlineResponse200244TwoFourGhzSettingsDot11ax)`

SetDot11ax sets Dot11ax field to given value.

### HasDot11ax

`func (o *InlineResponse200244TwoFourGhzSettings) HasDot11ax() bool`

HasDot11ax returns a boolean if a field has been set.

### GetRxsop

`func (o *InlineResponse200244TwoFourGhzSettings) GetRxsop() int32`

GetRxsop returns the Rxsop field if non-nil, zero value otherwise.

### GetRxsopOk

`func (o *InlineResponse200244TwoFourGhzSettings) GetRxsopOk() (*int32, bool)`

GetRxsopOk returns a tuple with the Rxsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxsop

`func (o *InlineResponse200244TwoFourGhzSettings) SetRxsop(v int32)`

SetRxsop sets Rxsop field to given value.

### HasRxsop

`func (o *InlineResponse200244TwoFourGhzSettings) HasRxsop() bool`

HasRxsop returns a boolean if a field has been set.

### SetRxsopNil

`func (o *InlineResponse200244TwoFourGhzSettings) SetRxsopNil(b bool)`

 SetRxsopNil sets the value for Rxsop to be an explicit nil

### UnsetRxsop
`func (o *InlineResponse200244TwoFourGhzSettings) UnsetRxsop()`

UnsetRxsop ensures that no value is present for Rxsop, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


