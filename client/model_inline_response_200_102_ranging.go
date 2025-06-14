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

// InlineResponse200102Ranging Ranging status and progress information
type InlineResponse200102Ranging struct {
	// Ranging status. Possible values: 'scheduled', 'in progress', 'error', 'finished', 'no neighbors'
	Status *string `json:"status,omitempty"`
	Completed *InlineResponse200102RangingCompleted `json:"completed,omitempty"`
}

// NewInlineResponse200102Ranging instantiates a new InlineResponse200102Ranging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200102Ranging() *InlineResponse200102Ranging {
	this := InlineResponse200102Ranging{}
	return &this
}

// NewInlineResponse200102RangingWithDefaults instantiates a new InlineResponse200102Ranging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200102RangingWithDefaults() *InlineResponse200102Ranging {
	this := InlineResponse200102Ranging{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse200102Ranging) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Ranging) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse200102Ranging) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse200102Ranging) SetStatus(v string) {
	o.Status = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *InlineResponse200102Ranging) GetCompleted() InlineResponse200102RangingCompleted {
	if o == nil || isNil(o.Completed) {
		var ret InlineResponse200102RangingCompleted
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200102Ranging) GetCompletedOk() (*InlineResponse200102RangingCompleted, bool) {
	if o == nil || isNil(o.Completed) {
    return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *InlineResponse200102Ranging) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given InlineResponse200102RangingCompleted and assigns it to the Completed field.
func (o *InlineResponse200102Ranging) SetCompleted(v InlineResponse200102RangingCompleted) {
	o.Completed = &v
}

func (o InlineResponse200102Ranging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200102Ranging struct {
	value *InlineResponse200102Ranging
	isSet bool
}

func (v NullableInlineResponse200102Ranging) Get() *InlineResponse200102Ranging {
	return v.value
}

func (v *NullableInlineResponse200102Ranging) Set(val *InlineResponse200102Ranging) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200102Ranging) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200102Ranging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200102Ranging(val *InlineResponse200102Ranging) *NullableInlineResponse200102Ranging {
	return &NullableInlineResponse200102Ranging{value: val, isSet: true}
}

func (v NullableInlineResponse200102Ranging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200102Ranging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


