# InlineResponse200101Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]InlineResponse200101Definitions**](InlineResponse200101Definitions.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**InlineResponse200101PerClientBandwidthLimits**](InlineResponse200101PerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **NullableInt32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**Priority** | Pointer to **string** |     A string, indicating the priority level for packets bound to your rule.     Can be &#39;low&#39;, &#39;normal&#39; or &#39;high&#39;. This field is supported for Security Appliance networks.  | [optional] 

## Methods

### NewInlineResponse200101Rules

`func NewInlineResponse200101Rules(definitions []InlineResponse200101Definitions, ) *InlineResponse200101Rules`

NewInlineResponse200101Rules instantiates a new InlineResponse200101Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200101RulesWithDefaults

`func NewInlineResponse200101RulesWithDefaults() *InlineResponse200101Rules`

NewInlineResponse200101RulesWithDefaults instantiates a new InlineResponse200101Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *InlineResponse200101Rules) GetDefinitions() []InlineResponse200101Definitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *InlineResponse200101Rules) GetDefinitionsOk() (*[]InlineResponse200101Definitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *InlineResponse200101Rules) SetDefinitions(v []InlineResponse200101Definitions)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *InlineResponse200101Rules) GetPerClientBandwidthLimits() InlineResponse200101PerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *InlineResponse200101Rules) GetPerClientBandwidthLimitsOk() (*InlineResponse200101PerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *InlineResponse200101Rules) SetPerClientBandwidthLimits(v InlineResponse200101PerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *InlineResponse200101Rules) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *InlineResponse200101Rules) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *InlineResponse200101Rules) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *InlineResponse200101Rules) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *InlineResponse200101Rules) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### SetDscpTagValueNil

`func (o *InlineResponse200101Rules) SetDscpTagValueNil(b bool)`

 SetDscpTagValueNil sets the value for DscpTagValue to be an explicit nil

### UnsetDscpTagValue
`func (o *InlineResponse200101Rules) UnsetDscpTagValue()`

UnsetDscpTagValue ensures that no value is present for DscpTagValue, not even an explicit nil
### GetPriority

`func (o *InlineResponse200101Rules) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse200101Rules) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse200101Rules) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse200101Rules) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


