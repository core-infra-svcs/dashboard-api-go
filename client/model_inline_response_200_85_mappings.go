/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20085Mappings struct for InlineResponse20085Mappings
type InlineResponse20085Mappings struct {
	// The Differentiated Services Code Point (DSCP) tag in the IP header that will be mapped to a particular Class-of-Service (CoS) queue. Value can be in the range of 0 to 63 inclusive.
	Dscp *int32 `json:"dscp,omitempty"`
	// The actual layer-2 CoS queue the DSCP value is mapped to. These are not bits set on outgoing frames. Value can be in the range of 0 to 5 inclusive.
	Cos *int32 `json:"cos,omitempty"`
	// Label for the mapping (optional).
	Title *string `json:"title,omitempty"`
}

// NewInlineResponse20085Mappings instantiates a new InlineResponse20085Mappings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20085Mappings() *InlineResponse20085Mappings {
	this := InlineResponse20085Mappings{}
	return &this
}

// NewInlineResponse20085MappingsWithDefaults instantiates a new InlineResponse20085Mappings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20085MappingsWithDefaults() *InlineResponse20085Mappings {
	this := InlineResponse20085Mappings{}
	return &this
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *InlineResponse20085Mappings) GetDscp() int32 {
	if o == nil || isNil(o.Dscp) {
		var ret int32
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Mappings) GetDscpOk() (*int32, bool) {
	if o == nil || isNil(o.Dscp) {
    return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *InlineResponse20085Mappings) HasDscp() bool {
	if o != nil && !isNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given int32 and assigns it to the Dscp field.
func (o *InlineResponse20085Mappings) SetDscp(v int32) {
	o.Dscp = &v
}

// GetCos returns the Cos field value if set, zero value otherwise.
func (o *InlineResponse20085Mappings) GetCos() int32 {
	if o == nil || isNil(o.Cos) {
		var ret int32
		return ret
	}
	return *o.Cos
}

// GetCosOk returns a tuple with the Cos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Mappings) GetCosOk() (*int32, bool) {
	if o == nil || isNil(o.Cos) {
    return nil, false
	}
	return o.Cos, true
}

// HasCos returns a boolean if a field has been set.
func (o *InlineResponse20085Mappings) HasCos() bool {
	if o != nil && !isNil(o.Cos) {
		return true
	}

	return false
}

// SetCos gets a reference to the given int32 and assigns it to the Cos field.
func (o *InlineResponse20085Mappings) SetCos(v int32) {
	o.Cos = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *InlineResponse20085Mappings) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20085Mappings) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *InlineResponse20085Mappings) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *InlineResponse20085Mappings) SetTitle(v string) {
	o.Title = &v
}

func (o InlineResponse20085Mappings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !isNil(o.Cos) {
		toSerialize["cos"] = o.Cos
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20085Mappings struct {
	value *InlineResponse20085Mappings
	isSet bool
}

func (v NullableInlineResponse20085Mappings) Get() *InlineResponse20085Mappings {
	return v.value
}

func (v *NullableInlineResponse20085Mappings) Set(val *InlineResponse20085Mappings) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20085Mappings) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20085Mappings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20085Mappings(val *InlineResponse20085Mappings) *NullableInlineResponse20085Mappings {
	return &NullableInlineResponse20085Mappings{value: val, isSet: true}
}

func (v NullableInlineResponse20085Mappings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20085Mappings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

