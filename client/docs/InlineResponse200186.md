# InlineResponse200186

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200186Rules**](InlineResponse200186Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 

## Methods

### NewInlineResponse200186

`func NewInlineResponse200186() *InlineResponse200186`

NewInlineResponse200186 instantiates a new InlineResponse200186 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200186WithDefaults

`func NewInlineResponse200186WithDefaults() *InlineResponse200186`

NewInlineResponse200186WithDefaults instantiates a new InlineResponse200186 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200186) GetRules() []InlineResponse200186Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200186) GetRulesOk() (*[]InlineResponse200186Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200186) SetRules(v []InlineResponse200186Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200186) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


