/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200118 struct for InlineResponse200118
type InlineResponse200118 struct {
	// Name of the switch
	Name *string `json:"name,omitempty"`
	// Serial of the switch
	Serial *string `json:"serial,omitempty"`
	// MAC address of the switch
	Mac *string `json:"mac,omitempty"`
	// Model of the switch
	Model *string `json:"model,omitempty"`
	Network *OrganizationsOrganizationIdSwitchPortsBySwitchNetwork `json:"network,omitempty"`
	// Ports belonging to the switch
	Ports []OrganizationsOrganizationIdSwitchPortsBySwitchPorts `json:"ports,omitempty"`
}

// NewInlineResponse200118 instantiates a new InlineResponse200118 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200118() *InlineResponse200118 {
	this := InlineResponse200118{}
	return &this
}

// NewInlineResponse200118WithDefaults instantiates a new InlineResponse200118 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200118WithDefaults() *InlineResponse200118 {
	this := InlineResponse200118{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200118) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200118) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200118) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200118) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200118) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200118) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200118) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200118) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200118) SetMac(v string) {
	o.Mac = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200118) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200118) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200118) SetModel(v string) {
	o.Model = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200118) GetNetwork() OrganizationsOrganizationIdSwitchPortsBySwitchNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdSwitchPortsBySwitchNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetNetworkOk() (*OrganizationsOrganizationIdSwitchPortsBySwitchNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200118) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdSwitchPortsBySwitchNetwork and assigns it to the Network field.
func (o *InlineResponse200118) SetNetwork(v OrganizationsOrganizationIdSwitchPortsBySwitchNetwork) {
	o.Network = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse200118) GetPorts() []OrganizationsOrganizationIdSwitchPortsBySwitchPorts {
	if o == nil || isNil(o.Ports) {
		var ret []OrganizationsOrganizationIdSwitchPortsBySwitchPorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200118) GetPortsOk() ([]OrganizationsOrganizationIdSwitchPortsBySwitchPorts, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse200118) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []OrganizationsOrganizationIdSwitchPortsBySwitchPorts and assigns it to the Ports field.
func (o *InlineResponse200118) SetPorts(v []OrganizationsOrganizationIdSwitchPortsBySwitchPorts) {
	o.Ports = v
}

func (o InlineResponse200118) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200118 struct {
	value *InlineResponse200118
	isSet bool
}

func (v NullableInlineResponse200118) Get() *InlineResponse200118 {
	return v.value
}

func (v *NullableInlineResponse200118) Set(val *InlineResponse200118) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200118) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200118) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200118(val *InlineResponse200118) *NullableInlineResponse200118 {
	return &NullableInlineResponse200118{value: val, isSet: true}
}

func (v NullableInlineResponse200118) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200118) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

