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

// InlineResponse200227CountsByStatus byStatus
type InlineResponse200227CountsByStatus struct {
	// number of uplinks that are active and working
	Active *int32 `json:"active,omitempty"`
	// number of uplinks that are working but on standby
	Ready *int32 `json:"ready,omitempty"`
	// number of uplinks that were working but have failed
	Failed *int32 `json:"failed,omitempty"`
	// number of uplinks currently connecting
	Connecting *int32 `json:"connecting,omitempty"`
	// number of uplinks currently where nothing is plugged in
	NotConnected *int32 `json:"notConnected,omitempty"`
}

// NewInlineResponse200227CountsByStatus instantiates a new InlineResponse200227CountsByStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200227CountsByStatus() *InlineResponse200227CountsByStatus {
	this := InlineResponse200227CountsByStatus{}
	return &this
}

// NewInlineResponse200227CountsByStatusWithDefaults instantiates a new InlineResponse200227CountsByStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200227CountsByStatusWithDefaults() *InlineResponse200227CountsByStatus {
	this := InlineResponse200227CountsByStatus{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *InlineResponse200227CountsByStatus) GetActive() int32 {
	if o == nil || isNil(o.Active) {
		var ret int32
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227CountsByStatus) GetActiveOk() (*int32, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *InlineResponse200227CountsByStatus) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given int32 and assigns it to the Active field.
func (o *InlineResponse200227CountsByStatus) SetActive(v int32) {
	o.Active = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *InlineResponse200227CountsByStatus) GetReady() int32 {
	if o == nil || isNil(o.Ready) {
		var ret int32
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227CountsByStatus) GetReadyOk() (*int32, bool) {
	if o == nil || isNil(o.Ready) {
    return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *InlineResponse200227CountsByStatus) HasReady() bool {
	if o != nil && !isNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given int32 and assigns it to the Ready field.
func (o *InlineResponse200227CountsByStatus) SetReady(v int32) {
	o.Ready = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *InlineResponse200227CountsByStatus) GetFailed() int32 {
	if o == nil || isNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227CountsByStatus) GetFailedOk() (*int32, bool) {
	if o == nil || isNil(o.Failed) {
    return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *InlineResponse200227CountsByStatus) HasFailed() bool {
	if o != nil && !isNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *InlineResponse200227CountsByStatus) SetFailed(v int32) {
	o.Failed = &v
}

// GetConnecting returns the Connecting field value if set, zero value otherwise.
func (o *InlineResponse200227CountsByStatus) GetConnecting() int32 {
	if o == nil || isNil(o.Connecting) {
		var ret int32
		return ret
	}
	return *o.Connecting
}

// GetConnectingOk returns a tuple with the Connecting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227CountsByStatus) GetConnectingOk() (*int32, bool) {
	if o == nil || isNil(o.Connecting) {
    return nil, false
	}
	return o.Connecting, true
}

// HasConnecting returns a boolean if a field has been set.
func (o *InlineResponse200227CountsByStatus) HasConnecting() bool {
	if o != nil && !isNil(o.Connecting) {
		return true
	}

	return false
}

// SetConnecting gets a reference to the given int32 and assigns it to the Connecting field.
func (o *InlineResponse200227CountsByStatus) SetConnecting(v int32) {
	o.Connecting = &v
}

// GetNotConnected returns the NotConnected field value if set, zero value otherwise.
func (o *InlineResponse200227CountsByStatus) GetNotConnected() int32 {
	if o == nil || isNil(o.NotConnected) {
		var ret int32
		return ret
	}
	return *o.NotConnected
}

// GetNotConnectedOk returns a tuple with the NotConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200227CountsByStatus) GetNotConnectedOk() (*int32, bool) {
	if o == nil || isNil(o.NotConnected) {
    return nil, false
	}
	return o.NotConnected, true
}

// HasNotConnected returns a boolean if a field has been set.
func (o *InlineResponse200227CountsByStatus) HasNotConnected() bool {
	if o != nil && !isNil(o.NotConnected) {
		return true
	}

	return false
}

// SetNotConnected gets a reference to the given int32 and assigns it to the NotConnected field.
func (o *InlineResponse200227CountsByStatus) SetNotConnected(v int32) {
	o.NotConnected = &v
}

func (o InlineResponse200227CountsByStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !isNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !isNil(o.Connecting) {
		toSerialize["connecting"] = o.Connecting
	}
	if !isNil(o.NotConnected) {
		toSerialize["notConnected"] = o.NotConnected
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200227CountsByStatus struct {
	value *InlineResponse200227CountsByStatus
	isSet bool
}

func (v NullableInlineResponse200227CountsByStatus) Get() *InlineResponse200227CountsByStatus {
	return v.value
}

func (v *NullableInlineResponse200227CountsByStatus) Set(val *InlineResponse200227CountsByStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200227CountsByStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200227CountsByStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200227CountsByStatus(val *InlineResponse200227CountsByStatus) *NullableInlineResponse200227CountsByStatus {
	return &NullableInlineResponse200227CountsByStatus{value: val, isSet: true}
}

func (v NullableInlineResponse200227CountsByStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200227CountsByStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


