/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 May, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.33.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 Quality and resolution for MV52 camera models.
type NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
	Resolution string `json:"resolution"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52(quality string, resolution string) *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52WithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52WithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52{}
	return &this
}

// GetQuality returns the Quality field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) SetResolution(v string) {
	o.Resolution = v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) Get() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52 {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV52) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


