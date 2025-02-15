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

// InlineResponse200105 struct for InlineResponse200105
type InlineResponse200105 struct {
	// Alert identifier. Value can be empty
	Id *string `json:"id,omitempty"`
	// Category of the alert
	Category *string `json:"category,omitempty"`
	// Alert type
	Type *string `json:"type,omitempty"`
	// Severity of the alert
	Severity *string `json:"severity,omitempty"`
	Scope *NetworksNetworkIdHealthAlertsScope `json:"scope,omitempty"`
}

// NewInlineResponse200105 instantiates a new InlineResponse200105 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200105() *InlineResponse200105 {
	this := InlineResponse200105{}
	return &this
}

// NewInlineResponse200105WithDefaults instantiates a new InlineResponse200105 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200105WithDefaults() *InlineResponse200105 {
	this := InlineResponse200105{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200105) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200105) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200105) SetId(v string) {
	o.Id = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *InlineResponse200105) GetCategory() string {
	if o == nil || isNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetCategoryOk() (*string, bool) {
	if o == nil || isNil(o.Category) {
    return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *InlineResponse200105) HasCategory() bool {
	if o != nil && !isNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *InlineResponse200105) SetCategory(v string) {
	o.Category = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse200105) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse200105) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse200105) SetType(v string) {
	o.Type = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *InlineResponse200105) GetSeverity() string {
	if o == nil || isNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetSeverityOk() (*string, bool) {
	if o == nil || isNil(o.Severity) {
    return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *InlineResponse200105) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *InlineResponse200105) SetSeverity(v string) {
	o.Severity = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *InlineResponse200105) GetScope() NetworksNetworkIdHealthAlertsScope {
	if o == nil || isNil(o.Scope) {
		var ret NetworksNetworkIdHealthAlertsScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200105) GetScopeOk() (*NetworksNetworkIdHealthAlertsScope, bool) {
	if o == nil || isNil(o.Scope) {
    return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *InlineResponse200105) HasScope() bool {
	if o != nil && !isNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given NetworksNetworkIdHealthAlertsScope and assigns it to the Scope field.
func (o *InlineResponse200105) SetScope(v NetworksNetworkIdHealthAlertsScope) {
	o.Scope = &v
}

func (o InlineResponse200105) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !isNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200105 struct {
	value *InlineResponse200105
	isSet bool
}

func (v NullableInlineResponse200105) Get() *InlineResponse200105 {
	return v.value
}

func (v *NullableInlineResponse200105) Set(val *InlineResponse200105) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200105) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200105) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200105(val *InlineResponse200105) *NullableInlineResponse200105 {
	return &NullableInlineResponse200105{value: val, isSet: true}
}

func (v NullableInlineResponse200105) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200105) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


