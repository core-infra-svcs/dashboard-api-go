/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses struct for OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses
type OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses struct {
	// Type of address for the device uplink. Available options are: ipv4, ipv6.
	Protocol *string `json:"protocol,omitempty"`
	// Indicates how the device uplink address is assigned. Available options are: static, dynamic.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// Device uplink address.
	Address *string `json:"address,omitempty"`
	// Device uplink gateway address.
	Gateway *string `json:"gateway,omitempty"`
	Public *OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic `json:"public,omitempty"`
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses{}
	return &this
}

// NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddressesWithDefaults instantiates a new OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddressesWithDefaults() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses {
	this := OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetAssignmentMode() string {
	if o == nil || isNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.AssignmentMode) {
    return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) HasAssignmentMode() bool {
	if o != nil && !isNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) SetGateway(v string) {
	o.Gateway = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetPublic() OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic {
	if o == nil || isNil(o.Public) {
		var ret OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) GetPublicOk() (*OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic, bool) {
	if o == nil || isNil(o.Public) {
    return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) HasPublic() bool {
	if o != nil && !isNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic and assigns it to the Public field.
func (o *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) SetPublic(v OrganizationsOrganizationIdDevicesUplinksAddressesByDevicePublic) {
	o.Public = &v
}

func (o OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses struct {
	value *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses
	isSet bool
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) Get() *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) Set(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses(val *OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses {
	return &NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdDevicesUplinksAddressesByDeviceAddresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


