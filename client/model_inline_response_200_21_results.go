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

// InlineResponse20021Results struct for InlineResponse20021Results
type InlineResponse20021Results struct {
	// The port for which the test was performed.
	Port *string `json:"port,omitempty"`
	// The current status of the port. If the cable test is still being performed on the port, \"in-progress\" is used. If an error occurred during the cable test, \"error\" is used and the error property will be populated.
	Status *string `json:"status,omitempty"`
	// Speed in Mbps.  A speed of 0 indicates the port is down or the port speed is automatic.
	SpeedMbps *int32 `json:"speedMbps,omitempty"`
	// If an error occurred during the cable test, the error message will be populated here.
	Error *string `json:"error,omitempty"`
	// Results for each twisted pair within the cable.
	Pairs []InlineResponse20021Pairs `json:"pairs,omitempty"`
}

// NewInlineResponse20021Results instantiates a new InlineResponse20021Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20021Results() *InlineResponse20021Results {
	this := InlineResponse20021Results{}
	return &this
}

// NewInlineResponse20021ResultsWithDefaults instantiates a new InlineResponse20021Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20021ResultsWithDefaults() *InlineResponse20021Results {
	this := InlineResponse20021Results{}
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse20021Results) GetPort() string {
	if o == nil || isNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Results) GetPortOk() (*string, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse20021Results) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *InlineResponse20021Results) SetPort(v string) {
	o.Port = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InlineResponse20021Results) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Results) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InlineResponse20021Results) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InlineResponse20021Results) SetStatus(v string) {
	o.Status = &v
}

// GetSpeedMbps returns the SpeedMbps field value if set, zero value otherwise.
func (o *InlineResponse20021Results) GetSpeedMbps() int32 {
	if o == nil || isNil(o.SpeedMbps) {
		var ret int32
		return ret
	}
	return *o.SpeedMbps
}

// GetSpeedMbpsOk returns a tuple with the SpeedMbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Results) GetSpeedMbpsOk() (*int32, bool) {
	if o == nil || isNil(o.SpeedMbps) {
    return nil, false
	}
	return o.SpeedMbps, true
}

// HasSpeedMbps returns a boolean if a field has been set.
func (o *InlineResponse20021Results) HasSpeedMbps() bool {
	if o != nil && !isNil(o.SpeedMbps) {
		return true
	}

	return false
}

// SetSpeedMbps gets a reference to the given int32 and assigns it to the SpeedMbps field.
func (o *InlineResponse20021Results) SetSpeedMbps(v int32) {
	o.SpeedMbps = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InlineResponse20021Results) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Results) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InlineResponse20021Results) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *InlineResponse20021Results) SetError(v string) {
	o.Error = &v
}

// GetPairs returns the Pairs field value if set, zero value otherwise.
func (o *InlineResponse20021Results) GetPairs() []InlineResponse20021Pairs {
	if o == nil || isNil(o.Pairs) {
		var ret []InlineResponse20021Pairs
		return ret
	}
	return o.Pairs
}

// GetPairsOk returns a tuple with the Pairs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20021Results) GetPairsOk() ([]InlineResponse20021Pairs, bool) {
	if o == nil || isNil(o.Pairs) {
    return nil, false
	}
	return o.Pairs, true
}

// HasPairs returns a boolean if a field has been set.
func (o *InlineResponse20021Results) HasPairs() bool {
	if o != nil && !isNil(o.Pairs) {
		return true
	}

	return false
}

// SetPairs gets a reference to the given []InlineResponse20021Pairs and assigns it to the Pairs field.
func (o *InlineResponse20021Results) SetPairs(v []InlineResponse20021Pairs) {
	o.Pairs = v
}

func (o InlineResponse20021Results) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.SpeedMbps) {
		toSerialize["speedMbps"] = o.SpeedMbps
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !isNil(o.Pairs) {
		toSerialize["pairs"] = o.Pairs
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20021Results struct {
	value *InlineResponse20021Results
	isSet bool
}

func (v NullableInlineResponse20021Results) Get() *InlineResponse20021Results {
	return v.value
}

func (v *NullableInlineResponse20021Results) Set(val *InlineResponse20021Results) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20021Results) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20021Results) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20021Results(val *InlineResponse20021Results) *NullableInlineResponse20021Results {
	return &NullableInlineResponse20021Results{value: val, isSet: true}
}

func (v NullableInlineResponse20021Results) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20021Results) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


