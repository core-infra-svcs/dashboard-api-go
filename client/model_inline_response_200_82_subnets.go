/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20082Subnets struct for InlineResponse20082Subnets
type InlineResponse20082Subnets struct {
	// Serial Number of the MG.
	Serial *string `json:"serial,omitempty"`
	// Name of the MG.
	Name *string `json:"name,omitempty"`
	// Appliance IP of the MG device.
	ApplianceIp *string `json:"applianceIp,omitempty"`
	// Subnet of MG device.
	Subnet *string `json:"subnet,omitempty"`
}

// NewInlineResponse20082Subnets instantiates a new InlineResponse20082Subnets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20082Subnets() *InlineResponse20082Subnets {
	this := InlineResponse20082Subnets{}
	return &this
}

// NewInlineResponse20082SubnetsWithDefaults instantiates a new InlineResponse20082Subnets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20082SubnetsWithDefaults() *InlineResponse20082Subnets {
	this := InlineResponse20082Subnets{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20082Subnets) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082Subnets) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20082Subnets) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20082Subnets) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20082Subnets) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082Subnets) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20082Subnets) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20082Subnets) SetName(v string) {
	o.Name = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *InlineResponse20082Subnets) GetApplianceIp() string {
	if o == nil || isNil(o.ApplianceIp) {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082Subnets) GetApplianceIpOk() (*string, bool) {
	if o == nil || isNil(o.ApplianceIp) {
    return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *InlineResponse20082Subnets) HasApplianceIp() bool {
	if o != nil && !isNil(o.ApplianceIp) {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *InlineResponse20082Subnets) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineResponse20082Subnets) GetSubnet() string {
	if o == nil || isNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20082Subnets) GetSubnetOk() (*string, bool) {
	if o == nil || isNil(o.Subnet) {
    return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineResponse20082Subnets) HasSubnet() bool {
	if o != nil && !isNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineResponse20082Subnets) SetSubnet(v string) {
	o.Subnet = &v
}

func (o InlineResponse20082Subnets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ApplianceIp) {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	if !isNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20082Subnets struct {
	value *InlineResponse20082Subnets
	isSet bool
}

func (v NullableInlineResponse20082Subnets) Get() *InlineResponse20082Subnets {
	return v.value
}

func (v *NullableInlineResponse20082Subnets) Set(val *InlineResponse20082Subnets) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20082Subnets) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20082Subnets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20082Subnets(val *InlineResponse20082Subnets) *NullableInlineResponse20082Subnets {
	return &NullableInlineResponse20082Subnets{value: val, isSet: true}
}

func (v NullableInlineResponse20082Subnets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20082Subnets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


