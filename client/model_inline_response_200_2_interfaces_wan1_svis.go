/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 December, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.40.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2002InterfacesWan1Svis SVI settings by protocol.
type InlineResponse2002InterfacesWan1Svis struct {
	Ipv4 *InlineResponse2002InterfacesWan1SvisIpv4 `json:"ipv4,omitempty"`
	Ipv6 *InlineResponse2002InterfacesWan1SvisIpv6 `json:"ipv6,omitempty"`
}

// NewInlineResponse2002InterfacesWan1Svis instantiates a new InlineResponse2002InterfacesWan1Svis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002InterfacesWan1Svis() *InlineResponse2002InterfacesWan1Svis {
	this := InlineResponse2002InterfacesWan1Svis{}
	return &this
}

// NewInlineResponse2002InterfacesWan1SvisWithDefaults instantiates a new InlineResponse2002InterfacesWan1Svis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002InterfacesWan1SvisWithDefaults() *InlineResponse2002InterfacesWan1Svis {
	this := InlineResponse2002InterfacesWan1Svis{}
	return &this
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *InlineResponse2002InterfacesWan1Svis) GetIpv4() InlineResponse2002InterfacesWan1SvisIpv4 {
	if o == nil || isNil(o.Ipv4) {
		var ret InlineResponse2002InterfacesWan1SvisIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002InterfacesWan1Svis) GetIpv4Ok() (*InlineResponse2002InterfacesWan1SvisIpv4, bool) {
	if o == nil || isNil(o.Ipv4) {
    return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *InlineResponse2002InterfacesWan1Svis) HasIpv4() bool {
	if o != nil && !isNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given InlineResponse2002InterfacesWan1SvisIpv4 and assigns it to the Ipv4 field.
func (o *InlineResponse2002InterfacesWan1Svis) SetIpv4(v InlineResponse2002InterfacesWan1SvisIpv4) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *InlineResponse2002InterfacesWan1Svis) GetIpv6() InlineResponse2002InterfacesWan1SvisIpv6 {
	if o == nil || isNil(o.Ipv6) {
		var ret InlineResponse2002InterfacesWan1SvisIpv6
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002InterfacesWan1Svis) GetIpv6Ok() (*InlineResponse2002InterfacesWan1SvisIpv6, bool) {
	if o == nil || isNil(o.Ipv6) {
    return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *InlineResponse2002InterfacesWan1Svis) HasIpv6() bool {
	if o != nil && !isNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given InlineResponse2002InterfacesWan1SvisIpv6 and assigns it to the Ipv6 field.
func (o *InlineResponse2002InterfacesWan1Svis) SetIpv6(v InlineResponse2002InterfacesWan1SvisIpv6) {
	o.Ipv6 = &v
}

func (o InlineResponse2002InterfacesWan1Svis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !isNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2002InterfacesWan1Svis struct {
	value *InlineResponse2002InterfacesWan1Svis
	isSet bool
}

func (v NullableInlineResponse2002InterfacesWan1Svis) Get() *InlineResponse2002InterfacesWan1Svis {
	return v.value
}

func (v *NullableInlineResponse2002InterfacesWan1Svis) Set(val *InlineResponse2002InterfacesWan1Svis) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2002InterfacesWan1Svis) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2002InterfacesWan1Svis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2002InterfacesWan1Svis(val *InlineResponse2002InterfacesWan1Svis) *NullableInlineResponse2002InterfacesWan1Svis {
	return &NullableInlineResponse2002InterfacesWan1Svis{value: val, isSet: true}
}

func (v NullableInlineResponse2002InterfacesWan1Svis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2002InterfacesWan1Svis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


