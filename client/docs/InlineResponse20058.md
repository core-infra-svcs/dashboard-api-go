# InlineResponse20058

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Intrusion detection mode | [optional] 
**IdsRulesets** | Pointer to **string** | Intrusion detection ruleset | [optional] 
**ProtectedNetworks** | Pointer to [**InlineResponse20058ProtectedNetworks**](InlineResponse20058ProtectedNetworks.md) |  | [optional] 

## Methods

### NewInlineResponse20058

`func NewInlineResponse20058() *InlineResponse20058`

NewInlineResponse20058 instantiates a new InlineResponse20058 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20058WithDefaults

`func NewInlineResponse20058WithDefaults() *InlineResponse20058`

NewInlineResponse20058WithDefaults instantiates a new InlineResponse20058 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineResponse20058) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineResponse20058) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineResponse20058) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineResponse20058) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIdsRulesets

`func (o *InlineResponse20058) GetIdsRulesets() string`

GetIdsRulesets returns the IdsRulesets field if non-nil, zero value otherwise.

### GetIdsRulesetsOk

`func (o *InlineResponse20058) GetIdsRulesetsOk() (*string, bool)`

GetIdsRulesetsOk returns a tuple with the IdsRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsRulesets

`func (o *InlineResponse20058) SetIdsRulesets(v string)`

SetIdsRulesets sets IdsRulesets field to given value.

### HasIdsRulesets

`func (o *InlineResponse20058) HasIdsRulesets() bool`

HasIdsRulesets returns a boolean if a field has been set.

### GetProtectedNetworks

`func (o *InlineResponse20058) GetProtectedNetworks() InlineResponse20058ProtectedNetworks`

GetProtectedNetworks returns the ProtectedNetworks field if non-nil, zero value otherwise.

### GetProtectedNetworksOk

`func (o *InlineResponse20058) GetProtectedNetworksOk() (*InlineResponse20058ProtectedNetworks, bool)`

GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNetworks

`func (o *InlineResponse20058) SetProtectedNetworks(v InlineResponse20058ProtectedNetworks)`

SetProtectedNetworks sets ProtectedNetworks field to given value.

### HasProtectedNetworks

`func (o *InlineResponse20058) HasProtectedNetworks() bool`

HasProtectedNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


