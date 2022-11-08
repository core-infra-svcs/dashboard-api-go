/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject160 struct for InlineObject160
type InlineObject160 struct {
	// The name of the Identity PSK
	Name *string `json:"name,omitempty"`
	// The passphrase for client authentication
	Passphrase *string `json:"passphrase,omitempty"`
	// The group policy to be applied to clients
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewInlineObject160 instantiates a new InlineObject160 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject160() *InlineObject160 {
	this := InlineObject160{}
	return &this
}

// NewInlineObject160WithDefaults instantiates a new InlineObject160 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject160WithDefaults() *InlineObject160 {
	this := InlineObject160{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject160) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject160) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject160) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject160) SetName(v string) {
	o.Name = &v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineObject160) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject160) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineObject160) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineObject160) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *InlineObject160) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject160) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *InlineObject160) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *InlineObject160) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o InlineObject160) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Passphrase) {
		toSerialize["passphrase"] = o.Passphrase
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject160 struct {
	value *InlineObject160
	isSet bool
}

func (v NullableInlineObject160) Get() *InlineObject160 {
	return v.value
}

func (v *NullableInlineObject160) Set(val *InlineObject160) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject160) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject160) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject160(val *InlineObject160) *NullableInlineObject160 {
	return &NullableInlineObject160{value: val, isSet: true}
}

func (v NullableInlineObject160) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject160) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


