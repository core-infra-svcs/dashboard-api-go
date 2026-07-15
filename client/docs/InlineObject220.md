# InlineObject220

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200248Rules**](InlineResponse200248Rules.md) | An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration. | [optional] 

## Methods

### NewInlineObject220

`func NewInlineObject220() *InlineObject220`

NewInlineObject220 instantiates a new InlineObject220 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject220WithDefaults

`func NewInlineObject220WithDefaults() *InlineObject220`

NewInlineObject220WithDefaults instantiates a new InlineObject220 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject220) GetRules() []InlineResponse200248Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject220) GetRulesOk() (*[]InlineResponse200248Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject220) SetRules(v []InlineResponse200248Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject220) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


