/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 05 October, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.26.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules struct for NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules
type NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules struct {
	//     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required. 
	Definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions `json:"definitions"`
	PerClientBandwidthLimits *NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits `json:"perClientBandwidthLimits,omitempty"`
	//     The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint. 
	DscpTagValue *int32 `json:"dscpTagValue,omitempty"`
	//     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means 'Do not set PCP tag'. 
	PcpTagValue *int32 `json:"pcpTagValue,omitempty"`
	//     A string, indicating the priority level for packets bound to your rule.     Can be 'low', 'normal' or 'high'. 
	Priority *string `json:"priority,omitempty"`
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules(definitions []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules{}
	this.Definitions = definitions
	return &this
}

// NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesWithDefaults instantiates a new NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRulesWithDefaults() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules {
	this := NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules{}
	return &this
}

// GetDefinitions returns the Definitions field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDefinitions() []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions {
	if o == nil {
		var ret []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions
		return ret
	}

	return o.Definitions
}

// GetDefinitionsOk returns a tuple with the Definitions field value
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDefinitionsOk() (*[]NetworksNetworkIdApplianceTrafficShapingRulesDefinitions, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Definitions, true
}

// SetDefinitions sets field value
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetDefinitions(v []NetworksNetworkIdApplianceTrafficShapingRulesDefinitions) {
	o.Definitions = v
}

// GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPerClientBandwidthLimits() NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits {
	if o == nil || o.PerClientBandwidthLimits == nil {
		var ret NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits
		return ret
	}
	return *o.PerClientBandwidthLimits
}

// GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPerClientBandwidthLimitsOk() (*NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits, bool) {
	if o == nil || o.PerClientBandwidthLimits == nil {
		return nil, false
	}
	return o.PerClientBandwidthLimits, true
}

// HasPerClientBandwidthLimits returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasPerClientBandwidthLimits() bool {
	if o != nil && o.PerClientBandwidthLimits != nil {
		return true
	}

	return false
}

// SetPerClientBandwidthLimits gets a reference to the given NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits and assigns it to the PerClientBandwidthLimits field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetPerClientBandwidthLimits(v NetworksNetworkIdApplianceTrafficShapingRulesPerClientBandwidthLimits) {
	o.PerClientBandwidthLimits = &v
}

// GetDscpTagValue returns the DscpTagValue field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDscpTagValue() int32 {
	if o == nil || o.DscpTagValue == nil {
		var ret int32
		return ret
	}
	return *o.DscpTagValue
}

// GetDscpTagValueOk returns a tuple with the DscpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetDscpTagValueOk() (*int32, bool) {
	if o == nil || o.DscpTagValue == nil {
		return nil, false
	}
	return o.DscpTagValue, true
}

// HasDscpTagValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasDscpTagValue() bool {
	if o != nil && o.DscpTagValue != nil {
		return true
	}

	return false
}

// SetDscpTagValue gets a reference to the given int32 and assigns it to the DscpTagValue field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetDscpTagValue(v int32) {
	o.DscpTagValue = &v
}

// GetPcpTagValue returns the PcpTagValue field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPcpTagValue() int32 {
	if o == nil || o.PcpTagValue == nil {
		var ret int32
		return ret
	}
	return *o.PcpTagValue
}

// GetPcpTagValueOk returns a tuple with the PcpTagValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPcpTagValueOk() (*int32, bool) {
	if o == nil || o.PcpTagValue == nil {
		return nil, false
	}
	return o.PcpTagValue, true
}

// HasPcpTagValue returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasPcpTagValue() bool {
	if o != nil && o.PcpTagValue != nil {
		return true
	}

	return false
}

// SetPcpTagValue gets a reference to the given int32 and assigns it to the PcpTagValue field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetPcpTagValue(v int32) {
	o.PcpTagValue = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) SetPriority(v string) {
	o.Priority = &v
}

func (o NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["definitions"] = o.Definitions
	}
	if o.PerClientBandwidthLimits != nil {
		toSerialize["perClientBandwidthLimits"] = o.PerClientBandwidthLimits
	}
	if o.DscpTagValue != nil {
		toSerialize["dscpTagValue"] = o.DscpTagValue
	}
	if o.PcpTagValue != nil {
		toSerialize["pcpTagValue"] = o.PcpTagValue
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	return json.Marshal(toSerialize)
}

type NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules struct {
	value *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules
	isSet bool
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) Get() *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules {
	return v.value
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) Set(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules(val *NetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules {
	return &NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules{value: val, isSet: true}
}

func (v NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworksNetworkIdGroupPoliciesFirewallAndTrafficShapingTrafficShapingRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


