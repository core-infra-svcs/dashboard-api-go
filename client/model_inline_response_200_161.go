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

// InlineResponse200161 struct for InlineResponse200161
type InlineResponse200161 struct {
	// The device MAC address.
	Mac *string `json:"mac,omitempty"`
	// The device name.
	Name *string `json:"name,omitempty"`
	Network *OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork `json:"network,omitempty"`
	// Device product type.
	ProductType *string `json:"productType,omitempty"`
	// The device serial number.
	Serial *string `json:"serial,omitempty"`
	// List of custom tags for the device.
	Tags []string `json:"tags,omitempty"`
	// Information for the device's AC power supplies.
	Slots []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots `json:"slots,omitempty"`
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

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200161) GetNetwork() OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetNetworkOk() (*OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200161) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork and assigns it to the Network field.
func (o *InlineResponse200161) SetNetwork(v OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceNetwork) {
	o.Network = &v
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

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200161) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200161) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200161) SetTags(v []string) {
	o.Tags = v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *InlineResponse200161) GetSlots() []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots {
	if o == nil || isNil(o.Slots) {
		var ret []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200161) GetSlotsOk() ([]OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots, bool) {
	if o == nil || isNil(o.Slots) {
    return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *InlineResponse200161) HasSlots() bool {
	if o != nil && !isNil(o.Slots) {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots and assigns it to the Slots field.
func (o *InlineResponse200161) SetSlots(v []OrganizationsOrganizationIdDevicesPowerModulesStatusesByDeviceSlots) {
	o.Slots = v
}

func (o InlineResponse200161) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Slots) {
		toSerialize["slots"] = o.Slots
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


