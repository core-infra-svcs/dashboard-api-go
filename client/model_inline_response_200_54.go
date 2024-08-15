/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 August, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.49.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20054 struct for InlineResponse20054
type InlineResponse20054 struct {
	// Static delegated prefix id.
	StaticDelegatedPrefixId *string `json:"staticDelegatedPrefixId,omitempty"`
	// IPv6 prefix/prefix length.
	Prefix *string `json:"prefix,omitempty"`
	Origin *NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin `json:"origin,omitempty"`
	// Identifying description for the prefix.
	Description *string `json:"description,omitempty"`
	// Prefix creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Prefix Updated time.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewInlineResponse20054 instantiates a new InlineResponse20054 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20054() *InlineResponse20054 {
	this := InlineResponse20054{}
	return &this
}

// NewInlineResponse20054WithDefaults instantiates a new InlineResponse20054 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20054WithDefaults() *InlineResponse20054 {
	this := InlineResponse20054{}
	return &this
}

// GetStaticDelegatedPrefixId returns the StaticDelegatedPrefixId field value if set, zero value otherwise.
func (o *InlineResponse20054) GetStaticDelegatedPrefixId() string {
	if o == nil || isNil(o.StaticDelegatedPrefixId) {
		var ret string
		return ret
	}
	return *o.StaticDelegatedPrefixId
}

// GetStaticDelegatedPrefixIdOk returns a tuple with the StaticDelegatedPrefixId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetStaticDelegatedPrefixIdOk() (*string, bool) {
	if o == nil || isNil(o.StaticDelegatedPrefixId) {
    return nil, false
	}
	return o.StaticDelegatedPrefixId, true
}

// HasStaticDelegatedPrefixId returns a boolean if a field has been set.
func (o *InlineResponse20054) HasStaticDelegatedPrefixId() bool {
	if o != nil && !isNil(o.StaticDelegatedPrefixId) {
		return true
	}

	return false
}

// SetStaticDelegatedPrefixId gets a reference to the given string and assigns it to the StaticDelegatedPrefixId field.
func (o *InlineResponse20054) SetStaticDelegatedPrefixId(v string) {
	o.StaticDelegatedPrefixId = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *InlineResponse20054) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *InlineResponse20054) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *InlineResponse20054) SetPrefix(v string) {
	o.Prefix = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *InlineResponse20054) GetOrigin() NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin {
	if o == nil || isNil(o.Origin) {
		var ret NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetOriginOk() (*NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin, bool) {
	if o == nil || isNil(o.Origin) {
    return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *InlineResponse20054) HasOrigin() bool {
	if o != nil && !isNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin and assigns it to the Origin field.
func (o *InlineResponse20054) SetOrigin(v NetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin) {
	o.Origin = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InlineResponse20054) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InlineResponse20054) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InlineResponse20054) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse20054) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse20054) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse20054) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *InlineResponse20054) GetUpdatedAt() time.Time {
	if o == nil || isNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20054) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.UpdatedAt) {
    return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *InlineResponse20054) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *InlineResponse20054) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o InlineResponse20054) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StaticDelegatedPrefixId) {
		toSerialize["staticDelegatedPrefixId"] = o.StaticDelegatedPrefixId
	}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20054 struct {
	value *InlineResponse20054
	isSet bool
}

func (v NullableInlineResponse20054) Get() *InlineResponse20054 {
	return v.value
}

func (v *NullableInlineResponse20054) Set(val *InlineResponse20054) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20054) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20054) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20054(val *InlineResponse20054) *NullableInlineResponse20054 {
	return &NullableInlineResponse20054{value: val, isSet: true}
}

func (v NullableInlineResponse20054) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20054) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


