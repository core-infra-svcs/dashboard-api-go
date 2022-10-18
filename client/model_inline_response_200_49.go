/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20049 struct for InlineResponse20049
type InlineResponse20049 struct {
	// Switch serial.
	Serial *string `json:"serial,omitempty"`
	// Switch name.
	Name *string `json:"name,omitempty"`
	// Url link to switch.
	Url *string `json:"url,omitempty"`
	// Whether this switch supports Dynamic ARP Inspection.
	SupportsInspection *bool `json:"supportsInspection,omitempty"`
	// Whether this switch has a trusted DAI port. Always false if supportsInspection is false.
	HasTrustedPort *bool `json:"hasTrustedPort,omitempty"`
}

// NewInlineResponse20049 instantiates a new InlineResponse20049 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20049() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// NewInlineResponse20049WithDefaults instantiates a new InlineResponse20049 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20049WithDefaults() *InlineResponse20049 {
	this := InlineResponse20049{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20049) GetSerial() string {
	if o == nil || o.Serial == nil {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetSerialOk() (*string, bool) {
	if o == nil || o.Serial == nil {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20049) HasSerial() bool {
	if o != nil && o.Serial != nil {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20049) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20049) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20049) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20049) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20049) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20049) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20049) SetUrl(v string) {
	o.Url = &v
}

// GetSupportsInspection returns the SupportsInspection field value if set, zero value otherwise.
func (o *InlineResponse20049) GetSupportsInspection() bool {
	if o == nil || o.SupportsInspection == nil {
		var ret bool
		return ret
	}
	return *o.SupportsInspection
}

// GetSupportsInspectionOk returns a tuple with the SupportsInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetSupportsInspectionOk() (*bool, bool) {
	if o == nil || o.SupportsInspection == nil {
		return nil, false
	}
	return o.SupportsInspection, true
}

// HasSupportsInspection returns a boolean if a field has been set.
func (o *InlineResponse20049) HasSupportsInspection() bool {
	if o != nil && o.SupportsInspection != nil {
		return true
	}

	return false
}

// SetSupportsInspection gets a reference to the given bool and assigns it to the SupportsInspection field.
func (o *InlineResponse20049) SetSupportsInspection(v bool) {
	o.SupportsInspection = &v
}

// GetHasTrustedPort returns the HasTrustedPort field value if set, zero value otherwise.
func (o *InlineResponse20049) GetHasTrustedPort() bool {
	if o == nil || o.HasTrustedPort == nil {
		var ret bool
		return ret
	}
	return *o.HasTrustedPort
}

// GetHasTrustedPortOk returns a tuple with the HasTrustedPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20049) GetHasTrustedPortOk() (*bool, bool) {
	if o == nil || o.HasTrustedPort == nil {
		return nil, false
	}
	return o.HasTrustedPort, true
}

// HasHasTrustedPort returns a boolean if a field has been set.
func (o *InlineResponse20049) HasHasTrustedPort() bool {
	if o != nil && o.HasTrustedPort != nil {
		return true
	}

	return false
}

// SetHasTrustedPort gets a reference to the given bool and assigns it to the HasTrustedPort field.
func (o *InlineResponse20049) SetHasTrustedPort(v bool) {
	o.HasTrustedPort = &v
}

func (o InlineResponse20049) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Serial != nil {
		toSerialize["serial"] = o.Serial
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.SupportsInspection != nil {
		toSerialize["supportsInspection"] = o.SupportsInspection
	}
	if o.HasTrustedPort != nil {
		toSerialize["hasTrustedPort"] = o.HasTrustedPort
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20049 struct {
	value *InlineResponse20049
	isSet bool
}

func (v NullableInlineResponse20049) Get() *InlineResponse20049 {
	return v.value
}

func (v *NullableInlineResponse20049) Set(val *InlineResponse20049) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20049) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20049) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20049(val *InlineResponse20049) *NullableInlineResponse20049 {
	return &NullableInlineResponse20049{value: val, isSet: true}
}

func (v NullableInlineResponse20049) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20049) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


