/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage Properties for setting a new image.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage struct {
	// The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
	// The file contents (a base 64 encoded string) of your new image.
	Contents *string `json:"contents,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImageWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImageWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
    return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) SetFormat(v string) {
	o.Format = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) GetContents() string {
	if o == nil || isNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) GetContentsOk() (*string, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) SetContents(v string) {
	o.Contents = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashImageImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


