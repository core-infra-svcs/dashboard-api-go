/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 07 February, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20086 struct for InlineResponse20086
type InlineResponse20086 struct {
	Alerts *InlineResponse20086Alerts `json:"alerts,omitempty"`
	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	DefaultPolicy *string `json:"defaultPolicy,omitempty"`
	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set       to allow.An empty array will clear the entries.
	BlockedServers []string `json:"blockedServers,omitempty"`
	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set       to block.An empty array will clear the entries.
	AllowedServers []string `json:"allowedServers,omitempty"`
	ArpInspection *InlineResponse20086ArpInspection `json:"arpInspection,omitempty"`
}

// NewInlineResponse20086 instantiates a new InlineResponse20086 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20086() *InlineResponse20086 {
	this := InlineResponse20086{}
	return &this
}

// NewInlineResponse20086WithDefaults instantiates a new InlineResponse20086 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20086WithDefaults() *InlineResponse20086 {
	this := InlineResponse20086{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *InlineResponse20086) GetAlerts() InlineResponse20086Alerts {
	if o == nil || isNil(o.Alerts) {
		var ret InlineResponse20086Alerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetAlertsOk() (*InlineResponse20086Alerts, bool) {
	if o == nil || isNil(o.Alerts) {
    return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *InlineResponse20086) HasAlerts() bool {
	if o != nil && !isNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given InlineResponse20086Alerts and assigns it to the Alerts field.
func (o *InlineResponse20086) SetAlerts(v InlineResponse20086Alerts) {
	o.Alerts = &v
}

// GetDefaultPolicy returns the DefaultPolicy field value if set, zero value otherwise.
func (o *InlineResponse20086) GetDefaultPolicy() string {
	if o == nil || isNil(o.DefaultPolicy) {
		var ret string
		return ret
	}
	return *o.DefaultPolicy
}

// GetDefaultPolicyOk returns a tuple with the DefaultPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetDefaultPolicyOk() (*string, bool) {
	if o == nil || isNil(o.DefaultPolicy) {
    return nil, false
	}
	return o.DefaultPolicy, true
}

// HasDefaultPolicy returns a boolean if a field has been set.
func (o *InlineResponse20086) HasDefaultPolicy() bool {
	if o != nil && !isNil(o.DefaultPolicy) {
		return true
	}

	return false
}

// SetDefaultPolicy gets a reference to the given string and assigns it to the DefaultPolicy field.
func (o *InlineResponse20086) SetDefaultPolicy(v string) {
	o.DefaultPolicy = &v
}

// GetBlockedServers returns the BlockedServers field value if set, zero value otherwise.
func (o *InlineResponse20086) GetBlockedServers() []string {
	if o == nil || isNil(o.BlockedServers) {
		var ret []string
		return ret
	}
	return o.BlockedServers
}

// GetBlockedServersOk returns a tuple with the BlockedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetBlockedServersOk() ([]string, bool) {
	if o == nil || isNil(o.BlockedServers) {
    return nil, false
	}
	return o.BlockedServers, true
}

// HasBlockedServers returns a boolean if a field has been set.
func (o *InlineResponse20086) HasBlockedServers() bool {
	if o != nil && !isNil(o.BlockedServers) {
		return true
	}

	return false
}

// SetBlockedServers gets a reference to the given []string and assigns it to the BlockedServers field.
func (o *InlineResponse20086) SetBlockedServers(v []string) {
	o.BlockedServers = v
}

// GetAllowedServers returns the AllowedServers field value if set, zero value otherwise.
func (o *InlineResponse20086) GetAllowedServers() []string {
	if o == nil || isNil(o.AllowedServers) {
		var ret []string
		return ret
	}
	return o.AllowedServers
}

// GetAllowedServersOk returns a tuple with the AllowedServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetAllowedServersOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedServers) {
    return nil, false
	}
	return o.AllowedServers, true
}

// HasAllowedServers returns a boolean if a field has been set.
func (o *InlineResponse20086) HasAllowedServers() bool {
	if o != nil && !isNil(o.AllowedServers) {
		return true
	}

	return false
}

// SetAllowedServers gets a reference to the given []string and assigns it to the AllowedServers field.
func (o *InlineResponse20086) SetAllowedServers(v []string) {
	o.AllowedServers = v
}

// GetArpInspection returns the ArpInspection field value if set, zero value otherwise.
func (o *InlineResponse20086) GetArpInspection() InlineResponse20086ArpInspection {
	if o == nil || isNil(o.ArpInspection) {
		var ret InlineResponse20086ArpInspection
		return ret
	}
	return *o.ArpInspection
}

// GetArpInspectionOk returns a tuple with the ArpInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20086) GetArpInspectionOk() (*InlineResponse20086ArpInspection, bool) {
	if o == nil || isNil(o.ArpInspection) {
    return nil, false
	}
	return o.ArpInspection, true
}

// HasArpInspection returns a boolean if a field has been set.
func (o *InlineResponse20086) HasArpInspection() bool {
	if o != nil && !isNil(o.ArpInspection) {
		return true
	}

	return false
}

// SetArpInspection gets a reference to the given InlineResponse20086ArpInspection and assigns it to the ArpInspection field.
func (o *InlineResponse20086) SetArpInspection(v InlineResponse20086ArpInspection) {
	o.ArpInspection = &v
}

func (o InlineResponse20086) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !isNil(o.DefaultPolicy) {
		toSerialize["defaultPolicy"] = o.DefaultPolicy
	}
	if !isNil(o.BlockedServers) {
		toSerialize["blockedServers"] = o.BlockedServers
	}
	if !isNil(o.AllowedServers) {
		toSerialize["allowedServers"] = o.AllowedServers
	}
	if !isNil(o.ArpInspection) {
		toSerialize["arpInspection"] = o.ArpInspection
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20086 struct {
	value *InlineResponse20086
	isSet bool
}

func (v NullableInlineResponse20086) Get() *InlineResponse20086 {
	return v.value
}

func (v *NullableInlineResponse20086) Set(val *InlineResponse20086) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20086) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20086) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20086(val *InlineResponse20086) *NullableInlineResponse20086 {
	return &NullableInlineResponse20086{value: val, isSet: true}
}

func (v NullableInlineResponse20086) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20086) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


