/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 May, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.46.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineObject189 struct for InlineObject189
type InlineObject189 struct {
	// The name of the Identity PSK
	Name string `json:"name"`
	// The passphrase for client authentication. If left blank, one will be auto-generated.
	Passphrase *string `json:"passphrase,omitempty"`
	// The group policy to be applied to clients
	GroupPolicyId string `json:"groupPolicyId"`
	// Timestamp for when the Identity PSK expires. Will not expire if left blank.
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// NewInlineObject189 instantiates a new InlineObject189 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject189(name string, groupPolicyId string) *InlineObject189 {
	this := InlineObject189{}
	this.Name = name
	this.GroupPolicyId = groupPolicyId
	return &this
}

// NewInlineObject189WithDefaults instantiates a new InlineObject189 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject189WithDefaults() *InlineObject189 {
	this := InlineObject189{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject189) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject189) SetName(v string) {
	o.Name = v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineObject189) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineObject189) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineObject189) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value
func (o *InlineObject189) GetGroupPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupPolicyId, true
}

// SetGroupPolicyId sets field value
func (o *InlineObject189) SetGroupPolicyId(v string) {
	o.GroupPolicyId = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *InlineObject189) GetExpiresAt() time.Time {
	if o == nil || isNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject189) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpiresAt) {
    return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *InlineObject189) HasExpiresAt() bool {
	if o != nil && !isNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *InlineObject189) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

func (o InlineObject189) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if true {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	if !isNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject189 struct {
	value *InlineObject189
	isSet bool
}

func (v NullableInlineObject189) Get() *InlineObject189 {
	return v.value
}

func (v *NullableInlineObject189) Set(val *InlineObject189) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject189) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject189) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject189(val *InlineObject189) *NullableInlineObject189 {
	return &NullableInlineObject189{value: val, isSet: true}
}

func (v NullableInlineObject189) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject189) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


