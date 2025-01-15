/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse200260 struct for InlineResponse200260
type InlineResponse200260 struct {
	// ID of Early Access Feature
	Id *string `json:"id,omitempty"`
	// Name of Early Access Feature
	ShortName *string `json:"shortName,omitempty"`
	// Networks assigned to the Early Access Feature
	LimitScopeToNetworks []InlineResponse200260LimitScopeToNetworks `json:"limitScopeToNetworks,omitempty"`
	OptOutEligibility *InlineResponse200260OptOutEligibility `json:"optOutEligibility,omitempty"`
	// Time when Early Access Feature was created
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}

// NewInlineResponse200260 instantiates a new InlineResponse200260 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200260() *InlineResponse200260 {
	this := InlineResponse200260{}
	return &this
}

// NewInlineResponse200260WithDefaults instantiates a new InlineResponse200260 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200260WithDefaults() *InlineResponse200260 {
	this := InlineResponse200260{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200260) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200260) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200260) SetId(v string) {
	o.Id = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *InlineResponse200260) GetShortName() string {
	if o == nil || isNil(o.ShortName) {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260) GetShortNameOk() (*string, bool) {
	if o == nil || isNil(o.ShortName) {
    return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *InlineResponse200260) HasShortName() bool {
	if o != nil && !isNil(o.ShortName) {
		return true
	}

	return false
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *InlineResponse200260) SetShortName(v string) {
	o.ShortName = &v
}

// GetLimitScopeToNetworks returns the LimitScopeToNetworks field value if set, zero value otherwise.
func (o *InlineResponse200260) GetLimitScopeToNetworks() []InlineResponse200260LimitScopeToNetworks {
	if o == nil || isNil(o.LimitScopeToNetworks) {
		var ret []InlineResponse200260LimitScopeToNetworks
		return ret
	}
	return o.LimitScopeToNetworks
}

// GetLimitScopeToNetworksOk returns a tuple with the LimitScopeToNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260) GetLimitScopeToNetworksOk() ([]InlineResponse200260LimitScopeToNetworks, bool) {
	if o == nil || isNil(o.LimitScopeToNetworks) {
    return nil, false
	}
	return o.LimitScopeToNetworks, true
}

// HasLimitScopeToNetworks returns a boolean if a field has been set.
func (o *InlineResponse200260) HasLimitScopeToNetworks() bool {
	if o != nil && !isNil(o.LimitScopeToNetworks) {
		return true
	}

	return false
}

// SetLimitScopeToNetworks gets a reference to the given []InlineResponse200260LimitScopeToNetworks and assigns it to the LimitScopeToNetworks field.
func (o *InlineResponse200260) SetLimitScopeToNetworks(v []InlineResponse200260LimitScopeToNetworks) {
	o.LimitScopeToNetworks = v
}

// GetOptOutEligibility returns the OptOutEligibility field value if set, zero value otherwise.
func (o *InlineResponse200260) GetOptOutEligibility() InlineResponse200260OptOutEligibility {
	if o == nil || isNil(o.OptOutEligibility) {
		var ret InlineResponse200260OptOutEligibility
		return ret
	}
	return *o.OptOutEligibility
}

// GetOptOutEligibilityOk returns a tuple with the OptOutEligibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260) GetOptOutEligibilityOk() (*InlineResponse200260OptOutEligibility, bool) {
	if o == nil || isNil(o.OptOutEligibility) {
    return nil, false
	}
	return o.OptOutEligibility, true
}

// HasOptOutEligibility returns a boolean if a field has been set.
func (o *InlineResponse200260) HasOptOutEligibility() bool {
	if o != nil && !isNil(o.OptOutEligibility) {
		return true
	}

	return false
}

// SetOptOutEligibility gets a reference to the given InlineResponse200260OptOutEligibility and assigns it to the OptOutEligibility field.
func (o *InlineResponse200260) SetOptOutEligibility(v InlineResponse200260OptOutEligibility) {
	o.OptOutEligibility = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200260) GetCreatedAt() time.Time {
	if o == nil || isNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200260) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *InlineResponse200260) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o InlineResponse200260) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200260 struct {
	value *InlineResponse200260
	isSet bool
}

func (v NullableInlineResponse200260) Get() *InlineResponse200260 {
	return v.value
}

func (v *NullableInlineResponse200260) Set(val *InlineResponse200260) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200260) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200260) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200260(val *InlineResponse200260) *NullableInlineResponse200260 {
	return &NullableInlineResponse200260{value: val, isSet: true}
}

func (v NullableInlineResponse200260) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200260) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


