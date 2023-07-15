/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 July, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.35.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse2003Results Results of the ping request.
type InlineResponse2003Results struct {
	// Number of packets sent
	Sent *int32 `json:"sent,omitempty"`
	// Number of packets received
	Received *int32 `json:"received,omitempty"`
	Loss *InlineResponse2003ResultsLoss `json:"loss,omitempty"`
	Latencies *InlineResponse2003ResultsLatencies `json:"latencies,omitempty"`
	// Received packets
	Replies []InlineResponse2003ResultsReplies `json:"replies,omitempty"`
}

// NewInlineResponse2003Results instantiates a new InlineResponse2003Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003Results() *InlineResponse2003Results {
	this := InlineResponse2003Results{}
	return &this
}

// NewInlineResponse2003ResultsWithDefaults instantiates a new InlineResponse2003Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003ResultsWithDefaults() *InlineResponse2003Results {
	this := InlineResponse2003Results{}
	return &this
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *InlineResponse2003Results) GetSent() int32 {
	if o == nil || isNil(o.Sent) {
		var ret int32
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Results) GetSentOk() (*int32, bool) {
	if o == nil || isNil(o.Sent) {
    return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *InlineResponse2003Results) HasSent() bool {
	if o != nil && !isNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given int32 and assigns it to the Sent field.
func (o *InlineResponse2003Results) SetSent(v int32) {
	o.Sent = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *InlineResponse2003Results) GetReceived() int32 {
	if o == nil || isNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Results) GetReceivedOk() (*int32, bool) {
	if o == nil || isNil(o.Received) {
    return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *InlineResponse2003Results) HasReceived() bool {
	if o != nil && !isNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *InlineResponse2003Results) SetReceived(v int32) {
	o.Received = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *InlineResponse2003Results) GetLoss() InlineResponse2003ResultsLoss {
	if o == nil || isNil(o.Loss) {
		var ret InlineResponse2003ResultsLoss
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Results) GetLossOk() (*InlineResponse2003ResultsLoss, bool) {
	if o == nil || isNil(o.Loss) {
    return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *InlineResponse2003Results) HasLoss() bool {
	if o != nil && !isNil(o.Loss) {
		return true
	}

	return false
}

// SetLoss gets a reference to the given InlineResponse2003ResultsLoss and assigns it to the Loss field.
func (o *InlineResponse2003Results) SetLoss(v InlineResponse2003ResultsLoss) {
	o.Loss = &v
}

// GetLatencies returns the Latencies field value if set, zero value otherwise.
func (o *InlineResponse2003Results) GetLatencies() InlineResponse2003ResultsLatencies {
	if o == nil || isNil(o.Latencies) {
		var ret InlineResponse2003ResultsLatencies
		return ret
	}
	return *o.Latencies
}

// GetLatenciesOk returns a tuple with the Latencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Results) GetLatenciesOk() (*InlineResponse2003ResultsLatencies, bool) {
	if o == nil || isNil(o.Latencies) {
    return nil, false
	}
	return o.Latencies, true
}

// HasLatencies returns a boolean if a field has been set.
func (o *InlineResponse2003Results) HasLatencies() bool {
	if o != nil && !isNil(o.Latencies) {
		return true
	}

	return false
}

// SetLatencies gets a reference to the given InlineResponse2003ResultsLatencies and assigns it to the Latencies field.
func (o *InlineResponse2003Results) SetLatencies(v InlineResponse2003ResultsLatencies) {
	o.Latencies = &v
}

// GetReplies returns the Replies field value if set, zero value otherwise.
func (o *InlineResponse2003Results) GetReplies() []InlineResponse2003ResultsReplies {
	if o == nil || isNil(o.Replies) {
		var ret []InlineResponse2003ResultsReplies
		return ret
	}
	return o.Replies
}

// GetRepliesOk returns a tuple with the Replies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Results) GetRepliesOk() ([]InlineResponse2003ResultsReplies, bool) {
	if o == nil || isNil(o.Replies) {
    return nil, false
	}
	return o.Replies, true
}

// HasReplies returns a boolean if a field has been set.
func (o *InlineResponse2003Results) HasReplies() bool {
	if o != nil && !isNil(o.Replies) {
		return true
	}

	return false
}

// SetReplies gets a reference to the given []InlineResponse2003ResultsReplies and assigns it to the Replies field.
func (o *InlineResponse2003Results) SetReplies(v []InlineResponse2003ResultsReplies) {
	o.Replies = v
}

func (o InlineResponse2003Results) MarshalJSON() ([]byte, error) {
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

type NullableInlineResponse2003Results struct {
	value *InlineResponse2003Results
	isSet bool
}

func (v NullableInlineResponse2003Results) Get() *InlineResponse2003Results {
	return v.value
}

func (v *NullableInlineResponse2003Results) Set(val *InlineResponse2003Results) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003Results) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003Results) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003Results(val *InlineResponse2003Results) *NullableInlineResponse2003Results {
	return &NullableInlineResponse2003Results{value: val, isSet: true}
}

func (v NullableInlineResponse2003Results) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003Results) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


