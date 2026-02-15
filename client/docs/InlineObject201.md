# InlineObject201

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200213Rules**](InlineResponse200213Rules.md) | An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration. | [optional] 

## Methods

### NewInlineObject201

`func NewInlineObject201() *InlineObject201`

NewInlineObject201 instantiates a new InlineObject201 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject201WithDefaults

`func NewInlineObject201WithDefaults() *InlineObject201`

NewInlineObject201WithDefaults instantiates a new InlineObject201 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineObject201) GetRules() []InlineResponse200213Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineObject201) GetRulesOk() (*[]InlineResponse200213Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineObject201) SetRules(v []InlineResponse200213Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineObject201) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


