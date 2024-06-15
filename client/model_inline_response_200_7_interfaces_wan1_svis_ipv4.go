/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2007InterfacesWan1SvisIpv4 IPv4 settings for static/dynamic mode.
type InlineResponse2007InterfacesWan1SvisIpv4 struct {
	// The assignment mode for this SVI. Applies only when PPPoE is disabled.
	AssignmentMode *string `json:"assignmentMode,omitempty"`
	// IP address and subnet mask when in static mode.
	Address *string `json:"address,omitempty"`
	// Gateway IP address when in static mode.
	Gateway *string `json:"gateway,omitempty"`
	Nameservers *InlineResponse2007InterfacesWan1SvisIpv4Nameservers `json:"nameservers,omitempty"`
}

// NewInlineResponse2007InterfacesWan1SvisIpv4 instantiates a new InlineResponse2007InterfacesWan1SvisIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007InterfacesWan1SvisIpv4() *InlineResponse2007InterfacesWan1SvisIpv4 {
	this := InlineResponse2007InterfacesWan1SvisIpv4{}
	return &this
}

// NewInlineResponse2007InterfacesWan1SvisIpv4WithDefaults instantiates a new InlineResponse2007InterfacesWan1SvisIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007InterfacesWan1SvisIpv4WithDefaults() *InlineResponse2007InterfacesWan1SvisIpv4 {
	this := InlineResponse2007InterfacesWan1SvisIpv4{}
	return &this
}

// GetAssignmentMode returns the AssignmentMode field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetAssignmentMode() string {
	if o == nil || isNil(o.AssignmentMode) {
		var ret string
		return ret
	}
	return *o.AssignmentMode
}

// GetAssignmentModeOk returns a tuple with the AssignmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetAssignmentModeOk() (*string, bool) {
	if o == nil || isNil(o.AssignmentMode) {
    return nil, false
	}
	return o.AssignmentMode, true
}

// HasAssignmentMode returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) HasAssignmentMode() bool {
	if o != nil && !isNil(o.AssignmentMode) {
		return true
	}

	return false
}

// SetAssignmentMode gets a reference to the given string and assigns it to the AssignmentMode field.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) SetAssignmentMode(v string) {
	o.AssignmentMode = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) SetAddress(v string) {
	o.Address = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetGateway() string {
	if o == nil || isNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetGatewayOk() (*string, bool) {
	if o == nil || isNil(o.Gateway) {
    return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) HasGateway() bool {
	if o != nil && !isNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) SetGateway(v string) {
	o.Gateway = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetNameservers() InlineResponse2007InterfacesWan1SvisIpv4Nameservers {
	if o == nil || isNil(o.Nameservers) {
		var ret InlineResponse2007InterfacesWan1SvisIpv4Nameservers
		return ret
	}
	return *o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) GetNameserversOk() (*InlineResponse2007InterfacesWan1SvisIpv4Nameservers, bool) {
	if o == nil || isNil(o.Nameservers) {
    return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) HasNameservers() bool {
	if o != nil && !isNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given InlineResponse2007InterfacesWan1SvisIpv4Nameservers and assigns it to the Nameservers field.
func (o *InlineResponse2007InterfacesWan1SvisIpv4) SetNameservers(v InlineResponse2007InterfacesWan1SvisIpv4Nameservers) {
	o.Nameservers = &v
}

func (o InlineResponse2007InterfacesWan1SvisIpv4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AssignmentMode) {
		toSerialize["assignmentMode"] = o.AssignmentMode
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !isNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !isNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2007InterfacesWan1SvisIpv4 struct {
	value *InlineResponse2007InterfacesWan1SvisIpv4
	isSet bool
}

func (v NullableInlineResponse2007InterfacesWan1SvisIpv4) Get() *InlineResponse2007InterfacesWan1SvisIpv4 {
	return v.value
}

func (v *NullableInlineResponse2007InterfacesWan1SvisIpv4) Set(val *InlineResponse2007InterfacesWan1SvisIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2007InterfacesWan1SvisIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2007InterfacesWan1SvisIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2007InterfacesWan1SvisIpv4(val *InlineResponse2007InterfacesWan1SvisIpv4) *NullableInlineResponse2007InterfacesWan1SvisIpv4 {
	return &NullableInlineResponse2007InterfacesWan1SvisIpv4{value: val, isSet: true}
}

func (v NullableInlineResponse2007InterfacesWan1SvisIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2007InterfacesWan1SvisIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


