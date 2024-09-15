/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20043BasicServiceSets struct for InlineResponse20043BasicServiceSets
type InlineResponse20043BasicServiceSets struct {
	// Name of wireless network
	SsidName *string `json:"ssidName,omitempty"`
	// Unique identifier of wireless network
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// Status of wireless network
	Enabled *bool `json:"enabled,omitempty"`
	// Frequency range used by wireless network
	Band *string `json:"band,omitempty"`
	// Unique identifier of wireless access point
	Bssid *string `json:"bssid,omitempty"`
	// Frequency channel used by wireless network
	Channel *int32 `json:"channel,omitempty"`
	// Width of frequency channel used by wireless network
	ChannelWidth *string `json:"channelWidth,omitempty"`
	// Strength of wireless signal
	Power *string `json:"power,omitempty"`
	// Whether the SSID is advertised or hidden
	Visible *bool `json:"visible,omitempty"`
	// Whether the SSID is broadcasting based on an availability schedule
	Broadcasting *bool `json:"broadcasting,omitempty"`
}

// NewInlineResponse20043BasicServiceSets instantiates a new InlineResponse20043BasicServiceSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20043BasicServiceSets() *InlineResponse20043BasicServiceSets {
	this := InlineResponse20043BasicServiceSets{}
	return &this
}

// NewInlineResponse20043BasicServiceSetsWithDefaults instantiates a new InlineResponse20043BasicServiceSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20043BasicServiceSetsWithDefaults() *InlineResponse20043BasicServiceSets {
	this := InlineResponse20043BasicServiceSets{}
	return &this
}

// GetSsidName returns the SsidName field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetSsidName() string {
	if o == nil || isNil(o.SsidName) {
		var ret string
		return ret
	}
	return *o.SsidName
}

// GetSsidNameOk returns a tuple with the SsidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetSsidNameOk() (*string, bool) {
	if o == nil || isNil(o.SsidName) {
    return nil, false
	}
	return o.SsidName, true
}

// HasSsidName returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasSsidName() bool {
	if o != nil && !isNil(o.SsidName) {
		return true
	}

	return false
}

// SetSsidName gets a reference to the given string and assigns it to the SsidName field.
func (o *InlineResponse20043BasicServiceSets) SetSsidName(v string) {
	o.SsidName = &v
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetSsidNumber() int32 {
	if o == nil || isNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetSsidNumberOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNumber) {
    return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasSsidNumber() bool {
	if o != nil && !isNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *InlineResponse20043BasicServiceSets) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *InlineResponse20043BasicServiceSets) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetBand() string {
	if o == nil || isNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetBandOk() (*string, bool) {
	if o == nil || isNil(o.Band) {
    return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasBand() bool {
	if o != nil && !isNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *InlineResponse20043BasicServiceSets) SetBand(v string) {
	o.Band = &v
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetBssid() string {
	if o == nil || isNil(o.Bssid) {
		var ret string
		return ret
	}
	return *o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetBssidOk() (*string, bool) {
	if o == nil || isNil(o.Bssid) {
    return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasBssid() bool {
	if o != nil && !isNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given string and assigns it to the Bssid field.
func (o *InlineResponse20043BasicServiceSets) SetBssid(v string) {
	o.Bssid = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *InlineResponse20043BasicServiceSets) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetChannelWidth() string {
	if o == nil || isNil(o.ChannelWidth) {
		var ret string
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetChannelWidthOk() (*string, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given string and assigns it to the ChannelWidth field.
func (o *InlineResponse20043BasicServiceSets) SetChannelWidth(v string) {
	o.ChannelWidth = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetPower() string {
	if o == nil || isNil(o.Power) {
		var ret string
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetPowerOk() (*string, bool) {
	if o == nil || isNil(o.Power) {
    return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasPower() bool {
	if o != nil && !isNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given string and assigns it to the Power field.
func (o *InlineResponse20043BasicServiceSets) SetPower(v string) {
	o.Power = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetVisible() bool {
	if o == nil || isNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetVisibleOk() (*bool, bool) {
	if o == nil || isNil(o.Visible) {
    return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasVisible() bool {
	if o != nil && !isNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *InlineResponse20043BasicServiceSets) SetVisible(v bool) {
	o.Visible = &v
}

// GetBroadcasting returns the Broadcasting field value if set, zero value otherwise.
func (o *InlineResponse20043BasicServiceSets) GetBroadcasting() bool {
	if o == nil || isNil(o.Broadcasting) {
		var ret bool
		return ret
	}
	return *o.Broadcasting
}

// GetBroadcastingOk returns a tuple with the Broadcasting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20043BasicServiceSets) GetBroadcastingOk() (*bool, bool) {
	if o == nil || isNil(o.Broadcasting) {
    return nil, false
	}
	return o.Broadcasting, true
}

// HasBroadcasting returns a boolean if a field has been set.
func (o *InlineResponse20043BasicServiceSets) HasBroadcasting() bool {
	if o != nil && !isNil(o.Broadcasting) {
		return true
	}

	return false
}

// SetBroadcasting gets a reference to the given bool and assigns it to the Broadcasting field.
func (o *InlineResponse20043BasicServiceSets) SetBroadcasting(v bool) {
	o.Broadcasting = &v
}

func (o InlineResponse20043BasicServiceSets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SsidName) {
		toSerialize["ssidName"] = o.SsidName
	}
	if !isNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !isNil(o.Bssid) {
		toSerialize["bssid"] = o.Bssid
	}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !isNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !isNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !isNil(o.Broadcasting) {
		toSerialize["broadcasting"] = o.Broadcasting
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20043BasicServiceSets struct {
	value *InlineResponse20043BasicServiceSets
	isSet bool
}

func (v NullableInlineResponse20043BasicServiceSets) Get() *InlineResponse20043BasicServiceSets {
	return v.value
}

func (v *NullableInlineResponse20043BasicServiceSets) Set(val *InlineResponse20043BasicServiceSets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20043BasicServiceSets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20043BasicServiceSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20043BasicServiceSets(val *InlineResponse20043BasicServiceSets) *NullableInlineResponse20043BasicServiceSets {
	return &NullableInlineResponse20043BasicServiceSets{value: val, isSet: true}
}

func (v NullableInlineResponse20043BasicServiceSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20043BasicServiceSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


