/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 June, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.59.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesSchedulingThursday The schedule object for Thursday.
type NetworksNetworkIdGroupPoliciesSchedulingThursday struct {
	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active *bool `json:"active,omitempty"`
	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From *string `json:"from,omitempty"`
	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To *string `json:"to,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesSchedulingThursday instantiates a new NetworksNetworkIdGroupPoliciesSchedulingThursday object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesSchedulingThursday() *NetworksNetworkIdGroupPoliciesSchedulingThursday {
	this := NetworksNetworkIdGroupPoliciesSchedulingThursday{}
	return &this
}

// NewNetworksNetworkIdGroupPoliciesSchedulingThursdayWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesSchedulingThursday object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesSchedulingThursdayWithDefaults() *NetworksNetworkIdGroupPoliciesSchedulingThursday {
	this := NetworksNetworkIdGroupPoliciesSchedulingThursday{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) GetActive() bool {
	if o == nil || isNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) GetActiveOk() (*bool, bool) {
	if o == nil || isNil(o.Active) {
    return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) HasActive() bool {
	if o != nil && !isNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) SetActive(v bool) {
	o.Active = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) GetFrom() string {
	if o == nil || isNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) GetFromOk() (*string, bool) {
	if o == nil || isNil(o.From) {
    return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) HasFrom() bool {
	if o != nil && !isNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) GetTo() string {
	if o == nil || isNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) GetToOk() (*string, bool) {
	if o == nil || isNil(o.To) {
    return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) HasTo() bool {
	if o != nil && !isNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *NetworksNetworkIdGroupPoliciesSchedulingThursday) SetTo(v string) {
	o.To = &v
}

func (o NetworksNetworkIdGroupPoliciesSchedulingThursday) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !isNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !isNil(o.To) {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesSchedulingThursday struct {
	value *NetworksNetworkIdGroupPoliciesSchedulingThursday
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesSchedulingThursday) Get() *NetworksNetworkIdGroupPoliciesSchedulingThursday {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesSchedulingThursday) Set(val *NetworksNetworkIdGroupPoliciesSchedulingThursday) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesSchedulingThursday) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesSchedulingThursday) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesSchedulingThursday(val *NetworksNetworkIdGroupPoliciesSchedulingThursday) *NullableNetworksNetworkIdGroupPoliciesSchedulingThursday {
	return &NullableNetworksNetworkIdGroupPoliciesSchedulingThursday{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesSchedulingThursday) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesSchedulingThursday) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


