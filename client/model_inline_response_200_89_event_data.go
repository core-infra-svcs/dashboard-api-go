/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20089EventData An object containing more data related to the event.
type InlineResponse20089EventData struct {
	// The radio band number the client is trying to connect to.
	Radio *string `json:"radio,omitempty"`
	// The virtual access point (VAP) number the client is connecting to.
	Vap *string `json:"vap,omitempty"`
	// The client's MAC address
	ClientMac *string `json:"client_mac,omitempty"`
	// The client's IP address
	ClientIp *string `json:"client_ip,omitempty"`
	// The radio channel the client is connecting to.
	Channel *string `json:"channel,omitempty"`
	// The current received signal strength indication (RSSI) of the client connected to an AP.
	Rssi *string `json:"rssi,omitempty"`
	// The association ID of the client.
	Aid *string `json:"aid,omitempty"`
}

// NewInlineResponse20089EventData instantiates a new InlineResponse20089EventData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20089EventData() *InlineResponse20089EventData {
	this := InlineResponse20089EventData{}
	return &this
}

// NewInlineResponse20089EventDataWithDefaults instantiates a new InlineResponse20089EventData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20089EventDataWithDefaults() *InlineResponse20089EventData {
	this := InlineResponse20089EventData{}
	return &this
}

// GetRadio returns the Radio field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetRadio() string {
	if o == nil || isNil(o.Radio) {
		var ret string
		return ret
	}
	return *o.Radio
}

// GetRadioOk returns a tuple with the Radio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetRadioOk() (*string, bool) {
	if o == nil || isNil(o.Radio) {
    return nil, false
	}
	return o.Radio, true
}

// HasRadio returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasRadio() bool {
	if o != nil && !isNil(o.Radio) {
		return true
	}

	return false
}

// SetRadio gets a reference to the given string and assigns it to the Radio field.
func (o *InlineResponse20089EventData) SetRadio(v string) {
	o.Radio = &v
}

// GetVap returns the Vap field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetVap() string {
	if o == nil || isNil(o.Vap) {
		var ret string
		return ret
	}
	return *o.Vap
}

// GetVapOk returns a tuple with the Vap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetVapOk() (*string, bool) {
	if o == nil || isNil(o.Vap) {
    return nil, false
	}
	return o.Vap, true
}

// HasVap returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasVap() bool {
	if o != nil && !isNil(o.Vap) {
		return true
	}

	return false
}

// SetVap gets a reference to the given string and assigns it to the Vap field.
func (o *InlineResponse20089EventData) SetVap(v string) {
	o.Vap = &v
}

// GetClientMac returns the ClientMac field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetClientMac() string {
	if o == nil || isNil(o.ClientMac) {
		var ret string
		return ret
	}
	return *o.ClientMac
}

// GetClientMacOk returns a tuple with the ClientMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetClientMacOk() (*string, bool) {
	if o == nil || isNil(o.ClientMac) {
    return nil, false
	}
	return o.ClientMac, true
}

// HasClientMac returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasClientMac() bool {
	if o != nil && !isNil(o.ClientMac) {
		return true
	}

	return false
}

// SetClientMac gets a reference to the given string and assigns it to the ClientMac field.
func (o *InlineResponse20089EventData) SetClientMac(v string) {
	o.ClientMac = &v
}

// GetClientIp returns the ClientIp field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetClientIp() string {
	if o == nil || isNil(o.ClientIp) {
		var ret string
		return ret
	}
	return *o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetClientIpOk() (*string, bool) {
	if o == nil || isNil(o.ClientIp) {
    return nil, false
	}
	return o.ClientIp, true
}

// HasClientIp returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasClientIp() bool {
	if o != nil && !isNil(o.ClientIp) {
		return true
	}

	return false
}

// SetClientIp gets a reference to the given string and assigns it to the ClientIp field.
func (o *InlineResponse20089EventData) SetClientIp(v string) {
	o.ClientIp = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetChannel() string {
	if o == nil || isNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetChannelOk() (*string, bool) {
	if o == nil || isNil(o.Channel) {
    return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasChannel() bool {
	if o != nil && !isNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *InlineResponse20089EventData) SetChannel(v string) {
	o.Channel = &v
}

// GetRssi returns the Rssi field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetRssi() string {
	if o == nil || isNil(o.Rssi) {
		var ret string
		return ret
	}
	return *o.Rssi
}

// GetRssiOk returns a tuple with the Rssi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetRssiOk() (*string, bool) {
	if o == nil || isNil(o.Rssi) {
    return nil, false
	}
	return o.Rssi, true
}

// HasRssi returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasRssi() bool {
	if o != nil && !isNil(o.Rssi) {
		return true
	}

	return false
}

// SetRssi gets a reference to the given string and assigns it to the Rssi field.
func (o *InlineResponse20089EventData) SetRssi(v string) {
	o.Rssi = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *InlineResponse20089EventData) GetAid() string {
	if o == nil || isNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20089EventData) GetAidOk() (*string, bool) {
	if o == nil || isNil(o.Aid) {
    return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *InlineResponse20089EventData) HasAid() bool {
	if o != nil && !isNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *InlineResponse20089EventData) SetAid(v string) {
	o.Aid = &v
}

func (o InlineResponse20089EventData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Radio) {
		toSerialize["radio"] = o.Radio
	}
	if !isNil(o.Vap) {
		toSerialize["vap"] = o.Vap
	}
	if !isNil(o.ClientMac) {
		toSerialize["client_mac"] = o.ClientMac
	}
	if !isNil(o.ClientIp) {
		toSerialize["client_ip"] = o.ClientIp
	}
	if !isNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !isNil(o.Rssi) {
		toSerialize["rssi"] = o.Rssi
	}
	if !isNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20089EventData struct {
	value *InlineResponse20089EventData
	isSet bool
}

func (v NullableInlineResponse20089EventData) Get() *InlineResponse20089EventData {
	return v.value
}

func (v *NullableInlineResponse20089EventData) Set(val *InlineResponse20089EventData) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20089EventData) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20089EventData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20089EventData(val *InlineResponse20089EventData) *NullableInlineResponse20089EventData {
	return &NullableInlineResponse20089EventData{value: val, isSet: true}
}

func (v NullableInlineResponse20089EventData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20089EventData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


