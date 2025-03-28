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

// InlineResponse20025Results Results of the ping request.
type InlineResponse20025Results struct {
	// Number of packets sent
	Sent *int32 `json:"sent,omitempty"`
	// Number of packets received
	Received *int32 `json:"received,omitempty"`
	Loss *InlineResponse20025ResultsLoss `json:"loss,omitempty"`
	Latencies *InlineResponse20025ResultsLatencies `json:"latencies,omitempty"`
	// Received packets
	Replies []InlineResponse20025ResultsReplies `json:"replies,omitempty"`
}

// NewInlineResponse20025Results instantiates a new InlineResponse20025Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20025Results() *InlineResponse20025Results {
	this := InlineResponse20025Results{}
	return &this
}

// NewInlineResponse20025ResultsWithDefaults instantiates a new InlineResponse20025Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20025ResultsWithDefaults() *InlineResponse20025Results {
	this := InlineResponse20025Results{}
	return &this
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse20025Results) GetSent() int32 {
	if o == nil || isNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Results) GetSentOk() (*int32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse20025Results) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *InlineResponse20025Results) SetSent(v int32) {
	o.Sent = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *InlineResponse20025Results) GetReceived() int32 {
	if o == nil || isNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Results) GetReceivedOk() (*int32, bool) {
	if o == nil || isNil(o.Received) {
    return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *InlineResponse20025Results) HasReceived() bool {
	if o != nil && !isNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *InlineResponse20025Results) SetReceived(v int32) {
	o.Received = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *InlineResponse20025Results) GetLoss() InlineResponse20025ResultsLoss {
	if o == nil || isNil(o.Loss) {
		var ret InlineResponse20025ResultsLoss
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Results) GetLossOk() (*InlineResponse20025ResultsLoss, bool) {
	if o == nil || isNil(o.Loss) {
    return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *InlineResponse20025Results) HasLoss() bool {
	if o != nil && !isNil(o.Loss) {
		return true
	}

	return false
}

// SetLoss gets a reference to the given InlineResponse20025ResultsLoss and assigns it to the Loss field.
func (o *InlineResponse20025Results) SetLoss(v InlineResponse20025ResultsLoss) {
	o.Loss = &v
}

// GetLatencies returns the Latencies field value if set, zero value otherwise.
func (o *InlineResponse20025Results) GetLatencies() InlineResponse20025ResultsLatencies {
	if o == nil || isNil(o.Latencies) {
		var ret InlineResponse20025ResultsLatencies
		return ret
	}
	return *o.Latencies
}

// GetLatenciesOk returns a tuple with the Latencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Results) GetLatenciesOk() (*InlineResponse20025ResultsLatencies, bool) {
	if o == nil || isNil(o.Latencies) {
    return nil, false
	}
	return o.Latencies, true
}

// HasLatencies returns a boolean if a field has been set.
func (o *InlineResponse20025Results) HasLatencies() bool {
	if o != nil && !isNil(o.Latencies) {
		return true
	}

	return false
}

// SetLatencies gets a reference to the given InlineResponse20025ResultsLatencies and assigns it to the Latencies field.
func (o *InlineResponse20025Results) SetLatencies(v InlineResponse20025ResultsLatencies) {
	o.Latencies = &v
}

// GetReplies returns the Replies field value if set, zero value otherwise.
func (o *InlineResponse20025Results) GetReplies() []InlineResponse20025ResultsReplies {
	if o == nil || isNil(o.Replies) {
		var ret []InlineResponse20025ResultsReplies
		return ret
	}
	return o.Replies
}

// GetRepliesOk returns a tuple with the Replies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20025Results) GetRepliesOk() ([]InlineResponse20025ResultsReplies, bool) {
	if o == nil || isNil(o.Replies) {
    return nil, false
	}
	return o.Replies, true
}

// HasReplies returns a boolean if a field has been set.
func (o *InlineResponse20025Results) HasReplies() bool {
	if o != nil && !isNil(o.Replies) {
		return true
	}

	return false
}

// SetReplies gets a reference to the given []InlineResponse20025ResultsReplies and assigns it to the Replies field.
func (o *InlineResponse20025Results) SetReplies(v []InlineResponse20025ResultsReplies) {
	o.Replies = v
}

func (o InlineResponse20025Results) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !isNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !isNil(o.Loss) {
		toSerialize["loss"] = o.Loss
	}
	if !isNil(o.Latencies) {
		toSerialize["latencies"] = o.Latencies
	}
	if !isNil(o.Replies) {
		toSerialize["replies"] = o.Replies
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20025Results struct {
	value *InlineResponse20025Results
	isSet bool
}

func (v NullableInlineResponse20025Results) Get() *InlineResponse20025Results {
	return v.value
}

func (v *NullableInlineResponse20025Results) Set(val *InlineResponse20025Results) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20025Results) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20025Results) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20025Results(val *InlineResponse20025Results) *NullableInlineResponse20025Results {
	return &NullableInlineResponse20025Results{value: val, isSet: true}
}

func (v NullableInlineResponse20025Results) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20025Results) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


