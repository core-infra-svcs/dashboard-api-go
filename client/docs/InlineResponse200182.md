# InlineResponse200182

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | The currency code of this node group&#39;s billing plans | [optional] 
**Plans** | Pointer to [**[]InlineResponse200182Plans**](InlineResponse200182Plans.md) | Array of billing plans in the node group. (Can configure a maximum of 5) | [optional] 

## Methods

### NewInlineResponse200182

`func NewInlineResponse200182() *InlineResponse200182`

NewInlineResponse200182 instantiates a new InlineResponse200182 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200182WithDefaults

`func NewInlineResponse200182WithDefaults() *InlineResponse200182`

NewInlineResponse200182WithDefaults instantiates a new InlineResponse200182 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *InlineResponse200182) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineResponse200182) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineResponse200182) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *InlineResponse200182) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPlans

`func (o *InlineResponse200182) GetPlans() []InlineResponse200182Plans`

GetPlans returns the Plans field if non-nil, zero value otherwise.

### GetPlansOk

`func (o *InlineResponse200182) GetPlansOk() (*[]InlineResponse200182Plans, bool)`

GetPlansOk returns a tuple with the Plans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlans

`func (o *InlineResponse200182) SetPlans(v []InlineResponse200182Plans)`

SetPlans sets Plans field to given value.

### HasPlans

`func (o *InlineResponse200182) HasPlans() bool`

HasPlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


