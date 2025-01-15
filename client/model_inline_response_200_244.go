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

// InlineResponse200244 struct for InlineResponse200244
type InlineResponse200244 struct {
	// The ID of the network or config template to copy configuration from
	Id *string `json:"id,omitempty"`
	// The name of the configuration template
	Name *string `json:"name,omitempty"`
	// The product types of the configuration template
	ProductTypes []string `json:"productTypes,omitempty"`
	// The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template
	TimeZone *string `json:"timeZone,omitempty"`
}

// NewInlineResponse200244 instantiates a new InlineResponse200244 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200244() *InlineResponse200244 {
	this := InlineResponse200244{}
	return &this
}

// NewInlineResponse200244WithDefaults instantiates a new InlineResponse200244 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200244WithDefaults() *InlineResponse200244 {
	this := InlineResponse200244{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200244) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200244) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200244) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200244) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200244) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200244) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200244) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200244) SetName(v string) {
	o.Name = &v
}

// GetProductTypes returns the ProductTypes field value if set, zero value otherwise.
func (o *InlineResponse200244) GetProductTypes() []string {
	if o == nil || isNil(o.ProductTypes) {
		var ret []string
		return ret
	}
	return o.ProductTypes
}

// GetProductTypesOk returns a tuple with the ProductTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200244) GetProductTypesOk() ([]string, bool) {
	if o == nil || isNil(o.ProductTypes) {
    return nil, false
	}
	return o.ProductTypes, true
}

// HasProductTypes returns a boolean if a field has been set.
func (o *InlineResponse200244) HasProductTypes() bool {
	if o != nil && !isNil(o.ProductTypes) {
		return true
	}

	return false
}

// SetProductTypes gets a reference to the given []string and assigns it to the ProductTypes field.
func (o *InlineResponse200244) SetProductTypes(v []string) {
	o.ProductTypes = v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *InlineResponse200244) GetTimeZone() string {
	if o == nil || isNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200244) GetTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.TimeZone) {
    return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *InlineResponse200244) HasTimeZone() bool {
	if o != nil && !isNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *InlineResponse200244) SetTimeZone(v string) {
	o.TimeZone = &v
}

func (o InlineResponse200244) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.ProductTypes) {
		toSerialize["productTypes"] = o.ProductTypes
	}
	if !isNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200244 struct {
	value *InlineResponse200244
	isSet bool
}

func (v NullableInlineResponse200244) Get() *InlineResponse200244 {
	return v.value
}

func (v *NullableInlineResponse200244) Set(val *InlineResponse200244) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200244) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200244) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200244(val *InlineResponse200244) *NullableInlineResponse200244 {
	return &NullableInlineResponse200244{value: val, isSet: true}
}

func (v NullableInlineResponse200244) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200244) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


