/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200359CountsConnections Wireless LAN controller associated access point counts
type InlineResponse200359CountsConnections struct {
	// Wireless LAN controller associated total access point count
	Total *int32 `json:"total,omitempty"`
	ByStatus *InlineResponse200359CountsConnectionsByStatus `json:"byStatus,omitempty"`
}

// NewInlineResponse200359CountsConnections instantiates a new InlineResponse200359CountsConnections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200359CountsConnections() *InlineResponse200359CountsConnections {
	this := InlineResponse200359CountsConnections{}
	return &this
}

// NewInlineResponse200359CountsConnectionsWithDefaults instantiates a new InlineResponse200359CountsConnections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200359CountsConnectionsWithDefaults() *InlineResponse200359CountsConnections {
	this := InlineResponse200359CountsConnections{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *InlineResponse200359CountsConnections) GetTotal() int32 {
	if o == nil || isNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359CountsConnections) GetTotalOk() (*int32, bool) {
	if o == nil || isNil(o.Total) {
    return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *InlineResponse200359CountsConnections) HasTotal() bool {
	if o != nil && !isNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *InlineResponse200359CountsConnections) SetTotal(v int32) {
	o.Total = &v
}

// GetByStatus returns the ByStatus field value if set, zero value otherwise.
func (o *InlineResponse200359CountsConnections) GetByStatus() InlineResponse200359CountsConnectionsByStatus {
	if o == nil || isNil(o.ByStatus) {
		var ret InlineResponse200359CountsConnectionsByStatus
		return ret
	}
	return *o.ByStatus
}

// GetByStatusOk returns a tuple with the ByStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200359CountsConnections) GetByStatusOk() (*InlineResponse200359CountsConnectionsByStatus, bool) {
	if o == nil || isNil(o.ByStatus) {
    return nil, false
	}
	return o.ByStatus, true
}

// HasByStatus returns a boolean if a field has been set.
func (o *InlineResponse200359CountsConnections) HasByStatus() bool {
	if o != nil && !isNil(o.ByStatus) {
		return true
	}

	return false
}

// SetByStatus gets a reference to the given InlineResponse200359CountsConnectionsByStatus and assigns it to the ByStatus field.
func (o *InlineResponse200359CountsConnections) SetByStatus(v InlineResponse200359CountsConnectionsByStatus) {
	o.ByStatus = &v
}

func (o InlineResponse200359CountsConnections) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !isNil(o.ByStatus) {
		toSerialize["byStatus"] = o.ByStatus
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200359CountsConnections struct {
	value *InlineResponse200359CountsConnections
	isSet bool
}

func (v NullableInlineResponse200359CountsConnections) Get() *InlineResponse200359CountsConnections {
	return v.value
}

func (v *NullableInlineResponse200359CountsConnections) Set(val *InlineResponse200359CountsConnections) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200359CountsConnections) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200359CountsConnections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200359CountsConnections(val *InlineResponse200359CountsConnections) *NullableInlineResponse200359CountsConnections {
	return &NullableInlineResponse200359CountsConnections{value: val, isSet: true}
}

func (v NullableInlineResponse200359CountsConnections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200359CountsConnections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


