# InlineResponse200362Cap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **int32** | The total bandwidth available for the specified term in megabytes. | [optional] 
**Threshold** | Pointer to **float32** | The data usage threshold at which the rule should be triggered. This is expressed as a float between 0.01 and 1.0 (where 0.01 is equal to 1% of total and 1.0 is equal to 100%). By default this will be 1.0. | [optional] 
**Term** | Pointer to [**InlineResponse200362CapTerm**](InlineResponse200362CapTerm.md) |  | [optional] 

## Methods

### NewInlineResponse200362Cap

`func NewInlineResponse200362Cap() *InlineResponse200362Cap`

NewInlineResponse200362Cap instantiates a new InlineResponse200362Cap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200362CapWithDefaults

`func NewInlineResponse200362CapWithDefaults() *InlineResponse200362Cap`

NewInlineResponse200362CapWithDefaults instantiates a new InlineResponse200362Cap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *InlineResponse200362Cap) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InlineResponse200362Cap) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *InlineResponse200362Cap) SetValue(v int32)`

SetValue sets Value field to given value.

### HasValue

`func (o *InlineResponse200362Cap) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetThreshold

`func (o *InlineResponse200362Cap) GetThreshold() float32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *InlineResponse200362Cap) GetThresholdOk() (*float32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *InlineResponse200362Cap) SetThreshold(v float32)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *InlineResponse200362Cap) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetTerm

`func (o *InlineResponse200362Cap) GetTerm() InlineResponse200362CapTerm`

GetTerm returns the Term field if non-nil, zero value otherwise.

### GetTermOk

`func (o *InlineResponse200362Cap) GetTermOk() (*InlineResponse200362CapTerm, bool)`

GetTermOk returns a tuple with the Term field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerm

`func (o *InlineResponse200362Cap) SetTerm(v InlineResponse200362CapTerm)`

SetTerm sets Term field to given value.

### HasTerm

`func (o *InlineResponse200362Cap) HasTerm() bool`

HasTerm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


