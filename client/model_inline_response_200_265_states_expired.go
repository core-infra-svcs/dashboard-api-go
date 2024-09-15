/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 September, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.50.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse200265StatesExpired Data for expired licenses
type InlineResponse200265StatesExpired struct {
	// The number of expired licenses
	Count *int32 `json:"count,omitempty"`
}

// NewInlineResponse200265StatesExpired instantiates a new InlineResponse200265StatesExpired object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200265StatesExpired() *InlineResponse200265StatesExpired {
	this := InlineResponse200265StatesExpired{}
	return &this
}

// NewInlineResponse200265StatesExpiredWithDefaults instantiates a new InlineResponse200265StatesExpired object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200265StatesExpiredWithDefaults() *InlineResponse200265StatesExpired {
	this := InlineResponse200265StatesExpired{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *InlineResponse200265StatesExpired) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200265StatesExpired) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
    return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *InlineResponse200265StatesExpired) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *InlineResponse200265StatesExpired) SetCount(v int32) {
	o.Count = &v
}

func (o InlineResponse200265StatesExpired) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200265StatesExpired struct {
	value *InlineResponse200265StatesExpired
	isSet bool
}

func (v NullableInlineResponse200265StatesExpired) Get() *InlineResponse200265StatesExpired {
	return v.value
}

func (v *NullableInlineResponse200265StatesExpired) Set(val *InlineResponse200265StatesExpired) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200265StatesExpired) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200265StatesExpired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200265StatesExpired(val *InlineResponse200265StatesExpired) *NullableInlineResponse200265StatesExpired {
	return &NullableInlineResponse200265StatesExpired{value: val, isSet: true}
}

func (v NullableInlineResponse200265StatesExpired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200265StatesExpired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


