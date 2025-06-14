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

// InlineObject302 struct for InlineObject302
type InlineObject302 struct {
	Network OrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork `json:"network"`
	// Receiver Url
	Url string `json:"url"`
	// Scanning API Version
	Version string `json:"version"`
	Radio OrganizationsOrganizationIdWirelessLocationScanningReceiversRadio `json:"radio"`
	// Secret Value for Receiver
	SharedSecret string `json:"sharedSecret"`
}

// NewInlineObject302 instantiates a new InlineObject302 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject302(network OrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork, url string, version string, radio OrganizationsOrganizationIdWirelessLocationScanningReceiversRadio, sharedSecret string) *InlineObject302 {
	this := InlineObject302{}
	this.Network = network
	this.Url = url
	this.Version = version
	this.Radio = radio
	this.SharedSecret = sharedSecret
	return &this
}

// NewInlineObject302WithDefaults instantiates a new InlineObject302 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject302WithDefaults() *InlineObject302 {
	this := InlineObject302{}
	return &this
}

// GetNetwork returns the Network field value
func (o *InlineObject302) GetNetwork() OrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork {
	if o == nil {
		var ret OrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *InlineObject302) GetNetworkOk() (*OrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *InlineObject302) SetNetwork(v OrganizationsOrganizationIdWirelessLocationScanningReceiversNetwork) {
	o.Network = v
}

// GetUrl returns the Url field value
func (o *InlineObject302) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *InlineObject302) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *InlineObject302) SetUrl(v string) {
	o.Url = v
}

// GetVersion returns the Version field value
func (o *InlineObject302) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *InlineObject302) GetVersionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *InlineObject302) SetVersion(v string) {
	o.Version = v
}

// GetRadio returns the Radio field value
func (o *InlineObject302) GetRadio() OrganizationsOrganizationIdWirelessLocationScanningReceiversRadio {
	if o == nil {
		var ret OrganizationsOrganizationIdWirelessLocationScanningReceiversRadio
		return ret
	}

	return o.Radio
}

// GetRadioOk returns a tuple with the Radio field value
// and a boolean to check if the value has been set.
func (o *InlineObject302) GetRadioOk() (*OrganizationsOrganizationIdWirelessLocationScanningReceiversRadio, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Radio, true
}

// SetRadio sets field value
func (o *InlineObject302) SetRadio(v OrganizationsOrganizationIdWirelessLocationScanningReceiversRadio) {
	o.Radio = v
}

// GetSharedSecret returns the SharedSecret field value
func (o *InlineObject302) GetSharedSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SharedSecret
}

// GetSharedSecretOk returns a tuple with the SharedSecret field value
// and a boolean to check if the value has been set.
func (o *InlineObject302) GetSharedSecretOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SharedSecret, true
}

// SetSharedSecret sets field value
func (o *InlineObject302) SetSharedSecret(v string) {
	o.SharedSecret = v
}

func (o InlineObject302) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["radio"] = o.Radio
	}
	if true {
		toSerialize["sharedSecret"] = o.SharedSecret
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject302 struct {
	value *InlineObject302
	isSet bool
}

func (v NullableInlineObject302) Get() *InlineObject302 {
	return v.value
}

func (v *NullableInlineObject302) Set(val *InlineObject302) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject302) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject302) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject302(val *InlineObject302) *NullableInlineObject302 {
	return &NullableInlineObject302{value: val, isSet: true}
}

func (v NullableInlineObject302) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject302) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


