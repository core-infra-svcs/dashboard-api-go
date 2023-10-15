/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20068 struct for InlineResponse20068
type InlineResponse20068 struct {
	// When the Meraki record for the wlanList was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The Meraki managed Id of the wlanList record.
	Id *string `json:"id,omitempty"`
	// An XML string containing the WLAN List for the device.
	Xml *string `json:"xml,omitempty"`
}

// NewInlineResponse20068 instantiates a new InlineResponse20068 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20068() *InlineResponse20068 {
	this := InlineResponse20068{}
	return &this
}

// NewInlineResponse20068WithDefaults instantiates a new InlineResponse20068 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20068WithDefaults() *InlineResponse20068 {
	this := InlineResponse20068{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InlineResponse20068) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InlineResponse20068) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *InlineResponse20068) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20068) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20068) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20068) SetId(v string) {
	o.Id = &v
}

// GetXml returns the Xml field value if set, zero value otherwise.
func (o *InlineResponse20068) GetXml() string {
	if o == nil || isNil(o.Xml) {
		var ret string
		return ret
	}
	return *o.Xml
}

// GetXmlOk returns a tuple with the Xml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetXmlOk() (*string, bool) {
	if o == nil || isNil(o.Xml) {
    return nil, false
	}
	return o.Xml, true
}

// HasXml returns a boolean if a field has been set.
func (o *InlineResponse20068) HasXml() bool {
	if o != nil && !isNil(o.Xml) {
		return true
	}

	return false
}

// SetXml gets a reference to the given string and assigns it to the Xml field.
func (o *InlineResponse20068) SetXml(v string) {
	o.Xml = &v
}

func (o InlineResponse20068) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20068 struct {
	value *InlineResponse20068
	isSet bool
}

func (v NullableInlineResponse20068) Get() *InlineResponse20068 {
	return v.value
}

func (v *NullableInlineResponse20068) Set(val *InlineResponse20068) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20068) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20068) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20068(val *InlineResponse20068) *NullableInlineResponse20068 {
	return &NullableInlineResponse20068{value: val, isSet: true}
}

func (v NullableInlineResponse20068) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20068) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


