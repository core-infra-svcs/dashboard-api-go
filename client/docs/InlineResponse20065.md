# InlineResponse20065

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Intrusion detection mode | [optional] 
**IdsRulesets** | Pointer to **string** | Intrusion detection ruleset | [optional] 
**ProtectedNetworks** | Pointer to [**InlineResponse20065ProtectedNetworks**](InlineResponse20065ProtectedNetworks.md) |  | [optional] 

## Methods

### NewInlineResponse20065

`func NewInlineResponse20065() *InlineResponse20065`

NewInlineResponse20065 instantiates a new InlineResponse20065 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20065WithDefaults

`func NewInlineResponse20065WithDefaults() *InlineResponse20065`

NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20065) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20065) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20065) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20065) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIdsRulesets

`func (o *InlineResponse20065) GetIdsRulesets() string`

GetIdsRulesets returns the IdsRulesets field if non-nil, zero value otherwise.

### GetIdsRulesetsOk

`func (o *InlineResponse20065) GetIdsRulesetsOk() (*string, bool)`

GetIdsRulesetsOk returns a tuple with the IdsRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsRulesets

`func (o *InlineResponse20065) SetIdsRulesets(v string)`

SetIdsRulesets sets IdsRulesets field to given value.

### HasIdsRulesets

`func (o *InlineResponse20065) HasIdsRulesets() bool`

HasIdsRulesets returns a boolean if a field has been set.

### GetProtectedNetworks

`func (o *InlineResponse20065) GetProtectedNetworks() InlineResponse20065ProtectedNetworks`

GetProtectedNetworks returns the ProtectedNetworks field if non-nil, zero value otherwise.

### GetProtectedNetworksOk

`func (o *InlineResponse20065) GetProtectedNetworksOk() (*InlineResponse20065ProtectedNetworks, bool)`

GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNetworks

`func (o *InlineResponse20065) SetProtectedNetworks(v InlineResponse20065ProtectedNetworks)`

SetProtectedNetworks sets ProtectedNetworks field to given value.

### HasProtectedNetworks

`func (o *InlineResponse20065) HasProtectedNetworks() bool`

HasProtectedNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


