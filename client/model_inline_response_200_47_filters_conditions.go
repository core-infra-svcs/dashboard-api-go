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

// InlineResponse20047FiltersConditions struct for InlineResponse20047FiltersConditions
type InlineResponse20047FiltersConditions struct {
	// Type of condition
	Type *string `json:"type,omitempty"`
	// Unit
	Unit *string `json:"unit,omitempty"`
	// Duration
	Duration *int32 `json:"duration,omitempty"`
	// Direction
	Direction *string `json:"direction,omitempty"`
	// Threshold
	Threshold *float32 `json:"threshold,omitempty"`
}

// NewInlineResponse20047FiltersConditions instantiates a new InlineResponse20047FiltersConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20047FiltersConditions() *InlineResponse20047FiltersConditions {
	this := InlineResponse20047FiltersConditions{}
	return &this
}

// NewInlineResponse20047FiltersConditionsWithDefaults instantiates a new InlineResponse20047FiltersConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20047FiltersConditionsWithDefaults() *InlineResponse20047FiltersConditions {
	this := InlineResponse20047FiltersConditions{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20047FiltersConditions) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047FiltersConditions) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20047FiltersConditions) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20047FiltersConditions) SetType(v string) {
	o.Type = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *InlineResponse20047FiltersConditions) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047FiltersConditions) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *InlineResponse20047FiltersConditions) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *InlineResponse20047FiltersConditions) SetUnit(v string) {
	o.Unit = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse20047FiltersConditions) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047FiltersConditions) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse20047FiltersConditions) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse20047FiltersConditions) SetDuration(v int32) {
	o.Duration = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *InlineResponse20047FiltersConditions) GetDirection() string {
	if o == nil || isNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047FiltersConditions) GetDirectionOk() (*string, bool) {
	if o == nil || isNil(o.Direction) {
    return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *InlineResponse20047FiltersConditions) HasDirection() bool {
	if o != nil && !isNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *InlineResponse20047FiltersConditions) SetDirection(v string) {
	o.Direction = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *InlineResponse20047FiltersConditions) GetThreshold() float32 {
	if o == nil || isNil(o.Threshold) {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047FiltersConditions) GetThresholdOk() (*float32, bool) {
	if o == nil || isNil(o.Threshold) {
    return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *InlineResponse20047FiltersConditions) HasThreshold() bool {
	if o != nil && !isNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *InlineResponse20047FiltersConditions) SetThreshold(v float32) {
	o.Threshold = &v
}

func (o InlineResponse20047FiltersConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !isNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !isNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !isNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20047FiltersConditions struct {
	value *InlineResponse20047FiltersConditions
	isSet bool
}

func (v NullableInlineResponse20047FiltersConditions) Get() *InlineResponse20047FiltersConditions {
	return v.value
}

func (v *NullableInlineResponse20047FiltersConditions) Set(val *InlineResponse20047FiltersConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20047FiltersConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20047FiltersConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20047FiltersConditions(val *InlineResponse20047FiltersConditions) *NullableInlineResponse20047FiltersConditions {
	return &NullableInlineResponse20047FiltersConditions{value: val, isSet: true}
}

func (v NullableInlineResponse20047FiltersConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20047FiltersConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


