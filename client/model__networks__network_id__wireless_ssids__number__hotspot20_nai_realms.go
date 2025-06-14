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

// NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms struct for NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms
type NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms struct {
	// The format for the realm ('1' or '0')
	Format *string `json:"format,omitempty"`
	// The name of the realm
	Realm *string `json:"realm,omitempty"`
	// An array of EAP methods for the realm.
	Methods []NetworksNetworkIdWirelessSsidsNumberHotspot20Methods `json:"methods,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms() *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealmsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealmsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) GetFormat() string {
	if o == nil || isNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) GetFormatOk() (*string, bool) {
	if o == nil || isNil(o.Format) {
    return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) HasFormat() bool {
	if o != nil && !isNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) SetFormat(v string) {
	o.Format = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) GetRealm() string {
	if o == nil || isNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) GetRealmOk() (*string, bool) {
	if o == nil || isNil(o.Realm) {
    return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) HasRealm() bool {
	if o != nil && !isNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) SetRealm(v string) {
	o.Realm = &v
}

// GetMethods returns the Methods field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) GetMethods() []NetworksNetworkIdWirelessSsidsNumberHotspot20Methods {
	if o == nil || isNil(o.Methods) {
		var ret []NetworksNetworkIdWirelessSsidsNumberHotspot20Methods
		return ret
	}
	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) GetMethodsOk() ([]NetworksNetworkIdWirelessSsidsNumberHotspot20Methods, bool) {
	if o == nil || isNil(o.Methods) {
    return nil, false
	}
	return o.Methods, true
}

// HasMethods returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) HasMethods() bool {
	if o != nil && !isNil(o.Methods) {
		return true
	}

	return false
}

// SetMethods gets a reference to the given []NetworksNetworkIdWirelessSsidsNumberHotspot20Methods and assigns it to the Methods field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) SetMethods(v []NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) {
	o.Methods = v
}

func (o NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !isNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !isNil(o.Methods) {
		toSerialize["methods"] = o.Methods
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms struct {
	value *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) Get() *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) Set(val *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms(val *NetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms {
	return &NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20NaiRealms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


