/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20069Switches struct for InlineResponse20069Switches
type InlineResponse20069Switches struct {
	// Switch serial number
	Serial *string `json:"serial,omitempty"`
	// Switch alternative management IP. To remove a prior IP setting, provide an empty string
	AlternateManagementIp *string `json:"alternateManagementIp,omitempty"`
	// Switch subnet mask must be in IP format. Only and must be specified for Polaris switches
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Switch gateway must be in IP format. Only and must be specified for Polaris switches
	Gateway *string `json:"gateway,omitempty"`
}

// NewInlineResponse20069Switches instantiates a new InlineResponse20069Switches object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20069Switches() *InlineResponse20069Switches {
	this := InlineResponse20069Switches{}
	return &this
}

// NewInlineResponse20069SwitchesWithDefaults instantiates a new InlineResponse20069Switches object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20069SwitchesWithDefaults() *InlineResponse20069Switches {
	this := InlineResponse20069Switches{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20069Switches) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Switches) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20069Switches) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20069Switches) SetSerial(v string) {
	o.Serial = &v
}

// GetAlternateManagementIp returns the AlternateManagementIp field value if set, zero value otherwise.
func (o *InlineResponse20069Switches) GetAlternateManagementIp() string {
	if o == nil || isNil(o.AlternateManagementIp) {
		var ret string
		return ret
	}
	return *o.AlternateManagementIp
}

// GetAlternateManagementIpOk returns a tuple with the AlternateManagementIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Switches) GetAlternateManagementIpOk() (*string, bool) {
	if o == nil || isNil(o.AlternateManagementIp) {
    return nil, false
	}
	return o.AlternateManagementIp, true
}

// HasAlternateManagementIp returns a boolean if a field has been set.
func (o *InlineResponse20069Switches) HasAlternateManagementIp() bool {
	if o != nil && !isNil(o.AlternateManagementIp) {
		return true
	}

	return false
}

// SetAlternateManagementIp gets a reference to the given string and assigns it to the AlternateManagementIp field.
func (o *InlineResponse20069Switches) SetAlternateManagementIp(v string) {
	o.AlternateManagementIp = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *InlineResponse20069Switches) GetSubnetMask() string {
	if o == nil || isNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Switches) GetSubnetMaskOk() (*string, bool) {
	if o == nil || isNil(o.SubnetMask) {
    return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *InlineResponse20069Switches) HasSubnetMask() bool {
	if o != nil && !isNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *InlineResponse20069Switches) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse20069Switches) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20069Switches) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse20069Switches) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse20069Switches) SetGateway(v string) {
	o.Gateway = &v
}

func (o InlineResponse20069Switches) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.AlternateManagementIp) {
		toSerialize["alternateManagementIp"] = o.AlternateManagementIp
	}
	if !isNil(o.SubnetMask) {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20069Switches struct {
	value *InlineResponse20069Switches
	isSet bool
}

func (v NullableInlineResponse20069Switches) Get() *InlineResponse20069Switches {
	return v.value
}

func (v *NullableInlineResponse20069Switches) Set(val *InlineResponse20069Switches) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20069Switches) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20069Switches) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20069Switches(val *InlineResponse20069Switches) *NullableInlineResponse20069Switches {
	return &NullableInlineResponse20069Switches{value: val, isSet: true}
}

func (v NullableInlineResponse20069Switches) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20069Switches) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

