# InlineResponse20023

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20023Products**](InlineResponse20023Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20023Stages**](InlineResponse20023Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20022Reasons**](InlineResponse20022Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20023

`func NewInlineResponse20023() *InlineResponse20023`

NewInlineResponse20023 instantiates a new InlineResponse20023 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20023WithDefaults

`func NewInlineResponse20023WithDefaults() *InlineResponse20023`

NewInlineResponse20023WithDefaults instantiates a new InlineResponse20023 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20023) GetProducts() InlineResponse20023Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20023) GetProductsOk() (*InlineResponse20023Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20023) SetProducts(v InlineResponse20023Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20023) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20023) GetStages() []InlineResponse20023Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20023) GetStagesOk() (*[]InlineResponse20023Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20023) SetStages(v []InlineResponse20023Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20023) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20023) GetReasons() []InlineResponse20022Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20023) GetReasonsOk() (*[]InlineResponse20022Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20023) SetReasons(v []InlineResponse20022Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20023) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


