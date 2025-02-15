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

// InlineResponse200185 struct for InlineResponse200185
type InlineResponse200185 struct {
	// The number of failed association attempts
	Assoc *int32 `json:"assoc,omitempty"`
	// The number of failed authentication attempts
	Auth *int32 `json:"auth,omitempty"`
	// The number of failed DHCP attempts
	Dhcp *int32 `json:"dhcp,omitempty"`
	// The number of failed DNS attempts
	Dns *int32 `json:"dns,omitempty"`
	// The number of successful connection attempts
	Success *int32 `json:"success,omitempty"`
}

// NewInlineResponse200185 instantiates a new InlineResponse200185 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200185() *InlineResponse200185 {
	this := InlineResponse200185{}
	return &this
}

// NewInlineResponse200185WithDefaults instantiates a new InlineResponse200185 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200185WithDefaults() *InlineResponse200185 {
	this := InlineResponse200185{}
	return &this
}

// GetAssoc returns the Assoc field value if set, zero value otherwise.
func (o *InlineResponse200185) GetAssoc() int32 {
	if o == nil || isNil(o.Assoc) {
		var ret int32
		return ret
	}
	return *o.Assoc
}

// GetAssocOk returns a tuple with the Assoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetAssocOk() (*int32, bool) {
	if o == nil || isNil(o.Assoc) {
    return nil, false
	}
	return o.Assoc, true
}

// HasAssoc returns a boolean if a field has been set.
func (o *InlineResponse200185) HasAssoc() bool {
	if o != nil && !isNil(o.Assoc) {
		return true
	}

	return false
}

// SetAssoc gets a reference to the given int32 and assigns it to the Assoc field.
func (o *InlineResponse200185) SetAssoc(v int32) {
	o.Assoc = &v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *InlineResponse200185) GetAuth() int32 {
	if o == nil || isNil(o.Auth) {
		var ret int32
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetAuthOk() (*int32, bool) {
	if o == nil || isNil(o.Auth) {
    return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *InlineResponse200185) HasAuth() bool {
	if o != nil && !isNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given int32 and assigns it to the Auth field.
func (o *InlineResponse200185) SetAuth(v int32) {
	o.Auth = &v
}

// GetDhcp returns the Dhcp field value if set, zero value otherwise.
func (o *InlineResponse200185) GetDhcp() int32 {
	if o == nil || isNil(o.Dhcp) {
		var ret int32
		return ret
	}
	return *o.Dhcp
}

// GetDhcpOk returns a tuple with the Dhcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetDhcpOk() (*int32, bool) {
	if o == nil || isNil(o.Dhcp) {
    return nil, false
	}
	return o.Dhcp, true
}

// HasDhcp returns a boolean if a field has been set.
func (o *InlineResponse200185) HasDhcp() bool {
	if o != nil && !isNil(o.Dhcp) {
		return true
	}

	return false
}

// SetDhcp gets a reference to the given int32 and assigns it to the Dhcp field.
func (o *InlineResponse200185) SetDhcp(v int32) {
	o.Dhcp = &v
}

// GetDns returns the Dns field value if set, zero value otherwise.
func (o *InlineResponse200185) GetDns() int32 {
	if o == nil || isNil(o.Dns) {
		var ret int32
		return ret
	}
	return *o.Dns
}

// GetDnsOk returns a tuple with the Dns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetDnsOk() (*int32, bool) {
	if o == nil || isNil(o.Dns) {
    return nil, false
	}
	return o.Dns, true
}

// HasDns returns a boolean if a field has been set.
func (o *InlineResponse200185) HasDns() bool {
	if o != nil && !isNil(o.Dns) {
		return true
	}

	return false
}

// SetDns gets a reference to the given int32 and assigns it to the Dns field.
func (o *InlineResponse200185) SetDns(v int32) {
	o.Dns = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *InlineResponse200185) GetSuccess() int32 {
	if o == nil || isNil(o.Success) {
		var ret int32
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200185) GetSuccessOk() (*int32, bool) {
	if o == nil || isNil(o.Success) {
    return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *InlineResponse200185) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given int32 and assigns it to the Success field.
func (o *InlineResponse200185) SetSuccess(v int32) {
	o.Success = &v
}

func (o InlineResponse200185) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Assoc) {
		toSerialize["assoc"] = o.Assoc
	}
	if !isNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !isNil(o.Dhcp) {
		toSerialize["dhcp"] = o.Dhcp
	}
	if !isNil(o.Dns) {
		toSerialize["dns"] = o.Dns
	}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200185 struct {
	value *InlineResponse200185
	isSet bool
}

func (v NullableInlineResponse200185) Get() *InlineResponse200185 {
	return v.value
}

func (v *NullableInlineResponse200185) Set(val *InlineResponse200185) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200185) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200185) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200185(val *InlineResponse200185) *NullableInlineResponse200185 {
	return &NullableInlineResponse200185{value: val, isSet: true}
}

func (v NullableInlineResponse200185) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200185) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


