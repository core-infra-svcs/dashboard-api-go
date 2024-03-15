/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 March, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.44.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdSettingsLocalStatusPage A hash of Local Status page(s)' authentication options applied to the Network.
type NetworksNetworkIdSettingsLocalStatusPage struct {
	Authentication *NetworksNetworkIdSettingsLocalStatusPageAuthentication `json:"authentication,omitempty"`
}

// NewNetworksNetworkIdSettingsLocalStatusPage instantiates a new NetworksNetworkIdSettingsLocalStatusPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSettingsLocalStatusPage() *NetworksNetworkIdSettingsLocalStatusPage {
	this := NetworksNetworkIdSettingsLocalStatusPage{}
	return &this
}

// NewNetworksNetworkIdSettingsLocalStatusPageWithDefaults instantiates a new NetworksNetworkIdSettingsLocalStatusPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSettingsLocalStatusPageWithDefaults() *NetworksNetworkIdSettingsLocalStatusPage {
	this := NetworksNetworkIdSettingsLocalStatusPage{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *NetworksNetworkIdSettingsLocalStatusPage) GetAuthentication() NetworksNetworkIdSettingsLocalStatusPageAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret NetworksNetworkIdSettingsLocalStatusPageAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSettingsLocalStatusPage) GetAuthenticationOk() (*NetworksNetworkIdSettingsLocalStatusPageAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *NetworksNetworkIdSettingsLocalStatusPage) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NetworksNetworkIdSettingsLocalStatusPageAuthentication and assigns it to the Authentication field.
func (o *NetworksNetworkIdSettingsLocalStatusPage) SetAuthentication(v NetworksNetworkIdSettingsLocalStatusPageAuthentication) {
	o.Authentication = &v
}

func (o NetworksNetworkIdSettingsLocalStatusPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSettingsLocalStatusPage struct {
	value *NetworksNetworkIdSettingsLocalStatusPage
	isSet bool
}

func (v NullableNetworksNetworkIdSettingsLocalStatusPage) Get() *NetworksNetworkIdSettingsLocalStatusPage {
	return v.value
}

func (v *NullableNetworksNetworkIdSettingsLocalStatusPage) Set(val *NetworksNetworkIdSettingsLocalStatusPage) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSettingsLocalStatusPage) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSettingsLocalStatusPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSettingsLocalStatusPage(val *NetworksNetworkIdSettingsLocalStatusPage) *NullableNetworksNetworkIdSettingsLocalStatusPage {
	return &NullableNetworksNetworkIdSettingsLocalStatusPage{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSettingsLocalStatusPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSettingsLocalStatusPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


