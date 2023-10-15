/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineObject155 struct for InlineObject155
type InlineObject155 struct {
	// The name of the new template
	Name string `json:"name"`
	// The liquid template used for the body of the webhook message. Either `body` or `bodyFile` must be specified.
	Body *string `json:"body,omitempty"`
	// The liquid template used with the webhook headers.
	Headers []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 `json:"headers,omitempty"`
	// A file containing liquid template used for the body of the webhook message. Either `body` or `bodyFile` must be specified.
	BodyFile *string `json:"bodyFile,omitempty"`
	// A file containing the liquid template used with the webhook headers.
	HeadersFile *string `json:"headersFile,omitempty"`
}

// NewInlineObject155 instantiates a new InlineObject155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject155(name string) *InlineObject155 {
	this := InlineObject155{}
	this.Name = name
	return &this
}

// NewInlineObject155WithDefaults instantiates a new InlineObject155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject155WithDefaults() *InlineObject155 {
	this := InlineObject155{}
	return &this
}

// GetName returns the Name field value
func (o *InlineObject155) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *InlineObject155) SetName(v string) {
	o.Name = v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *InlineObject155) GetBody() string {
	if o == nil || isNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetBodyOk() (*string, bool) {
	if o == nil || isNil(o.Body) {
    return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *InlineObject155) HasBody() bool {
	if o != nil && !isNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *InlineObject155) SetBody(v string) {
	o.Body = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *InlineObject155) GetHeaders() []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 {
	if o == nil || isNil(o.Headers) {
		var ret []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetHeadersOk() ([]NetworksNetworkIdWebhooksPayloadTemplatesHeaders1, bool) {
	if o == nil || isNil(o.Headers) {
    return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *InlineObject155) HasHeaders() bool {
	if o != nil && !isNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1 and assigns it to the Headers field.
func (o *InlineObject155) SetHeaders(v []NetworksNetworkIdWebhooksPayloadTemplatesHeaders1) {
	o.Headers = v
}

// GetBodyFile returns the BodyFile field value if set, zero value otherwise.
func (o *InlineObject155) GetBodyFile() string {
	if o == nil || isNil(o.BodyFile) {
		var ret string
		return ret
	}
	return *o.BodyFile
}

// GetBodyFileOk returns a tuple with the BodyFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetBodyFileOk() (*string, bool) {
	if o == nil || isNil(o.BodyFile) {
    return nil, false
	}
	return o.BodyFile, true
}

// HasBodyFile returns a boolean if a field has been set.
func (o *InlineObject155) HasBodyFile() bool {
	if o != nil && !isNil(o.BodyFile) {
		return true
	}

	return false
}

// SetBodyFile gets a reference to the given string and assigns it to the BodyFile field.
func (o *InlineObject155) SetBodyFile(v string) {
	o.BodyFile = &v
}

// GetHeadersFile returns the HeadersFile field value if set, zero value otherwise.
func (o *InlineObject155) GetHeadersFile() string {
	if o == nil || isNil(o.HeadersFile) {
		var ret string
		return ret
	}
	return *o.HeadersFile
}

// GetHeadersFileOk returns a tuple with the HeadersFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject155) GetHeadersFileOk() (*string, bool) {
	if o == nil || isNil(o.HeadersFile) {
    return nil, false
	}
	return o.HeadersFile, true
}

// HasHeadersFile returns a boolean if a field has been set.
func (o *InlineObject155) HasHeadersFile() bool {
	if o != nil && !isNil(o.HeadersFile) {
		return true
	}

	return false
}

// SetHeadersFile gets a reference to the given string and assigns it to the HeadersFile field.
func (o *InlineObject155) SetHeadersFile(v string) {
	o.HeadersFile = &v
}

func (o InlineObject155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
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

type NullableInlineObject155 struct {
	value *InlineObject155
	isSet bool
}

func (v NullableInlineObject155) Get() *InlineObject155 {
	return v.value
}

func (v *NullableInlineObject155) Set(val *InlineObject155) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject155) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject155(val *InlineObject155) *NullableInlineObject155 {
	return &NullableInlineObject155{value: val, isSet: true}
}

func (v NullableInlineObject155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


