/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200161 struct for InlineResponse200161
type InlineResponse200161 struct {
	// Name of the Device
	Name *string `json:"name,omitempty"`
	// Serial of the Device
	Serial *string `json:"serial,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// The product type
	ProductType *string `json:"productType,omitempty"`
	VlanProfile *NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile `json:"vlanProfile,omitempty"`
	Stack *NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack `json:"stack,omitempty"`
}

// NewInlineResponse200161 instantiates a new InlineResponse200161 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200161() *InlineResponse200161 {
	this := InlineResponse200161{}
	return &this
}

// NewInlineResponse200161WithDefaults instantiates a new InlineResponse200161 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200161WithDefaults() *InlineResponse200161 {
	this := InlineResponse200161{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200161) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200161) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200161) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200161) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200161) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200161) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200161) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200161) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200161) SetMac(v string) {
	o.Mac = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse200161) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse200161) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse200161) SetProductType(v string) {
	o.ProductType = &v
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *InlineResponse200161) GetVlanProfile() NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile {
	if o == nil || isNil(o.VlanProfile) {
		var ret NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetVlanProfileOk() (*NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile, bool) {
	if o == nil || isNil(o.VlanProfile) {
    return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *InlineResponse200161) HasVlanProfile() bool {
	if o != nil && !isNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile and assigns it to the VlanProfile field.
func (o *InlineResponse200161) SetVlanProfile(v NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) {
	o.VlanProfile = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *InlineResponse200161) GetStack() NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack {
	if o == nil || isNil(o.Stack) {
		var ret NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetStackOk() (*NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack, bool) {
	if o == nil || isNil(o.Stack) {
    return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *InlineResponse200161) HasStack() bool {
	if o != nil && !isNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack and assigns it to the Stack field.
func (o *InlineResponse200161) SetStack(v NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) {
	o.Stack = &v
}

func (o InlineResponse200161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.VlanProfile) {
		toSerialize["vlanProfile"] = o.VlanProfile
	}
	if !isNil(o.Stack) {
		toSerialize["stack"] = o.Stack
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200161 struct {
	value *InlineResponse200161
	isSet bool
}

func (v NullableInlineResponse200161) Get() *InlineResponse200161 {
	return v.value
}

func (v *NullableInlineResponse200161) Set(val *InlineResponse200161) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200161) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200161) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200161(val *InlineResponse200161) *NullableInlineResponse200161 {
	return &NullableInlineResponse200161{value: val, isSet: true}
}

func (v NullableInlineResponse200161) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200161) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


