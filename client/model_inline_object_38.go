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

// InlineObject38 struct for InlineObject38
type InlineObject38 struct {
	// The subnet of the single LAN configuration
	Subnet *string `json:"subnet,omitempty"`
	// The appliance IP address of the single LAN
	ApplianceIp *string `json:"applianceIp,omitempty"`
}

// NewInlineObject38 instantiates a new InlineObject38 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject38() *InlineObject38 {
	this := InlineObject38{}
	return &this
}

// NewInlineObject38WithDefaults instantiates a new InlineObject38 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject38WithDefaults() *InlineObject38 {
	this := InlineObject38{}
	return &this
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *InlineObject38) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject38) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *InlineObject38) HasSubnet() bool {
	if o != nil && o.Subnet != nil {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *InlineObject38) SetSubnet(v string) {
	o.Subnet = &v
}

// GetApplianceIp returns the ApplianceIp field value if set, zero value otherwise.
func (o *InlineObject38) GetApplianceIp() string {
	if o == nil || o.ApplianceIp == nil {
		var ret string
		return ret
	}
	return *o.ApplianceIp
}

// GetApplianceIpOk returns a tuple with the ApplianceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject38) GetApplianceIpOk() (*string, bool) {
	if o == nil || o.ApplianceIp == nil {
		return nil, false
	}
	return o.ApplianceIp, true
}

// HasApplianceIp returns a boolean if a field has been set.
func (o *InlineObject38) HasApplianceIp() bool {
	if o != nil && o.ApplianceIp != nil {
		return true
	}

	return false
}

// SetApplianceIp gets a reference to the given string and assigns it to the ApplianceIp field.
func (o *InlineObject38) SetApplianceIp(v string) {
	o.ApplianceIp = &v
}

func (o InlineObject38) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.ApplianceIp != nil {
		toSerialize["applianceIp"] = o.ApplianceIp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject38 struct {
	value *InlineObject38
	isSet bool
}

func (v NullableInlineObject38) Get() *InlineObject38 {
	return v.value
}

func (v *NullableInlineObject38) Set(val *InlineObject38) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject38) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject38) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject38(val *InlineObject38) *NullableInlineObject38 {
	return &NullableInlineObject38{value: val, isSet: true}
}

func (v NullableInlineObject38) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject38) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

