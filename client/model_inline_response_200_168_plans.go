/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 03 July, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.48.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200168Plans struct for InlineResponse200168Plans
type InlineResponse200168Plans struct {
	// The id of the pricing plan to update.
	Id *string `json:"id,omitempty"`
	// The price of the billing plan.
	Price *float32 `json:"price,omitempty"`
	BandwidthLimits *InlineResponse200168BandwidthLimits `json:"bandwidthLimits,omitempty"`
	// The time limit of the pricing plan in minutes.
	TimeLimit *string `json:"timeLimit,omitempty"`
}

// NewInlineResponse200168Plans instantiates a new InlineResponse200168Plans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200168Plans() *InlineResponse200168Plans {
	this := InlineResponse200168Plans{}
	return &this
}

// NewInlineResponse200168PlansWithDefaults instantiates a new InlineResponse200168Plans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200168PlansWithDefaults() *InlineResponse200168Plans {
	this := InlineResponse200168Plans{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200168Plans) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168Plans) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200168Plans) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200168Plans) SetId(v string) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *InlineResponse200168Plans) GetPrice() float32 {
	if o == nil || isNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168Plans) GetPriceOk() (*float32, bool) {
	if o == nil || isNil(o.Price) {
    return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *InlineResponse200168Plans) HasPrice() bool {
	if o != nil && !isNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *InlineResponse200168Plans) SetPrice(v float32) {
	o.Price = &v
}

// GetBandwidthLimits returns the BandwidthLimits field value if set, zero value otherwise.
func (o *InlineResponse200168Plans) GetBandwidthLimits() InlineResponse200168BandwidthLimits {
	if o == nil || isNil(o.BandwidthLimits) {
		var ret InlineResponse200168BandwidthLimits
		return ret
	}
	return *o.BandwidthLimits
}

// GetBandwidthLimitsOk returns a tuple with the BandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168Plans) GetBandwidthLimitsOk() (*InlineResponse200168BandwidthLimits, bool) {
	if o == nil || isNil(o.BandwidthLimits) {
    return nil, false
	}
	return o.BandwidthLimits, true
}

// HasBandwidthLimits returns a boolean if a field has been set.
func (o *InlineResponse200168Plans) HasBandwidthLimits() bool {
	if o != nil && !isNil(o.BandwidthLimits) {
		return true
	}

	return false
}

// SetBandwidthLimits gets a reference to the given InlineResponse200168BandwidthLimits and assigns it to the BandwidthLimits field.
func (o *InlineResponse200168Plans) SetBandwidthLimits(v InlineResponse200168BandwidthLimits) {
	o.BandwidthLimits = &v
}

// GetTimeLimit returns the TimeLimit field value if set, zero value otherwise.
func (o *InlineResponse200168Plans) GetTimeLimit() string {
	if o == nil || isNil(o.TimeLimit) {
		var ret string
		return ret
	}
	return *o.TimeLimit
}

// GetTimeLimitOk returns a tuple with the TimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200168Plans) GetTimeLimitOk() (*string, bool) {
	if o == nil || isNil(o.TimeLimit) {
    return nil, false
	}
	return o.TimeLimit, true
}

// HasTimeLimit returns a boolean if a field has been set.
func (o *InlineResponse200168Plans) HasTimeLimit() bool {
	if o != nil && !isNil(o.TimeLimit) {
		return true
	}

	return false
}

// SetTimeLimit gets a reference to the given string and assigns it to the TimeLimit field.
func (o *InlineResponse200168Plans) SetTimeLimit(v string) {
	o.TimeLimit = &v
}

func (o InlineResponse200168Plans) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !isNil(o.BandwidthLimits) {
		toSerialize["bandwidthLimits"] = o.BandwidthLimits
	}
	if !isNil(o.TimeLimit) {
		toSerialize["timeLimit"] = o.TimeLimit
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200168Plans struct {
	value *InlineResponse200168Plans
	isSet bool
}

func (v NullableInlineResponse200168Plans) Get() *InlineResponse200168Plans {
	return v.value
}

func (v *NullableInlineResponse200168Plans) Set(val *InlineResponse200168Plans) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200168Plans) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200168Plans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200168Plans(val *InlineResponse200168Plans) *NullableInlineResponse200168Plans {
	return &NullableInlineResponse200168Plans{value: val, isSet: true}
}

func (v NullableInlineResponse200168Plans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200168Plans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


