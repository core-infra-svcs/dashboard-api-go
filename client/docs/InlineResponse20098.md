# InlineResponse20098

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20098Products**](InlineResponse20098Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20098Stages**](InlineResponse20098Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20097Reasons**](InlineResponse20097Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20098

`func NewInlineResponse20098() *InlineResponse20098`

NewInlineResponse20098 instantiates a new InlineResponse20098 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20098WithDefaults

`func NewInlineResponse20098WithDefaults() *InlineResponse20098`

NewInlineResponse20098WithDefaults instantiates a new InlineResponse20098 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20098) GetProducts() InlineResponse20098Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20098) GetProductsOk() (*InlineResponse20098Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20098) SetProducts(v InlineResponse20098Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20098) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20098) GetStages() []InlineResponse20098Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20098) GetStagesOk() (*[]InlineResponse20098Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20098) SetStages(v []InlineResponse20098Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20098) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20098) GetReasons() []InlineResponse20097Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20098) GetReasonsOk() (*[]InlineResponse20097Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20098) SetReasons(v []InlineResponse20097Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20098) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


