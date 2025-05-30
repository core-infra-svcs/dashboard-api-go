# InlineResponse20047BasicServiceSets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidName** | Pointer to **string** | Name of wireless network | [optional] 
**SsidNumber** | Pointer to **int32** | Unique identifier of wireless network | [optional] 
**Enabled** | Pointer to **bool** | Status of wireless network | [optional] 
**Band** | Pointer to **string** | Frequency range used by wireless network | [optional] 
**Bssid** | Pointer to **string** | Unique identifier of wireless access point | [optional] 
**Channel** | Pointer to **int32** | Frequency channel used by wireless network | [optional] 
**ChannelWidth** | Pointer to **string** | Width of frequency channel used by wireless network | [optional] 
**Power** | Pointer to **string** | Strength of wireless signal | [optional] 
**Visible** | Pointer to **bool** | Whether the SSID is advertised or hidden | [optional] 
**Broadcasting** | Pointer to **bool** | Whether the SSID is broadcasting based on an availability schedule | [optional] 

## Methods

### NewInlineResponse20047BasicServiceSets

`func NewInlineResponse20047BasicServiceSets() *InlineResponse20047BasicServiceSets`

NewInlineResponse20047BasicServiceSets instantiates a new InlineResponse20047BasicServiceSets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20047BasicServiceSetsWithDefaults

`func NewInlineResponse20047BasicServiceSetsWithDefaults() *InlineResponse20047BasicServiceSets`

NewInlineResponse20047BasicServiceSetsWithDefaults instantiates a new InlineResponse20047BasicServiceSets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidName

`func (o *InlineResponse20047BasicServiceSets) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *InlineResponse20047BasicServiceSets) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *InlineResponse20047BasicServiceSets) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *InlineResponse20047BasicServiceSets) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetSsidNumber

`func (o *InlineResponse20047BasicServiceSets) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *InlineResponse20047BasicServiceSets) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *InlineResponse20047BasicServiceSets) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *InlineResponse20047BasicServiceSets) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetEnabled

`func (o *InlineResponse20047BasicServiceSets) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse20047BasicServiceSets) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse20047BasicServiceSets) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse20047BasicServiceSets) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBand

`func (o *InlineResponse20047BasicServiceSets) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *InlineResponse20047BasicServiceSets) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *InlineResponse20047BasicServiceSets) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *InlineResponse20047BasicServiceSets) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetBssid

`func (o *InlineResponse20047BasicServiceSets) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *InlineResponse20047BasicServiceSets) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *InlineResponse20047BasicServiceSets) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *InlineResponse20047BasicServiceSets) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetChannel

`func (o *InlineResponse20047BasicServiceSets) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *InlineResponse20047BasicServiceSets) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *InlineResponse20047BasicServiceSets) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *InlineResponse20047BasicServiceSets) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelWidth

`func (o *InlineResponse20047BasicServiceSets) GetChannelWidth() string`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *InlineResponse20047BasicServiceSets) GetChannelWidthOk() (*string, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *InlineResponse20047BasicServiceSets) SetChannelWidth(v string)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *InlineResponse20047BasicServiceSets) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetPower

`func (o *InlineResponse20047BasicServiceSets) GetPower() string`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *InlineResponse20047BasicServiceSets) GetPowerOk() (*string, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *InlineResponse20047BasicServiceSets) SetPower(v string)`

SetPower sets Power field to given value.

### HasPower

`func (o *InlineResponse20047BasicServiceSets) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetVisible

`func (o *InlineResponse20047BasicServiceSets) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InlineResponse20047BasicServiceSets) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InlineResponse20047BasicServiceSets) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *InlineResponse20047BasicServiceSets) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetBroadcasting

`func (o *InlineResponse20047BasicServiceSets) GetBroadcasting() bool`

GetBroadcasting returns the Broadcasting field if non-nil, zero value otherwise.

### GetBroadcastingOk

`func (o *InlineResponse20047BasicServiceSets) GetBroadcastingOk() (*bool, bool)`

GetBroadcastingOk returns a tuple with the Broadcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasting

`func (o *InlineResponse20047BasicServiceSets) SetBroadcasting(v bool)`

SetBroadcasting sets Broadcasting field to given value.

### HasBroadcasting

`func (o *InlineResponse20047BasicServiceSets) HasBroadcasting() bool`

HasBroadcasting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


