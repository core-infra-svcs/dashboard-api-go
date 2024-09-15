/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200131 struct for InlineResponse200131
type InlineResponse200131 struct {
	// When the Meraki record for the wlanList was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The Meraki managed Id of the wlanList record.
	Id *string `json:"id,omitempty"`
	// An XML string containing the WLAN List for the device.
	Xml *string `json:"xml,omitempty"`
}

// NewInlineResponse200131 instantiates a new InlineResponse200131 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200131() *InlineResponse200131 {
	this := InlineResponse200131{}
	return &this
}

// NewInlineResponse200131WithDefaults instantiates a new InlineResponse200131 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200131WithDefaults() *InlineResponse200131 {
	this := InlineResponse200131{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse200131) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse200131) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse200131) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200131) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200131) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200131) SetId(v string) {
	o.Id = &v
}

// GetXml returns the Xml field value if set, zero value otherwise.
func (o *InlineResponse200131) GetXml() string {
	if o == nil || isNil(o.Xml) {
		var ret string
		return ret
	}
	return *o.Xml
}

// GetXmlOk returns a tuple with the Xml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200131) GetXmlOk() (*string, bool) {
	if o == nil || isNil(o.Xml) {
    return nil, false
	}
	return o.Xml, true
}

// HasXml returns a boolean if a field has been set.
func (o *InlineResponse200131) HasXml() bool {
	if o != nil && !isNil(o.Xml) {
		return true
	}

	return false
}

// SetXml gets a reference to the given string and assigns it to the Xml field.
func (o *InlineResponse200131) SetXml(v string) {
	o.Xml = &v
}

func (o InlineResponse200131) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Xml) {
		toSerialize["xml"] = o.Xml
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200131 struct {
	value *InlineResponse200131
	isSet bool
}

func (v NullableInlineResponse200131) Get() *InlineResponse200131 {
	return v.value
}

func (v *NullableInlineResponse200131) Set(val *InlineResponse200131) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200131) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200131) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200131(val *InlineResponse200131) *NullableInlineResponse200131 {
	return &NullableInlineResponse200131{value: val, isSet: true}
}

func (v NullableInlineResponse200131) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200131) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


