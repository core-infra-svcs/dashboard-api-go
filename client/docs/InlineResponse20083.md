# InlineResponse20083

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Intrusion detection mode | [optional] 
**IdsRulesets** | Pointer to **string** | Intrusion detection ruleset | [optional] 
**ProtectedNetworks** | Pointer to [**InlineResponse20083ProtectedNetworks**](InlineResponse20083ProtectedNetworks.md) |  | [optional] 

## Methods

### NewInlineResponse20083

`func NewInlineResponse20083() *InlineResponse20083`

NewInlineResponse20083 instantiates a new InlineResponse20083 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20083WithDefaults

`func NewInlineResponse20083WithDefaults() *InlineResponse20083`

NewInlineResponse20083WithDefaults instantiates a new InlineResponse20083 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20083) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20083) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20083) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20083) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIdsRulesets

`func (o *InlineResponse20083) GetIdsRulesets() string`

GetIdsRulesets returns the IdsRulesets field if non-nil, zero value otherwise.

### GetIdsRulesetsOk

`func (o *InlineResponse20083) GetIdsRulesetsOk() (*string, bool)`

GetIdsRulesetsOk returns a tuple with the IdsRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsRulesets

`func (o *InlineResponse20083) SetIdsRulesets(v string)`

SetIdsRulesets sets IdsRulesets field to given value.

### HasIdsRulesets

`func (o *InlineResponse20083) HasIdsRulesets() bool`

HasIdsRulesets returns a boolean if a field has been set.

### GetProtectedNetworks

`func (o *InlineResponse20083) GetProtectedNetworks() InlineResponse20083ProtectedNetworks`

GetProtectedNetworks returns the ProtectedNetworks field if non-nil, zero value otherwise.

### GetProtectedNetworksOk

`func (o *InlineResponse20083) GetProtectedNetworksOk() (*InlineResponse20083ProtectedNetworks, bool)`

GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNetworks

`func (o *InlineResponse20083) SetProtectedNetworks(v InlineResponse20083ProtectedNetworks)`

SetProtectedNetworks sets ProtectedNetworks field to given value.

### HasProtectedNetworks

`func (o *InlineResponse20083) HasProtectedNetworks() bool`

HasProtectedNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


