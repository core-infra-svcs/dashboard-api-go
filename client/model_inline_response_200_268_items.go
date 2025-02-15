/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200268Items struct for InlineResponse200268Items
type InlineResponse200268Items struct {
	// Unique serial number for the device
	Serial *string `json:"serial,omitempty"`
	// Model of the device
	Model *string `json:"model,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
	// MAC address of the device
	Mac *string `json:"mac,omitempty"`
	// List of custom tags for the device
	Tags []string `json:"tags,omitempty"`
	// The total RAM size provisioned on the device, in kB
	Provisioned *int32 `json:"provisioned,omitempty"`
	Used *InlineResponse200268Used `json:"used,omitempty"`
	Free *InlineResponse200268Free `json:"free,omitempty"`
	Network *InlineResponse200268Network `json:"network,omitempty"`
	// Time interval snapshots of system memory utilization on the device with the most recent snapshot first
	Intervals []InlineResponse200268Intervals `json:"intervals,omitempty"`
}

// NewInlineResponse200268Items instantiates a new InlineResponse200268Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200268Items() *InlineResponse200268Items {
	this := InlineResponse200268Items{}
	return &this
}

// NewInlineResponse200268ItemsWithDefaults instantiates a new InlineResponse200268Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200268ItemsWithDefaults() *InlineResponse200268Items {
	this := InlineResponse200268Items{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetSerial() string {
	if o == nil || isNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetSerialOk() (*string, bool) {
	if o == nil || isNil(o.Serial) {
    return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasSerial() bool {
	if o != nil && !isNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *InlineResponse200268Items) SetSerial(v string) {
	o.Serial = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetModel() string {
	if o == nil || isNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetModelOk() (*string, bool) {
	if o == nil || isNil(o.Model) {
    return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasModel() bool {
	if o != nil && !isNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *InlineResponse200268Items) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200268Items) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *InlineResponse200268Items) SetMac(v string) {
	o.Mac = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetTags() []string {
	if o == nil || isNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetTagsOk() ([]string, bool) {
	if o == nil || isNil(o.Tags) {
    return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasTags() bool {
	if o != nil && !isNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *InlineResponse200268Items) SetTags(v []string) {
	o.Tags = v
}

// GetProvisioned returns the Provisioned field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetProvisioned() int32 {
	if o == nil || isNil(o.Provisioned) {
		var ret int32
		return ret
	}
	return *o.Provisioned
}

// GetProvisionedOk returns a tuple with the Provisioned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetProvisionedOk() (*int32, bool) {
	if o == nil || isNil(o.Provisioned) {
    return nil, false
	}
	return o.Provisioned, true
}

// HasProvisioned returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasProvisioned() bool {
	if o != nil && !isNil(o.Provisioned) {
		return true
	}

	return false
}

// SetProvisioned gets a reference to the given int32 and assigns it to the Provisioned field.
func (o *InlineResponse200268Items) SetProvisioned(v int32) {
	o.Provisioned = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetUsed() InlineResponse200268Used {
	if o == nil || isNil(o.Used) {
		var ret InlineResponse200268Used
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetUsedOk() (*InlineResponse200268Used, bool) {
	if o == nil || isNil(o.Used) {
    return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasUsed() bool {
	if o != nil && !isNil(o.Used) {
		return true
	}

	return false
}

// SetUsed gets a reference to the given InlineResponse200268Used and assigns it to the Used field.
func (o *InlineResponse200268Items) SetUsed(v InlineResponse200268Used) {
	o.Used = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetFree() InlineResponse200268Free {
	if o == nil || isNil(o.Free) {
		var ret InlineResponse200268Free
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetFreeOk() (*InlineResponse200268Free, bool) {
	if o == nil || isNil(o.Free) {
    return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasFree() bool {
	if o != nil && !isNil(o.Free) {
		return true
	}

	return false
}

// SetFree gets a reference to the given InlineResponse200268Free and assigns it to the Free field.
func (o *InlineResponse200268Items) SetFree(v InlineResponse200268Free) {
	o.Free = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetNetwork() InlineResponse200268Network {
	if o == nil || isNil(o.Network) {
		var ret InlineResponse200268Network
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetNetworkOk() (*InlineResponse200268Network, bool) {
	if o == nil || isNil(o.Network) {
    return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasNetwork() bool {
	if o != nil && !isNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given InlineResponse200268Network and assigns it to the Network field.
func (o *InlineResponse200268Items) SetNetwork(v InlineResponse200268Network) {
	o.Network = &v
}

// GetIntervals returns the Intervals field value if set, zero value otherwise.
func (o *InlineResponse200268Items) GetIntervals() []InlineResponse200268Intervals {
	if o == nil || isNil(o.Intervals) {
		var ret []InlineResponse200268Intervals
		return ret
	}
	return o.Intervals
}

// GetIntervalsOk returns a tuple with the Intervals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200268Items) GetIntervalsOk() ([]InlineResponse200268Intervals, bool) {
	if o == nil || isNil(o.Intervals) {
    return nil, false
	}
	return o.Intervals, true
}

// HasIntervals returns a boolean if a field has been set.
func (o *InlineResponse200268Items) HasIntervals() bool {
	if o != nil && !isNil(o.Intervals) {
		return true
	}

	return false
}

// SetIntervals gets a reference to the given []InlineResponse200268Intervals and assigns it to the Intervals field.
func (o *InlineResponse200268Items) SetIntervals(v []InlineResponse200268Intervals) {
	o.Intervals = v
}

func (o InlineResponse200268Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !isNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !isNil(o.Provisioned) {
		toSerialize["provisioned"] = o.Provisioned
	}
	if !isNil(o.Used) {
		toSerialize["used"] = o.Used
	}
	if !isNil(o.Free) {
		toSerialize["free"] = o.Free
	}
	if !isNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !isNil(o.Intervals) {
		toSerialize["intervals"] = o.Intervals
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200268Items struct {
	value *InlineResponse200268Items
	isSet bool
}

func (v NullableInlineResponse200268Items) Get() *InlineResponse200268Items {
	return v.value
}

func (v *NullableInlineResponse200268Items) Set(val *InlineResponse200268Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200268Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200268Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200268Items(val *InlineResponse200268Items) *NullableInlineResponse200268Items {
	return &NullableInlineResponse200268Items{value: val, isSet: true}
}

func (v NullableInlineResponse200268Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200268Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


