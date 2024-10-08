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

// InlineResponse200253 struct for InlineResponse200253
type InlineResponse200253 struct {
	// ID of Early Access Feature
	Id *string `json:"id,omitempty"`
	// Name of Early Access Feature
	ShortName *string `json:"shortName,omitempty"`
	// Networks assigned to the Early Access Feature
	LimitScopeToNetworks []InlineResponse200253LimitScopeToNetworks `json:"limitScopeToNetworks,omitempty"`
	OptOutEligibility *InlineResponse200253OptOutEligibility `json:"optOutEligibility,omitempty"`
	// Time when Early Access Feature was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewInlineResponse200253 instantiates a new InlineResponse200253 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200253() *InlineResponse200253 {
	this := InlineResponse200253{}
	return &this
}

// NewInlineResponse200253WithDefaults instantiates a new InlineResponse200253 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200253WithDefaults() *InlineResponse200253 {
	this := InlineResponse200253{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200253) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200253) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200253) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200253) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse200253) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200253) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse200253) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse200253) SetShortName(v string) {
	o.ShortName = &v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *InlineResponse200253) GetLimitScopeToNetworks() []InlineResponse200253LimitScopeToNetworks {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []InlineResponse200253LimitScopeToNetworks
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200253) GetLimitScopeToNetworksOk() ([]InlineResponse200253LimitScopeToNetworks, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *InlineResponse200253) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []InlineResponse200253LimitScopeToNetworks and assigns it to the LimitScopeToNetworks field.
func (o *InlineResponse200253) SetLimitScopeToNetworks(v []InlineResponse200253LimitScopeToNetworks) {
	o.LimitScopeToNetworks = v
}

// GetOptOutEligibility returns the OptOutEligibility field value if set, zero value otherwise.
func (o *InlineResponse200253) GetOptOutEligibility() InlineResponse200253OptOutEligibility {
	if o == nil || isNil(o.OptOutEligibility) {
		var ret InlineResponse200253OptOutEligibility
		return ret
	}
	return *o.OptOutEligibility
}

// GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200253) GetOptOutEligibilityOk() (*InlineResponse200253OptOutEligibility, bool) {
	if o == nil || isNil(o.OptOutEligibility) {
    return nil, false
	}
	return o.OptOutEligibility, true
}

// HasOptOutEligibility returns a boolean if a field has been set.
func (o *InlineResponse200253) HasOptOutEligibility() bool {
	if o != nil && !isNil(o.OptOutEligibility) {
		return true
	}

	return false
}

// SetOptOutEligibility gets a reference to the given InlineResponse200253OptOutEligibility and assigns it to the OptOutEligibility field.
func (o *InlineResponse200253) SetOptOutEligibility(v InlineResponse200253OptOutEligibility) {
	o.OptOutEligibility = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200253) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200253) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200253) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200253) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o InlineResponse200253) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200253 struct {
	value *InlineResponse200253
	isSet bool
}

func (v NullableInlineResponse200253) Get() *InlineResponse200253 {
	return v.value
}

func (v *NullableInlineResponse200253) Set(val *InlineResponse200253) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200253) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200253) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200253(val *InlineResponse200253) *NullableInlineResponse200253 {
	return &NullableInlineResponse200253{value: val, isSet: true}
}

func (v NullableInlineResponse200253) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200253) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


