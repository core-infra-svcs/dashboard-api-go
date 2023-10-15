/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings Video quality and resolution settings for all the camera models.
type NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings struct {
	MV21MV71 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 `json:"MV21/MV71,omitempty"`
	MV12MV22MV72 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 `json:"MV12/MV22/MV72,omitempty"`
	MV32 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV32 `json:"MV32,omitempty"`
	MV33 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33 `json:"MV33,omitempty"`
	MV12WE *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12WE `json:"MV12WE,omitempty"`
	MV13 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV13 `json:"MV13,omitempty"`
	MV22XMV72X *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV22XMV72X `json:"MV22X/MV72X,omitempty"`
	MV52 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 `json:"MV52,omitempty"`
	MV63 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63 `json:"MV63,omitempty"`
	MV93 *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93 `json:"MV93,omitempty"`
	MV63X *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63X `json:"MV63X,omitempty"`
	MV93X *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93X `json:"MV93X,omitempty"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings{}
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsWithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsWithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings{}
	return &this
}

// GetMV21MV71 returns the MV21MV71 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV21MV71() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 {
	if o == nil || isNil(o.MV21MV71) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71
		return ret
	}
	return *o.MV21MV71
}

// GetMV21MV71Ok returns a tuple with the MV21MV71 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV21MV71Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71, bool) {
	if o == nil || isNil(o.MV21MV71) {
    return nil, false
	}
	return o.MV21MV71, true
}

// HasMV21MV71 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV21MV71() bool {
	if o != nil && !isNil(o.MV21MV71) {
		return true
	}

	return false
}

// SetMV21MV71 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 and assigns it to the MV21MV71 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV21MV71(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) {
	o.MV21MV71 = &v
}

// GetMV12MV22MV72 returns the MV12MV22MV72 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV12MV22MV72() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 {
	if o == nil || isNil(o.MV12MV22MV72) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72
		return ret
	}
	return *o.MV12MV22MV72
}

// GetMV12MV22MV72Ok returns a tuple with the MV12MV22MV72 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV12MV22MV72Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72, bool) {
	if o == nil || isNil(o.MV12MV22MV72) {
    return nil, false
	}
	return o.MV12MV22MV72, true
}

// HasMV12MV22MV72 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV12MV22MV72() bool {
	if o != nil && !isNil(o.MV12MV22MV72) {
		return true
	}

	return false
}

// SetMV12MV22MV72 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 and assigns it to the MV12MV22MV72 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV12MV22MV72(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) {
	o.MV12MV22MV72 = &v
}

// GetMV32 returns the MV32 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV32() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV32 {
	if o == nil || isNil(o.MV32) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV32
		return ret
	}
	return *o.MV32
}

// GetMV32Ok returns a tuple with the MV32 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV32Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV32, bool) {
	if o == nil || isNil(o.MV32) {
    return nil, false
	}
	return o.MV32, true
}

// HasMV32 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV32() bool {
	if o != nil && !isNil(o.MV32) {
		return true
	}

	return false
}

// SetMV32 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV32 and assigns it to the MV32 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV32(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV32) {
	o.MV32 = &v
}

// GetMV33 returns the MV33 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV33() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33 {
	if o == nil || isNil(o.MV33) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33
		return ret
	}
	return *o.MV33
}

// GetMV33Ok returns a tuple with the MV33 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV33Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33, bool) {
	if o == nil || isNil(o.MV33) {
    return nil, false
	}
	return o.MV33, true
}

// HasMV33 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV33() bool {
	if o != nil && !isNil(o.MV33) {
		return true
	}

	return false
}

// SetMV33 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33 and assigns it to the MV33 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV33(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33) {
	o.MV33 = &v
}

// GetMV12WE returns the MV12WE field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV12WE() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12WE {
	if o == nil || isNil(o.MV12WE) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12WE
		return ret
	}
	return *o.MV12WE
}

// GetMV12WEOk returns a tuple with the MV12WE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV12WEOk() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12WE, bool) {
	if o == nil || isNil(o.MV12WE) {
    return nil, false
	}
	return o.MV12WE, true
}

// HasMV12WE returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV12WE() bool {
	if o != nil && !isNil(o.MV12WE) {
		return true
	}

	return false
}

// SetMV12WE gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12WE and assigns it to the MV12WE field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV12WE(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12WE) {
	o.MV12WE = &v
}

// GetMV13 returns the MV13 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV13() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV13 {
	if o == nil || isNil(o.MV13) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV13
		return ret
	}
	return *o.MV13
}

// GetMV13Ok returns a tuple with the MV13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV13Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV13, bool) {
	if o == nil || isNil(o.MV13) {
    return nil, false
	}
	return o.MV13, true
}

// HasMV13 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV13() bool {
	if o != nil && !isNil(o.MV13) {
		return true
	}

	return false
}

// SetMV13 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV13 and assigns it to the MV13 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV13(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV13) {
	o.MV13 = &v
}

// GetMV22XMV72X returns the MV22XMV72X field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV22XMV72X() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV22XMV72X {
	if o == nil || isNil(o.MV22XMV72X) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV22XMV72X
		return ret
	}
	return *o.MV22XMV72X
}

// GetMV22XMV72XOk returns a tuple with the MV22XMV72X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV22XMV72XOk() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV22XMV72X, bool) {
	if o == nil || isNil(o.MV22XMV72X) {
    return nil, false
	}
	return o.MV22XMV72X, true
}

// HasMV22XMV72X returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV22XMV72X() bool {
	if o != nil && !isNil(o.MV22XMV72X) {
		return true
	}

	return false
}

// SetMV22XMV72X gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV22XMV72X and assigns it to the MV22XMV72X field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV22XMV72X(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV22XMV72X) {
	o.MV22XMV72X = &v
}

// GetMV52 returns the MV52 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV52() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 {
	if o == nil || isNil(o.MV52) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52
		return ret
	}
	return *o.MV52
}

// GetMV52Ok returns a tuple with the MV52 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV52Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52, bool) {
	if o == nil || isNil(o.MV52) {
    return nil, false
	}
	return o.MV52, true
}

// HasMV52 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV52() bool {
	if o != nil && !isNil(o.MV52) {
		return true
	}

	return false
}

// SetMV52 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 and assigns it to the MV52 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV52(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) {
	o.MV52 = &v
}

// GetMV63 returns the MV63 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV63() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63 {
	if o == nil || isNil(o.MV63) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63
		return ret
	}
	return *o.MV63
}

// GetMV63Ok returns a tuple with the MV63 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV63Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63, bool) {
	if o == nil || isNil(o.MV63) {
    return nil, false
	}
	return o.MV63, true
}

// HasMV63 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV63() bool {
	if o != nil && !isNil(o.MV63) {
		return true
	}

	return false
}

// SetMV63 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63 and assigns it to the MV63 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV63(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63) {
	o.MV63 = &v
}

// GetMV93 returns the MV93 field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV93() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93 {
	if o == nil || isNil(o.MV93) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93
		return ret
	}
	return *o.MV93
}

// GetMV93Ok returns a tuple with the MV93 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV93Ok() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93, bool) {
	if o == nil || isNil(o.MV93) {
    return nil, false
	}
	return o.MV93, true
}

// HasMV93 returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV93() bool {
	if o != nil && !isNil(o.MV93) {
		return true
	}

	return false
}

// SetMV93 gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93 and assigns it to the MV93 field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV93(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93) {
	o.MV93 = &v
}

// GetMV63X returns the MV63X field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV63X() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63X {
	if o == nil || isNil(o.MV63X) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63X
		return ret
	}
	return *o.MV63X
}

// GetMV63XOk returns a tuple with the MV63X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV63XOk() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63X, bool) {
	if o == nil || isNil(o.MV63X) {
    return nil, false
	}
	return o.MV63X, true
}

// HasMV63X returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV63X() bool {
	if o != nil && !isNil(o.MV63X) {
		return true
	}

	return false
}

// SetMV63X gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63X and assigns it to the MV63X field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV63X(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV63X) {
	o.MV63X = &v
}

// GetMV93X returns the MV93X field value if set, zero value otherwise.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV93X() NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93X {
	if o == nil || isNil(o.MV93X) {
		var ret NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93X
		return ret
	}
	return *o.MV93X
}

// GetMV93XOk returns a tuple with the MV93X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) GetMV93XOk() (*NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93X, bool) {
	if o == nil || isNil(o.MV93X) {
    return nil, false
	}
	return o.MV93X, true
}

// HasMV93X returns a boolean if a field has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) HasMV93X() bool {
	if o != nil && !isNil(o.MV93X) {
		return true
	}

	return false
}

// SetMV93X gets a reference to the given NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93X and assigns it to the MV93X field.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) SetMV93X(v NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV93X) {
	o.MV93X = &v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MV21MV71) {
		toSerialize["MV21/MV71"] = o.MV21MV71
	}
	if !isNil(o.MV12MV22MV72) {
		toSerialize["MV12/MV22/MV72"] = o.MV12MV22MV72
	}
	if !isNil(o.MV32) {
		toSerialize["MV32"] = o.MV32
	}
	if !isNil(o.MV33) {
		toSerialize["MV33"] = o.MV33
	}
	if !isNil(o.MV12WE) {
		toSerialize["MV12WE"] = o.MV12WE
	}
	if !isNil(o.MV13) {
		toSerialize["MV13"] = o.MV13
	}
	if !isNil(o.MV22XMV72X) {
		toSerialize["MV22X/MV72X"] = o.MV22XMV72X
	}
	if !isNil(o.MV52) {
		toSerialize["MV52"] = o.MV52
	}
	if !isNil(o.MV63) {
		toSerialize["MV63"] = o.MV63
	}
	if !isNil(o.MV93) {
		toSerialize["MV93"] = o.MV93
	}
	if !isNil(o.MV63X) {
		toSerialize["MV63X"] = o.MV63X
	}
	if !isNil(o.MV93X) {
		toSerialize["MV93X"] = o.MV93X
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) Get() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


