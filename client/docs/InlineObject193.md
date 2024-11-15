# InlineObject193

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200196Rules**](InlineResponse200196Rules.md) | An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration. | [optional] 

## Methods

### NewInlineObject193

`func NewInlineObject193() *InlineObject193`

NewInlineObject193 instantiates a new InlineObject193 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject193WithDefaults

`func NewInlineObject193WithDefaults() *InlineObject193`

NewInlineObject193WithDefaults instantiates a new InlineObject193 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject193) GetRules() []InlineResponse200196Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject193) GetRulesOk() (*[]InlineResponse200196Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject193) SetRules(v []InlineResponse200196Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject193) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


