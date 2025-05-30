/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 May, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.58.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200280OptOutEligibility Descriptions of the early access feature
type InlineResponse200280OptOutEligibility struct {
	// Condition flag to opt out from the feature
	Eligible *bool `json:"eligible,omitempty"`
	// User friendly message regarding opt-out eligibility
	Reason *string `json:"reason,omitempty"`
	Help *InlineResponse200280OptOutEligibilityHelp `json:"help,omitempty"`
}

// NewInlineResponse200280OptOutEligibility instantiates a new InlineResponse200280OptOutEligibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200280OptOutEligibility() *InlineResponse200280OptOutEligibility {
	this := InlineResponse200280OptOutEligibility{}
	return &this
}

// NewInlineResponse200280OptOutEligibilityWithDefaults instantiates a new InlineResponse200280OptOutEligibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200280OptOutEligibilityWithDefaults() *InlineResponse200280OptOutEligibility {
	this := InlineResponse200280OptOutEligibility{}
	return &this
}

// GetEligible returns the Eligible field value if set, zero value otherwise.
func (o *InlineResponse200280OptOutEligibility) GetEligible() bool {
	if o == nil || isNil(o.Eligible) {
		var ret bool
		return ret
	}
	return *o.Eligible
}

// GetEligibleOk returns a tuple with the Eligible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280OptOutEligibility) GetEligibleOk() (*bool, bool) {
	if o == nil || isNil(o.Eligible) {
    return nil, false
	}
	return o.Eligible, true
}

// HasEligible returns a boolean if a field has been set.
func (o *InlineResponse200280OptOutEligibility) HasEligible() bool {
	if o != nil && !isNil(o.Eligible) {
		return true
	}

	return false
}

// SetEligible gets a reference to the given bool and assigns it to the Eligible field.
func (o *InlineResponse200280OptOutEligibility) SetEligible(v bool) {
	o.Eligible = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *InlineResponse200280OptOutEligibility) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280OptOutEligibility) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *InlineResponse200280OptOutEligibility) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *InlineResponse200280OptOutEligibility) SetReason(v string) {
	o.Reason = &v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *InlineResponse200280OptOutEligibility) GetHelp() InlineResponse200280OptOutEligibilityHelp {
	if o == nil || isNil(o.Help) {
		var ret InlineResponse200280OptOutEligibilityHelp
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200280OptOutEligibility) GetHelpOk() (*InlineResponse200280OptOutEligibilityHelp, bool) {
	if o == nil || isNil(o.Help) {
    return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *InlineResponse200280OptOutEligibility) HasHelp() bool {
	if o != nil && !isNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given InlineResponse200280OptOutEligibilityHelp and assigns it to the Help field.
func (o *InlineResponse200280OptOutEligibility) SetHelp(v InlineResponse200280OptOutEligibilityHelp) {
	o.Help = &v
}

func (o InlineResponse200280OptOutEligibility) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse200280OptOutEligibility struct {
	value *InlineResponse200280OptOutEligibility
	isSet bool
}

func (v NullableInlineResponse200280OptOutEligibility) Get() *InlineResponse200280OptOutEligibility {
	return v.value
}

func (v *NullableInlineResponse200280OptOutEligibility) Set(val *InlineResponse200280OptOutEligibility) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200280OptOutEligibility) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200280OptOutEligibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200280OptOutEligibility(val *InlineResponse200280OptOutEligibility) *NullableInlineResponse200280OptOutEligibility {
	return &NullableInlineResponse200280OptOutEligibility{value: val, isSet: true}
}

func (v NullableInlineResponse200280OptOutEligibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200280OptOutEligibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


