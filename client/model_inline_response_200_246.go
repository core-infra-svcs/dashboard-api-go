/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200246 struct for InlineResponse200246
type InlineResponse200246 struct {
	// ID of Early Access Feature
	Id *string `json:"id,omitempty"`
	// Name of Early Access Feature
	ShortName *string `json:"shortName,omitempty"`
	// Networks assigned to the Early Access Feature
	LimitScopeToNetworks []InlineResponse200246LimitScopeToNetworks `json:"limitScopeToNetworks,omitempty"`
	OptOutEligibility *InlineResponse200246OptOutEligibility `json:"optOutEligibility,omitempty"`
	// Time when Early Access Feature was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewInlineResponse200246 instantiates a new InlineResponse200246 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200246() *InlineResponse200246 {
	this := InlineResponse200246{}
	return &this
}

// NewInlineResponse200246WithDefaults instantiates a new InlineResponse200246 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200246WithDefaults() *InlineResponse200246 {
	this := InlineResponse200246{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200246) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200246) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200246) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse200246) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse200246) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse200246) SetShortName(v string) {
	o.ShortName = &v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *InlineResponse200246) GetLimitScopeToNetworks() []InlineResponse200246LimitScopeToNetworks {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []InlineResponse200246LimitScopeToNetworks
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246) GetLimitScopeToNetworksOk() ([]InlineResponse200246LimitScopeToNetworks, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *InlineResponse200246) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []InlineResponse200246LimitScopeToNetworks and assigns it to the LimitScopeToNetworks field.
func (o *InlineResponse200246) SetLimitScopeToNetworks(v []InlineResponse200246LimitScopeToNetworks) {
	o.LimitScopeToNetworks = v
}

// GetOptOutEligibility returns the OptOutEligibility field value if set, zero value otherwise.
func (o *InlineResponse200246) GetOptOutEligibility() InlineResponse200246OptOutEligibility {
	if o == nil || isNil(o.OptOutEligibility) {
		var ret InlineResponse200246OptOutEligibility
		return ret
	}
	return *o.OptOutEligibility
}

// GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246) GetOptOutEligibilityOk() (*InlineResponse200246OptOutEligibility, bool) {
	if o == nil || isNil(o.OptOutEligibility) {
    return nil, false
	}
	return o.OptOutEligibility, true
}

// HasOptOutEligibility returns a boolean if a field has been set.
func (o *InlineResponse200246) HasOptOutEligibility() bool {
	if o != nil && !isNil(o.OptOutEligibility) {
		return true
	}

	return false
}

// SetOptOutEligibility gets a reference to the given InlineResponse200246OptOutEligibility and assigns it to the OptOutEligibility field.
func (o *InlineResponse200246) SetOptOutEligibility(v InlineResponse200246OptOutEligibility) {
	o.OptOutEligibility = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200246) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200246) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200246) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200246) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o InlineResponse200246) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ShortName) {
		toSerialize["shortName"] = o.ShortName
	}
	if !isNil(o.LimitScopeToNetworks) {
		toSerialize["limitScopeToNetworks"] = o.LimitScopeToNetworks
	}
	if !isNil(o.OptOutEligibility) {
		toSerialize["optOutEligibility"] = o.OptOutEligibility
	}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200246 struct {
	value *InlineResponse200246
	isSet bool
}

func (v NullableInlineResponse200246) Get() *InlineResponse200246 {
	return v.value
}

func (v *NullableInlineResponse200246) Set(val *InlineResponse200246) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200246) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200246) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200246(val *InlineResponse200246) *NullableInlineResponse200246 {
	return &NullableInlineResponse200246{value: val, isSet: true}
}

func (v NullableInlineResponse200246) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200246) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

