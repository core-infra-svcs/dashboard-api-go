# InlineResponse20092Rules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definitions** | [**[]InlineResponse20092Definitions**](InlineResponse20092Definitions.md) |     A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.  | 
**PerClientBandwidthLimits** | Pointer to [**InlineResponse20092PerClientBandwidthLimits**](InlineResponse20092PerClientBandwidthLimits.md) |  | [optional] 
**DscpTagValue** | Pointer to **NullableInt32** |     The DSCP tag applied by your rule. null means &#39;Do not change DSCP tag&#39;.     For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.  | [optional] 
**Priority** | Pointer to **string** |     A string, indicating the priority level for packets bound to your rule.     Can be &#39;low&#39;, &#39;normal&#39; or &#39;high&#39;.  | [optional] 

## Methods

### NewInlineResponse20092Rules

`func NewInlineResponse20092Rules(definitions []InlineResponse20092Definitions, ) *InlineResponse20092Rules`

NewInlineResponse20092Rules instantiates a new InlineResponse20092Rules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20092RulesWithDefaults

`func NewInlineResponse20092RulesWithDefaults() *InlineResponse20092Rules`

NewInlineResponse20092RulesWithDefaults instantiates a new InlineResponse20092Rules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinitions

`func (o *InlineResponse20092Rules) GetDefinitions() []InlineResponse20092Definitions`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *InlineResponse20092Rules) GetDefinitionsOk() (*[]InlineResponse20092Definitions, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *InlineResponse20092Rules) SetDefinitions(v []InlineResponse20092Definitions)`

SetDefinitions sets Definitions field to given value.


### GetPerClientBandwidthLimits

`func (o *InlineResponse20092Rules) GetPerClientBandwidthLimits() InlineResponse20092PerClientBandwidthLimits`

GetPerClientBandwidthLimits returns the PerClientBandwidthLimits field if non-nil, zero value otherwise.

### GetPerClientBandwidthLimitsOk

`func (o *InlineResponse20092Rules) GetPerClientBandwidthLimitsOk() (*InlineResponse20092PerClientBandwidthLimits, bool)`

GetPerClientBandwidthLimitsOk returns a tuple with the PerClientBandwidthLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerClientBandwidthLimits

`func (o *InlineResponse20092Rules) SetPerClientBandwidthLimits(v InlineResponse20092PerClientBandwidthLimits)`

SetPerClientBandwidthLimits sets PerClientBandwidthLimits field to given value.

### HasPerClientBandwidthLimits

`func (o *InlineResponse20092Rules) HasPerClientBandwidthLimits() bool`

HasPerClientBandwidthLimits returns a boolean if a field has been set.

### GetDscpTagValue

`func (o *InlineResponse20092Rules) GetDscpTagValue() int32`

GetDscpTagValue returns the DscpTagValue field if non-nil, zero value otherwise.

### GetDscpTagValueOk

`func (o *InlineResponse20092Rules) GetDscpTagValueOk() (*int32, bool)`

GetDscpTagValueOk returns a tuple with the DscpTagValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpTagValue

`func (o *InlineResponse20092Rules) SetDscpTagValue(v int32)`

SetDscpTagValue sets DscpTagValue field to given value.

### HasDscpTagValue

`func (o *InlineResponse20092Rules) HasDscpTagValue() bool`

HasDscpTagValue returns a boolean if a field has been set.

### SetDscpTagValueNil

`func (o *InlineResponse20092Rules) SetDscpTagValueNil(b bool)`

 SetDscpTagValueNil sets the value for DscpTagValue to be an explicit nil

### UnsetDscpTagValue
`func (o *InlineResponse20092Rules) UnsetDscpTagValue()`

UnsetDscpTagValue ensures that no value is present for DscpTagValue, not even an explicit nil
### GetPriority

`func (o *InlineResponse20092Rules) GetPriority() string`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *InlineResponse20092Rules) GetPriorityOk() (*string, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *InlineResponse20092Rules) SetPriority(v string)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *InlineResponse20092Rules) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


