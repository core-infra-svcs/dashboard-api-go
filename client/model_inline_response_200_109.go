/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 March, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.56.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200109 struct for InlineResponse200109
type InlineResponse200109 struct {
	// ID of the MQTT Broker.
	Id *string `json:"id,omitempty"`
	// Name of the MQTT Broker.
	Name *string `json:"name,omitempty"`
	// Host name/IP address where the MQTT broker runs.
	Host *string `json:"host,omitempty"`
	// Host port though which the MQTT broker can be reached.
	Port *int32 `json:"port,omitempty"`
	Security *NetworksNetworkIdMqttBrokersSecurity `json:"security,omitempty"`
	Authentication *NetworksNetworkIdMqttBrokersAuthentication `json:"authentication,omitempty"`
}

// NewInlineResponse200109 instantiates a new InlineResponse200109 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200109() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// NewInlineResponse200109WithDefaults instantiates a new InlineResponse200109 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200109WithDefaults() *InlineResponse200109 {
	this := InlineResponse200109{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse200109) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse200109) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse200109) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse200109) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse200109) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse200109) SetName(v string) {
	o.Name = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *InlineResponse200109) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *InlineResponse200109) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *InlineResponse200109) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse200109) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse200109) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InlineResponse200109) SetPort(v int32) {
	o.Port = &v
}

// GetSecurity returns the Security field value if set, zero value otherwise.
func (o *InlineResponse200109) GetSecurity() NetworksNetworkIdMqttBrokersSecurity {
	if o == nil || isNil(o.Security) {
		var ret NetworksNetworkIdMqttBrokersSecurity
		return ret
	}
	return *o.Security
}

// GetSecurityOk returns a tuple with the Security field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetSecurityOk() (*NetworksNetworkIdMqttBrokersSecurity, bool) {
	if o == nil || isNil(o.Security) {
    return nil, false
	}
	return o.Security, true
}

// HasSecurity returns a boolean if a field has been set.
func (o *InlineResponse200109) HasSecurity() bool {
	if o != nil && !isNil(o.Security) {
		return true
	}

	return false
}

// SetSecurity gets a reference to the given NetworksNetworkIdMqttBrokersSecurity and assigns it to the Security field.
func (o *InlineResponse200109) SetSecurity(v NetworksNetworkIdMqttBrokersSecurity) {
	o.Security = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *InlineResponse200109) GetAuthentication() NetworksNetworkIdMqttBrokersAuthentication {
	if o == nil || isNil(o.Authentication) {
		var ret NetworksNetworkIdMqttBrokersAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200109) GetAuthenticationOk() (*NetworksNetworkIdMqttBrokersAuthentication, bool) {
	if o == nil || isNil(o.Authentication) {
    return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *InlineResponse200109) HasAuthentication() bool {
	if o != nil && !isNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NetworksNetworkIdMqttBrokersAuthentication and assigns it to the Authentication field.
func (o *InlineResponse200109) SetAuthentication(v NetworksNetworkIdMqttBrokersAuthentication) {
	o.Authentication = &v
}

func (o InlineResponse200109) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Security) {
		toSerialize["security"] = o.Security
	}
	if !isNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200109 struct {
	value *InlineResponse200109
	isSet bool
}

func (v NullableInlineResponse200109) Get() *InlineResponse200109 {
	return v.value
}

func (v *NullableInlineResponse200109) Set(val *InlineResponse200109) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200109) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200109) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200109(val *InlineResponse200109) *NullableInlineResponse200109 {
	return &NullableInlineResponse200109{value: val, isSet: true}
}

func (v NullableInlineResponse200109) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200109) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


