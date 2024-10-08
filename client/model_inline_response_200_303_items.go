/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200303Items struct for InlineResponse200303Items
type InlineResponse200303Items struct {
	Network *InlineResponse2019Network `json:"network,omitempty"`
	// Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)
	RuleId *string `json:"ruleId,omitempty"`
	// Indicates whether or not clients are allowed to        connect to rogue SSIDs by default. (blocked by default)
	Type *string `json:"type,omitempty"`
	// Updated at timestamp
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Created at timestamp
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	Match *InlineResponse200303Match `json:"match,omitempty"`
}

// NewInlineResponse200303Items instantiates a new InlineResponse200303Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200303Items() *InlineResponse200303Items {
	this := InlineResponse200303Items{}
	return &this
}

// NewInlineResponse200303ItemsWithDefaults instantiates a new InlineResponse200303Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200303ItemsWithDefaults() *InlineResponse200303Items {
	this := InlineResponse200303Items{}
	return &this
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200303Items) GetNetwork() InlineResponse2019Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse2019Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200303Items) GetNetworkOk() (*InlineResponse2019Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200303Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse2019Network and assigns it to the Network field.
func (o *InlineResponse200303Items) SetNetwork(v InlineResponse2019Network) {
	o.Network = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *InlineResponse200303Items) GetRuleId() string {
	if o == nil || isNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200303Items) GetRuleIdOk() (*string, bool) {
	if o == nil || isNil(o.RuleId) {
    return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *InlineResponse200303Items) HasRuleId() bool {
	if o != nil && !isNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *InlineResponse200303Items) SetRuleId(v string) {
	o.RuleId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200303Items) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200303Items) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200303Items) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200303Items) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse200303Items) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200303Items) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse200303Items) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InlineResponse200303Items) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200303Items) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200303Items) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200303Items) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200303Items) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *InlineResponse200303Items) GetMatch() InlineResponse200303Match {
	if o == nil || isNil(o.Match) {
		var ret InlineResponse200303Match
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200303Items) GetMatchOk() (*InlineResponse200303Match, bool) {
	if o == nil || isNil(o.Match) {
    return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *InlineResponse200303Items) HasMatch() bool {
	if o != nil && !isNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given InlineResponse200303Match and assigns it to the Match field.
func (o *InlineResponse200303Items) SetMatch(v InlineResponse200303Match) {
	o.Match = &v
}

func (o InlineResponse200303Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200303Items struct {
	value *InlineResponse200303Items
	isSet bool
}

func (v NullableInlineResponse200303Items) Get() *InlineResponse200303Items {
	return v.value
}

func (v *NullableInlineResponse200303Items) Set(val *InlineResponse200303Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200303Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200303Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200303Items(val *InlineResponse200303Items) *NullableInlineResponse200303Items {
	return &NullableInlineResponse200303Items{value: val, isSet: true}
}

func (v NullableInlineResponse200303Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200303Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


