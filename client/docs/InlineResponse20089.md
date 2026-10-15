# InlineResponse20089

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Intrusion detection mode | [optional] 
**IdsRulesets** | Pointer to **string** | Intrusion detection ruleset | [optional] 
**Policy** | Pointer to [**NullableInlineResponse20089Policy**](InlineResponse20089Policy.md) |  | [optional] 
**ProtectedNetworks** | Pointer to [**InlineResponse20089ProtectedNetworks**](InlineResponse20089ProtectedNetworks.md) |  | [optional] 

## Methods

### NewInlineResponse20089

`func NewInlineResponse20089() *InlineResponse20089`

NewInlineResponse20089 instantiates a new InlineResponse20089 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20089WithDefaults

`func NewInlineResponse20089WithDefaults() *InlineResponse20089`

NewInlineResponse20089WithDefaults instantiates a new InlineResponse20089 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20089) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20089) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20089) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20089) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIdsRulesets

`func (o *InlineResponse20089) GetIdsRulesets() string`

GetIdsRulesets returns the IdsRulesets field if non-nil, zero value otherwise.

### GetIdsRulesetsOk

`func (o *InlineResponse20089) GetIdsRulesetsOk() (*string, bool)`

GetIdsRulesetsOk returns a tuple with the IdsRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsRulesets

`func (o *InlineResponse20089) SetIdsRulesets(v string)`

SetIdsRulesets sets IdsRulesets field to given value.

### HasIdsRulesets

`func (o *InlineResponse20089) HasIdsRulesets() bool`

HasIdsRulesets returns a boolean if a field has been set.

### GetPolicy

`func (o *InlineResponse20089) GetPolicy() InlineResponse20089Policy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *InlineResponse20089) GetPolicyOk() (*InlineResponse20089Policy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *InlineResponse20089) SetPolicy(v InlineResponse20089Policy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *InlineResponse20089) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicyNil

`func (o *InlineResponse20089) SetPolicyNil(b bool)`

 SetPolicyNil sets the value for Policy to be an explicit nil

### UnsetPolicy
`func (o *InlineResponse20089) UnsetPolicy()`

UnsetPolicy ensures that no value is present for Policy, not even an explicit nil
### GetProtectedNetworks

`func (o *InlineResponse20089) GetProtectedNetworks() InlineResponse20089ProtectedNetworks`

GetProtectedNetworks returns the ProtectedNetworks field if non-nil, zero value otherwise.

### GetProtectedNetworksOk

`func (o *InlineResponse20089) GetProtectedNetworksOk() (*InlineResponse20089ProtectedNetworks, bool)`

GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNetworks

`func (o *InlineResponse20089) SetProtectedNetworks(v InlineResponse20089ProtectedNetworks)`

SetProtectedNetworks sets ProtectedNetworks field to given value.

### HasProtectedNetworks

`func (o *InlineResponse20089) HasProtectedNetworks() bool`

HasProtectedNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


