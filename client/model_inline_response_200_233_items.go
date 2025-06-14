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

// InlineResponse200233Items struct for InlineResponse200233Items
type InlineResponse200233Items struct {
	// SLA policy ID
	Id *string `json:"id,omitempty"`
	// SLA policy name
	Name *string `json:"name,omitempty"`
	// Endpoint for testing SLA
	Uri *string `json:"uri,omitempty"`
	Ipsec *InlineResponse200233Ipsec `json:"ipsec,omitempty"`
}

// NewInlineResponse200233Items instantiates a new InlineResponse200233Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200233Items() *InlineResponse200233Items {
	this := InlineResponse200233Items{}
	return &this
}

// NewInlineResponse200233ItemsWithDefaults instantiates a new InlineResponse200233Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200233ItemsWithDefaults() *InlineResponse200233Items {
	this := InlineResponse200233Items{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200233Items) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200233Items) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200233Items) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200233Items) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200233Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200233Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200233Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200233Items) SetName(v string) {
	o.Name = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *InlineResponse200233Items) GetUri() string {
	if o == nil || isNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200233Items) GetUriOk() (*string, bool) {
	if o == nil || isNil(o.Uri) {
    return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *InlineResponse200233Items) HasUri() bool {
	if o != nil && !isNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *InlineResponse200233Items) SetUri(v string) {
	o.Uri = &v
}

// GetIpsec returns the Ipsec field value if set, zero value otherwise.
func (o *InlineResponse200233Items) GetIpsec() InlineResponse200233Ipsec {
	if o == nil || isNil(o.Ipsec) {
		var ret InlineResponse200233Ipsec
		return ret
	}
	return *o.Ipsec
}

// GetIpsecOk returns a tuple with the Ipsec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200233Items) GetIpsecOk() (*InlineResponse200233Ipsec, bool) {
	if o == nil || isNil(o.Ipsec) {
    return nil, false
	}
	return o.Ipsec, true
}

// HasIpsec returns a boolean if a field has been set.
func (o *InlineResponse200233Items) HasIpsec() bool {
	if o != nil && !isNil(o.Ipsec) {
		return true
	}

	return false
}

// SetIpsec gets a reference to the given InlineResponse200233Ipsec and assigns it to the Ipsec field.
func (o *InlineResponse200233Items) SetIpsec(v InlineResponse200233Ipsec) {
	o.Ipsec = &v
}

func (o InlineResponse200233Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !isNil(o.Ipsec) {
		toSerialize["ipsec"] = o.Ipsec
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200233Items struct {
	value *InlineResponse200233Items
	isSet bool
}

func (v NullableInlineResponse200233Items) Get() *InlineResponse200233Items {
	return v.value
}

func (v *NullableInlineResponse200233Items) Set(val *InlineResponse200233Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200233Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200233Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200233Items(val *InlineResponse200233Items) *NullableInlineResponse200233Items {
	return &NullableInlineResponse200233Items{value: val, isSet: true}
}

func (v NullableInlineResponse200233Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200233Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


