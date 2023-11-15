/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink struct for OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink
type OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink struct {
	// Uplink serial
	Serial *string `json:"serial,omitempty"`
	// Uplink name
	Interface *string `json:"interface,omitempty"`
	// Bytes sent
	Sent *int32 `json:"sent,omitempty"`
	// Bytes received
	Received *int32 `json:"received,omitempty"`
}

// NewOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink instantiates a new OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink() *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	this := OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink{}
	return &this
}

// NewOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplinkWithDefaults instantiates a new OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplinkWithDefaults() *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	this := OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) SetSerial(v string) {
	o.Serial = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetInterface() string {
	if o == nil || isNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetInterfaceOk() (*string, bool) {
	if o == nil || isNil(o.Interface) {
    return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) HasInterface() bool {
	if o != nil && !isNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) SetInterface(v string) {
	o.Interface = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetSent() int32 {
	if o == nil || isNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetSentOk() (*int32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) SetSent(v int32) {
	o.Sent = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetReceived() int32 {
	if o == nil || isNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) GetReceivedOk() (*int32, bool) {
	if o == nil || isNil(o.Received) {
    return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) HasReceived() bool {
	if o != nil && !isNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) SetReceived(v int32) {
	o.Received = &v
}

func (o OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink struct {
	value *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink
	isSet bool
}

func (v NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) Get() *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	return v.value
}

func (v *NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) Set(val *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink(val *OrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) *NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink {
	return &NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink{value: val, isSet: true}
}

func (v NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsOrganizationIdApplianceUplinksUsageByNetworkByUplink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

