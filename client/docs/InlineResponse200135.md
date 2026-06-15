# InlineResponse200135

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse200135Products**](InlineResponse200135Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse200135Stages**](InlineResponse200135Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse200134Reasons**](InlineResponse200134Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse200135

`func NewInlineResponse200135() *InlineResponse200135`

NewInlineResponse200135 instantiates a new InlineResponse200135 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse200135WithDefaults

`func NewInlineResponse200135WithDefaults() *InlineResponse200135`

NewInlineResponse200135WithDefaults instantiates a new InlineResponse200135 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse200135) GetProducts() InlineResponse200135Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse200135) GetProductsOk() (*InlineResponse200135Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse200135) SetProducts(v InlineResponse200135Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse200135) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse200135) GetStages() []InlineResponse200135Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse200135) GetStagesOk() (*[]InlineResponse200135Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse200135) SetStages(v []InlineResponse200135Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse200135) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse200135) GetReasons() []InlineResponse200134Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse200135) GetReasonsOk() (*[]InlineResponse200134Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse200135) SetReasons(v []InlineResponse200134Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse200135) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


