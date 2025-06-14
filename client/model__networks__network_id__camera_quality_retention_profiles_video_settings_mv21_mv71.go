/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 Quality and resolution for MV21/MV71 camera models.
type NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720'.
	Resolution string `json:"resolution"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71(quality string, resolution string) *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71WithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71WithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71{}
	return &this
}

// GetQuality returns the Quality field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) SetResolution(v string) {
	o.Resolution = v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) Get() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71 {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV21MV71) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


