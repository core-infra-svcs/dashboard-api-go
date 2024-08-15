# InlineResponse200169

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency code of this node group&#39;s billing plans | [optional] 
**Plans** | Pointer to [**[]InlineResponse200169Plans**](InlineResponse200169Plans.md) | Array of billing plans in the node group. (Can configure a maximum of 5) | [optional] 

## Methods

### NewInlineResponse200169

`func NewInlineResponse200169() *InlineResponse200169`

NewInlineResponse200169 instantiates a new InlineResponse200169 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200169WithDefaults

`func NewInlineResponse200169WithDefaults() *InlineResponse200169`

NewInlineResponse200169WithDefaults instantiates a new InlineResponse200169 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InlineResponse200169) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineResponse200169) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineResponse200169) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineResponse200169) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *InlineResponse200169) GetPlans() []InlineResponse200169Plans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InlineResponse200169) GetPlansOk() (*[]InlineResponse200169Plans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InlineResponse200169) SetPlans(v []InlineResponse200169Plans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InlineResponse200169) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


