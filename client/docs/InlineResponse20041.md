# InlineResponse20041

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20041Products**](InlineResponse20041Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20041Stages**](InlineResponse20041Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20040Reasons**](InlineResponse20040Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20041

`func NewInlineResponse20041() *InlineResponse20041`

NewInlineResponse20041 instantiates a new InlineResponse20041 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20041WithDefaults

`func NewInlineResponse20041WithDefaults() *InlineResponse20041`

NewInlineResponse20041WithDefaults instantiates a new InlineResponse20041 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20041) GetProducts() InlineResponse20041Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20041) GetProductsOk() (*InlineResponse20041Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20041) SetProducts(v InlineResponse20041Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20041) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20041) GetStages() []InlineResponse20041Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20041) GetStagesOk() (*[]InlineResponse20041Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20041) SetStages(v []InlineResponse20041Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20041) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20041) GetReasons() []InlineResponse20040Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20041) GetReasonsOk() (*[]InlineResponse20040Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20041) SetReasons(v []InlineResponse20040Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20041) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


