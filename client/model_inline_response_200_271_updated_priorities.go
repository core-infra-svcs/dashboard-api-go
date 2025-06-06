/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200271UpdatedPriorities struct for InlineResponse200271UpdatedPriorities
type InlineResponse200271UpdatedPriorities struct {
	// Id of the scheduled packet capture
	PcapScheduleConfigurationId *string `json:"pcapScheduleConfigurationId,omitempty"`
	// Priority value
	Priority *int32 `json:"priority,omitempty"`
}

// NewInlineResponse200271UpdatedPriorities instantiates a new InlineResponse200271UpdatedPriorities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200271UpdatedPriorities() *InlineResponse200271UpdatedPriorities {
	this := InlineResponse200271UpdatedPriorities{}
	return &this
}

// NewInlineResponse200271UpdatedPrioritiesWithDefaults instantiates a new InlineResponse200271UpdatedPriorities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200271UpdatedPrioritiesWithDefaults() *InlineResponse200271UpdatedPriorities {
	this := InlineResponse200271UpdatedPriorities{}
	return &this
}

// GetPcapScheduleConfigurationId returns the PcapScheduleConfigurationId field value if set, zero value otherwise.
func (o *InlineResponse200271UpdatedPriorities) GetPcapScheduleConfigurationId() string {
	if o == nil || isNil(o.PcapScheduleConfigurationId) {
		var ret string
		return ret
	}
	return *o.PcapScheduleConfigurationId
}

// GetPcapScheduleConfigurationIdOk returns a tuple with the PcapScheduleConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200271UpdatedPriorities) GetPcapScheduleConfigurationIdOk() (*string, bool) {
	if o == nil || isNil(o.PcapScheduleConfigurationId) {
    return nil, false
	}
	return o.PcapScheduleConfigurationId, true
}

// HasPcapScheduleConfigurationId returns a boolean if a field has been set.
func (o *InlineResponse200271UpdatedPriorities) HasPcapScheduleConfigurationId() bool {
	if o != nil && !isNil(o.PcapScheduleConfigurationId) {
		return true
	}

	return false
}

// SetPcapScheduleConfigurationId gets a reference to the given string and assigns it to the PcapScheduleConfigurationId field.
func (o *InlineResponse200271UpdatedPriorities) SetPcapScheduleConfigurationId(v string) {
	o.PcapScheduleConfigurationId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InlineResponse200271UpdatedPriorities) GetPriority() int32 {
	if o == nil || isNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200271UpdatedPriorities) GetPriorityOk() (*int32, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InlineResponse200271UpdatedPriorities) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *InlineResponse200271UpdatedPriorities) SetPriority(v int32) {
	o.Priority = &v
}

func (o InlineResponse200271UpdatedPriorities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PcapScheduleConfigurationId) {
		toSerialize["pcapScheduleConfigurationId"] = o.PcapScheduleConfigurationId
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200271UpdatedPriorities struct {
	value *InlineResponse200271UpdatedPriorities
	isSet bool
}

func (v NullableInlineResponse200271UpdatedPriorities) Get() *InlineResponse200271UpdatedPriorities {
	return v.value
}

func (v *NullableInlineResponse200271UpdatedPriorities) Set(val *InlineResponse200271UpdatedPriorities) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200271UpdatedPriorities) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200271UpdatedPriorities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200271UpdatedPriorities(val *InlineResponse200271UpdatedPriorities) *NullableInlineResponse200271UpdatedPriorities {
	return &NullableInlineResponse200271UpdatedPriorities{value: val, isSet: true}
}

func (v NullableInlineResponse200271UpdatedPriorities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200271UpdatedPriorities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


