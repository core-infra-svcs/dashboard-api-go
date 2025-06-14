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

// InlineResponse200378Intervals struct for InlineResponse200378Intervals
type InlineResponse200378Intervals struct {
	// The start time interval snapshots of the query range with iso8601 format
	StartTs *string `json:"startTs,omitempty"`
	// The end time interval snapshots of the query range with iso8601 format
	EndTs *string `json:"endTs,omitempty"`
	Overall *InlineResponse200378Overall `json:"overall,omitempty"`
	// The usage data on the interfaces of the wireless LAN controller
	ByInterface []InlineResponse200378ByInterface `json:"byInterface,omitempty"`
}

// NewInlineResponse200378Intervals instantiates a new InlineResponse200378Intervals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200378Intervals() *InlineResponse200378Intervals {
	this := InlineResponse200378Intervals{}
	return &this
}

// NewInlineResponse200378IntervalsWithDefaults instantiates a new InlineResponse200378Intervals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200378IntervalsWithDefaults() *InlineResponse200378Intervals {
	this := InlineResponse200378Intervals{}
	return &this
}

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse200378Intervals) GetStartTs() string {
	if o == nil || isNil(o.StartTs) {
		var ret string
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200378Intervals) GetStartTsOk() (*string, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse200378Intervals) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given string and assigns it to the StartTs field.
func (o *InlineResponse200378Intervals) SetStartTs(v string) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse200378Intervals) GetEndTs() string {
	if o == nil || isNil(o.EndTs) {
		var ret string
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200378Intervals) GetEndTsOk() (*string, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse200378Intervals) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given string and assigns it to the EndTs field.
func (o *InlineResponse200378Intervals) SetEndTs(v string) {
	o.EndTs = &v
}

// GetOverall returns the Overall field value if set, zero value otherwise.
func (o *InlineResponse200378Intervals) GetOverall() InlineResponse200378Overall {
	if o == nil || isNil(o.Overall) {
		var ret InlineResponse200378Overall
		return ret
	}
	return *o.Overall
}

// GetOverallOk returns a tuple with the Overall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200378Intervals) GetOverallOk() (*InlineResponse200378Overall, bool) {
	if o == nil || isNil(o.Overall) {
    return nil, false
	}
	return o.Overall, true
}

// HasOverall returns a boolean if a field has been set.
func (o *InlineResponse200378Intervals) HasOverall() bool {
	if o != nil && !isNil(o.Overall) {
		return true
	}

	return false
}

// SetOverall gets a reference to the given InlineResponse200378Overall and assigns it to the Overall field.
func (o *InlineResponse200378Intervals) SetOverall(v InlineResponse200378Overall) {
	o.Overall = &v
}

// GetByInterface returns the ByInterface field value if set, zero value otherwise.
func (o *InlineResponse200378Intervals) GetByInterface() []InlineResponse200378ByInterface {
	if o == nil || isNil(o.ByInterface) {
		var ret []InlineResponse200378ByInterface
		return ret
	}
	return o.ByInterface
}

// GetByInterfaceOk returns a tuple with the ByInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200378Intervals) GetByInterfaceOk() ([]InlineResponse200378ByInterface, bool) {
	if o == nil || isNil(o.ByInterface) {
    return nil, false
	}
	return o.ByInterface, true
}

// HasByInterface returns a boolean if a field has been set.
func (o *InlineResponse200378Intervals) HasByInterface() bool {
	if o != nil && !isNil(o.ByInterface) {
		return true
	}

	return false
}

// SetByInterface gets a reference to the given []InlineResponse200378ByInterface and assigns it to the ByInterface field.
func (o *InlineResponse200378Intervals) SetByInterface(v []InlineResponse200378ByInterface) {
	o.ByInterface = v
}

func (o InlineResponse200378Intervals) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200378Intervals struct {
	value *InlineResponse200378Intervals
	isSet bool
}

func (v NullableInlineResponse200378Intervals) Get() *InlineResponse200378Intervals {
	return v.value
}

func (v *NullableInlineResponse200378Intervals) Set(val *InlineResponse200378Intervals) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200378Intervals) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200378Intervals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200378Intervals(val *InlineResponse200378Intervals) *NullableInlineResponse200378Intervals {
	return &NullableInlineResponse200378Intervals{value: val, isSet: true}
}

func (v NullableInlineResponse200378Intervals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200378Intervals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


