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

// InlineResponse20037ReservedIpRanges struct for InlineResponse20037ReservedIpRanges
type InlineResponse20037ReservedIpRanges struct {
	// The starting IP address of the reserved IP range
	Start *string `json:"start,omitempty"`
	// The ending IP address of the reserved IP range
	End *string `json:"end,omitempty"`
	// The comment for the reserved IP range
	Comment *string `json:"comment,omitempty"`
}

// NewInlineResponse20037ReservedIpRanges instantiates a new InlineResponse20037ReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20037ReservedIpRanges() *InlineResponse20037ReservedIpRanges {
	this := InlineResponse20037ReservedIpRanges{}
	return &this
}

// NewInlineResponse20037ReservedIpRangesWithDefaults instantiates a new InlineResponse20037ReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20037ReservedIpRangesWithDefaults() *InlineResponse20037ReservedIpRanges {
	this := InlineResponse20037ReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *InlineResponse20037ReservedIpRanges) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20037ReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *InlineResponse20037ReservedIpRanges) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *InlineResponse20037ReservedIpRanges) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *InlineResponse20037ReservedIpRanges) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20037ReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *InlineResponse20037ReservedIpRanges) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *InlineResponse20037ReservedIpRanges) SetEnd(v string) {
	o.End = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20037ReservedIpRanges) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20037ReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20037ReservedIpRanges) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20037ReservedIpRanges) SetComment(v string) {
	o.Comment = &v
}

func (o InlineResponse20037ReservedIpRanges) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !isNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !isNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20037ReservedIpRanges struct {
	value *InlineResponse20037ReservedIpRanges
	isSet bool
}

func (v NullableInlineResponse20037ReservedIpRanges) Get() *InlineResponse20037ReservedIpRanges {
	return v.value
}

func (v *NullableInlineResponse20037ReservedIpRanges) Set(val *InlineResponse20037ReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20037ReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20037ReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20037ReservedIpRanges(val *InlineResponse20037ReservedIpRanges) *NullableInlineResponse20037ReservedIpRanges {
	return &NullableInlineResponse20037ReservedIpRanges{value: val, isSet: true}
}

func (v NullableInlineResponse20037ReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20037ReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


