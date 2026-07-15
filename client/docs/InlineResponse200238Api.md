# InlineResponse200238Api

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether to enable push API for scanning events. Analytics must be enabled to enable scanning API. | [optional] 
**Validator** | Pointer to [**InlineResponse200238ApiValidator**](InlineResponse200238ApiValidator.md) |  | [optional] 

## Methods

### NewInlineResponse200238Api

`func NewInlineResponse200238Api() *InlineResponse200238Api`

NewInlineResponse200238Api instantiates a new InlineResponse200238Api object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200238ApiWithDefaults

`func NewInlineResponse200238ApiWithDefaults() *InlineResponse200238Api`

NewInlineResponse200238ApiWithDefaults instantiates a new InlineResponse200238Api object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *InlineResponse200238Api) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InlineResponse200238Api) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InlineResponse200238Api) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InlineResponse200238Api) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetValidator

`func (o *InlineResponse200238Api) GetValidator() InlineResponse200238ApiValidator`

GetValidator returns the Validator field if non-nil, zero value otherwise.

### GetValidatorOk

`func (o *InlineResponse200238Api) GetValidatorOk() (*InlineResponse200238ApiValidator, bool)`

GetValidatorOk returns a tuple with the Validator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidator

`func (o *InlineResponse200238Api) SetValidator(v InlineResponse200238ApiValidator)`

SetValidator sets Validator field to given value.

### HasValidator

`func (o *InlineResponse200238Api) HasValidator() bool`

HasValidator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


