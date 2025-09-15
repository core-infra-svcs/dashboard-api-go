# InlineObject61

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Set mode to &#39;disabled&#39;/&#39;detection&#39;/&#39;prevention&#39; (optional - omitting will leave current config unchanged) | [optional] 
**IdsRulesets** | Pointer to **string** | Set the detection ruleset &#39;connectivity&#39;/&#39;balanced&#39;/&#39;security&#39; (optional - omitting will leave current config unchanged). Default value is &#39;balanced&#39; if none currently saved | [optional] 
**ProtectedNetworks** | Pointer to [**NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks**](NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks.md) |  | [optional] 

## Methods

### NewInlineObject61

`func NewInlineObject61() *InlineObject61`

NewInlineObject61 instantiates a new InlineObject61 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject61WithDefaults

`func NewInlineObject61WithDefaults() *InlineObject61`

NewInlineObject61WithDefaults instantiates a new InlineObject61 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *InlineObject61) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *InlineObject61) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *InlineObject61) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *InlineObject61) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetIdsRulesets

`func (o *InlineObject61) GetIdsRulesets() string`

GetIdsRulesets returns the IdsRulesets field if non-nil, zero value otherwise.

### GetIdsRulesetsOk

`func (o *InlineObject61) GetIdsRulesetsOk() (*string, bool)`

GetIdsRulesetsOk returns a tuple with the IdsRulesets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdsRulesets

`func (o *InlineObject61) SetIdsRulesets(v string)`

SetIdsRulesets sets IdsRulesets field to given value.

### HasIdsRulesets

`func (o *InlineObject61) HasIdsRulesets() bool`

HasIdsRulesets returns a boolean if a field has been set.

### GetProtectedNetworks

`func (o *InlineObject61) GetProtectedNetworks() NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks`

GetProtectedNetworks returns the ProtectedNetworks field if non-nil, zero value otherwise.

### GetProtectedNetworksOk

`func (o *InlineObject61) GetProtectedNetworksOk() (*NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks, bool)`

GetProtectedNetworksOk returns a tuple with the ProtectedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedNetworks

`func (o *InlineObject61) SetProtectedNetworks(v NetworksNetworkIdApplianceSecurityIntrusionProtectedNetworks)`

SetProtectedNetworks sets ProtectedNetworks field to given value.

### HasProtectedNetworks

`func (o *InlineObject61) HasProtectedNetworks() bool`

HasProtectedNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


