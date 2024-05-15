# InlineResponse20090

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20090Products**](InlineResponse20090Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20090Stages**](InlineResponse20090Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20089Reasons**](InlineResponse20089Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20090

`func NewInlineResponse20090() *InlineResponse20090`

NewInlineResponse20090 instantiates a new InlineResponse20090 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20090WithDefaults

`func NewInlineResponse20090WithDefaults() *InlineResponse20090`

NewInlineResponse20090WithDefaults instantiates a new InlineResponse20090 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20090) GetProducts() InlineResponse20090Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20090) GetProductsOk() (*InlineResponse20090Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20090) SetProducts(v InlineResponse20090Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20090) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20090) GetStages() []InlineResponse20090Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20090) GetStagesOk() (*[]InlineResponse20090Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20090) SetStages(v []InlineResponse20090Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20090) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20090) GetReasons() []InlineResponse20089Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20090) GetReasonsOk() (*[]InlineResponse20089Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20090) SetReasons(v []InlineResponse20089Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20090) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


