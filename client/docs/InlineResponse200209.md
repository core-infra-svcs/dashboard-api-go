# InlineResponse200209

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**InlineResponse200209Exception**](InlineResponse200209Exception.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200209Rules**](InlineResponse200209Rules.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewInlineResponse200209

`func NewInlineResponse200209() *InlineResponse200209`

NewInlineResponse200209 instantiates a new InlineResponse200209 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200209WithDefaults

`func NewInlineResponse200209WithDefaults() *InlineResponse200209`

NewInlineResponse200209WithDefaults instantiates a new InlineResponse200209 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200209) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200209) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200209) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200209) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *InlineResponse200209) GetException() InlineResponse200209Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineResponse200209) GetExceptionOk() (*InlineResponse200209Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineResponse200209) SetException(v InlineResponse200209Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineResponse200209) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200209) GetRules() []InlineResponse200209Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200209) GetRulesOk() (*[]InlineResponse200209Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200209) SetRules(v []InlineResponse200209Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200209) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


