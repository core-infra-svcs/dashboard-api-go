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

// InlineResponse200200Rules struct for InlineResponse200200Rules
type InlineResponse200200Rules struct {
	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. 
	Definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions `json:"definitions"`
	PerClientBandwidthLimits *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`
	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint. 
	DscpTagValue *int32 `json:"dscpTagValue,omitempty"`
	//     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'. 
	PcpTagValue *int32 `json:"pcpTagValue,omitempty"`
}

// NewInlineResponse200200Rules instantiates a new InlineResponse200200Rules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse200200Rules(definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) *InlineResponse200200Rules {
	this := InlineResponse200200Rules{}
	this.Definitions = definitions
	return &this
}

// NewInlineResponse200200RulesWithDefaults instantiates a new InlineResponse200200Rules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse200200RulesWithDefaults() *InlineResponse200200Rules {
	this := InlineResponse200200Rules{}
	return &this
}

// GetDefinitions returns the Definitions field value
func (o *InlineResponse200200Rules) GetDefinitions() []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions {
	if o == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions
		return ret
	}

	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetDefinitionsOk() ([]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, bool) {
	if o == nil {
    return nil, false
	}
	return o.Definitions, true
}

// SetDefinitions sets field value
func (o *InlineResponse200200Rules) SetDefinitions(v []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) {
	o.Definitions = v
}

// GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field value if set, zero value otherwise.
func (o *InlineResponse200200Rules) GetPerClientBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	if o == nil || isNil(o.PerClientBandwidthLimits) {
		var ret NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits
		return ret
	}
	return *o.PerClientBandwidthLimits
}

// GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetPerClientBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits, bool) {
	if o == nil || isNil(o.PerClientBandwidthLimits) {
    return nil, false
	}
	return o.PerClientBandwidthLimits, true
}

// HasPerClientBandwidthLimits returns a boolean if a field has been set.
func (o *InlineResponse200200Rules) HasPerClientBandwidthLimits() bool {
	if o != nil && !isNil(o.PerClientBandwidthLimits) {
		return true
	}

	return false
}

// SetPerClientBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits and assigns it to the PerClientBandwidthLimits field.
func (o *InlineResponse200200Rules) SetPerClientBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) {
	o.PerClientBandwidthLimits = &v
}

// GetDscpTagValue returns the DscpTagValue field value if set, zero value otherwise.
func (o *InlineResponse200200Rules) GetDscpTagValue() int32 {
	if o == nil || isNil(o.DscpTagValue) {
		var ret int32
		return ret
	}
	return *o.DscpTagValue
}

// GetDscpTagValueOk returns a tuple with the DscpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetDscpTagValueOk() (*int32, bool) {
	if o == nil || isNil(o.DscpTagValue) {
    return nil, false
	}
	return o.DscpTagValue, true
}

// HasDscpTagValue returns a boolean if a field has been set.
func (o *InlineResponse200200Rules) HasDscpTagValue() bool {
	if o != nil && !isNil(o.DscpTagValue) {
		return true
	}

	return false
}

// SetDscpTagValue gets a reference to the given int32 and assigns it to the DscpTagValue field.
func (o *InlineResponse200200Rules) SetDscpTagValue(v int32) {
	o.DscpTagValue = &v
}

// GetPcpTagValue returns the PcpTagValue field value if set, zero value otherwise.
func (o *InlineResponse200200Rules) GetPcpTagValue() int32 {
	if o == nil || isNil(o.PcpTagValue) {
		var ret int32
		return ret
	}
	return *o.PcpTagValue
}

// GetPcpTagValueOk returns a tuple with the PcpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse200200Rules) GetPcpTagValueOk() (*int32, bool) {
	if o == nil || isNil(o.PcpTagValue) {
    return nil, false
	}
	return o.PcpTagValue, true
}

// HasPcpTagValue returns a boolean if a field has been set.
func (o *InlineResponse200200Rules) HasPcpTagValue() bool {
	if o != nil && !isNil(o.PcpTagValue) {
		return true
	}

	return false
}

// SetPcpTagValue gets a reference to the given int32 and assigns it to the PcpTagValue field.
func (o *InlineResponse200200Rules) SetPcpTagValue(v int32) {
	o.PcpTagValue = &v
}

func (o InlineResponse200200Rules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["definitions"] = o.Definitions
	}
	if !isNil(o.PerClientBandwidthLimits) {
		toSerialize["perClientBandwidthLimits"] = o.PerClientBandwidthLimits
	}
	if !isNil(o.DscpTagValue) {
		toSerialize["dscpTagValue"] = o.DscpTagValue
	}
	if !isNil(o.PcpTagValue) {
		toSerialize["pcpTagValue"] = o.PcpTagValue
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse200200Rules struct {
	value *InlineResponse200200Rules
	isSet bool
}

func (v NullableInlineResponse200200Rules) Get() *InlineResponse200200Rules {
	return v.value
}

func (v *NullableInlineResponse200200Rules) Set(val *InlineResponse200200Rules) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse200200Rules) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse200200Rules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse200200Rules(val *InlineResponse200200Rules) *NullableInlineResponse200200Rules {
	return &NullableInlineResponse200200Rules{value: val, isSet: true}
}

func (v NullableInlineResponse200200Rules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse200200Rules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


