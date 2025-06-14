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

// InlineResponse200236Uplinks struct for InlineResponse200236Uplinks
type InlineResponse200236Uplinks struct {
	// Uplink Interface Name
	Interface *string `json:"interface,omitempty"`
	// Uplink IP address (in IP or CIDR notation)
	PublicIp *string `json:"publicIp,omitempty"`
}

// NewInlineResponse200236Uplinks instantiates a new InlineResponse200236Uplinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200236Uplinks() *InlineResponse200236Uplinks {
	this := InlineResponse200236Uplinks{}
	return &this
}

// NewInlineResponse200236UplinksWithDefaults instantiates a new InlineResponse200236Uplinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200236UplinksWithDefaults() *InlineResponse200236Uplinks {
	this := InlineResponse200236Uplinks{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *InlineResponse200236Uplinks) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200236Uplinks) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *InlineResponse200236Uplinks) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *InlineResponse200236Uplinks) SetInterface(v string) {
	o.Interface = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *InlineResponse200236Uplinks) GetPublicIp() string {
	if o == nil || isNil(o.PublicIp) {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200236Uplinks) GetPublicIpOk() (*string, bool) {
	if o == nil || isNil(o.PublicIp) {
    return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *InlineResponse200236Uplinks) HasPublicIp() bool {
	if o != nil && !isNil(o.PublicIp) {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *InlineResponse200236Uplinks) SetPublicIp(v string) {
	o.PublicIp = &v
}

func (o InlineResponse200236Uplinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.PublicIp) {
		toSerialize["publicIp"] = o.PublicIp
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200236Uplinks struct {
	value *InlineResponse200236Uplinks
	isSet bool
}

func (v NullableInlineResponse200236Uplinks) Get() *InlineResponse200236Uplinks {
	return v.value
}

func (v *NullableInlineResponse200236Uplinks) Set(val *InlineResponse200236Uplinks) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200236Uplinks) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200236Uplinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200236Uplinks(val *InlineResponse200236Uplinks) *NullableInlineResponse200236Uplinks {
	return &NullableInlineResponse200236Uplinks{value: val, isSet: true}
}

func (v NullableInlineResponse200236Uplinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200236Uplinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


