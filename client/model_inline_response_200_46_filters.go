/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 November, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.52.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// InlineResponse20046Filters A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.
type InlineResponse20046Filters struct {
	// Conditions
	Conditions []InlineResponse20046FiltersConditions `json:"conditions,omitempty"`
	// Failure Type
	FailureType *string `json:"failureType,omitempty"`
	// Loopback Window (in sec)
	LookbackWindow *int32 `json:"lookbackWindow,omitempty"`
	// Min Duration
	MinDuration *int32 `json:"minDuration,omitempty"`
	// Name
	Name *string `json:"name,omitempty"`
	// Period
	Period *int32 `json:"period,omitempty"`
	// Priority
	Priority *string `json:"priority,omitempty"`
	// Regex
	Regex *string `json:"regex,omitempty"`
	// Selector
	Selector *string `json:"selector,omitempty"`
	// Serials
	Serials []string `json:"serials,omitempty"`
	// SSID Number
	SsidNum *int32 `json:"ssidNum,omitempty"`
	// Tag
	Tag *string `json:"tag,omitempty"`
	// Threshold
	Threshold *int32 `json:"threshold,omitempty"`
	// Timeout
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewInlineResponse20046Filters instantiates a new InlineResponse20046Filters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20046Filters() *InlineResponse20046Filters {
	this := InlineResponse20046Filters{}
	return &this
}

// NewInlineResponse20046FiltersWithDefaults instantiates a new InlineResponse20046Filters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20046FiltersWithDefaults() *InlineResponse20046Filters {
	this := InlineResponse20046Filters{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetConditions() []InlineResponse20046FiltersConditions {
	if o == nil || isNil(o.Conditions) {
		var ret []InlineResponse20046FiltersConditions
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetConditionsOk() ([]InlineResponse20046FiltersConditions, bool) {
	if o == nil || isNil(o.Conditions) {
    return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasConditions() bool {
	if o != nil && !isNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []InlineResponse20046FiltersConditions and assigns it to the Conditions field.
func (o *InlineResponse20046Filters) SetConditions(v []InlineResponse20046FiltersConditions) {
	o.Conditions = v
}

// GetFailureType returns the FailureType field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetFailureType() string {
	if o == nil || isNil(o.FailureType) {
		var ret string
		return ret
	}
	return *o.FailureType
}

// GetFailureTypeOk returns a tuple with the FailureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetFailureTypeOk() (*string, bool) {
	if o == nil || isNil(o.FailureType) {
    return nil, false
	}
	return o.FailureType, true
}

// HasFailureType returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasFailureType() bool {
	if o != nil && !isNil(o.FailureType) {
		return true
	}

	return false
}

// SetFailureType gets a reference to the given string and assigns it to the FailureType field.
func (o *InlineResponse20046Filters) SetFailureType(v string) {
	o.FailureType = &v
}

// GetLookbackWindow returns the LookbackWindow field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetLookbackWindow() int32 {
	if o == nil || isNil(o.LookbackWindow) {
		var ret int32
		return ret
	}
	return *o.LookbackWindow
}

// GetLookbackWindowOk returns a tuple with the LookbackWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetLookbackWindowOk() (*int32, bool) {
	if o == nil || isNil(o.LookbackWindow) {
    return nil, false
	}
	return o.LookbackWindow, true
}

// HasLookbackWindow returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasLookbackWindow() bool {
	if o != nil && !isNil(o.LookbackWindow) {
		return true
	}

	return false
}

// SetLookbackWindow gets a reference to the given int32 and assigns it to the LookbackWindow field.
func (o *InlineResponse20046Filters) SetLookbackWindow(v int32) {
	o.LookbackWindow = &v
}

// GetMinDuration returns the MinDuration field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetMinDuration() int32 {
	if o == nil || isNil(o.MinDuration) {
		var ret int32
		return ret
	}
	return *o.MinDuration
}

// GetMinDurationOk returns a tuple with the MinDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetMinDurationOk() (*int32, bool) {
	if o == nil || isNil(o.MinDuration) {
    return nil, false
	}
	return o.MinDuration, true
}

// HasMinDuration returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasMinDuration() bool {
	if o != nil && !isNil(o.MinDuration) {
		return true
	}

	return false
}

