# InlineResponse200196TwoFourGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPower** | Pointer to **int32** | Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30. | [optional] 
**MinPower** | Pointer to **int32** | Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5. | [optional] 
**MinBitrate** | Pointer to **float32** | Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of &#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. Defaults to 11. | [optional] 
**ValidAutoChannels** | Pointer to **[]int32** | Sets valid auto channels for 2.4Ghz band. Can be one of &#39;1&#39;, &#39;6&#39; or &#39;11&#39;. Defaults to [1, 6, 11]. | [optional] 
**AxEnabled** | Pointer to **bool** | Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true. | [optional] 
**Rxsop** | Pointer to **int32** | The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default. | [optional] 

## Methods

### NewInlineResponse200196TwoFourGhzSettings

`func NewInlineResponse200196TwoFourGhzSettings() *InlineResponse200196TwoFourGhzSettings`

NewInlineResponse200196TwoFourGhzSettings instantiates a new InlineResponse200196TwoFourGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200196TwoFourGhzSettingsWithDefaults

`func NewInlineResponse200196TwoFourGhzSettingsWithDefaults() *InlineResponse200196TwoFourGhzSettings`

NewInlineResponse200196TwoFourGhzSettingsWithDefaults instantiates a new InlineResponse200196TwoFourGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPower

`func (o *InlineResponse200196TwoFourGhzSettings) GetMaxPower() int32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *InlineResponse200196TwoFourGhzSettings) GetMaxPowerOk() (*int32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *InlineResponse200196TwoFourGhzSettings) SetMaxPower(v int32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *InlineResponse200196TwoFourGhzSettings) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *InlineResponse200196TwoFourGhzSettings) GetMinPower() int32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *InlineResponse200196TwoFourGhzSettings) GetMinPowerOk() (*int32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *InlineResponse200196TwoFourGhzSettings) SetMinPower(v int32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *InlineResponse200196TwoFourGhzSettings) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse200196TwoFourGhzSettings) GetMinBitrate() float32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse200196TwoFourGhzSettings) GetMinBitrateOk() (*float32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse200196TwoFourGhzSettings) SetMinBitrate(v float32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineResponse200196TwoFourGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetValidAutoChannels

`func (o *InlineResponse200196TwoFourGhzSettings) GetValidAutoChannels() []int32`

GetValidAutoChannels returns the ValidAutoChannels field if non-nil, zero value otherwise.

### GetValidAutoChannelsOk

`func (o *InlineResponse200196TwoFourGhzSettings) GetValidAutoChannelsOk() (*[]int32, bool)`

GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAutoChannels

`func (o *InlineResponse200196TwoFourGhzSettings) SetValidAutoChannels(v []int32)`

SetValidAutoChannels sets ValidAutoChannels field to given value.

### HasValidAutoChannels

`func (o *InlineResponse200196TwoFourGhzSettings) HasValidAutoChannels() bool`

HasValidAutoChannels returns a boolean if a field has been set.

### GetAxEnabled

`func (o *InlineResponse200196TwoFourGhzSettings) GetAxEnabled() bool`

GetAxEnabled returns the AxEnabled field if non-nil, zero value otherwise.

### GetAxEnabledOk

`func (o *InlineResponse200196TwoFourGhzSettings) GetAxEnabledOk() (*bool, bool)`

GetAxEnabledOk returns a tuple with the AxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxEnabled

`func (o *InlineResponse200196TwoFourGhzSettings) SetAxEnabled(v bool)`

SetAxEnabled sets AxEnabled field to given value.

### HasAxEnabled

`func (o *InlineResponse200196TwoFourGhzSettings) HasAxEnabled() bool`

HasAxEnabled returns a boolean if a field has been set.

### GetRxsop

`func (o *InlineResponse200196TwoFourGhzSettings) GetRxsop() int32`

GetRxsop returns the Rxsop field if non-nil, zero value otherwise.

### GetRxsopOk

`func (o *InlineResponse200196TwoFourGhzSettings) GetRxsopOk() (*int32, bool)`

GetRxsopOk returns a tuple with the Rxsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxsop

`func (o *InlineResponse200196TwoFourGhzSettings) SetRxsop(v int32)`

SetRxsop sets Rxsop field to given value.

### HasRxsop

`func (o *InlineResponse200196TwoFourGhzSettings) HasRxsop() bool`

HasRxsop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


