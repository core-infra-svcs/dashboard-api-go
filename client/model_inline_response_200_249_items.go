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

// InlineResponse200249Items struct for InlineResponse200249Items
type InlineResponse200249Items struct {
	// Service provider name.
	Name *string `json:"name,omitempty"`
	Logo *InlineResponse200249Logo `json:"logo,omitempty"`
	// Indicates if service provider is the bootstrap provider.
	IsBootstrap *bool `json:"isBootstrap,omitempty"`
	Terms *InlineResponse200249Terms `json:"terms,omitempty"`
}

// NewInlineResponse200249Items instantiates a new InlineResponse200249Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200249Items() *InlineResponse200249Items {
	this := InlineResponse200249Items{}
	return &this
}

// NewInlineResponse200249ItemsWithDefaults instantiates a new InlineResponse200249Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200249ItemsWithDefaults() *InlineResponse200249Items {
	this := InlineResponse200249Items{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200249Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200249Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200249Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200249Items) SetName(v string) {
	o.Name = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *InlineResponse200249Items) GetLogo() InlineResponse200249Logo {
	if o == nil || isNil(o.Logo) {
		var ret InlineResponse200249Logo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200249Items) GetLogoOk() (*InlineResponse200249Logo, bool) {
	if o == nil || isNil(o.Logo) {
    return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *InlineResponse200249Items) HasLogo() bool {
	if o != nil && !isNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given InlineResponse200249Logo and assigns it to the Logo field.
func (o *InlineResponse200249Items) SetLogo(v InlineResponse200249Logo) {
	o.Logo = &v
}

// GetIsBootstrap returns the IsBootstrap field value if set, zero value otherwise.
func (o *InlineResponse200249Items) GetIsBootstrap() bool {
	if o == nil || isNil(o.IsBootstrap) {
		var ret bool
		return ret
	}
	return *o.IsBootstrap
}

// GetIsBootstrapOk returns a tuple with the IsBootstrap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200249Items) GetIsBootstrapOk() (*bool, bool) {
	if o == nil || isNil(o.IsBootstrap) {
    return nil, false
	}
	return o.IsBootstrap, true
}

// HasIsBootstrap returns a boolean if a field has been set.
func (o *InlineResponse200249Items) HasIsBootstrap() bool {
	if o != nil && !isNil(o.IsBootstrap) {
		return true
	}

	return false
}

// SetIsBootstrap gets a reference to the given bool and assigns it to the IsBootstrap field.
func (o *InlineResponse200249Items) SetIsBootstrap(v bool) {
	o.IsBootstrap = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *InlineResponse200249Items) GetTerms() InlineResponse200249Terms {
	if o == nil || isNil(o.Terms) {
		var ret InlineResponse200249Terms
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200249Items) GetTermsOk() (*InlineResponse200249Terms, bool) {
	if o == nil || isNil(o.Terms) {
    return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *InlineResponse200249Items) HasTerms() bool {
	if o != nil && !isNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given InlineResponse200249Terms and assigns it to the Terms field.
func (o *InlineResponse200249Items) SetTerms(v InlineResponse200249Terms) {
	o.Terms = &v
}

func (o InlineResponse200249Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	if !isNil(o.IsBootstrap) {
		toSerialize["isBootstrap"] = o.IsBootstrap
	}
	if !isNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200249Items struct {
	value *InlineResponse200249Items
	isSet bool
}

func (v NullableInlineResponse200249Items) Get() *InlineResponse200249Items {
	return v.value
}

func (v *NullableInlineResponse200249Items) Set(val *InlineResponse200249Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200249Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200249Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200249Items(val *InlineResponse200249Items) *NullableInlineResponse200249Items {
	return &NullableInlineResponse200249Items{value: val, isSet: true}
}

func (v NullableInlineResponse200249Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200249Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


