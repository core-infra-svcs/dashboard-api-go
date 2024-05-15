/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200162 struct for InlineResponse200162
type InlineResponse200162 struct {
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

// NewInlineResponse200162 instantiates a new InlineResponse200162 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200162() *InlineResponse200162 {
	this := InlineResponse200162{}
	return &this
}

// NewInlineResponse200162WithDefaults instantiates a new InlineResponse200162 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200162WithDefaults() *InlineResponse200162 {
	this := InlineResponse200162{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200162) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200162) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200162) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200162) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200162) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200162) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200162) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200162) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200162) SetMac(v string) {
	o.Mac = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse200162) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse200162) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse200162) SetProductType(v string) {
	o.ProductType = &v
}

// GetVlanProfile returns the VlanProfile field value if set, zero value otherwise.
func (o *InlineResponse200162) GetVlanProfile() NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile {
	if o == nil || isNil(o.VlanProfile) {
		var ret NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile
		return ret
	}
	return *o.VlanProfile
}

// GetVlanProfileOk returns a tuple with the VlanProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetVlanProfileOk() (*NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile, bool) {
	if o == nil || isNil(o.VlanProfile) {
    return nil, false
	}
	return o.VlanProfile, true
}

// HasVlanProfile returns a boolean if a field has been set.
func (o *InlineResponse200162) HasVlanProfile() bool {
	if o != nil && !isNil(o.VlanProfile) {
		return true
	}

	return false
}

// SetVlanProfile gets a reference to the given NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile and assigns it to the VlanProfile field.
func (o *InlineResponse200162) SetVlanProfile(v NetworksNetworkIdVlanProfilesAssignmentsByDeviceVlanProfile) {
	o.VlanProfile = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *InlineResponse200162) GetStack() NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack {
	if o == nil || isNil(o.Stack) {
		var ret NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetStackOk() (*NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack, bool) {
	if o == nil || isNil(o.Stack) {
    return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *InlineResponse200162) HasStack() bool {
	if o != nil && !isNil(o.Stack) {
		return true
	}

	return false
}

// SetStack gets a reference to the given NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack and assigns it to the Stack field.
func (o *InlineResponse200162) SetStack(v NetworksNetworkIdVlanProfilesAssignmentsByDeviceStack) {
	o.Stack = &v
}

func (o InlineResponse200162) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200162 struct {
	value *InlineResponse200162
	isSet bool
}

func (v NullableInlineResponse200162) Get() *InlineResponse200162 {
	return v.value
}

func (v *NullableInlineResponse200162) Set(val *InlineResponse200162) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200162) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200162) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200162(val *InlineResponse200162) *NullableInlineResponse200162 {
	return &NullableInlineResponse200162{value: val, isSet: true}
}

func (v NullableInlineResponse200162) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200162) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


