# InlineResponse200198Api

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable push API for scanning events. Analytics must be enabled to enable scanning API. | [optional] 
**Validator** | Pointer to [**InlineResponse200198ApiValidator**](InlineResponse200198ApiValidator.md) |  | [optional] 

## Methods

### NewInlineResponse200198Api

`func NewInlineResponse200198Api() *InlineResponse200198Api`

NewInlineResponse200198Api instantiates a new InlineResponse200198Api object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200198ApiWithDefaults

`func NewInlineResponse200198ApiWithDefaults() *InlineResponse200198Api`

NewInlineResponse200198ApiWithDefaults instantiates a new InlineResponse200198Api object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200198Api) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200198Api) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200198Api) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200198Api) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetValidator

`func (o *InlineResponse200198Api) GetValidator() InlineResponse200198ApiValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *InlineResponse200198Api) GetValidatorOk() (*InlineResponse200198ApiValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *InlineResponse200198Api) SetValidator(v InlineResponse200198ApiValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *InlineResponse200198Api) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


