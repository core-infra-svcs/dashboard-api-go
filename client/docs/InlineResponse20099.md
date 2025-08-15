# InlineResponse20099

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20099Products**](InlineResponse20099Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20099Stages**](InlineResponse20099Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20098Reasons**](InlineResponse20098Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20099

`func NewInlineResponse20099() *InlineResponse20099`

NewInlineResponse20099 instantiates a new InlineResponse20099 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20099WithDefaults

`func NewInlineResponse20099WithDefaults() *InlineResponse20099`

NewInlineResponse20099WithDefaults instantiates a new InlineResponse20099 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20099) GetProducts() InlineResponse20099Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20099) GetProductsOk() (*InlineResponse20099Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20099) SetProducts(v InlineResponse20099Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20099) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20099) GetStages() []InlineResponse20099Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20099) GetStagesOk() (*[]InlineResponse20099Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20099) SetStages(v []InlineResponse20099Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20099) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20099) GetReasons() []InlineResponse20098Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20099) GetReasonsOk() (*[]InlineResponse20098Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20099) SetReasons(v []InlineResponse20098Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20099) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


