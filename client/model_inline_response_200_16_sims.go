/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20016Sims struct for InlineResponse20016Sims
type InlineResponse20016Sims struct {
	// SIM slot being configured. Must be 'sim1' on single-sim devices. Use 'sim3' for eSIM.
	Slot *string `json:"slot,omitempty"`
	// If true, this SIM is activated on platform bootup. It must be true on single-SIM devices and is a required field for dual-SIM MGs unless it is being configured using 'simOrdering'.
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// APN configurations. If empty, the default APN will be used.
	Apns []InlineResponse20016Apns `json:"apns,omitempty"`
}

// NewInlineResponse20016Sims instantiates a new InlineResponse20016Sims object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016Sims() *InlineResponse20016Sims {
	this := InlineResponse20016Sims{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// NewInlineResponse20016SimsWithDefaults instantiates a new InlineResponse20016Sims object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016SimsWithDefaults() *InlineResponse20016Sims {
	this := InlineResponse20016Sims{}
	var isPrimary bool = false
	this.IsPrimary = &isPrimary
	return &this
}

// GetSlot returns the Slot field value if set, zero value otherwise.
func (o *InlineResponse20016Sims) GetSlot() string {
	if o == nil || isNil(o.Slot) {
		var ret string
		return ret
	}
	return *o.Slot
}

// GetSlotOk returns a tuple with the Slot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Sims) GetSlotOk() (*string, bool) {
	if o == nil || isNil(o.Slot) {
    return nil, false
	}
	return o.Slot, true
}

// HasSlot returns a boolean if a field has been set.
func (o *InlineResponse20016Sims) HasSlot() bool {
	if o != nil && !isNil(o.Slot) {
		return true
	}

	return false
}

// SetSlot gets a reference to the given string and assigns it to the Slot field.
func (o *InlineResponse20016Sims) SetSlot(v string) {
	o.Slot = &v
}

// GetIsPrimary returns the IsPrimary field value if set, zero value otherwise.
func (o *InlineResponse20016Sims) GetIsPrimary() bool {
	if o == nil || isNil(o.IsPrimary) {
		var ret bool
		return ret
	}
	return *o.IsPrimary
}

// GetIsPrimaryOk returns a tuple with the IsPrimary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Sims) GetIsPrimaryOk() (*bool, bool) {
	if o == nil || isNil(o.IsPrimary) {
    return nil, false
	}
	return o.IsPrimary, true
}

// HasIsPrimary returns a boolean if a field has been set.
func (o *InlineResponse20016Sims) HasIsPrimary() bool {
	if o != nil && !isNil(o.IsPrimary) {
		return true
	}

	return false
}

// SetIsPrimary gets a reference to the given bool and assigns it to the IsPrimary field.
func (o *InlineResponse20016Sims) SetIsPrimary(v bool) {
	o.IsPrimary = &v
}

// GetApns returns the Apns field value if set, zero value otherwise.
func (o *InlineResponse20016Sims) GetApns() []InlineResponse20016Apns {
	if o == nil || isNil(o.Apns) {
		var ret []InlineResponse20016Apns
		return ret
	}
	return o.Apns
}

// GetApnsOk returns a tuple with the Apns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016Sims) GetApnsOk() ([]InlineResponse20016Apns, bool) {
	if o == nil || isNil(o.Apns) {
    return nil, false
	}
	return o.Apns, true
}

// HasApns returns a boolean if a field has been set.
func (o *InlineResponse20016Sims) HasApns() bool {
	if o != nil && !isNil(o.Apns) {
		return true
	}

	return false
}

// SetApns gets a reference to the given []InlineResponse20016Apns and assigns it to the Apns field.
func (o *InlineResponse20016Sims) SetApns(v []InlineResponse20016Apns) {
	o.Apns = v
}

func (o InlineResponse20016Sims) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Slot) {
		toSerialize["slot"] = o.Slot
	}
	if !isNil(o.IsPrimary) {
		toSerialize["isPrimary"] = o.IsPrimary
	}
	if !isNil(o.Apns) {
		toSerialize["apns"] = o.Apns
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20016Sims struct {
	value *InlineResponse20016Sims
	isSet bool
}

func (v NullableInlineResponse20016Sims) Get() *InlineResponse20016Sims {
	return v.value
}

func (v *NullableInlineResponse20016Sims) Set(val *InlineResponse20016Sims) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016Sims) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016Sims) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016Sims(val *InlineResponse20016Sims) *NullableInlineResponse20016Sims {
	return &NullableInlineResponse20016Sims{value: val, isSet: true}
}

func (v NullableInlineResponse20016Sims) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016Sims) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


