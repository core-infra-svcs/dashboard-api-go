/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200333Items struct for InlineResponse200333Items
type InlineResponse200333Items struct {
	// The name of the switch.
	Name *string `json:"name,omitempty"`
	// The serial number of the switch.
	Serial *string `json:"serial,omitempty"`
	// The MAC address of the switch.
	Mac *string `json:"mac,omitempty"`
	Network *InlineResponse200329Network `json:"network,omitempty"`
	// The model of the switch.
	Model *string `json:"model,omitempty"`
	// Ports belonging to the switch with LLDP/CDP discovery info.
	Ports []InlineResponse200333Ports `json:"ports,omitempty"`
}

// NewInlineResponse200333Items instantiates a new InlineResponse200333Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200333Items() *InlineResponse200333Items {
	this := InlineResponse200333Items{}
	return &this
}

// NewInlineResponse200333ItemsWithDefaults instantiates a new InlineResponse200333Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200333ItemsWithDefaults() *InlineResponse200333Items {
	this := InlineResponse200333Items{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200333Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200333Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200333Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200333Items) SetName(v string) {
	o.Name = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200333Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200333Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200333Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200333Items) SetSerial(v string) {
	o.Serial = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200333Items) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200333Items) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200333Items) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200333Items) SetMac(v string) {
	o.Mac = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200333Items) GetNetwork() InlineResponse200329Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200329Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200333Items) GetNetworkOk() (*InlineResponse200329Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200333Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200329Network and assigns it to the Network field.
func (o *InlineResponse200333Items) SetNetwork(v InlineResponse200329Network) {
	o.Network = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200333Items) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200333Items) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200333Items) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200333Items) SetModel(v string) {
	o.Model = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse200333Items) GetPorts() []InlineResponse200333Ports {
	if o == nil || isNil(o.Ports) {
		var ret []InlineResponse200333Ports
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200333Items) GetPortsOk() ([]InlineResponse200333Ports, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse200333Items) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []InlineResponse200333Ports and assigns it to the Ports field.
func (o *InlineResponse200333Items) SetPorts(v []InlineResponse200333Ports) {
	o.Ports = v
}

func (o InlineResponse200333Items) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200333Items struct {
	value *InlineResponse200333Items
	isSet bool
}

func (v NullableInlineResponse200333Items) Get() *InlineResponse200333Items {
	return v.value
}

func (v *NullableInlineResponse200333Items) Set(val *InlineResponse200333Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200333Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200333Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200333Items(val *InlineResponse200333Items) *NullableInlineResponse200333Items {
	return &NullableInlineResponse200333Items{value: val, isSet: true}
}

func (v NullableInlineResponse200333Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200333Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


