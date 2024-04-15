/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 April, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.45.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20044FiltersConditions struct for InlineResponse20044FiltersConditions
type InlineResponse20044FiltersConditions struct {
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

// NewInlineResponse20044FiltersConditions instantiates a new InlineResponse20044FiltersConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20044FiltersConditions() *InlineResponse20044FiltersConditions {
	this := InlineResponse20044FiltersConditions{}
	return &this
}

// NewInlineResponse20044FiltersConditionsWithDefaults instantiates a new InlineResponse20044FiltersConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20044FiltersConditionsWithDefaults() *InlineResponse20044FiltersConditions {
	this := InlineResponse20044FiltersConditions{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse20044FiltersConditions) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044FiltersConditions) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse20044FiltersConditions) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse20044FiltersConditions) SetType(v string) {
	o.Type = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *InlineResponse20044FiltersConditions) GetUnit() string {
	if o == nil || isNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044FiltersConditions) GetUnitOk() (*string, bool) {
	if o == nil || isNil(o.Unit) {
    return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *InlineResponse20044FiltersConditions) HasUnit() bool {
	if o != nil && !isNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *InlineResponse20044FiltersConditions) SetUnit(v string) {
	o.Unit = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *InlineResponse20044FiltersConditions) GetDuration() int32 {
	if o == nil || isNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044FiltersConditions) GetDurationOk() (*int32, bool) {
	if o == nil || isNil(o.Duration) {
    return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *InlineResponse20044FiltersConditions) HasDuration() bool {
	if o != nil && !isNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *InlineResponse20044FiltersConditions) SetDuration(v int32) {
	o.Duration = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *InlineResponse20044FiltersConditions) GetDirection() string {
	if o == nil || isNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044FiltersConditions) GetDirectionOk() (*string, bool) {
	if o == nil || isNil(o.Direction) {
    return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *InlineResponse20044FiltersConditions) HasDirection() bool {
	if o != nil && !isNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *InlineResponse20044FiltersConditions) SetDirection(v string) {
	o.Direction = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *InlineResponse20044FiltersConditions) GetThreshold() float32 {
	if o == nil || isNil(o.Threshold) {
		var ret float32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20044FiltersConditions) GetThresholdOk() (*float32, bool) {
	if o == nil || isNil(o.Threshold) {
    return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *InlineResponse20044FiltersConditions) HasThreshold() bool {
	if o != nil && !isNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float32 and assigns it to the Threshold field.
func (o *InlineResponse20044FiltersConditions) SetThreshold(v float32) {
	o.Threshold = &v
}

func (o InlineResponse20044FiltersConditions) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20044FiltersConditions struct {
	value *InlineResponse20044FiltersConditions
	isSet bool
}

func (v NullableInlineResponse20044FiltersConditions) Get() *InlineResponse20044FiltersConditions {
	return v.value
}

func (v *NullableInlineResponse20044FiltersConditions) Set(val *InlineResponse20044FiltersConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20044FiltersConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20044FiltersConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20044FiltersConditions(val *InlineResponse20044FiltersConditions) *NullableInlineResponse20044FiltersConditions {
	return &NullableInlineResponse20044FiltersConditions{value: val, isSet: true}
}

func (v NullableInlineResponse20044FiltersConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20044FiltersConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


