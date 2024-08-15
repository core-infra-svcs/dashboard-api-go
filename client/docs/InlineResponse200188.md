# InlineResponse200188

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If true, Bonjour forwarding is enabled on the SSID. | [optional] 
**Exception** | Pointer to [**InlineResponse200188Exception**](InlineResponse200188Exception.md) |  | [optional] 
**Rules** | Pointer to [**[]InlineResponse200188Rules**](InlineResponse200188Rules.md) | Bonjour forwarding rules | [optional] 

## Methods

### NewInlineResponse200188

`func NewInlineResponse200188() *InlineResponse200188`

NewInlineResponse200188 instantiates a new InlineResponse200188 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200188WithDefaults

`func NewInlineResponse200188WithDefaults() *InlineResponse200188`

NewInlineResponse200188WithDefaults instantiates a new InlineResponse200188 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200188) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200188) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200188) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200188) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetException

`func (o *InlineResponse200188) GetException() InlineResponse200188Exception`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *InlineResponse200188) GetExceptionOk() (*InlineResponse200188Exception, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *InlineResponse200188) SetException(v InlineResponse200188Exception)`

SetException sets Exception field to given value.

### HasException

`func (o *InlineResponse200188) HasException() bool`

HasException returns a boolean if a field has been set.

### GetRules

`func (o *InlineResponse200188) GetRules() []InlineResponse200188Rules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *InlineResponse200188) GetRulesOk() (*[]InlineResponse200188Rules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *InlineResponse200188) SetRules(v []InlineResponse200188Rules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *InlineResponse200188) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


