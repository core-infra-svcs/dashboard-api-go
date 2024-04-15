# InlineObject184

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200186Rules**](InlineResponse200186Rules.md) | An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration. | [optional] 

## Methods

### NewInlineObject184

`func NewInlineObject184() *InlineObject184`

NewInlineObject184 instantiates a new InlineObject184 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject184WithDefaults

`func NewInlineObject184WithDefaults() *InlineObject184`

NewInlineObject184WithDefaults instantiates a new InlineObject184 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject184) GetRules() []InlineResponse200186Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject184) GetRulesOk() (*[]InlineResponse200186Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject184) SetRules(v []InlineResponse200186Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject184) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


