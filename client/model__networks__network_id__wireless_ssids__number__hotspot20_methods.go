/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 May, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.21.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdWirelessSsidsNumberHotspot20Methods struct for NetworksNetworkIdWirelessSsidsNumberHotspot20Methods
type NetworksNetworkIdWirelessSsidsNumberHotspot20Methods struct {
	// ID of method
	Id *string `json:"id,omitempty"`
	// The authentication types for the method. These should be formatted as an object with the EAP method category in camelcase as the key and the list of types as the value (nonEapInnerAuthentication: Reserved, PAP, CHAP, MSCHAP, MSCHAPV2; eapInnerAuthentication: EAP-TLS, EAP-SIM, EAP-AKA, EAP-TTLS with MSCHAPv2; credentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, none, Reserved, Vendor Specific; tunneledEapMethodCredentials: SIM, USIM, NFC Secure Element, Hardware Token, Softoken, Certificate, username/password, Reserved, Anonymous, Vendor Specific)
	AuthenticationTypes map[string]interface{} `json:"authenticationTypes,omitempty"`
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20Methods instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20Methods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20Methods() *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20Methods{}
	return &this
}

// NewNetworksNetworkIdWirelessSsidsNumberHotspot20MethodsWithDefaults instantiates a new NetworksNetworkIdWirelessSsidsNumberHotspot20Methods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdWirelessSsidsNumberHotspot20MethodsWithDefaults() *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods {
	this := NetworksNetworkIdWirelessSsidsNumberHotspot20Methods{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) SetId(v string) {
	o.Id = &v
}

// GetAuthenticationTypes returns the AuthenticationTypes field value if set, zero value otherwise.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) GetAuthenticationTypes() map[string]interface{} {
	if o == nil || o.AuthenticationTypes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AuthenticationTypes
}

// GetAuthenticationTypesOk returns a tuple with the AuthenticationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) GetAuthenticationTypesOk() (map[string]interface{}, bool) {
	if o == nil || o.AuthenticationTypes == nil {
		return nil, false
	}
	return o.AuthenticationTypes, true
}

// HasAuthenticationTypes returns a boolean if a field has been set.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) HasAuthenticationTypes() bool {
	if o != nil && o.AuthenticationTypes != nil {
		return true
	}

	return false
}

// SetAuthenticationTypes gets a reference to the given map[string]interface{} and assigns it to the AuthenticationTypes field.
func (o *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) SetAuthenticationTypes(v map[string]interface{}) {
	o.AuthenticationTypes = v
}

func (o NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AuthenticationTypes != nil {
		toSerialize["authenticationTypes"] = o.AuthenticationTypes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods struct {
	value *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods
	isSet bool
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods) Get() *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods {
	return v.value
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods) Set(val *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods(val *NetworksNetworkIdWirelessSsidsNumberHotspot20Methods) *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods {
	return &NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdWirelessSsidsNumberHotspot20Methods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