// SetMinDuration gets a reference to the given int32 and assigns it to the MinDuration field.
func (o *InlineResponse20046Filters) SetMinDuration(v int32) {
	o.MinDuration = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse20046Filters) SetName(v string) {
	o.Name = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetPeriod() int32 {
	if o == nil || isNil(o.Period) {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetPeriodOk() (*int32, bool) {
	if o == nil || isNil(o.Period) {
    return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasPeriod() bool {
	if o != nil && !isNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *InlineResponse20046Filters) SetPeriod(v int32) {
	o.Period = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetPriority() string {
	if o == nil || isNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetPriorityOk() (*string, bool) {
	if o == nil || isNil(o.Priority) {
    return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *InlineResponse20046Filters) SetPriority(v string) {
	o.Priority = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetRegex() string {
	if o == nil || isNil(o.Regex) {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetRegexOk() (*string, bool) {
	if o == nil || isNil(o.Regex) {
    return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasRegex() bool {
	if o != nil && !isNil(o.Regex) {
		return true
	}

	return false
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *InlineResponse20046Filters) SetRegex(v string) {
	o.Regex = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetSelector() string {
	if o == nil || isNil(o.Selector) {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetSelectorOk() (*string, bool) {
	if o == nil || isNil(o.Selector) {
    return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasSelector() bool {
	if o != nil && !isNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *InlineResponse20046Filters) SetSelector(v string) {
	o.Selector = &v
}

// GetSerials returns the Serials field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetSerials() []string {
	if o == nil || isNil(o.Serials) {
		var ret []string
		return ret
	}
	return o.Serials
}

// GetSerialsOk returns a tuple with the Serials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetSerialsOk() ([]string, bool) {
	if o == nil || isNil(o.Serials) {
    return nil, false
	}
	return o.Serials, true
}

// HasSerials returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasSerials() bool {
	if o != nil && !isNil(o.Serials) {
		return true
	}

	return false
}

// SetSerials gets a reference to the given []string and assigns it to the Serials field.
func (o *InlineResponse20046Filters) SetSerials(v []string) {
	o.Serials = v
}

// GetSsidNum returns the SsidNum field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetSsidNum() int32 {
	if o == nil || isNil(o.SsidNum) {
		var ret int32
		return ret
	}
	return *o.SsidNum
}

// GetSsidNumOk returns a tuple with the SsidNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetSsidNumOk() (*int32, bool) {
	if o == nil || isNil(o.SsidNum) {
    return nil, false
	}
	return o.SsidNum, true
}

// HasSsidNum returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasSsidNum() bool {
	if o != nil && !isNil(o.SsidNum) {
		return true
	}

	return false
}

// SetSsidNum gets a reference to the given int32 and assigns it to the SsidNum field.
func (o *InlineResponse20046Filters) SetSsidNum(v int32) {
	o.SsidNum = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
    return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *InlineResponse20046Filters) SetTag(v string) {
	o.Tag = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetThreshold() int32 {
	if o == nil || isNil(o.Threshold) {
		var ret int32
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetThresholdOk() (*int32, bool) {
	if o == nil || isNil(o.Threshold) {
    return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasThreshold() bool {
	if o != nil && !isNil(o.Threshold) {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given int32 and assigns it to the Threshold field.
func (o *InlineResponse20046Filters) SetThreshold(v int32) {
	o.Threshold = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *InlineResponse20046Filters) GetTimeout() int32 {
	if o == nil || isNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20046Filters) GetTimeoutOk() (*int32, bool) {
	if o == nil || isNil(o.Timeout) {
    return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *InlineResponse20046Filters) HasTimeout() bool {
	if o != nil && !isNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *InlineResponse20046Filters) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o InlineResponse20046Filters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !isNil(o.FailureType) {
		toSerialize["failureType"] = o.FailureType
	}
	if !isNil(o.LookbackWindow) {
		toSerialize["lookbackWindow"] = o.LookbackWindow
	}
	if !isNil(o.MinDuration) {
		toSerialize["minDuration"] = o.MinDuration
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.Regex) {
		toSerialize["regex"] = o.Regex
	}
	if !isNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	if !isNil(o.Serials) {
		toSerialize["serials"] = o.Serials
	}
	if !isNil(o.SsidNum) {
		toSerialize["ssidNum"] = o.SsidNum
	}
	if !isNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !isNil(o.Threshold) {
		toSerialize["threshold"] = o.Threshold
	}
	if !isNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse20046Filters struct {
	value *InlineResponse20046Filters
	isSet bool
}

func (v NullableInlineResponse20046Filters) Get() *InlineResponse20046Filters {
	return v.value
}

func (v *NullableInlineResponse20046Filters) Set(val *InlineResponse20046Filters) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse20046Filters) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse20046Filters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse20046Filters(val *InlineResponse20046Filters) *NullableInlineResponse20046Filters {
	return &NullableInlineResponse20046Filters{value: val, isSet: true}
}

func (v NullableInlineResponse20046Filters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse20046Filters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


