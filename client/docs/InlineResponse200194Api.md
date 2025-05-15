# InlineResponse200194Api

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable push API for scanning events. Analytics must be enabled to enable scanning API. | [optional] 
**Validator** | Pointer to [**InlineResponse200194ApiValidator**](InlineResponse200194ApiValidator.md) |  | [optional] 

## Methods

### NewInlineResponse200194Api

`func NewInlineResponse200194Api() *InlineResponse200194Api`

NewInlineResponse200194Api instantiates a new InlineResponse200194Api object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200194ApiWithDefaults

`func NewInlineResponse200194ApiWithDefaults() *InlineResponse200194Api`

NewInlineResponse200194ApiWithDefaults instantiates a new InlineResponse200194Api object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200194Api) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200194Api) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200194Api) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200194Api) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetValidator

`func (o *InlineResponse200194Api) GetValidator() InlineResponse200194ApiValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *InlineResponse200194Api) GetValidatorOk() (*InlineResponse200194ApiValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *InlineResponse200194Api) SetValidator(v InlineResponse200194ApiValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *InlineResponse200194Api) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


