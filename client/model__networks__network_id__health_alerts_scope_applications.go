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

// NetworksNetworkIdHealthAlertsScopeApplications struct for NetworksNetworkIdHealthAlertsScopeApplications
type NetworksNetworkIdHealthAlertsScopeApplications struct {
	// URL to the application
	Url *string `json:"url,omitempty"`
	// Name of the application
	Name *string `json:"name,omitempty"`
}

// NewNetworksNetworkIdHealthAlertsScopeApplications instantiates a new NetworksNetworkIdHealthAlertsScopeApplications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdHealthAlertsScopeApplications() *NetworksNetworkIdHealthAlertsScopeApplications {
	this := NetworksNetworkIdHealthAlertsScopeApplications{}
	return &this
}

// NewNetworksNetworkIdHealthAlertsScopeApplicationsWithDefaults instantiates a new NetworksNetworkIdHealthAlertsScopeApplications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdHealthAlertsScopeApplicationsWithDefaults() *NetworksNetworkIdHealthAlertsScopeApplications {
	this := NetworksNetworkIdHealthAlertsScopeApplications{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdHealthAlertsScopeApplications) SetName(v string) {
	o.Name = &v
}

func (o NetworksNetworkIdHealthAlertsScopeApplications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdHealthAlertsScopeApplications struct {
	value *NetworksNetworkIdHealthAlertsScopeApplications
	isSet bool
}

func (v NullableNetworksNetworkIdHealthAlertsScopeApplications) Get() *NetworksNetworkIdHealthAlertsScopeApplications {
	return v.value
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeApplications) Set(val *NetworksNetworkIdHealthAlertsScopeApplications) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdHealthAlertsScopeApplications) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeApplications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdHealthAlertsScopeApplications(val *NetworksNetworkIdHealthAlertsScopeApplications) *NullableNetworksNetworkIdHealthAlertsScopeApplications {
	return &NullableNetworksNetworkIdHealthAlertsScopeApplications{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdHealthAlertsScopeApplications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdHealthAlertsScopeApplications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


