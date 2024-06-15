/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 June, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.47.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20016ReservedIpRanges struct for InlineResponse20016ReservedIpRanges
type InlineResponse20016ReservedIpRanges struct {
	// Starting IP included in the reserved range of IPs
	Start *string `json:"start,omitempty"`
	// Ending IP included in the reserved range of IPs
	End *string `json:"end,omitempty"`
	// Comment explaining the reserved IP range
	Comment *string `json:"comment,omitempty"`
}

// NewInlineResponse20016ReservedIpRanges instantiates a new InlineResponse20016ReservedIpRanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20016ReservedIpRanges() *InlineResponse20016ReservedIpRanges {
	this := InlineResponse20016ReservedIpRanges{}
	return &this
}

// NewInlineResponse20016ReservedIpRangesWithDefaults instantiates a new InlineResponse20016ReservedIpRanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20016ReservedIpRangesWithDefaults() *InlineResponse20016ReservedIpRanges {
	this := InlineResponse20016ReservedIpRanges{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *InlineResponse20016ReservedIpRanges) GetStart() string {
	if o == nil || isNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016ReservedIpRanges) GetStartOk() (*string, bool) {
	if o == nil || isNil(o.Start) {
    return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *InlineResponse20016ReservedIpRanges) HasStart() bool {
	if o != nil && !isNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *InlineResponse20016ReservedIpRanges) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *InlineResponse20016ReservedIpRanges) GetEnd() string {
	if o == nil || isNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016ReservedIpRanges) GetEndOk() (*string, bool) {
	if o == nil || isNil(o.End) {
    return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *InlineResponse20016ReservedIpRanges) HasEnd() bool {
	if o != nil && !isNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *InlineResponse20016ReservedIpRanges) SetEnd(v string) {
	o.End = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *InlineResponse20016ReservedIpRanges) GetComment() string {
	if o == nil || isNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20016ReservedIpRanges) GetCommentOk() (*string, bool) {
	if o == nil || isNil(o.Comment) {
    return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *InlineResponse20016ReservedIpRanges) HasComment() bool {
	if o != nil && !isNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *InlineResponse20016ReservedIpRanges) SetComment(v string) {
	o.Comment = &v
}

func (o InlineResponse20016ReservedIpRanges) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20016ReservedIpRanges struct {
	value *InlineResponse20016ReservedIpRanges
	isSet bool
}

func (v NullableInlineResponse20016ReservedIpRanges) Get() *InlineResponse20016ReservedIpRanges {
	return v.value
}

func (v *NullableInlineResponse20016ReservedIpRanges) Set(val *InlineResponse20016ReservedIpRanges) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20016ReservedIpRanges) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20016ReservedIpRanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20016ReservedIpRanges(val *InlineResponse20016ReservedIpRanges) *NullableInlineResponse20016ReservedIpRanges {
	return &NullableInlineResponse20016ReservedIpRanges{value: val, isSet: true}
}

func (v NullableInlineResponse20016ReservedIpRanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20016ReservedIpRanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


