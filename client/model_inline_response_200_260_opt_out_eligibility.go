/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 January, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.54.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200260OptOutEligibility Descriptions of the early access feature
type InlineResponse200260OptOutEligibility struct {
	// Condition flag to opt out from the feature
	Eligible *bool `json:"eligible,omitempty"`
	// User friendly message regarding opt-out eligibility
	Reason *string `json:"reason,omitempty"`
	Help *InlineResponse200260OptOutEligibilityHelp `json:"help,omitempty"`
}

// NewInlineResponse200260OptOutEligibility instantiates a new InlineResponse200260OptOutEligibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200260OptOutEligibility() *InlineResponse200260OptOutEligibility {
	this := InlineResponse200260OptOutEligibility{}
	return &this
}

// NewInlineResponse200260OptOutEligibilityWithDefaults instantiates a new InlineResponse200260OptOutEligibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200260OptOutEligibilityWithDefaults() *InlineResponse200260OptOutEligibility {
	this := InlineResponse200260OptOutEligibility{}
	return &this
}

// GetEligible returns the Eligible field value if set, zero value otherwise.
func (o *InlineResponse200260OptOutEligibility) GetEligible() bool {
	if o == nil || isNil(o.Eligible) {
		var ret bool
		return ret
	}
	return *o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260OptOutEligibility) GetEligibleOk() (*bool, bool) {
	if o == nil || isNil(o.Eligible) {
    return nil, false
	}
	return o.Eligible, true
}

// HasEligible returns a boolean if a field has been set.
func (o *InlineResponse200260OptOutEligibility) HasEligible() bool {
	if o != nil && !isNil(o.Eligible) {
		return true
	}

	return false
}

// SetEligible gets a reference to the given bool and assigns it to the Eligible field.
func (o *InlineResponse200260OptOutEligibility) SetEligible(v bool) {
	o.Eligible = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InlineResponse200260OptOutEligibility) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260OptOutEligibility) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InlineResponse200260OptOutEligibility) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InlineResponse200260OptOutEligibility) SetReason(v string) {
	o.Reason = &v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *InlineResponse200260OptOutEligibility) GetHelp() InlineResponse200260OptOutEligibilityHelp {
	if o == nil || isNil(o.Help) {
		var ret InlineResponse200260OptOutEligibilityHelp
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200260OptOutEligibility) GetHelpOk() (*InlineResponse200260OptOutEligibilityHelp, bool) {
	if o == nil || isNil(o.Help) {
    return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *InlineResponse200260OptOutEligibility) HasHelp() bool {
	if o != nil && !isNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given InlineResponse200260OptOutEligibilityHelp and assigns it to the Help field.
func (o *InlineResponse200260OptOutEligibility) SetHelp(v InlineResponse200260OptOutEligibilityHelp) {
	o.Help = &v
}

func (o InlineResponse200260OptOutEligibility) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Eligible) {
		toSerialize["eligible"] = o.Eligible
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.Help) {
		toSerialize["help"] = o.Help
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200260OptOutEligibility struct {
	value *InlineResponse200260OptOutEligibility
	isSet bool
}

func (v NullableInlineResponse200260OptOutEligibility) Get() *InlineResponse200260OptOutEligibility {
	return v.value
}

func (v *NullableInlineResponse200260OptOutEligibility) Set(val *InlineResponse200260OptOutEligibility) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200260OptOutEligibility) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200260OptOutEligibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200260OptOutEligibility(val *InlineResponse200260OptOutEligibility) *NullableInlineResponse200260OptOutEligibility {
	return &NullableInlineResponse200260OptOutEligibility{value: val, isSet: true}
}

func (v NullableInlineResponse200260OptOutEligibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200260OptOutEligibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

