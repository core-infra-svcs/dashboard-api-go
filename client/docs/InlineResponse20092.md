# InlineResponse20092

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20092Products**](InlineResponse20092Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20092Stages**](InlineResponse20092Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20091Reasons**](InlineResponse20091Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20092

`func NewInlineResponse20092() *InlineResponse20092`

NewInlineResponse20092 instantiates a new InlineResponse20092 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20092WithDefaults

`func NewInlineResponse20092WithDefaults() *InlineResponse20092`

NewInlineResponse20092WithDefaults instantiates a new InlineResponse20092 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20092) GetProducts() InlineResponse20092Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20092) GetProductsOk() (*InlineResponse20092Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20092) SetProducts(v InlineResponse20092Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20092) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20092) GetStages() []InlineResponse20092Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20092) GetStagesOk() (*[]InlineResponse20092Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20092) SetStages(v []InlineResponse20092Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20092) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20092) GetReasons() []InlineResponse20091Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20092) GetReasonsOk() (*[]InlineResponse20091Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20092) SetReasons(v []InlineResponse20091Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20092) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


