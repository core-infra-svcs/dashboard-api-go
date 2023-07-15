/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20075 struct for InlineResponse20075
type InlineResponse20075 struct {
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

// NewInlineResponse20075 instantiates a new InlineResponse20075 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20075() *InlineResponse20075 {
	this := InlineResponse20075{}
	return &this
}

// NewInlineResponse20075WithDefaults instantiates a new InlineResponse20075 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20075WithDefaults() *InlineResponse20075 {
	this := InlineResponse20075{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20075) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20075) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20075) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20075) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20075) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20075) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *InlineResponse20075) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *InlineResponse20075) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *InlineResponse20075) SetUrl(v string) {
	o.Url = &v
}

// GetSupportsInspection returns the SupportsInspection field value if set, zero value otherwise.
func (o *InlineResponse20075) GetSupportsInspection() bool {
	if o == nil || isNil(o.SupportsInspection) {
		var ret bool
		return ret
	}
	return *o.SupportsInspection
}

// GetSupportsInspectionOk returns a tuple with the SupportsInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetSupportsInspectionOk() (*bool, bool) {
	if o == nil || isNil(o.SupportsInspection) {
    return nil, false
	}
	return o.SupportsInspection, true
}

// HasSupportsInspection returns a boolean if a field has been set.
func (o *InlineResponse20075) HasSupportsInspection() bool {
	if o != nil && !isNil(o.SupportsInspection) {
		return true
	}

	return false
}

// SetSupportsInspection gets a reference to the given bool and assigns it to the SupportsInspection field.
func (o *InlineResponse20075) SetSupportsInspection(v bool) {
	o.SupportsInspection = &v
}

// GetHasTrustedPort returns the HasTrustedPort field value if set, zero value otherwise.
func (o *InlineResponse20075) GetHasTrustedPort() bool {
	if o == nil || isNil(o.HasTrustedPort) {
		var ret bool
		return ret
	}
	return *o.HasTrustedPort
}

// GetHasTrustedPortOk returns a tuple with the HasTrustedPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20075) GetHasTrustedPortOk() (*bool, bool) {
	if o == nil || isNil(o.HasTrustedPort) {
    return nil, false
	}
	return o.HasTrustedPort, true
}

// HasHasTrustedPort returns a boolean if a field has been set.
func (o *InlineResponse20075) HasHasTrustedPort() bool {
	if o != nil && !isNil(o.HasTrustedPort) {
		return true
	}

	return false
}

// SetHasTrustedPort gets a reference to the given bool and assigns it to the HasTrustedPort field.
func (o *InlineResponse20075) SetHasTrustedPort(v bool) {
	o.HasTrustedPort = &v
}

func (o InlineResponse20075) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse20075 struct {
	value *InlineResponse20075
	isSet bool
}

func (v NullableInlineResponse20075) Get() *InlineResponse20075 {
	return v.value
}

func (v *NullableInlineResponse20075) Set(val *InlineResponse20075) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20075) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20075) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20075(val *InlineResponse20075) *NullableInlineResponse20075 {
	return &NullableInlineResponse20075{value: val, isSet: true}
}

func (v NullableInlineResponse20075) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20075) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


