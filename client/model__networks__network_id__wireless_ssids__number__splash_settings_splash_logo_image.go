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

// NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage Properties for setting a new image.
type NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage struct {
	// The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
	Format *string `json:"format,omitempty"`
	// The file contents (a base 64 encoded string) of your new logo.
	Contents *string `json:"contents,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImageWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImageWithDefaults() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage {
	this := NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
    return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) SetFormat(v string) {
	o.Format = &v
}

// GetContents returns the Contents field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) GetContents() string {
	if o == nil || isNil(o.Contents) {
		var ret string
		return ret
	}
	return *o.Contents
}

// GetContentsOk returns a tuple with the Contents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) GetContentsOk() (*string, bool) {
	if o == nil || isNil(o.Contents) {
    return nil, false
	}
	return o.Contents, true
}

// HasContents returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) HasContents() bool {
	if o != nil && !isNil(o.Contents) {
		return true
	}

	return false
}

// SetContents gets a reference to the given string and assigns it to the Contents field.
func (o *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) SetContents(v string) {
	o.Contents = &v
}

func (o NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.Contents) {
		toSerialize["contents"] = o.Contents
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage struct {
	value *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) Get() *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) Set(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage(val *NetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage {
	return &NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberSplashSettingsSplashLogoImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


