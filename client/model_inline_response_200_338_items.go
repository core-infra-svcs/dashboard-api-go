/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200338Items struct for InlineResponse200338Items
type InlineResponse200338Items struct {
	// The name of the switch.
	Name *string `json:"name,omitempty"`
	// The serial number of the switch.
	Serial *string `json:"serial,omitempty"`
	// The MAC address of the switch.
	Mac *string `json:"mac,omitempty"`
	Network *InlineResponse200335Network `json:"network,omitempty"`
	// The model of the switch.
	Model *string `json:"model,omitempty"`
	// The statuses of the ports on the switch.
	Ports []InlineResponse200338Ports `json:"ports,omitempty"`
}

// NewInlineResponse200338Items instantiates a new InlineResponse200338Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200338Items() *InlineResponse200338Items {
	this := InlineResponse200338Items{}
	return &this
}

// NewInlineResponse200338ItemsWithDefaults instantiates a new InlineResponse200338Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200338ItemsWithDefaults() *InlineResponse200338Items {
	this := InlineResponse200338Items{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200338Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200338Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200338Items) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200338Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200338Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200338Items) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200338Items) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Items) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200338Items) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200338Items) SetMac(v string) {
	o.Mac = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200338Items) GetNetwork() InlineResponse200335Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200335Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Items) GetNetworkOk() (*InlineResponse200335Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200338Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200335Network and assigns it to the Network field.
func (o *InlineResponse200338Items) SetNetwork(v InlineResponse200335Network) {
	o.Network = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200338Items) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Items) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200338Items) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200338Items) SetModel(v string) {
	o.Model = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse200338Items) GetPorts() []InlineResponse200338Ports {
	if o == nil || isNil(o.Ports) {
		var ret []InlineResponse200338Ports
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200338Items) GetPortsOk() ([]InlineResponse200338Ports, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse200338Items) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []InlineResponse200338Ports and assigns it to the Ports field.
func (o *InlineResponse200338Items) SetPorts(v []InlineResponse200338Ports) {
	o.Ports = v
}

func (o InlineResponse200338Items) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200338Items struct {
	value *InlineResponse200338Items
	isSet bool
}

func (v NullableInlineResponse200338Items) Get() *InlineResponse200338Items {
	return v.value
}

func (v *NullableInlineResponse200338Items) Set(val *InlineResponse200338Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200338Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200338Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200338Items(val *InlineResponse200338Items) *NullableInlineResponse200338Items {
	return &NullableInlineResponse200338Items{value: val, isSet: true}
}

func (v NullableInlineResponse200338Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200338Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


