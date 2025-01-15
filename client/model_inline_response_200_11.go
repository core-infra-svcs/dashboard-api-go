/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20011 struct for InlineResponse20011
type InlineResponse20011 struct {
	// The start time
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time
	EndTs *time.Time `json:"endTs,omitempty"`
	// The zone id
	ZoneId *int32 `json:"zoneId,omitempty"`
	// The number of entrances
	Entrances *int32 `json:"entrances,omitempty"`
	// The average count
	AverageCount *float32 `json:"averageCount,omitempty"`
}

// NewInlineResponse20011 instantiates a new InlineResponse20011 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20011() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// NewInlineResponse20011WithDefaults instantiates a new InlineResponse20011 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20011WithDefaults() *InlineResponse20011 {
	this := InlineResponse20011{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse20011) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse20011) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse20011) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse20011) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse20011) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse20011) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *InlineResponse20011) GetZoneId() int32 {
	if o == nil || isNil(o.ZoneId) {
		var ret int32
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetZoneIdOk() (*int32, bool) {
	if o == nil || isNil(o.ZoneId) {
    return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *InlineResponse20011) HasZoneId() bool {
	if o != nil && !isNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int32 and assigns it to the ZoneId field.
func (o *InlineResponse20011) SetZoneId(v int32) {
	o.ZoneId = &v
}

// GetEntrances returns the Entrances field value if set, zero value otherwise.
func (o *InlineResponse20011) GetEntrances() int32 {
	if o == nil || isNil(o.Entrances) {
		var ret int32
		return ret
	}
	return *o.Entrances
}

// GetEntrancesOk returns a tuple with the Entrances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetEntrancesOk() (*int32, bool) {
	if o == nil || isNil(o.Entrances) {
    return nil, false
	}
	return o.Entrances, true
}

// HasEntrances returns a boolean if a field has been set.
func (o *InlineResponse20011) HasEntrances() bool {
	if o != nil && !isNil(o.Entrances) {
		return true
	}

	return false
}

// SetEntrances gets a reference to the given int32 and assigns it to the Entrances field.
func (o *InlineResponse20011) SetEntrances(v int32) {
	o.Entrances = &v
}

// GetAverageCount returns the AverageCount field value if set, zero value otherwise.
func (o *InlineResponse20011) GetAverageCount() float32 {
	if o == nil || isNil(o.AverageCount) {
		var ret float32
		return ret
	}
	return *o.AverageCount
}

// GetAverageCountOk returns a tuple with the AverageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20011) GetAverageCountOk() (*float32, bool) {
	if o == nil || isNil(o.AverageCount) {
    return nil, false
	}
	return o.AverageCount, true
}

// HasAverageCount returns a boolean if a field has been set.
func (o *InlineResponse20011) HasAverageCount() bool {
	if o != nil && !isNil(o.AverageCount) {
		return true
	}

	return false
}

// SetAverageCount gets a reference to the given float32 and assigns it to the AverageCount field.
func (o *InlineResponse20011) SetAverageCount(v float32) {
	o.AverageCount = &v
}

func (o InlineResponse20011) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.ZoneId) {
		toSerialize["zoneId"] = o.ZoneId
	}
	if !isNil(o.Entrances) {
		toSerialize["entrances"] = o.Entrances
	}
	if !isNil(o.AverageCount) {
		toSerialize["averageCount"] = o.AverageCount
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20011 struct {
	value *InlineResponse20011
	isSet bool
}

func (v NullableInlineResponse20011) Get() *InlineResponse20011 {
	return v.value
}

func (v *NullableInlineResponse20011) Set(val *InlineResponse20011) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20011) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20011) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20011(val *InlineResponse20011) *NullableInlineResponse20011 {
	return &NullableInlineResponse20011{value: val, isSet: true}
}

func (v NullableInlineResponse20011) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20011) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


