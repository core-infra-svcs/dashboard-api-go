/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20040Addresses struct for InlineResponse20040Addresses
type InlineResponse20040Addresses struct {
	// The IP protocol used for the address
	Protocol *string `json:"protocol,omitempty"`
	// The type of address assignment. Either static or dynamic.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// The IP address configured for the alternate management interface
	Address *string `json:"address,omitempty"`
	// The gateway address configured for the alternate managment interface
	Gateway *string `json:"gateway,omitempty"`
	// The IPv6 prefix of the interface. Required if IPv6 object is included.
	Prefix *string `json:"prefix,omitempty"`
	Nameservers *DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers `json:"nameservers,omitempty"`
}

// NewInlineResponse20040Addresses instantiates a new InlineResponse20040Addresses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20040Addresses() *InlineResponse20040Addresses {
	this := InlineResponse20040Addresses{}
	return &this
}

// NewInlineResponse20040AddressesWithDefaults instantiates a new InlineResponse20040Addresses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20040AddressesWithDefaults() *InlineResponse20040Addresses {
	this := InlineResponse20040Addresses{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20040Addresses) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040Addresses) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20040Addresses) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20040Addresses) SetProtocol(v string) {
	o.Protocol = &v
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *InlineResponse20040Addresses) GetAssignmentMode() string {
	if o == nil || isNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040Addresses) GetAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.AssignmentMode) {
    return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *InlineResponse20040Addresses) HasAssignmentMode() bool {
	if o != nil && !isNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *InlineResponse20040Addresses) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse20040Addresses) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040Addresses) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse20040Addresses) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse20040Addresses) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse20040Addresses) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040Addresses) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse20040Addresses) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse20040Addresses) SetGateway(v string) {
	o.Gateway = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InlineResponse20040Addresses) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040Addresses) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InlineResponse20040Addresses) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InlineResponse20040Addresses) SetPrefix(v string) {
	o.Prefix = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *InlineResponse20040Addresses) GetNameservers() DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers {
	if o == nil || isNil(o.Nameservers) {
		var ret DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers
		return ret
	}
	return *o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20040Addresses) GetNameserversOk() (*DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers, bool) {
	if o == nil || isNil(o.Nameservers) {
    return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *InlineResponse20040Addresses) HasNameservers() bool {
	if o != nil && !isNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers and assigns it to the Nameservers field.
func (o *InlineResponse20040Addresses) SetNameservers(v DevicesSerialWirelessAlternateManagementInterfaceIpv6Nameservers) {
	o.Nameservers = &v
}

func (o InlineResponse20040Addresses) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20040Addresses struct {
	value *InlineResponse20040Addresses
	isSet bool
}

func (v NullableInlineResponse20040Addresses) Get() *InlineResponse20040Addresses {
	return v.value
}

func (v *NullableInlineResponse20040Addresses) Set(val *InlineResponse20040Addresses) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20040Addresses) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20040Addresses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20040Addresses(val *InlineResponse20040Addresses) *NullableInlineResponse20040Addresses {
	return &NullableInlineResponse20040Addresses{value: val, isSet: true}
}

func (v NullableInlineResponse20040Addresses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20040Addresses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


