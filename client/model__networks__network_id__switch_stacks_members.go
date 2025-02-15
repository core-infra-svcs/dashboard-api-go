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

// NetworksNetworkIdSwitchStacksMembers struct for NetworksNetworkIdSwitchStacksMembers
type NetworksNetworkIdSwitchStacksMembers struct {
	// Serial number of the device
	Serial *string `json:"serial,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// Role of the device
	Role *string `json:"role,omitempty"`
}

// NewNetworksNetworkIdSwitchStacksMembers instantiates a new NetworksNetworkIdSwitchStacksMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdSwitchStacksMembers() *NetworksNetworkIdSwitchStacksMembers {
	this := NetworksNetworkIdSwitchStacksMembers{}
	return &this
}

// NewNetworksNetworkIdSwitchStacksMembersWithDefaults instantiates a new NetworksNetworkIdSwitchStacksMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdSwitchStacksMembersWithDefaults() *NetworksNetworkIdSwitchStacksMembers {
	this := NetworksNetworkIdSwitchStacksMembers{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksMembers) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *NetworksNetworkIdSwitchStacksMembers) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksMembers) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworksNetworkIdSwitchStacksMembers) SetName(v string) {
	o.Name = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksMembers) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *NetworksNetworkIdSwitchStacksMembers) SetModel(v string) {
	o.Model = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksMembers) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *NetworksNetworkIdSwitchStacksMembers) SetMac(v string) {
	o.Mac = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *NetworksNetworkIdSwitchStacksMembers) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *NetworksNetworkIdSwitchStacksMembers) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *NetworksNetworkIdSwitchStacksMembers) SetRole(v string) {
	o.Role = &v
}

func (o NetworksNetworkIdSwitchStacksMembers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdSwitchStacksMembers struct {
	value *NetworksNetworkIdSwitchStacksMembers
	isSet bool
}

func (v NullableNetworksNetworkIdSwitchStacksMembers) Get() *NetworksNetworkIdSwitchStacksMembers {
	return v.value
}

func (v *NullableNetworksNetworkIdSwitchStacksMembers) Set(val *NetworksNetworkIdSwitchStacksMembers) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdSwitchStacksMembers) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdSwitchStacksMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdSwitchStacksMembers(val *NetworksNetworkIdSwitchStacksMembers) *NullableNetworksNetworkIdSwitchStacksMembers {
	return &NullableNetworksNetworkIdSwitchStacksMembers{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdSwitchStacksMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdSwitchStacksMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


