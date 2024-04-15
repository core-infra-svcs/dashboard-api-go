# InlineResponse200183

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**InlineResponse200183Exception**](InlineResponse200183Exception.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200183Rules**](InlineResponse200183Rules.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewInlineResponse200183

`func NewInlineResponse200183() *InlineResponse200183`

NewInlineResponse200183 instantiates a new InlineResponse200183 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200183WithDefaults

`func NewInlineResponse200183WithDefaults() *InlineResponse200183`

NewInlineResponse200183WithDefaults instantiates a new InlineResponse200183 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200183) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200183) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200183) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200183) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *InlineResponse200183) GetException() InlineResponse200183Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineResponse200183) GetExceptionOk() (*InlineResponse200183Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineResponse200183) SetException(v InlineResponse200183Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineResponse200183) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200183) GetRules() []InlineResponse200183Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200183) GetRulesOk() (*[]InlineResponse200183Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200183) SetRules(v []InlineResponse200183Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200183) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


