# InlineResponse20033

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**InlineResponse20033Products**](InlineResponse20033Products.md) |  | [optional] 
**Stages** | Pointer to [**[]InlineResponse20033Stages**](InlineResponse20033Stages.md) | The ordered stages in the network | [optional] 
**Reasons** | Pointer to [**[]InlineResponse20032Reasons**](InlineResponse20032Reasons.md) | Reasons for the rollback | [optional] 

## Methods

### NewInlineResponse20033

`func NewInlineResponse20033() *InlineResponse20033`

NewInlineResponse20033 instantiates a new InlineResponse20033 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20033WithDefaults

`func NewInlineResponse20033WithDefaults() *InlineResponse20033`

NewInlineResponse20033WithDefaults instantiates a new InlineResponse20033 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InlineResponse20033) GetProducts() InlineResponse20033Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InlineResponse20033) GetProductsOk() (*InlineResponse20033Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InlineResponse20033) SetProducts(v InlineResponse20033Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InlineResponse20033) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetStages

`func (o *InlineResponse20033) GetStages() []InlineResponse20033Stages`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *InlineResponse20033) GetStagesOk() (*[]InlineResponse20033Stages, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *InlineResponse20033) SetStages(v []InlineResponse20033Stages)`

SetStages sets Stages field to given value.

### HasStages

`func (o *InlineResponse20033) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetReasons

`func (o *InlineResponse20033) GetReasons() []InlineResponse20032Reasons`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *InlineResponse20033) GetReasonsOk() (*[]InlineResponse20032Reasons, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *InlineResponse20033) SetReasons(v []InlineResponse20032Reasons)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *InlineResponse20033) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


