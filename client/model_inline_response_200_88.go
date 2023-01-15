/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 January, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.29.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20088 struct for InlineResponse20088
type InlineResponse20088 struct {
	// The device MAC address.
	Mac *string `json:"mac,omitempty"`
	// The device name.
	Name *string `json:"name,omitempty"`
	Network *OrganizationsOrganizationIdDevicesAvailabilitiesNetwork `json:"network,omitempty"`
	// Device product type.
	ProductType *string `json:"productType,omitempty"`
	// The device serial number.
	Serial *string `json:"serial,omitempty"`
	// List of custom tags for the device.
	Tags []string `json:"tags,omitempty"`
	// List of device uplink addresses information.
	Uplinks []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks `json:"uplinks,omitempty"`
}

// NewInlineResponse20088 instantiates a new InlineResponse20088 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20088() *InlineResponse20088 {
	this := InlineResponse20088{}
	return &this
}

// NewInlineResponse20088WithDefaults instantiates a new InlineResponse20088 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20088WithDefaults() *InlineResponse20088 {
	this := InlineResponse20088{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse20088) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse20088) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse20088) SetMac(v string) {
	o.Mac = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20088) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20088) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20088) SetName(v string) {
	o.Name = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse20088) GetNetwork() OrganizationsOrganizationIdDevicesAvailabilitiesNetwork {
	if o == nil || isNil(o.Network) {
		var ret OrganizationsOrganizationIdDevicesAvailabilitiesNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetNetworkOk() (*OrganizationsOrganizationIdDevicesAvailabilitiesNetwork, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse20088) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given OrganizationsOrganizationIdDevicesAvailabilitiesNetwork and assigns it to the Network field.
func (o *InlineResponse20088) SetNetwork(v OrganizationsOrganizationIdDevicesAvailabilitiesNetwork) {
	o.Network = &v
}

// GetProductType returns the ProductType field value if set, zero value otherwise.
func (o *InlineResponse20088) GetProductType() string {
	if o == nil || isNil(o.ProductType) {
		var ret string
		return ret
	}
	return *o.ProductType
}

// GetProductTypeOk returns a tuple with the ProductType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetProductTypeOk() (*string, bool) {
	if o == nil || isNil(o.ProductType) {
    return nil, false
	}
	return o.ProductType, true
}

// HasProductType returns a boolean if a field has been set.
func (o *InlineResponse20088) HasProductType() bool {
	if o != nil && !isNil(o.ProductType) {
		return true
	}

	return false
}

// SetProductType gets a reference to the given string and assigns it to the ProductType field.
func (o *InlineResponse20088) SetProductType(v string) {
	o.ProductType = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse20088) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse20088) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse20088) SetSerial(v string) {
	o.Serial = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse20088) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse20088) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse20088) SetTags(v []string) {
	o.Tags = v
}

// GetUplinks returns the Uplinks field value if set, zero value otherwise.
func (o *InlineResponse20088) GetUplinks() []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks {
	if o == nil || isNil(o.Uplinks) {
		var ret []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks
		return ret
	}
	return o.Uplinks
}

// GetUplinksOk returns a tuple with the Uplinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetUplinksOk() ([]OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks, bool) {
	if o == nil || isNil(o.Uplinks) {
    return nil, false
	}
	return o.Uplinks, true
}

// HasUplinks returns a boolean if a field has been set.
func (o *InlineResponse20088) HasUplinks() bool {
	if o != nil && !isNil(o.Uplinks) {
		return true
	}

	return false
}

// SetUplinks gets a reference to the given []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks and assigns it to the Uplinks field.
func (o *InlineResponse20088) SetUplinks(v []OrganizationsOrganizationIdDevicesUplinksAddressesByDeviceUplinks) {
	o.Uplinks = v
}

func (o InlineResponse20088) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.ProductType) {
		toSerialize["productType"] = o.ProductType
	}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Uplinks) {
		toSerialize["uplinks"] = o.Uplinks
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20088 struct {
	value *InlineResponse20088
	isSet bool
}

func (v NullableInlineResponse20088) Get() *InlineResponse20088 {
	return v.value
}

func (v *NullableInlineResponse20088) Set(val *InlineResponse20088) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20088) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20088) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20088(val *InlineResponse20088) *NullableInlineResponse20088 {
	return &NullableInlineResponse20088{value: val, isSet: true}
}

func (v NullableInlineResponse20088) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20088) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


