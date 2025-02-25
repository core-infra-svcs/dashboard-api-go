# InlineResponse200197

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**InlineResponse200197Exception**](InlineResponse200197Exception.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200197Rules**](InlineResponse200197Rules.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewInlineResponse200197

`func NewInlineResponse200197() *InlineResponse200197`

NewInlineResponse200197 instantiates a new InlineResponse200197 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200197WithDefaults

`func NewInlineResponse200197WithDefaults() *InlineResponse200197`

NewInlineResponse200197WithDefaults instantiates a new InlineResponse200197 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200197) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200197) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200197) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200197) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *InlineResponse200197) GetException() InlineResponse200197Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineResponse200197) GetExceptionOk() (*InlineResponse200197Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineResponse200197) SetException(v InlineResponse200197Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineResponse200197) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200197) GetRules() []InlineResponse200197Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200197) GetRulesOk() (*[]InlineResponse200197Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200197) SetRules(v []InlineResponse200197Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200197) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


