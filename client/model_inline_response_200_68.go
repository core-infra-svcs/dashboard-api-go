/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20068 struct for InlineResponse20068
type InlineResponse20068 struct {
	// The start time of the query range
	StartTs *time.Time `json:"startTs,omitempty"`
	// The end time of the query range
	EndTs *time.Time `json:"endTs,omitempty"`
	// Average data rate in kilobytes-per-second
	AverageKbps *int32 `json:"averageKbps,omitempty"`
	// Download rate in kilobytes-per-second
	DownloadKbps *int32 `json:"downloadKbps,omitempty"`
	// Upload rate in kilobytes-per-second
	UploadKbps *int32 `json:"uploadKbps,omitempty"`
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

// GetStartTs returns the StartTs field value if set, zero value otherwise.
func (o *InlineResponse20068) GetStartTs() time.Time {
	if o == nil || isNil(o.StartTs) {
		var ret time.Time
		return ret
	}
	return *o.StartTs
}

// GetStartTsOk returns a tuple with the StartTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetStartTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTs) {
    return nil, false
	}
	return o.StartTs, true
}

// HasStartTs returns a boolean if a field has been set.
func (o *InlineResponse20068) HasStartTs() bool {
	if o != nil && !isNil(o.StartTs) {
		return true
	}

	return false
}

// SetStartTs gets a reference to the given time.Time and assigns it to the StartTs field.
func (o *InlineResponse20068) SetStartTs(v time.Time) {
	o.StartTs = &v
}

// GetEndTs returns the EndTs field value if set, zero value otherwise.
func (o *InlineResponse20068) GetEndTs() time.Time {
	if o == nil || isNil(o.EndTs) {
		var ret time.Time
		return ret
	}
	return *o.EndTs
}

// GetEndTsOk returns a tuple with the EndTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetEndTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.EndTs) {
    return nil, false
	}
	return o.EndTs, true
}

// HasEndTs returns a boolean if a field has been set.
func (o *InlineResponse20068) HasEndTs() bool {
	if o != nil && !isNil(o.EndTs) {
		return true
	}

	return false
}

// SetEndTs gets a reference to the given time.Time and assigns it to the EndTs field.
func (o *InlineResponse20068) SetEndTs(v time.Time) {
	o.EndTs = &v
}

// GetAverageKbps returns the AverageKbps field value if set, zero value otherwise.
func (o *InlineResponse20068) GetAverageKbps() int32 {
	if o == nil || isNil(o.AverageKbps) {
		var ret int32
		return ret
	}
	return *o.AverageKbps
}

// GetAverageKbpsOk returns a tuple with the AverageKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetAverageKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.AverageKbps) {
    return nil, false
	}
	return o.AverageKbps, true
}

// HasAverageKbps returns a boolean if a field has been set.
func (o *InlineResponse20068) HasAverageKbps() bool {
	if o != nil && !isNil(o.AverageKbps) {
		return true
	}

	return false
}

// SetAverageKbps gets a reference to the given int32 and assigns it to the AverageKbps field.
func (o *InlineResponse20068) SetAverageKbps(v int32) {
	o.AverageKbps = &v
}

// GetDownloadKbps returns the DownloadKbps field value if set, zero value otherwise.
func (o *InlineResponse20068) GetDownloadKbps() int32 {
	if o == nil || isNil(o.DownloadKbps) {
		var ret int32
		return ret
	}
	return *o.DownloadKbps
}

// GetDownloadKbpsOk returns a tuple with the DownloadKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetDownloadKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.DownloadKbps) {
    return nil, false
	}
	return o.DownloadKbps, true
}

// HasDownloadKbps returns a boolean if a field has been set.
func (o *InlineResponse20068) HasDownloadKbps() bool {
	if o != nil && !isNil(o.DownloadKbps) {
		return true
	}

	return false
}

// SetDownloadKbps gets a reference to the given int32 and assigns it to the DownloadKbps field.
func (o *InlineResponse20068) SetDownloadKbps(v int32) {
	o.DownloadKbps = &v
}

// GetUploadKbps returns the UploadKbps field value if set, zero value otherwise.
func (o *InlineResponse20068) GetUploadKbps() int32 {
	if o == nil || isNil(o.UploadKbps) {
		var ret int32
		return ret
	}
	return *o.UploadKbps
}

// GetUploadKbpsOk returns a tuple with the UploadKbps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20068) GetUploadKbpsOk() (*int32, bool) {
	if o == nil || isNil(o.UploadKbps) {
    return nil, false
	}
	return o.UploadKbps, true
}

// HasUploadKbps returns a boolean if a field has been set.
func (o *InlineResponse20068) HasUploadKbps() bool {
	if o != nil && !isNil(o.UploadKbps) {
		return true
	}

	return false
}

// SetUploadKbps gets a reference to the given int32 and assigns it to the UploadKbps field.
func (o *InlineResponse20068) SetUploadKbps(v int32) {
	o.UploadKbps = &v
}

func (o InlineResponse20068) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTs) {
		toSerialize["startTs"] = o.StartTs
	}
	if !isNil(o.EndTs) {
		toSerialize["endTs"] = o.EndTs
	}
	if !isNil(o.AverageKbps) {
		toSerialize["averageKbps"] = o.AverageKbps
	}
	if !isNil(o.DownloadKbps) {
		toSerialize["downloadKbps"] = o.DownloadKbps
	}
	if !isNil(o.UploadKbps) {
		toSerialize["uploadKbps"] = o.UploadKbps
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


