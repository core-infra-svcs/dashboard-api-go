/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 Quality and resolution for MV23 camera models.
type NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
	Resolution string `json:"resolution"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23(quality string, resolution string) *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23WithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23WithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23{}
	return &this
}

// GetQuality returns the Quality field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) SetResolution(v string) {
	o.Resolution = v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) Get() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23 {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV23) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


