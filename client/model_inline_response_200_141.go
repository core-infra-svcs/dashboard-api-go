/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200141 struct for InlineResponse200141
type InlineResponse200141 struct {
	// ID of a profile.
	Id *string `json:"id,omitempty"`
	// Name of a profile.
	Name *string `json:"name,omitempty"`
	// Description of a profile.
	Description *string `json:"description,omitempty"`
	// Scope of a profile.
	Scope *string `json:"scope,omitempty"`
	// Tags of a profile.
	Tags []string `json:"tags,omitempty"`
	// Payloads in the profile.
	PayloadTypes []string `json:"payloadTypes,omitempty"`
}

// NewInlineResponse200141 instantiates a new InlineResponse200141 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200141() *InlineResponse200141 {
	this := InlineResponse200141{}
	return &this
}

// NewInlineResponse200141WithDefaults instantiates a new InlineResponse200141 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200141WithDefaults() *InlineResponse200141 {
	this := InlineResponse200141{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200141) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200141) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200141) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200141) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200141) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200141) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse200141) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse200141) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse200141) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse200141) GetScope() string {
	if o == nil || isNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetScopeOk() (*string, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse200141) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *InlineResponse200141) SetScope(v string) {
	o.Scope = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200141) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200141) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200141) SetTags(v []string) {
	o.Tags = v
}

// GetPayloadTypes returns the PayloadTypes field value if set, zero value otherwise.
func (o *InlineResponse200141) GetPayloadTypes() []string {
	if o == nil || isNil(o.PayloadTypes) {
		var ret []string
		return ret
	}
	return o.PayloadTypes
}

// GetPayloadTypesOk returns a tuple with the PayloadTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200141) GetPayloadTypesOk() ([]string, bool) {
	if o == nil || isNil(o.PayloadTypes) {
    return nil, false
	}
	return o.PayloadTypes, true
}

// HasPayloadTypes returns a boolean if a field has been set.
func (o *InlineResponse200141) HasPayloadTypes() bool {
	if o != nil && !isNil(o.PayloadTypes) {
		return true
	}

	return false
}

// SetPayloadTypes gets a reference to the given []string and assigns it to the PayloadTypes field.
func (o *InlineResponse200141) SetPayloadTypes(v []string) {
	o.PayloadTypes = v
}

func (o InlineResponse200141) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.PayloadTypes) {
		toSerialize["payloadTypes"] = o.PayloadTypes
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200141 struct {
	value *InlineResponse200141
	isSet bool
}

func (v NullableInlineResponse200141) Get() *InlineResponse200141 {
	return v.value
}

func (v *NullableInlineResponse200141) Set(val *InlineResponse200141) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200141) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200141) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200141(val *InlineResponse200141) *NullableInlineResponse200141 {
	return &NullableInlineResponse200141{value: val, isSet: true}
}

func (v NullableInlineResponse200141) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200141) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


