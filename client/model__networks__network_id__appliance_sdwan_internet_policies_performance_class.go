/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass Performance class setting for uplink preference rule
type NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass struct {
	// Type of this performance class. Must be one of: 'builtin' or 'custom'
	Type *string `json:"type,omitempty"`
	// Name of builtin performance class. Must be present when performanceClass type is 'builtin' and value must be one of: 'VoIP'
	BuiltinPerformanceClassName *string `json:"builtinPerformanceClassName,omitempty"`
	// ID of created custom performance class, must be present when performanceClass type is \"custom\"
	CustomPerformanceClassId *string `json:"customPerformanceClassId,omitempty"`
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass() *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass{}
	return &this
}

// NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClassWithDefaults instantiates a new NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClassWithDefaults() *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass {
	this := NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetType(v string) {
	o.Type = &v
}

// GetBuiltinPerformanceClassName returns the BuiltinPerformanceClassName field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetBuiltinPerformanceClassName() string {
	if o == nil || isNil(o.BuiltinPerformanceClassName) {
		var ret string
		return ret
	}
	return *o.BuiltinPerformanceClassName
}

// GetBuiltinPerformanceClassNameOk returns a tuple with the BuiltinPerformanceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetBuiltinPerformanceClassNameOk() (*string, bool) {
	if o == nil || isNil(o.BuiltinPerformanceClassName) {
    return nil, false
	}
	return o.BuiltinPerformanceClassName, true
}

// HasBuiltinPerformanceClassName returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) HasBuiltinPerformanceClassName() bool {
	if o != nil && !isNil(o.BuiltinPerformanceClassName) {
		return true
	}

	return false
}

// SetBuiltinPerformanceClassName gets a reference to the given string and assigns it to the BuiltinPerformanceClassName field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetBuiltinPerformanceClassName(v string) {
	o.BuiltinPerformanceClassName = &v
}

// GetCustomPerformanceClassId returns the CustomPerformanceClassId field value if set, zero value otherwise.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetCustomPerformanceClassId() string {
	if o == nil || isNil(o.CustomPerformanceClassId) {
		var ret string
		return ret
	}
	return *o.CustomPerformanceClassId
}

// GetCustomPerformanceClassIdOk returns a tuple with the CustomPerformanceClassId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) GetCustomPerformanceClassIdOk() (*string, bool) {
	if o == nil || isNil(o.CustomPerformanceClassId) {
    return nil, false
	}
	return o.CustomPerformanceClassId, true
}

// HasCustomPerformanceClassId returns a boolean if a field has been set.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) HasCustomPerformanceClassId() bool {
	if o != nil && !isNil(o.CustomPerformanceClassId) {
		return true
	}

	return false
}

// SetCustomPerformanceClassId gets a reference to the given string and assigns it to the CustomPerformanceClassId field.
func (o *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) SetCustomPerformanceClassId(v string) {
	o.CustomPerformanceClassId = &v
}

func (o NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.BuiltinPerformanceClassName) {
		toSerialize["builtinPerformanceClassName"] = o.BuiltinPerformanceClassName
	}
	if !isNil(o.CustomPerformanceClassId) {
		toSerialize["customPerformanceClassId"] = o.CustomPerformanceClassId
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass struct {
	value *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass
	isSet bool
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) Get() *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass {
	return v.value
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) Set(val *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass(val *NetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass {
	return &NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdApplianceSdwanInternetPoliciesPerformanceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


