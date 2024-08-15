/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel struct for NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel
type NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel struct {
	// Model of the AP
	Model *string `json:"model,omitempty"`
	// Band to use for each flex radio. For example, ['6'] will set the AP's first flex radio to 6 GHz
	Bands []string `json:"bands,omitempty"`
}

// NewNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel instantiates a new NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel() *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel {
	this := NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel{}
	return &this
}

// NewNetworksNetworkIdWirelessRfProfilesFlexRadiosByModelWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessRfProfilesFlexRadiosByModelWithDefaults() *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel {
	this := NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel{}
	return &this
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) SetModel(v string) {
	o.Model = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) GetBands() []string {
	if o == nil || isNil(o.Bands) {
		var ret []string
		return ret
	}
	return o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) GetBandsOk() ([]string, bool) {
	if o == nil || isNil(o.Bands) {
    return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) HasBands() bool {
	if o != nil && !isNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given []string and assigns it to the Bands field.
func (o *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) SetBands(v []string) {
	o.Bands = v
}

func (o NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel struct {
	value *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) Get() *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) Set(val *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel(val *NetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) *NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel {
	return &NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessRfProfilesFlexRadiosByModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


