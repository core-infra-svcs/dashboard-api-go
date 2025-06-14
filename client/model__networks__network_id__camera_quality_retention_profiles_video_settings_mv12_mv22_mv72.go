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

// NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 Quality and resolution for MV12/MV22/MV72 camera models.
type NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
	Resolution string `json:"resolution"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72(quality string, resolution string) *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72WithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72WithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72{}
	return &this
}

// GetQuality returns the Quality field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) SetResolution(v string) {
	o.Resolution = v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) Get() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


