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

// InlineResponse200339Intervals struct for InlineResponse200339Intervals
type InlineResponse200339Intervals struct {
	// The start time interval snapshots of the query range with iso8601 format
	StartTs *string `json:"startTs,omitempty"`
	// The end time interval snapshots of the query range with iso8601 format
	EndTs *string `json:"endTs,omitempty"`
	Overall *InlineResponse200339Overall `json:"overall,omitempty"`
	// The usage data on the interfaces of the wireless LAN controller
	ByInterface []InlineResponse200339ByInterface `json:"byInterface,omitempty"`
}

// NewInlineResponse200339Intervals instantiates a new InlineResponse200339Intervals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200339Intervals() *InlineResponse200339Intervals {
	this := InlineResponse200339Intervals{}
	return &this
}

// NewInlineResponse200339IntervalsWithDefaults instantiates a new InlineResponse200339Intervals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200339IntervalsWithDefaults() *InlineResponse200339Intervals {
	this := InlineResponse200339Intervals{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200339Intervals) GetStartTs() string {
	if o == nil || isNil(o.StartTs) {
		var ret string
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Intervals) GetStartTsOk() (*string, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200339Intervals) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given string and assigns it to the StartTs field.
func (o *InlineResponse200339Intervals) SetStartTs(v string) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200339Intervals) GetEndTs() string {
	if o == nil || isNil(o.EndTs) {
		var ret string
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Intervals) GetEndTsOk() (*string, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200339Intervals) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given string and assigns it to the EndTs field.
func (o *InlineResponse200339Intervals) SetEndTs(v string) {
	o.EndTs = &v
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *InlineResponse200339Intervals) GetOverall() InlineResponse200339Overall {
	if o == nil || isNil(o.Overall) {
		var ret InlineResponse200339Overall
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Intervals) GetOverallOk() (*InlineResponse200339Overall, bool) {
	if o == nil || isNil(o.Overall) {
    return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *InlineResponse200339Intervals) HasOverall() bool {
	if o != nil && !isNil(o.Overall) {
		return true
	}

	return false
}

// SetOverall gets a reference to the given InlineResponse200339Overall and assigns it to the Overall field.
func (o *InlineResponse200339Intervals) SetOverall(v InlineResponse200339Overall) {
	o.Overall = &v
}

// GetByInterface returns the ByInterface field value if set, zero value otherwise.
func (o *InlineResponse200339Intervals) GetByInterface() []InlineResponse200339ByInterface {
	if o == nil || isNil(o.ByInterface) {
		var ret []InlineResponse200339ByInterface
		return ret
	}
	return o.ByInterface
}

// GetByInterfaceOk returns a tuple with the ByInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200339Intervals) GetByInterfaceOk() ([]InlineResponse200339ByInterface, bool) {
	if o == nil || isNil(o.ByInterface) {
    return nil, false
	}
	return o.ByInterface, true
}

// HasByInterface returns a boolean if a field has been set.
func (o *InlineResponse200339Intervals) HasByInterface() bool {
	if o != nil && !isNil(o.ByInterface) {
		return true
	}

	return false
}

// SetByInterface gets a reference to the given []InlineResponse200339ByInterface and assigns it to the ByInterface field.
func (o *InlineResponse200339Intervals) SetByInterface(v []InlineResponse200339ByInterface) {
	o.ByInterface = v
}

func (o InlineResponse200339Intervals) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.Overall) {
		toSerialize["overall"] = o.Overall
	}
	if !isNil(o.ByInterface) {
		toSerialize["byInterface"] = o.ByInterface
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200339Intervals struct {
	value *InlineResponse200339Intervals
	isSet bool
}

func (v NullableInlineResponse200339Intervals) Get() *InlineResponse200339Intervals {
	return v.value
}

func (v *NullableInlineResponse200339Intervals) Set(val *InlineResponse200339Intervals) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200339Intervals) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200339Intervals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200339Intervals(val *InlineResponse200339Intervals) *NullableInlineResponse200339Intervals {
	return &NullableInlineResponse200339Intervals{value: val, isSet: true}
}

func (v NullableInlineResponse200339Intervals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200339Intervals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


