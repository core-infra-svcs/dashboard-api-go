# InlineResponse200245Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]InlineResponse20092Definitions**](InlineResponse20092Definitions.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**InlineResponse20092PerClientBandwidthLimits**](InlineResponse20092PerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **NullableInt32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**PcpTagValue** | Pointer to **NullableInt32** |     The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).     null means &#39;Do not set PCP tag&#39;.  | [optional] 

## Methods

### NewInlineResponse200245Rules

`func NewInlineResponse200245Rules(definitions []InlineResponse20092Definitions, ) *InlineResponse200245Rules`

NewInlineResponse200245Rules instantiates a new InlineResponse200245Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200245RulesWithDefaults

`func NewInlineResponse200245RulesWithDefaults() *InlineResponse200245Rules`

NewInlineResponse200245RulesWithDefaults instantiates a new InlineResponse200245Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *InlineResponse200245Rules) GetDefinitions() []InlineResponse20092Definitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *InlineResponse200245Rules) GetDefinitionsOk() (*[]InlineResponse20092Definitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *InlineResponse200245Rules) SetDefinitions(v []InlineResponse20092Definitions)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *InlineResponse200245Rules) GetPerClientBandwidthLimits() InlineResponse20092PerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *InlineResponse200245Rules) GetPerClientBandwidthLimitsOk() (*InlineResponse20092PerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *InlineResponse200245Rules) SetPerClientBandwidthLimits(v InlineResponse20092PerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *InlineResponse200245Rules) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *InlineResponse200245Rules) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *InlineResponse200245Rules) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *InlineResponse200245Rules) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *InlineResponse200245Rules) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### SetDscpTagValueNil

`func (o *InlineResponse200245Rules) SetDscpTagValueNil(b bool)`

 SetDscpTagValueNil sets the value for DscpTagValue to be an explicit nil

### UnsetDscpTagValue
`func (o *InlineResponse200245Rules) UnsetDscpTagValue()`

UnsetDscpTagValue ensures that no value is present for DscpTagValue, not even an explicit nil
### GetPcpTagValue

`func (o *InlineResponse200245Rules) GetPcpTagValue() int32`

GetPcpTagValue returns the PcpTagValue field if non-nil, zero value otherwise.

### GetPcpTagValueOk

`func (o *InlineResponse200245Rules) GetPcpTagValueOk() (*int32, bool)`

GetPcpTagValueOk returns a tuple with the PcpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcpTagValue

`func (o *InlineResponse200245Rules) SetPcpTagValue(v int32)`

SetPcpTagValue sets PcpTagValue field to given value.

### HasPcpTagValue

`func (o *InlineResponse200245Rules) HasPcpTagValue() bool`

HasPcpTagValue returns a boolean if a field has been set.

### SetPcpTagValueNil

`func (o *InlineResponse200245Rules) SetPcpTagValueNil(b bool)`

 SetPcpTagValueNil sets the value for PcpTagValue to be an explicit nil

### UnsetPcpTagValue
`func (o *InlineResponse200245Rules) UnsetPcpTagValue()`

UnsetPcpTagValue ensures that no value is present for PcpTagValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


