/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 December, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.53.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject173 struct for InlineObject173
type InlineObject173 struct {
	// The name of the template
	Name *string `json:"name,omitempty"`
	// The liquid template used for the body of the webhook message.
	Body *string `json:"body,omitempty"`
	// The liquid template used with the webhook headers.
	Headers []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 `json:"headers,omitempty"`
	// A file containing liquid template used for the body of the webhook message.
	BodyFile *string `json:"bodyFile,omitempty"`
	// A file containing the liquid template used with the webhook headers.
	HeadersFile *string `json:"headersFile,omitempty"`
}

// NewInlineObject173 instantiates a new InlineObject173 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject173() *InlineObject173 {
	this := InlineObject173{}
	return &this
}

// NewInlineObject173WithDefaults instantiates a new InlineObject173 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject173WithDefaults() *InlineObject173 {
	this := InlineObject173{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineObject173) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject173) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineObject173) SetName(v string) {
	o.Name = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineObject173) GetBody() string {
	if o == nil || isNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetBodyOk() (*string, bool) {
	if o == nil || isNil(o.Body) {
    return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineObject173) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InlineObject173) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineObject173) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 {
	if o == nil || isNil(o.Headers) {
		var ret []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetHeadersOk() ([]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineObject173) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 and assigns it to the Headers field.
func (o *InlineObject173) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) {
	o.Headers = v
}

// GetBodyFile returns the BodyFile field value if set, zero value otherwise.
func (o *InlineObject173) GetBodyFile() string {
	if o == nil || isNil(o.BodyFile) {
		var ret string
		return ret
	}
	return *o.BodyFile
}

// GetBodyFileOk returns a tuple with the BodyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetBodyFileOk() (*string, bool) {
	if o == nil || isNil(o.BodyFile) {
    return nil, false
	}
	return o.BodyFile, true
}

// HasBodyFile returns a boolean if a field has been set.
func (o *InlineObject173) HasBodyFile() bool {
	if o != nil && !isNil(o.BodyFile) {
		return true
	}

	return false
}

// SetBodyFile gets a reference to the given string and assigns it to the BodyFile field.
func (o *InlineObject173) SetBodyFile(v string) {
	o.BodyFile = &v
}

// GetHeadersFile returns the HeadersFile field value if set, zero value otherwise.
func (o *InlineObject173) GetHeadersFile() string {
	if o == nil || isNil(o.HeadersFile) {
		var ret string
		return ret
	}
	return *o.HeadersFile
}

// GetHeadersFileOk returns a tuple with the HeadersFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject173) GetHeadersFileOk() (*string, bool) {
	if o == nil || isNil(o.HeadersFile) {
    return nil, false
	}
	return o.HeadersFile, true
}

// HasHeadersFile returns a boolean if a field has been set.
func (o *InlineObject173) HasHeadersFile() bool {
	if o != nil && !isNil(o.HeadersFile) {
		return true
	}

	return false
}

// SetHeadersFile gets a reference to the given string and assigns it to the HeadersFile field.
func (o *InlineObject173) SetHeadersFile(v string) {
	o.HeadersFile = &v
}

func (o InlineObject173) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	if !isNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !isNil(o.BodyFile) {
		toSerialize["bodyFile"] = o.BodyFile
	}
	if !isNil(o.HeadersFile) {
		toSerialize["headersFile"] = o.HeadersFile
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject173 struct {
	value *InlineObject173
	isSet bool
}

func (v NullableInlineObject173) Get() *InlineObject173 {
	return v.value
}

func (v *NullableInlineObject173) Set(val *InlineObject173) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject173) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject173) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject173(val *InlineObject173) *NullableInlineObject173 {
	return &NullableInlineObject173{value: val, isSet: true}
}

func (v NullableInlineObject173) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject173) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


