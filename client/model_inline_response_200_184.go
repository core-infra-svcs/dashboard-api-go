/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200184 struct for InlineResponse200184
type InlineResponse200184 struct {
	// AP port profile ID
	ProfileId *string `json:"profileId,omitempty"`
	// AP port profile name
	Name *string `json:"name,omitempty"`
	// Is default profile
	IsDefault *bool `json:"isDefault,omitempty"`
	// Ports config
	Ports []NetworksNetworkIdWirelessEthernetPortsProfilesPorts `json:"ports,omitempty"`
	// Usb ports config
	UsbPorts []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts `json:"usbPorts,omitempty"`
}

// NewInlineResponse200184 instantiates a new InlineResponse200184 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200184() *InlineResponse200184 {
	this := InlineResponse200184{}
	return &this
}

// NewInlineResponse200184WithDefaults instantiates a new InlineResponse200184 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200184WithDefaults() *InlineResponse200184 {
	this := InlineResponse200184{}
	return &this
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *InlineResponse200184) GetProfileId() string {
	if o == nil || isNil(o.ProfileId) {
		var ret string
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ProfileId) {
    return nil, false
	}
	return o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *InlineResponse200184) HasProfileId() bool {
	if o != nil && !isNil(o.ProfileId) {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given string and assigns it to the ProfileId field.
func (o *InlineResponse200184) SetProfileId(v string) {
	o.ProfileId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200184) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200184) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200184) SetName(v string) {
	o.Name = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *InlineResponse200184) GetIsDefault() bool {
	if o == nil || isNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetIsDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.IsDefault) {
    return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *InlineResponse200184) HasIsDefault() bool {
	if o != nil && !isNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *InlineResponse200184) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *InlineResponse200184) GetPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesPorts {
	if o == nil || isNil(o.Ports) {
		var ret []NetworksNetworkIdWirelessEthernetPortsProfilesPorts
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetPortsOk() ([]NetworksNetworkIdWirelessEthernetPortsProfilesPorts, bool) {
	if o == nil || isNil(o.Ports) {
    return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *InlineResponse200184) HasPorts() bool {
	if o != nil && !isNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []NetworksNetworkIdWirelessEthernetPortsProfilesPorts and assigns it to the Ports field.
func (o *InlineResponse200184) SetPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesPorts) {
	o.Ports = v
}

// GetUsbPorts returns the UsbPorts field value if set, zero value otherwise.
func (o *InlineResponse200184) GetUsbPorts() []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts {
	if o == nil || isNil(o.UsbPorts) {
		var ret []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts
		return ret
	}
	return o.UsbPorts
}

// GetUsbPortsOk returns a tuple with the UsbPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200184) GetUsbPortsOk() ([]NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts, bool) {
	if o == nil || isNil(o.UsbPorts) {
    return nil, false
	}
	return o.UsbPorts, true
}

// HasUsbPorts returns a boolean if a field has been set.
func (o *InlineResponse200184) HasUsbPorts() bool {
	if o != nil && !isNil(o.UsbPorts) {
		return true
	}

	return false
}

// SetUsbPorts gets a reference to the given []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts and assigns it to the UsbPorts field.
func (o *InlineResponse200184) SetUsbPorts(v []NetworksNetworkIdWirelessEthernetPortsProfilesUsbPorts) {
	o.UsbPorts = v
}

func (o InlineResponse200184) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProfileId) {
		toSerialize["profileId"] = o.ProfileId
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !isNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !isNil(o.UsbPorts) {
		toSerialize["usbPorts"] = o.UsbPorts
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200184 struct {
	value *InlineResponse200184
	isSet bool
}

func (v NullableInlineResponse200184) Get() *InlineResponse200184 {
	return v.value
}

func (v *NullableInlineResponse200184) Set(val *InlineResponse200184) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200184) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200184) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200184(val *InlineResponse200184) *NullableInlineResponse200184 {
	return &NullableInlineResponse200184{value: val, isSet: true}
}

func (v NullableInlineResponse200184) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200184) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


