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

// InlineResponse20055ApplicationCategories struct for InlineResponse20055ApplicationCategories
type InlineResponse20055ApplicationCategories struct {
	// The id of the category
	Id *string `json:"id,omitempty"`
	// The name of the category
	Name *string `json:"name,omitempty"`
	// Details of the associated applications
	Applications []InlineResponse20055Applications `json:"applications,omitempty"`
}

// NewInlineResponse20055ApplicationCategories instantiates a new InlineResponse20055ApplicationCategories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20055ApplicationCategories() *InlineResponse20055ApplicationCategories {
	this := InlineResponse20055ApplicationCategories{}
	return &this
}

// NewInlineResponse20055ApplicationCategoriesWithDefaults instantiates a new InlineResponse20055ApplicationCategories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20055ApplicationCategoriesWithDefaults() *InlineResponse20055ApplicationCategories {
	this := InlineResponse20055ApplicationCategories{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse20055ApplicationCategories) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055ApplicationCategories) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse20055ApplicationCategories) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse20055ApplicationCategories) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20055ApplicationCategories) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055ApplicationCategories) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20055ApplicationCategories) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20055ApplicationCategories) SetName(v string) {
	o.Name = &v
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *InlineResponse20055ApplicationCategories) GetApplications() []InlineResponse20055Applications {
	if o == nil || isNil(o.Applications) {
		var ret []InlineResponse20055Applications
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20055ApplicationCategories) GetApplicationsOk() ([]InlineResponse20055Applications, bool) {
	if o == nil || isNil(o.Applications) {
    return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *InlineResponse20055ApplicationCategories) HasApplications() bool {
	if o != nil && !isNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []InlineResponse20055Applications and assigns it to the Applications field.
func (o *InlineResponse20055ApplicationCategories) SetApplications(v []InlineResponse20055Applications) {
	o.Applications = v
}

func (o InlineResponse20055ApplicationCategories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20055ApplicationCategories struct {
	value *InlineResponse20055ApplicationCategories
	isSet bool
}

func (v NullableInlineResponse20055ApplicationCategories) Get() *InlineResponse20055ApplicationCategories {
	return v.value
}

func (v *NullableInlineResponse20055ApplicationCategories) Set(val *InlineResponse20055ApplicationCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20055ApplicationCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20055ApplicationCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20055ApplicationCategories(val *InlineResponse20055ApplicationCategories) *NullableInlineResponse20055ApplicationCategories {
	return &NullableInlineResponse20055ApplicationCategories{value: val, isSet: true}
}

func (v NullableInlineResponse20055ApplicationCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20055ApplicationCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


