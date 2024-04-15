/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200290Radio Wireless access point radio identifier.
type InlineResponse200290Radio struct {
	// Frequency range used for wireless communication.
	Band *string `json:"band,omitempty"`
	// Frequency channel used for wireless communication.
	Channel *int32 `json:"channel,omitempty"`
	// Width of frequency channel used for wireless communication.
	ChannelWidth *int32 `json:"channelWidth,omitempty"`
	// Strength of wireless signal.
	Power *int32 `json:"power,omitempty"`
	// Indicates whether or not this radio is currently broadcasting.
	IsBroadcasting *bool `json:"isBroadcasting,omitempty"`
	// The radio index.
	Index *string `json:"index,omitempty"`
}

// NewInlineResponse200290Radio instantiates a new InlineResponse200290Radio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200290Radio() *InlineResponse200290Radio {
	this := InlineResponse200290Radio{}
	return &this
}

// NewInlineResponse200290RadioWithDefaults instantiates a new InlineResponse200290Radio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200290RadioWithDefaults() *InlineResponse200290Radio {
	this := InlineResponse200290Radio{}
	return &this
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *InlineResponse200290Radio) GetBand() string {
	if o == nil || isNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Radio) GetBandOk() (*string, bool) {
	if o == nil || isNil(o.Band) {
    return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *InlineResponse200290Radio) HasBand() bool {
	if o != nil && !isNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *InlineResponse200290Radio) SetBand(v string) {
	o.Band = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineResponse200290Radio) GetChannel() int32 {
	if o == nil || isNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Radio) GetChannelOk() (*int32, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineResponse200290Radio) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *InlineResponse200290Radio) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *InlineResponse200290Radio) GetChannelWidth() int32 {
	if o == nil || isNil(o.ChannelWidth) {
		var ret int32
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Radio) GetChannelWidthOk() (*int32, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *InlineResponse200290Radio) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given int32 and assigns it to the ChannelWidth field.
func (o *InlineResponse200290Radio) SetChannelWidth(v int32) {
	o.ChannelWidth = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *InlineResponse200290Radio) GetPower() int32 {
	if o == nil || isNil(o.Power) {
		var ret int32
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Radio) GetPowerOk() (*int32, bool) {
	if o == nil || isNil(o.Power) {
    return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *InlineResponse200290Radio) HasPower() bool {
	if o != nil && !isNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given int32 and assigns it to the Power field.
func (o *InlineResponse200290Radio) SetPower(v int32) {
	o.Power = &v
}

// GetIsBroadcasting returns the IsBroadcasting field value if set, zero value otherwise.
func (o *InlineResponse200290Radio) GetIsBroadcasting() bool {
	if o == nil || isNil(o.IsBroadcasting) {
		var ret bool
		return ret
	}
	return *o.IsBroadcasting
}

// GetIsBroadcastingOk returns a tuple with the IsBroadcasting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Radio) GetIsBroadcastingOk() (*bool, bool) {
	if o == nil || isNil(o.IsBroadcasting) {
    return nil, false
	}
	return o.IsBroadcasting, true
}

// HasIsBroadcasting returns a boolean if a field has been set.
func (o *InlineResponse200290Radio) HasIsBroadcasting() bool {
	if o != nil && !isNil(o.IsBroadcasting) {
		return true
	}

	return false
}

// SetIsBroadcasting gets a reference to the given bool and assigns it to the IsBroadcasting field.
func (o *InlineResponse200290Radio) SetIsBroadcasting(v bool) {
	o.IsBroadcasting = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *InlineResponse200290Radio) GetIndex() string {
	if o == nil || isNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200290Radio) GetIndexOk() (*string, bool) {
	if o == nil || isNil(o.Index) {
    return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *InlineResponse200290Radio) HasIndex() bool {
	if o != nil && !isNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *InlineResponse200290Radio) SetIndex(v string) {
	o.Index = &v
}

func (o InlineResponse200290Radio) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Band) {
		toSerialize["band"] = o.Band
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
	if !isNil(o.IsBroadcasting) {
		toSerialize["isBroadcasting"] = o.IsBroadcasting
	}
	if !isNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200290Radio struct {
	value *InlineResponse200290Radio
	isSet bool
}

func (v NullableInlineResponse200290Radio) Get() *InlineResponse200290Radio {
	return v.value
}

func (v *NullableInlineResponse200290Radio) Set(val *InlineResponse200290Radio) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200290Radio) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200290Radio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200290Radio(val *InlineResponse200290Radio) *NullableInlineResponse200290Radio {
	return &NullableInlineResponse200290Radio{value: val, isSet: true}
}

func (v NullableInlineResponse200290Radio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200290Radio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


