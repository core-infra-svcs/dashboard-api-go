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

// InlineObject159 struct for InlineObject159
type InlineObject159 struct {
	// The name of the Identity PSK
	Name string `json:"name"`
	// The passphrase for client authentication. If left blank, one will be auto-generated.
	Passphrase *string `json:"passphrase,omitempty"`
	// The group policy to be applied to clients
	GroupPolicyId string `json:"groupPolicyId"`
}

// NewInlineObject159 instantiates a new InlineObject159 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject159(name string, groupPolicyId string) *InlineObject159 {
	this := InlineObject159{}
	this.Name = name
	this.GroupPolicyId = groupPolicyId
	return &this
}

// NewInlineObject159WithDefaults instantiates a new InlineObject159 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject159WithDefaults() *InlineObject159 {
	this := InlineObject159{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject159) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject159) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject159) SetName(v string) {
	o.Name = v
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *InlineObject159) GetPassphrase() string {
	if o == nil || isNil(o.Passphrase) {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject159) GetPassphraseOk() (*string, bool) {
	if o == nil || isNil(o.Passphrase) {
    return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *InlineObject159) HasPassphrase() bool {
	if o != nil && !isNil(o.Passphrase) {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *InlineObject159) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetGroupPolicyId returns the GroupPolicyId field value
func (o *InlineObject159) GetGroupPolicyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value
// and a boolean to check if the value has been set.
func (o *InlineObject159) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.GroupPolicyId, true
}

// SetGroupPolicyId sets field value
func (o *InlineObject159) SetGroupPolicyId(v string) {
	o.GroupPolicyId = v
}

func (o InlineObject159) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableInlineObject159 struct {
	value *InlineObject159
	isSet bool
}

func (v NullableInlineObject159) Get() *InlineObject159 {
	return v.value
}

func (v *NullableInlineObject159) Set(val *InlineObject159) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject159) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject159) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject159(val *InlineObject159) *NullableInlineObject159 {
	return &NullableInlineObject159{value: val, isSet: true}
}

func (v NullableInlineObject159) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject159) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


