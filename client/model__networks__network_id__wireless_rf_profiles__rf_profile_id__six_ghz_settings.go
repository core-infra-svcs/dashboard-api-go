/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs
type NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings struct {
	// Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30.
	MaxPower *int32 `json:"maxPower,omitempty"`
	// Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30.
	MinPower *int32 `json:"minPower,omitempty"`
	// Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'.
	ValidAutoChannels []int32 `json:"validAutoChannels,omitempty"`
	// Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'.
	ChannelWidth *string `json:"channelWidth,omitempty"`
	// The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	Rxsop *int32 `json:"rxsop,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings instantiates a new NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings() *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings {
	this := NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettingsWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettingsWithDefaults() *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings {
	this := NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings{}
	return &this
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetMaxPower() int32 {
	if o == nil || isNil(o.MaxPower) {
		var ret int32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetMaxPowerOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPower) {
    return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) HasMaxPower() bool {
	if o != nil && !isNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int32 and assigns it to the MaxPower field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) SetMaxPower(v int32) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetMinPower() int32 {
	if o == nil || isNil(o.MinPower) {
		var ret int32
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetMinPowerOk() (*int32, bool) {
	if o == nil || isNil(o.MinPower) {
    return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) HasMinPower() bool {
	if o != nil && !isNil(o.MinPower) {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int32 and assigns it to the MinPower field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) SetMinPower(v int32) {
	o.MinPower = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetValidAutoChannels returns the ValidAutoChannels field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetValidAutoChannels() []int32 {
	if o == nil || isNil(o.ValidAutoChannels) {
		var ret []int32
		return ret
	}
	return o.ValidAutoChannels
}

// GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetValidAutoChannelsOk() ([]int32, bool) {
	if o == nil || isNil(o.ValidAutoChannels) {
    return nil, false
	}
	return o.ValidAutoChannels, true
}

// HasValidAutoChannels returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) HasValidAutoChannels() bool {
	if o != nil && !isNil(o.ValidAutoChannels) {
		return true
	}

	return false
}

// SetValidAutoChannels gets a reference to the given []int32 and assigns it to the ValidAutoChannels field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) SetValidAutoChannels(v []int32) {
	o.ValidAutoChannels = v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetChannelWidth() string {
	if o == nil || isNil(o.ChannelWidth) {
		var ret string
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetChannelWidthOk() (*string, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given string and assigns it to the ChannelWidth field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) SetChannelWidth(v string) {
	o.ChannelWidth = &v
}

// GetRxsop returns the Rxsop field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetRxsop() int32 {
	if o == nil || isNil(o.Rxsop) {
		var ret int32
		return ret
	}
	return *o.Rxsop
}

// GetRxsopOk returns a tuple with the Rxsop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) GetRxsopOk() (*int32, bool) {
	if o == nil || isNil(o.Rxsop) {
    return nil, false
	}
	return o.Rxsop, true
}

// HasRxsop returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) HasRxsop() bool {
	if o != nil && !isNil(o.Rxsop) {
		return true
	}

	return false
}

// SetRxsop gets a reference to the given int32 and assigns it to the Rxsop field.
func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) SetRxsop(v int32) {
	o.Rxsop = &v
}

func (o NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxPower) {
		toSerialize["maxPower"] = o.MaxPower
	}
	if !isNil(o.MinPower) {
		toSerialize["minPower"] = o.MinPower
	}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.ValidAutoChannels) {
		toSerialize["validAutoChannels"] = o.ValidAutoChannels
	}
	if !isNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !isNil(o.Rxsop) {
		toSerialize["rxsop"] = o.Rxsop
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings struct {
	value *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) Get() *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) Set(val *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings(val *NetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings {
	return &NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesRfProfileIdSixGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


