# InlineResponse200529SsidAddressingBonjourForwardingRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the rule | [optional] 
**Vlan** | Pointer to [**InlineResponse200529SsidAddressingBonjourForwardingVlan**](InlineResponse200529SsidAddressingBonjourForwardingVlan.md) |  | [optional] 
**Services** | Pointer to **[]string** | Services included in the rule | [optional] 

## Methods

### NewInlineResponse200529SsidAddressingBonjourForwardingRules

`func NewInlineResponse200529SsidAddressingBonjourForwardingRules() *InlineResponse200529SsidAddressingBonjourForwardingRules`

NewInlineResponse200529SsidAddressingBonjourForwardingRules instantiates a new InlineResponse200529SsidAddressingBonjourForwardingRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200529SsidAddressingBonjourForwardingRulesWithDefaults

`func NewInlineResponse200529SsidAddressingBonjourForwardingRulesWithDefaults() *InlineResponse200529SsidAddressingBonjourForwardingRules`

NewInlineResponse200529SsidAddressingBonjourForwardingRulesWithDefaults instantiates a new InlineResponse200529SsidAddressingBonjourForwardingRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVlan

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) GetVlan() InlineResponse200529SsidAddressingBonjourForwardingVlan`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) GetVlanOk() (*InlineResponse200529SsidAddressingBonjourForwardingVlan, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) SetVlan(v InlineResponse200529SsidAddressingBonjourForwardingVlan)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetServices

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *InlineResponse200529SsidAddressingBonjourForwardingRules) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


