# InlineResponse200250

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]InlineResponse200250Rules**](InlineResponse200250Rules.md) | An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule). | [optional] 

## Methods

### NewInlineResponse200250

`func NewInlineResponse200250() *InlineResponse200250`

NewInlineResponse200250 instantiates a new InlineResponse200250 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200250WithDefaults

`func NewInlineResponse200250WithDefaults() *InlineResponse200250`

NewInlineResponse200250WithDefaults instantiates a new InlineResponse200250 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *InlineResponse200250) GetRules() []InlineResponse200250Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200250) GetRulesOk() (*[]InlineResponse200250Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200250) SetRules(v []InlineResponse200250Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200250) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


