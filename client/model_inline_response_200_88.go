/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 February, 2025 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.55.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// InlineResponse20088 struct for InlineResponse20088
type InlineResponse20088 struct {
	// The start time from which daily traffic data was collected
	Ts *time.Time `json:"ts,omitempty"`
	// The name of the application the client is connected to
	Application *string `json:"application,omitempty"`
	// The IP or web address the client is connected to
	Destination *string `json:"destination,omitempty"`
	// The client protocol
	Protocol *string `json:"protocol,omitempty"`
	// The port the client is connected to
	Port *int32 `json:"port,omitempty"`
	// Usage received by the client
	Recv *float32 `json:"recv,omitempty"`
	// Usage sent by the client
	Sent *float32 `json:"sent,omitempty"`
	// The number of flows the client has
	NumFlows *int32 `json:"numFlows,omitempty"`
	// The amount of seconds the client was active
	ActiveSeconds *int32 `json:"activeSeconds,omitempty"`
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

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *InlineResponse20088) GetTs() time.Time {
	if o == nil || isNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetTsOk() (*time.Time, bool) {
	if o == nil || isNil(o.Ts) {
    return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *InlineResponse20088) HasTs() bool {
	if o != nil && !isNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *InlineResponse20088) SetTs(v time.Time) {
	o.Ts = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *InlineResponse20088) GetApplication() string {
	if o == nil || isNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetApplicationOk() (*string, bool) {
	if o == nil || isNil(o.Application) {
    return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *InlineResponse20088) HasApplication() bool {
	if o != nil && !isNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *InlineResponse20088) SetApplication(v string) {
	o.Application = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *InlineResponse20088) GetDestination() string {
	if o == nil || isNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetDestinationOk() (*string, bool) {
	if o == nil || isNil(o.Destination) {
    return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *InlineResponse20088) HasDestination() bool {
	if o != nil && !isNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *InlineResponse20088) SetDestination(v string) {
	o.Destination = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *InlineResponse20088) GetProtocol() string {
	if o == nil || isNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetProtocolOk() (*string, bool) {
	if o == nil || isNil(o.Protocol) {
    return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *InlineResponse20088) HasProtocol() bool {
	if o != nil && !isNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *InlineResponse20088) SetProtocol(v string) {
	o.Protocol = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *InlineResponse20088) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *InlineResponse20088) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *InlineResponse20088) SetPort(v int32) {
	o.Port = &v
}

// GetRecv returns the Recv field value if set, zero value otherwise.
func (o *InlineResponse20088) GetRecv() float32 {
	if o == nil || isNil(o.Recv) {
		var ret float32
		return ret
	}
	return *o.Recv
}

// GetRecvOk returns a tuple with the Recv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetRecvOk() (*float32, bool) {
	if o == nil || isNil(o.Recv) {
    return nil, false
	}
	return o.Recv, true
}

// HasRecv returns a boolean if a field has been set.
func (o *InlineResponse20088) HasRecv() bool {
	if o != nil && !isNil(o.Recv) {
		return true
	}

	return false
}

// SetRecv gets a reference to the given float32 and assigns it to the Recv field.
func (o *InlineResponse20088) SetRecv(v float32) {
	o.Recv = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse20088) GetSent() float32 {
	if o == nil || isNil(o.Sent) {
		var ret float32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetSentOk() (*float32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse20088) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given float32 and assigns it to the Sent field.
func (o *InlineResponse20088) SetSent(v float32) {
	o.Sent = &v
}

// GetNumFlows returns the NumFlows field value if set, zero value otherwise.
func (o *InlineResponse20088) GetNumFlows() int32 {
	if o == nil || isNil(o.NumFlows) {
		var ret int32
		return ret
	}
	return *o.NumFlows
}

// GetNumFlowsOk returns a tuple with the NumFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetNumFlowsOk() (*int32, bool) {
	if o == nil || isNil(o.NumFlows) {
    return nil, false
	}
	return o.NumFlows, true
}

// HasNumFlows returns a boolean if a field has been set.
func (o *InlineResponse20088) HasNumFlows() bool {
	if o != nil && !isNil(o.NumFlows) {
		return true
	}

	return false
}

// SetNumFlows gets a reference to the given int32 and assigns it to the NumFlows field.
func (o *InlineResponse20088) SetNumFlows(v int32) {
	o.NumFlows = &v
}

// GetActiveSeconds returns the ActiveSeconds field value if set, zero value otherwise.
func (o *InlineResponse20088) GetActiveSeconds() int32 {
	if o == nil || isNil(o.ActiveSeconds) {
		var ret int32
		return ret
	}
	return *o.ActiveSeconds
}

// GetActiveSecondsOk returns a tuple with the ActiveSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20088) GetActiveSecondsOk() (*int32, bool) {
	if o == nil || isNil(o.ActiveSeconds) {
    return nil, false
	}
	return o.ActiveSeconds, true
}

// HasActiveSeconds returns a boolean if a field has been set.
func (o *InlineResponse20088) HasActiveSeconds() bool {
	if o != nil && !isNil(o.ActiveSeconds) {
		return true
	}

	return false
}

// SetActiveSeconds gets a reference to the given int32 and assigns it to the ActiveSeconds field.
func (o *InlineResponse20088) SetActiveSeconds(v int32) {
	o.ActiveSeconds = &v
}

func (o InlineResponse20088) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !isNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !isNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !isNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.Recv) {
		toSerialize["recv"] = o.Recv
	}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.NumFlows) {
		toSerialize["numFlows"] = o.NumFlows
	}
	if !isNil(o.ActiveSeconds) {
		toSerialize["activeSeconds"] = o.ActiveSeconds
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


