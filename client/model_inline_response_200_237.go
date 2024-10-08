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

// InlineResponse200237 struct for InlineResponse200237
type InlineResponse200237 struct {
	// The ID of the client
	ClientId *string `json:"clientId,omitempty"`
	// The MAC address of the client
	Mac *string `json:"mac,omitempty"`
	// Manufacturer of the client
	Manufacturer *string `json:"manufacturer,omitempty"`
	// The clients that appear on any networks within an organization
	Records []InlineResponse200237Records `json:"records,omitempty"`
}

// NewInlineResponse200237 instantiates a new InlineResponse200237 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200237() *InlineResponse200237 {
	this := InlineResponse200237{}
	return &this
}

// NewInlineResponse200237WithDefaults instantiates a new InlineResponse200237 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200237WithDefaults() *InlineResponse200237 {
	this := InlineResponse200237{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InlineResponse200237) GetClientId() string {
	if o == nil || isNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237) GetClientIdOk() (*string, bool) {
	if o == nil || isNil(o.ClientId) {
    return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InlineResponse200237) HasClientId() bool {
	if o != nil && !isNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InlineResponse200237) SetClientId(v string) {
	o.ClientId = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200237) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200237) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200237) SetMac(v string) {
	o.Mac = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *InlineResponse200237) GetManufacturer() string {
	if o == nil || isNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237) GetManufacturerOk() (*string, bool) {
	if o == nil || isNil(o.Manufacturer) {
    return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *InlineResponse200237) HasManufacturer() bool {
	if o != nil && !isNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *InlineResponse200237) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *InlineResponse200237) GetRecords() []InlineResponse200237Records {
	if o == nil || isNil(o.Records) {
		var ret []InlineResponse200237Records
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200237) GetRecordsOk() ([]InlineResponse200237Records, bool) {
	if o == nil || isNil(o.Records) {
    return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *InlineResponse200237) HasRecords() bool {
	if o != nil && !isNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []InlineResponse200237Records and assigns it to the Records field.
func (o *InlineResponse200237) SetRecords(v []InlineResponse200237Records) {
	o.Records = v
}

func (o InlineResponse200237) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !isNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200237 struct {
	value *InlineResponse200237
	isSet bool
}

func (v NullableInlineResponse200237) Get() *InlineResponse200237 {
	return v.value
}

func (v *NullableInlineResponse200237) Set(val *InlineResponse200237) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200237) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200237) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200237(val *InlineResponse200237) *NullableInlineResponse200237 {
	return &NullableInlineResponse200237{value: val, isSet: true}
}

func (v NullableInlineResponse200237) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200237) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


