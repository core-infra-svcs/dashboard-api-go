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

// InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 The count data for RJ45 ports, indexed by speed in Mb
type InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 struct {
	// The total number of active RJ45 ports
	Total *int32 `json:"total,omitempty"`
	// The number of active 10 Mbps RJ45 ports
	Var10 *int32 `json:"10,omitempty"`
	// The number of active 100 Mbps RJ45 ports
	Var100 *int32 `json:"100,omitempty"`
	// The number of active 1 Gbps RJ45 ports
	Var1000 *int32 `json:"1000,omitempty"`
	// The number of active 2 Gbps RJ45 ports
	Var2500 *int32 `json:"2500,omitempty"`
	// The number of active 5 Gbps RJ45 ports
	Var5000 *int32 `json:"5000,omitempty"`
	// The number of active 10 Gbps RJ45 ports
	Var10000 *int32 `json:"10000,omitempty"`
}

// NewInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 instantiates a new InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45() *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 {
	this := InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45{}
	return &this
}

// NewInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45WithDefaults instantiates a new InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45WithDefaults() *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 {
	this := InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetTotal(v int32) {
	o.Total = &v
}

// GetVar10 returns the Var10 field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar10() int32 {
	if o == nil || isNil(o.Var10) {
		var ret int32
		return ret
	}
	return *o.Var10
}

// GetVar10Ok returns a tuple with the Var10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar10Ok() (*int32, bool) {
	if o == nil || isNil(o.Var10) {
    return nil, false
	}
	return o.Var10, true
}

// HasVar10 returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasVar10() bool {
	if o != nil && !isNil(o.Var10) {
		return true
	}

	return false
}

// SetVar10 gets a reference to the given int32 and assigns it to the Var10 field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetVar10(v int32) {
	o.Var10 = &v
}

// GetVar100 returns the Var100 field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar100() int32 {
	if o == nil || isNil(o.Var100) {
		var ret int32
		return ret
	}
	return *o.Var100
}

// GetVar100Ok returns a tuple with the Var100 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar100Ok() (*int32, bool) {
	if o == nil || isNil(o.Var100) {
    return nil, false
	}
	return o.Var100, true
}

// HasVar100 returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasVar100() bool {
	if o != nil && !isNil(o.Var100) {
		return true
	}

	return false
}

// SetVar100 gets a reference to the given int32 and assigns it to the Var100 field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetVar100(v int32) {
	o.Var100 = &v
}

// GetVar1000 returns the Var1000 field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar1000() int32 {
	if o == nil || isNil(o.Var1000) {
		var ret int32
		return ret
	}
	return *o.Var1000
}

// GetVar1000Ok returns a tuple with the Var1000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar1000Ok() (*int32, bool) {
	if o == nil || isNil(o.Var1000) {
    return nil, false
	}
	return o.Var1000, true
}

// HasVar1000 returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasVar1000() bool {
	if o != nil && !isNil(o.Var1000) {
		return true
	}

	return false
}

// SetVar1000 gets a reference to the given int32 and assigns it to the Var1000 field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetVar1000(v int32) {
	o.Var1000 = &v
}

// GetVar2500 returns the Var2500 field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar2500() int32 {
	if o == nil || isNil(o.Var2500) {
		var ret int32
		return ret
	}
	return *o.Var2500
}

// GetVar2500Ok returns a tuple with the Var2500 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar2500Ok() (*int32, bool) {
	if o == nil || isNil(o.Var2500) {
    return nil, false
	}
	return o.Var2500, true
}

// HasVar2500 returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasVar2500() bool {
	if o != nil && !isNil(o.Var2500) {
		return true
	}

	return false
}

// SetVar2500 gets a reference to the given int32 and assigns it to the Var2500 field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetVar2500(v int32) {
	o.Var2500 = &v
}

// GetVar5000 returns the Var5000 field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar5000() int32 {
	if o == nil || isNil(o.Var5000) {
		var ret int32
		return ret
	}
	return *o.Var5000
}

// GetVar5000Ok returns a tuple with the Var5000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar5000Ok() (*int32, bool) {
	if o == nil || isNil(o.Var5000) {
    return nil, false
	}
	return o.Var5000, true
}

// HasVar5000 returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasVar5000() bool {
	if o != nil && !isNil(o.Var5000) {
		return true
	}

	return false
}

// SetVar5000 gets a reference to the given int32 and assigns it to the Var5000 field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetVar5000(v int32) {
	o.Var5000 = &v
}

// GetVar10000 returns the Var10000 field value if set, zero value otherwise.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar10000() int32 {
	if o == nil || isNil(o.Var10000) {
		var ret int32
		return ret
	}
	return *o.Var10000
}

// GetVar10000Ok returns a tuple with the Var10000 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) GetVar10000Ok() (*int32, bool) {
	if o == nil || isNil(o.Var10000) {
    return nil, false
	}
	return o.Var10000, true
}

// HasVar10000 returns a boolean if a field has been set.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) HasVar10000() bool {
	if o != nil && !isNil(o.Var10000) {
		return true
	}

	return false
}

// SetVar10000 gets a reference to the given int32 and assigns it to the Var10000 field.
func (o *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) SetVar10000(v int32) {
	o.Var10000 = &v
}

func (o InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.Var10) {
		toSerialize["10"] = o.Var10
	}
	if !isNil(o.Var100) {
		toSerialize["100"] = o.Var100
	}
	if !isNil(o.Var1000) {
		toSerialize["1000"] = o.Var1000
	}
	if !isNil(o.Var2500) {
		toSerialize["2500"] = o.Var2500
	}
	if !isNil(o.Var5000) {
		toSerialize["5000"] = o.Var5000
	}
	if !isNil(o.Var10000) {
		toSerialize["10000"] = o.Var10000
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 struct {
	value *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45
	isSet bool
}

func (v NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) Get() *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 {
	return v.value
}

func (v *NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) Set(val *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45(val *InlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) *NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45 {
	return &NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45{value: val, isSet: true}
}

func (v NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200337CountsByStatusActiveByMediaAndLinkSpeedRj45) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


