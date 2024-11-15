/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200150 struct for InlineResponse200150
type InlineResponse200150 struct {
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

// NewInlineResponse200150 instantiates a new InlineResponse200150 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200150() *InlineResponse200150 {
	this := InlineResponse200150{}
	return &this
}

// NewInlineResponse200150WithDefaults instantiates a new InlineResponse200150 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200150WithDefaults() *InlineResponse200150 {
	this := InlineResponse200150{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200150) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200150) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200150) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200150) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200150) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200150) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse200150) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse200150) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse200150) SetUrl(v string) {
	o.Url = &v
}

// GetSupportsInspection returns the SupportsInspection field value if set, zero value otherwise.
func (o *InlineResponse200150) GetSupportsInspection() bool {
	if o == nil || isNil(o.SupportsInspection) {
		var ret bool
		return ret
	}
	return *o.SupportsInspection
}

// GetSupportsInspectionOk returns a tuple with the SupportsInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetSupportsInspectionOk() (*bool, bool) {
	if o == nil || isNil(o.SupportsInspection) {
    return nil, false
	}
	return o.SupportsInspection, true
}

// HasSupportsInspection returns a boolean if a field has been set.
func (o *InlineResponse200150) HasSupportsInspection() bool {
	if o != nil && !isNil(o.SupportsInspection) {
		return true
	}

	return false
}

// SetSupportsInspection gets a reference to the given bool and assigns it to the SupportsInspection field.
func (o *InlineResponse200150) SetSupportsInspection(v bool) {
	o.SupportsInspection = &v
}

// GetHasTrustedPort returns the HasTrustedPort field value if set, zero value otherwise.
func (o *InlineResponse200150) GetHasTrustedPort() bool {
	if o == nil || isNil(o.HasTrustedPort) {
		var ret bool
		return ret
	}
	return *o.HasTrustedPort
}

// GetHasTrustedPortOk returns a tuple with the HasTrustedPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200150) GetHasTrustedPortOk() (*bool, bool) {
	if o == nil || isNil(o.HasTrustedPort) {
    return nil, false
	}
	return o.HasTrustedPort, true
}

// HasHasTrustedPort returns a boolean if a field has been set.
func (o *InlineResponse200150) HasHasTrustedPort() bool {
	if o != nil && !isNil(o.HasTrustedPort) {
		return true
	}

	return false
}

// SetHasTrustedPort gets a reference to the given bool and assigns it to the HasTrustedPort field.
func (o *InlineResponse200150) SetHasTrustedPort(v bool) {
	o.HasTrustedPort = &v
}

func (o InlineResponse200150) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.SupportsInspection) {
		toSerialize["supportsInspection"] = o.SupportsInspection
	}
	if !isNil(o.HasTrustedPort) {
		toSerialize["hasTrustedPort"] = o.HasTrustedPort
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200150 struct {
	value *InlineResponse200150
	isSet bool
}

func (v NullableInlineResponse200150) Get() *InlineResponse200150 {
	return v.value
}

func (v *NullableInlineResponse200150) Set(val *InlineResponse200150) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200150) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200150) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200150(val *InlineResponse200150) *NullableInlineResponse200150 {
	return &NullableInlineResponse200150{value: val, isSet: true}
}

func (v NullableInlineResponse200150) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200150) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


