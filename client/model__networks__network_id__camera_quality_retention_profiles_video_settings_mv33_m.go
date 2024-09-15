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

// NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M Quality and resolution for MV33M camera models.
type NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M struct {
	// Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
	Quality string `json:"quality"`
	// Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
	Resolution string `json:"resolution"`
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M(quality string, resolution string) *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M{}
	this.Quality = quality
	this.Resolution = resolution
	return &this
}

// NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33MWithDefaults instantiates a new NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33MWithDefaults() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M {
	this := NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M{}
	return &this
}

// GetQuality returns the Quality field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) GetQuality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quality
}

// GetQualityOk returns a tuple with the Quality field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) GetQualityOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Quality, true
}

// SetQuality sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) SetQuality(v string) {
	o.Quality = v
}

// GetResolution returns the Resolution field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) GetResolutionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) SetResolution(v string) {
	o.Resolution = v
}

func (o NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["quality"] = o.Quality
	}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M struct {
	value *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M
	isSet bool
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) Get() *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M {
	return v.value
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) Set(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M(val *NetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M {
	return &NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdCameraQualityRetentionProfilesVideoSettingsMV33M) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


