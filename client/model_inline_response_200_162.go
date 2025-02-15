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

// InlineResponse200162 struct for InlineResponse200162
type InlineResponse200162 struct {
	// The id.
	RendezvousPointId *string `json:"rendezvousPointId,omitempty"`
	// The serial.
	Serial *string `json:"serial,omitempty"`
	// The name of the interface to use.
	InterfaceName *string `json:"interfaceName,omitempty"`
	// The IP address of the interface to use.
	InterfaceIp *string `json:"interfaceIp,omitempty"`
	// 'Any', or the IP address of a multicast group.
	MulticastGroup *string `json:"multicastGroup,omitempty"`
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

// GetRendezvousPointId returns the RendezvousPointId field value if set, zero value otherwise.
func (o *InlineResponse200162) GetRendezvousPointId() string {
	if o == nil || isNil(o.RendezvousPointId) {
		var ret string
		return ret
	}
	return *o.RendezvousPointId
}

// GetRendezvousPointIdOk returns a tuple with the RendezvousPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetRendezvousPointIdOk() (*string, bool) {
	if o == nil || isNil(o.RendezvousPointId) {
    return nil, false
	}
	return o.RendezvousPointId, true
}

// HasRendezvousPointId returns a boolean if a field has been set.
func (o *InlineResponse200162) HasRendezvousPointId() bool {
	if o != nil && !isNil(o.RendezvousPointId) {
		return true
	}

	return false
}

// SetRendezvousPointId gets a reference to the given string and assigns it to the RendezvousPointId field.
func (o *InlineResponse200162) SetRendezvousPointId(v string) {
	o.RendezvousPointId = &v
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

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *InlineResponse200162) GetInterfaceName() string {
	if o == nil || isNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetInterfaceNameOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceName) {
    return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *InlineResponse200162) HasInterfaceName() bool {
	if o != nil && !isNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *InlineResponse200162) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetInterfaceIp returns the InterfaceIp field value if set, zero value otherwise.
func (o *InlineResponse200162) GetInterfaceIp() string {
	if o == nil || isNil(o.InterfaceIp) {
		var ret string
		return ret
	}
	return *o.InterfaceIp
}

// GetInterfaceIpOk returns a tuple with the InterfaceIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetInterfaceIpOk() (*string, bool) {
	if o == nil || isNil(o.InterfaceIp) {
    return nil, false
	}
	return o.InterfaceIp, true
}

// HasInterfaceIp returns a boolean if a field has been set.
func (o *InlineResponse200162) HasInterfaceIp() bool {
	if o != nil && !isNil(o.InterfaceIp) {
		return true
	}

	return false
}

// SetInterfaceIp gets a reference to the given string and assigns it to the InterfaceIp field.
func (o *InlineResponse200162) SetInterfaceIp(v string) {
	o.InterfaceIp = &v
}

// GetMulticastGroup returns the MulticastGroup field value if set, zero value otherwise.
func (o *InlineResponse200162) GetMulticastGroup() string {
	if o == nil || isNil(o.MulticastGroup) {
		var ret string
		return ret
	}
	return *o.MulticastGroup
}

// GetMulticastGroupOk returns a tuple with the MulticastGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200162) GetMulticastGroupOk() (*string, bool) {
	if o == nil || isNil(o.MulticastGroup) {
    return nil, false
	}
	return o.MulticastGroup, true
}

// HasMulticastGroup returns a boolean if a field has been set.
func (o *InlineResponse200162) HasMulticastGroup() bool {
	if o != nil && !isNil(o.MulticastGroup) {
		return true
	}

	return false
}

// SetMulticastGroup gets a reference to the given string and assigns it to the MulticastGroup field.
func (o *InlineResponse200162) SetMulticastGroup(v string) {
	o.MulticastGroup = &v
}

func (o InlineResponse200162) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RendezvousPointId) {
		toSerialize["rendezvousPointId"] = o.RendezvousPointId
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.InterfaceName) {
		toSerialize["interfaceName"] = o.InterfaceName
	}
	if !isNil(o.InterfaceIp) {
		toSerialize["interfaceIp"] = o.InterfaceIp
	}
	if !isNil(o.MulticastGroup) {
		toSerialize["multicastGroup"] = o.MulticastGroup
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


