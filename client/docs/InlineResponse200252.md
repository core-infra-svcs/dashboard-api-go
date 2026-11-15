# InlineResponse200252

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**InlineResponse200252Exception**](InlineResponse200252Exception.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200252Rules**](InlineResponse200252Rules.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewInlineResponse200252

`func NewInlineResponse200252() *InlineResponse200252`

NewInlineResponse200252 instantiates a new InlineResponse200252 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200252WithDefaults

`func NewInlineResponse200252WithDefaults() *InlineResponse200252`

NewInlineResponse200252WithDefaults instantiates a new InlineResponse200252 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200252) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200252) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200252) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200252) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *InlineResponse200252) GetException() InlineResponse200252Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineResponse200252) GetExceptionOk() (*InlineResponse200252Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineResponse200252) SetException(v InlineResponse200252Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineResponse200252) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200252) GetRules() []InlineResponse200252Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200252) GetRulesOk() (*[]InlineResponse200252Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200252) SetRules(v []InlineResponse200252Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200252) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


