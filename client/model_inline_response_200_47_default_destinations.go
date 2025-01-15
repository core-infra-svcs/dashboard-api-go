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

// InlineResponse20047DefaultDestinations The network-wide destinations for all alerts on the network.
type InlineResponse20047DefaultDestinations struct {
	// A list of emails that will receive the alert(s).
	Emails []string `json:"emails,omitempty"`
	// If true, then all network admins will receive emails.
	AllAdmins *bool `json:"allAdmins,omitempty"`
	// If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.
	Snmp *bool `json:"snmp,omitempty"`
	// A list of HTTP server IDs to send a Webhook to
	HttpServerIds []string `json:"httpServerIds,omitempty"`
}

// NewInlineResponse20047DefaultDestinations instantiates a new InlineResponse20047DefaultDestinations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20047DefaultDestinations() *InlineResponse20047DefaultDestinations {
	this := InlineResponse20047DefaultDestinations{}
	return &this
}

// NewInlineResponse20047DefaultDestinationsWithDefaults instantiates a new InlineResponse20047DefaultDestinations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20047DefaultDestinationsWithDefaults() *InlineResponse20047DefaultDestinations {
	this := InlineResponse20047DefaultDestinations{}
	return &this
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *InlineResponse20047DefaultDestinations) GetEmails() []string {
	if o == nil || isNil(o.Emails) {
		var ret []string
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047DefaultDestinations) GetEmailsOk() ([]string, bool) {
	if o == nil || isNil(o.Emails) {
    return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *InlineResponse20047DefaultDestinations) HasEmails() bool {
	if o != nil && !isNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []string and assigns it to the Emails field.
func (o *InlineResponse20047DefaultDestinations) SetEmails(v []string) {
	o.Emails = v
}

// GetAllAdmins returns the AllAdmins field value if set, zero value otherwise.
func (o *InlineResponse20047DefaultDestinations) GetAllAdmins() bool {
	if o == nil || isNil(o.AllAdmins) {
		var ret bool
		return ret
	}
	return *o.AllAdmins
}

// GetAllAdminsOk returns a tuple with the AllAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047DefaultDestinations) GetAllAdminsOk() (*bool, bool) {
	if o == nil || isNil(o.AllAdmins) {
    return nil, false
	}
	return o.AllAdmins, true
}

// HasAllAdmins returns a boolean if a field has been set.
func (o *InlineResponse20047DefaultDestinations) HasAllAdmins() bool {
	if o != nil && !isNil(o.AllAdmins) {
		return true
	}

	return false
}

// SetAllAdmins gets a reference to the given bool and assigns it to the AllAdmins field.
func (o *InlineResponse20047DefaultDestinations) SetAllAdmins(v bool) {
	o.AllAdmins = &v
}

// GetSnmp returns the Snmp field value if set, zero value otherwise.
func (o *InlineResponse20047DefaultDestinations) GetSnmp() bool {
	if o == nil || isNil(o.Snmp) {
		var ret bool
		return ret
	}
	return *o.Snmp
}

// GetSnmpOk returns a tuple with the Snmp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047DefaultDestinations) GetSnmpOk() (*bool, bool) {
	if o == nil || isNil(o.Snmp) {
    return nil, false
	}
	return o.Snmp, true
}

// HasSnmp returns a boolean if a field has been set.
func (o *InlineResponse20047DefaultDestinations) HasSnmp() bool {
	if o != nil && !isNil(o.Snmp) {
		return true
	}

	return false
}

// SetSnmp gets a reference to the given bool and assigns it to the Snmp field.
func (o *InlineResponse20047DefaultDestinations) SetSnmp(v bool) {
	o.Snmp = &v
}

// GetHttpServerIds returns the HttpServerIds field value if set, zero value otherwise.
func (o *InlineResponse20047DefaultDestinations) GetHttpServerIds() []string {
	if o == nil || isNil(o.HttpServerIds) {
		var ret []string
		return ret
	}
	return o.HttpServerIds
}

// GetHttpServerIdsOk returns a tuple with the HttpServerIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20047DefaultDestinations) GetHttpServerIdsOk() ([]string, bool) {
	if o == nil || isNil(o.HttpServerIds) {
    return nil, false
	}
	return o.HttpServerIds, true
}

// HasHttpServerIds returns a boolean if a field has been set.
func (o *InlineResponse20047DefaultDestinations) HasHttpServerIds() bool {
	if o != nil && !isNil(o.HttpServerIds) {
		return true
	}

	return false
}

// SetHttpServerIds gets a reference to the given []string and assigns it to the HttpServerIds field.
func (o *InlineResponse20047DefaultDestinations) SetHttpServerIds(v []string) {
	o.HttpServerIds = v
}

func (o InlineResponse20047DefaultDestinations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !isNil(o.AllAdmins) {
		toSerialize["allAdmins"] = o.AllAdmins
	}
	if !isNil(o.Snmp) {
		toSerialize["snmp"] = o.Snmp
	}
	if !isNil(o.HttpServerIds) {
		toSerialize["httpServerIds"] = o.HttpServerIds
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20047DefaultDestinations struct {
	value *InlineResponse20047DefaultDestinations
	isSet bool
}

func (v NullableInlineResponse20047DefaultDestinations) Get() *InlineResponse20047DefaultDestinations {
	return v.value
}

func (v *NullableInlineResponse20047DefaultDestinations) Set(val *InlineResponse20047DefaultDestinations) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20047DefaultDestinations) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20047DefaultDestinations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20047DefaultDestinations(val *InlineResponse20047DefaultDestinations) *NullableInlineResponse20047DefaultDestinations {
	return &NullableInlineResponse20047DefaultDestinations{value: val, isSet: true}
}

func (v NullableInlineResponse20047DefaultDestinations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20047DefaultDestinations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

