# InlineResponse200246

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**InlineResponse200246Exception**](InlineResponse200246Exception.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200246Rules**](InlineResponse200246Rules.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewInlineResponse200246

`func NewInlineResponse200246() *InlineResponse200246`

NewInlineResponse200246 instantiates a new InlineResponse200246 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200246WithDefaults

`func NewInlineResponse200246WithDefaults() *InlineResponse200246`

NewInlineResponse200246WithDefaults instantiates a new InlineResponse200246 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200246) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200246) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200246) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200246) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *InlineResponse200246) GetException() InlineResponse200246Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineResponse200246) GetExceptionOk() (*InlineResponse200246Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineResponse200246) SetException(v InlineResponse200246Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineResponse200246) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200246) GetRules() []InlineResponse200246Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200246) GetRulesOk() (*[]InlineResponse200246Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200246) SetRules(v []InlineResponse200246Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200246) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


