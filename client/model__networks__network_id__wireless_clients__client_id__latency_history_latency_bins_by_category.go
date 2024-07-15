/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory The latency buckets by category
type NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory struct {
	BackgroundTraffic *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBackgroundTraffic `json:"backgroundTraffic,omitempty"`
	BestEffortTraffic *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBestEffortTraffic `json:"bestEffortTraffic,omitempty"`
	VideoTraffic *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic `json:"videoTraffic,omitempty"`
	VoiceTraffic *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVoiceTraffic `json:"voiceTraffic,omitempty"`
}

// NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory instantiates a new NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory() *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory {
	this := NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory{}
	return &this
}

// NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryWithDefaults instantiates a new NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryWithDefaults() *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory {
	this := NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory{}
	return &this
}

// GetBackgroundTraffic returns the BackgroundTraffic field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetBackgroundTraffic() NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBackgroundTraffic {
	if o == nil || isNil(o.BackgroundTraffic) {
		var ret NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBackgroundTraffic
		return ret
	}
	return *o.BackgroundTraffic
}

// GetBackgroundTrafficOk returns a tuple with the BackgroundTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetBackgroundTrafficOk() (*NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBackgroundTraffic, bool) {
	if o == nil || isNil(o.BackgroundTraffic) {
    return nil, false
	}
	return o.BackgroundTraffic, true
}

// HasBackgroundTraffic returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) HasBackgroundTraffic() bool {
	if o != nil && !isNil(o.BackgroundTraffic) {
		return true
	}

	return false
}

// SetBackgroundTraffic gets a reference to the given NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBackgroundTraffic and assigns it to the BackgroundTraffic field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) SetBackgroundTraffic(v NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBackgroundTraffic) {
	o.BackgroundTraffic = &v
}

// GetBestEffortTraffic returns the BestEffortTraffic field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetBestEffortTraffic() NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBestEffortTraffic {
	if o == nil || isNil(o.BestEffortTraffic) {
		var ret NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBestEffortTraffic
		return ret
	}
	return *o.BestEffortTraffic
}

// GetBestEffortTrafficOk returns a tuple with the BestEffortTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetBestEffortTrafficOk() (*NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBestEffortTraffic, bool) {
	if o == nil || isNil(o.BestEffortTraffic) {
    return nil, false
	}
	return o.BestEffortTraffic, true
}

// HasBestEffortTraffic returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) HasBestEffortTraffic() bool {
	if o != nil && !isNil(o.BestEffortTraffic) {
		return true
	}

	return false
}

// SetBestEffortTraffic gets a reference to the given NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBestEffortTraffic and assigns it to the BestEffortTraffic field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) SetBestEffortTraffic(v NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryBestEffortTraffic) {
	o.BestEffortTraffic = &v
}

// GetVideoTraffic returns the VideoTraffic field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetVideoTraffic() NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic {
	if o == nil || isNil(o.VideoTraffic) {
		var ret NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic
		return ret
	}
	return *o.VideoTraffic
}

// GetVideoTrafficOk returns a tuple with the VideoTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetVideoTrafficOk() (*NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic, bool) {
	if o == nil || isNil(o.VideoTraffic) {
    return nil, false
	}
	return o.VideoTraffic, true
}

// HasVideoTraffic returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) HasVideoTraffic() bool {
	if o != nil && !isNil(o.VideoTraffic) {
		return true
	}

	return false
}

// SetVideoTraffic gets a reference to the given NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic and assigns it to the VideoTraffic field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) SetVideoTraffic(v NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVideoTraffic) {
	o.VideoTraffic = &v
}

// GetVoiceTraffic returns the VoiceTraffic field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetVoiceTraffic() NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVoiceTraffic {
	if o == nil || isNil(o.VoiceTraffic) {
		var ret NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVoiceTraffic
		return ret
	}
	return *o.VoiceTraffic
}

// GetVoiceTrafficOk returns a tuple with the VoiceTraffic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) GetVoiceTrafficOk() (*NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVoiceTraffic, bool) {
	if o == nil || isNil(o.VoiceTraffic) {
    return nil, false
	}
	return o.VoiceTraffic, true
}

// HasVoiceTraffic returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) HasVoiceTraffic() bool {
	if o != nil && !isNil(o.VoiceTraffic) {
		return true
	}

	return false
}

// SetVoiceTraffic gets a reference to the given NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVoiceTraffic and assigns it to the VoiceTraffic field.
func (o *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) SetVoiceTraffic(v NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategoryVoiceTraffic) {
	o.VoiceTraffic = &v
}

func (o NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BackgroundTraffic) {
		toSerialize["backgroundTraffic"] = o.BackgroundTraffic
	}
	if !isNil(o.BestEffortTraffic) {
		toSerialize["bestEffortTraffic"] = o.BestEffortTraffic
	}
	if !isNil(o.VideoTraffic) {
		toSerialize["videoTraffic"] = o.VideoTraffic
	}
	if !isNil(o.VoiceTraffic) {
		toSerialize["voiceTraffic"] = o.VoiceTraffic
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory struct {
	value *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) Get() *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) Set(val *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory(val *NetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory {
	return &NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessClientsClientIdLatencyHistoryLatencyBinsByCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

