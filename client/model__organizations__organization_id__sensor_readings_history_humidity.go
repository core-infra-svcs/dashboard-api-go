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

// OrganizationsOrganizationIdSensorReadingsHistoryHumidity Reading for the 'humidity' metric. This will only be present if the 'metric' property equals 'humidity'.
type OrganizationsOrganizationIdSensorReadingsHistoryHumidity struct {
	// Humidity reading in %RH.
	RelativePercentage *int32 `json:"relativePercentage,omitempty"`
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryHumidity instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryHumidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdSensorReadingsHistoryHumidity() *OrganizationsOrganizationIdSensorReadingsHistoryHumidity {
	this := OrganizationsOrganizationIdSensorReadingsHistoryHumidity{}
	return &this
}

// NewOrganizationsOrganizationIdSensorReadingsHistoryHumidityWithDefaults instantiates a new OrganizationsOrganizationIdSensorReadingsHistoryHumidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdSensorReadingsHistoryHumidityWithDefaults() *OrganizationsOrganizationIdSensorReadingsHistoryHumidity {
	this := OrganizationsOrganizationIdSensorReadingsHistoryHumidity{}
	return &this
}

// GetRelativePercentage returns the RelativePercentage field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryHumidity) GetRelativePercentage() int32 {
	if o == nil || o.RelativePercentage == nil {
		var ret int32
		return ret
	}
	return *o.RelativePercentage
}

// GetRelativePercentageOk returns a tuple with the RelativePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryHumidity) GetRelativePercentageOk() (*int32, bool) {
	if o == nil || o.RelativePercentage == nil {
		return nil, false
	}
	return o.RelativePercentage, true
}

// HasRelativePercentage returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryHumidity) HasRelativePercentage() bool {
	if o != nil && o.RelativePercentage != nil {
		return true
	}

	return false
}

// SetRelativePercentage gets a reference to the given int32 and assigns it to the RelativePercentage field.
func (o *OrganizationsOrganizationIdSensorReadingsHistoryHumidity) SetRelativePercentage(v int32) {
	o.RelativePercentage = &v
}

func (o OrganizationsOrganizationIdSensorReadingsHistoryHumidity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RelativePercentage != nil {
		toSerialize["relativePercentage"] = o.RelativePercentage
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity struct {
	value *OrganizationsOrganizationIdSensorReadingsHistoryHumidity
	isSet bool
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity) Get() *OrganizationsOrganizationIdSensorReadingsHistoryHumidity {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity) Set(val *OrganizationsOrganizationIdSensorReadingsHistoryHumidity) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity(val *OrganizationsOrganizationIdSensorReadingsHistoryHumidity) *NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity {
	return &NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdSensorReadingsHistoryHumidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


