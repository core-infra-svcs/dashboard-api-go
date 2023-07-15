# InlineResponse20087

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency code of this node group&#39;s billing plans | [optional] 
**Plans** | Pointer to [**[]InlineResponse20087Plans**](InlineResponse20087Plans.md) | Array of billing plans in the node group. (Can configure a maximum of 5) | [optional] 

## Methods

### NewInlineResponse20087

`func NewInlineResponse20087() *InlineResponse20087`

NewInlineResponse20087 instantiates a new InlineResponse20087 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20087WithDefaults

`func NewInlineResponse20087WithDefaults() *InlineResponse20087`

NewInlineResponse20087WithDefaults instantiates a new InlineResponse20087 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InlineResponse20087) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineResponse20087) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineResponse20087) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineResponse20087) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *InlineResponse20087) GetPlans() []InlineResponse20087Plans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InlineResponse20087) GetPlansOk() (*[]InlineResponse20087Plans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InlineResponse20087) SetPlans(v []InlineResponse20087Plans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InlineResponse20087) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


